# Section-4

## Stand up vulnerable and non-vulnerable JBOSS servers

## Stand up Attack Host with exploit tools

## Using Attack Host to exploit

## Destroying the environment

## Introducing Target, Attack Surface and Automated Testing Methodology
* Domain:
    * Domain points to an Apache Tomcat server.
    * Domain has a github org with members.
* We will scan all the repositories of this org and all the repositories of the org's members using `repo-supervisor`. Results will be stored in a Google BigQuery table.
* We will then run `wfuzz` to do a focussed bruteforcing for Apache Tomcat endpoints on that domain. Results will be stored in a Google BigQuery table.
* We will then use the secrets obtained from `repo-supervisor` and try to bruteforce the basic authentication mechanism of the Apache Tomcat endpoints obtained from `wfuzz`.
* If there is a match, we will get back results in Slack via an incoming webhook.
* Demo of doing all this automatically.

### Running tools locally
* [Running repo-supervisor and wfuzz locally and converting the results to store in BigQuery](data-converter/README.md)
* [Running wfuzz basic authentication bruteforcer combining the data from both the tools](wfuzz-basicauth-bruteforcer/README.md)

### Running tools on a K8S cluster
* Show k8s deployment of both wfuzz and reposupervisor pods simultaneously manually using yaml files.
* Once they finish, show k8s deployment of bruteforcer using yaml file
* If more time, introduce Kubebot and show how it can all be automated - using slack to start the workflow, using channels/queues to communicate b/w the tools and bruteforcer & sending results back to slack

### Sending a request from Kubebot for a target company from mobile using Kubebot
* API server receives the request
* API server drops a message in the queue to start repo-supervisor against that company’s github and waits for it to finish
* API server drops a message in the queue to start wfuzz against that company’s main domain and waits for it to finish
* Repo-supervisor finishes running and stores the results in BQ.
* API server checks the status of repo-supervisor container and sends back a success message in the channel.
* WFUZZ finishes running and stores the results in BQ.
* API server checks the status of wfuzz container and sends back a success message in the channel.
* API server waits for both success messages from the channel. Once received, drops a message in the queue to start the special worker
* Special worker queries WFUZZ dataset from BQ for all tomcat related endpoints - /manager, /admin, /console, etc.
* If found, Special worker queries Repo-supervisor dataset from BQ for all secrets.
* Special worker then tries to bruteforce the endpoint with the secret.
* If successful, special worker sends back the response to Slack.