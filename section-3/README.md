# Section-3

## Switching back context to Minikube

* `kubectl config use-context minikube`

## Running NMAP on the local K8S cluster

* `kubectl apply -f deployments/nmap-deployment.yaml`

## Google PubSub in action

* `virtualenv env`
* `. env/bin/activate`
* `pip install --upgrade google-cloud-pubsub`
* `gcloud config list` - Verify your account, project and active configuration are correctly setup
* `python scripts/createtopicandsub.py` - Creating the topic and subscription
* `python scripts/sendtotopic.py` - Sending the message to the topic
* `python scripts/listenfromsub.py` - Listening for that message from the subscription
* `python scripts/deletetopicandsub.py` - Deleting the topic and subscription

References:
* https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-python

## Convert NMAP data into BigQuery ingest-able format using a Data Converter

### Running Locally
* Create a BigQuery dataset `nmapds` and an empty table `nmap` in Google BigQuery from the GCP UI to store the processed nmap results with the following schema (all nullable):
```
ip:string
fqdn:string
port:string
protocol:string
service:string
version:string
```
* `nmap -Pn -p 1-1000 -oN google_results.nmap google.com` - running nmap locally
* Complete the `.env` file with the appropriate `PROJECT_ID`, `DATASET_NAME` and `TABLE_NAME`
* `go run dataconvert.go google_results.nmap` - Run the data converter locally

### Running on a K8S cluster
* Navigate to `IAM & Admin` -> `Service Accounts`. Create a key for the default Compute Engine Service Account and download the JSON key
* `kubectl create secret generic googlesecret --from-file=$(CREDS_FILEPATH)` - Create a secret with the value of the secret being the JSON credentials file downloaded above. We need this because the containers on the cluster need to authenticate to our K8S cluster to be able to create anything. We don't do this locally because our gcloud environment, by default, is already configured when we first set it up but we need it when running on a K8S cluster
* `kubectl get secrets` - Verify the secret was created
* `kubectl apply -f ../deployments/nmap-bq-deployment.yaml`

References:
* https://github.com/maaaaz/nmaptocsv

## Querying BigQuery for data

## Running Cronjobs
