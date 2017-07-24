# Section-3

## Table of Contents

<!-- TOC -->

- [Section-3](#section-3)
    - [Table of Contents](#table-of-contents)
    - [Overview](#overview)
    - [Switching back context to Minikube](#switching-back-context-to-minikube)
    - [Running NMAP on the local K8S cluster](#running-nmap-on-the-local-k8s-cluster)
    - [Google PubSub in action](#google-pubsub-in-action)
    - [Convert NMAP data into BigQuery ingest-able format using a Data Converter](#convert-nmap-data-into-bigquery-ingest-able-format-using-a-data-converter)
        - [Running Locally](#running-locally)
        - [Running on a K8S cluster](#running-on-a-k8s-cluster)
    - [Querying BigQuery](#querying-bigquery)
    - [Running Cronjobs](#running-cronjobs)
    - [Cleanup](#cleanup)

<!-- /TOC -->

## Overview
In this section, we will
1. Switch our `kubectl` context back from the remote K8S cluster (on GCP) to Minikube.
2. Run a `NMAP` deployment in our local K8S cluster inside Minikube.
3. Show an example of how `topics` and `subscriptions` work in Google PubSub.
4. Convert NMAP data into a format that can be ingested into Google BigQuery using a custom Data Converter.
5. Show how to run queries on Google BigQuery.
6. Run a cronjob by running NMAP at a specific schedule against a target.
7. Cleanup everything we setup above.


## Switching back context to Minikube

* `kubectl config use-context minikube`

Type `kubectl get nodes` and make sure it says `minikube` to confirm if the context has been switched properly. If this does not work, manually edit the `~/.kube/config` file and change the `current` context to say `minikube`.

## Running NMAP on the local K8S cluster

* `cd` into `section-3` directory and type `kubectl apply -f deployments/nmap-deployment.yaml`
* Delete the deployment by typing - `kubectl delete deployments --all`

## Google PubSub in action
Before we could go ahead with this section, we need to make sure we have the correct environment setup. In your GCP cloud console, navigate to `IAM & Admin` -> `Service Accounts`. Create a key for the default Compute Engine Service Account and download the JSON key. Set an environment variable `GOOGLE_APPLICATION_CREDENTIALS` with the value being the location where you save the JSON key.

Reference: [link](https://developers.google.com/identity/protocols/application-default-credentials)

1. `virtualenv env`
2. `. env/bin/activate`
3. `pip install --upgrade google-cloud-pubsub`
4. `gcloud config list` - Verify your account, project and active configuration are correctly setup
5. `python scripts/createtopicandsub.py` - Creating the topic and subscription. Verify that the topic and subscription were created by navigating to the GCP Cloud console and then `Pub/Sub`.
6. `python scripts/sendtotopic.py` - Sending the message to the topic
7. `python scripts/listenfromsub.py` - Listening for that message from the subscription
8. `python scripts/deletetopicandsub.py` - Deleting the topic and subscription

Reference: [link](https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-python)

## Convert NMAP data into BigQuery ingest-able format using a Data Converter

### Running Locally
1. Create a BigQuery dataset `nmapds` and an empty table `nmap` in Google BigQuery from the GCP UI to store the processed nmap results with the following schema (all nullable):
```
ip:string
fqdn:string
port:string
protocol:string
service:string
version:string
```
2. Install `NMAP` if you don't already have it and run NMAP with the following command - `nmap -Pn -p 1-1000 -oN google_results.nmap google.com`.
3. `cd` into the `data-converter` directory. Complete the `.env.sample` file with the appropriate `PROJECT_ID` and copy it to `.env` in the same directory.
4. In that folder, type `go get cloud.google.com/go/bigquery` and then `go get github.com/subosito/gotenv` and finally `go run dataconvert.go ../google_results.nmap` - This command basically converts the nmap output into a BigQuery ingest-able format (csv) and uploads the data into the `nmap` table of the `nmapds` dataset.
5. Verify that the data was successfully uploaded to BigQuery by navigating to GCP Cloud console -> BigQuery -> `nmap` table.

### Running on a local K8S cluster (Minikube)
1. `kubectl create secret generic googlesecret --from-file=$(CREDS_FILEPATH)` - Create a secret with the value of the secret being the JSON credentials file downloaded above. We need this because the containers on the cluster need to authenticate to our K8S cluster to be able to create anything.
2. `kubectl get secrets` - Verify the secret was created.
3. `cd` back into the `section-3` directory.
4. Make sure the environment values in the `deployments/nmap-bq-pod.yaml` deployment file are accurate. The following values need to be changed - [PROJECT_ID](https://github.com/devsecops/defcon-workshop/blob/master/section-3/deployments/nmap-bq-pod.yaml#L14) and [GOOGLE_APPLICATION_CREDENTIALS].(https://github.com/devsecops/defcon-workshop/blob/master/section-3/deployments/nmap-bq-pod.yaml#L20)
5. Now, in order to automate what we did in the section before this i.e. run nmap and then run the data converter to upload the nmap results to BQ all together, type `kubectl apply -f deployments/nmap-bq-pod.yaml`
6. Verify that there are new rows in the BQ table now after the Pod completes running.

Reference: [link](https://github.com/maaaaz/nmaptocsv)


## Querying BigQuery

* From your GCP cloud console, navigate to BigQuery and run the below query:
```
SELECT ip, port FROM nmapds.nmap
WHERE ip IS NOT NULL AND port IS NOT NULL
GROUP BY ip, port
```

## Running Cronjobs

1. Make sure the environment values in the `deployments/nmap-cronjob.yaml` deployment file are accurate.
2. `kubectl apply -f deployments/nmap-cronjob.yaml` - Start the cronjob
3. `kubectl get cronjobs --watch` - Watch the status of the cronjob
4. Whenever the cronjob kicks in, type `minikube dashboard` and notice how a new job appears, runs and terminates. You can also refresh the BigQuery table to verify new rows are added on every run.

## Cleanup
1. `kubectl delete cronjobs --all`
2. `kubectl delete pods --all`
3. `kubectl delete jobs --all`
4. `bq rm nmapds.nmap` and `bq rm -r -f nmapds` - Delete the BigQuery dataset `nmapds` and table `nmap`
5. `deactivate` to come out of the virtualenv
