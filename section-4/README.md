# Section-4

## Stand up vulnerable and non-vulnerable JBOSS servers

## Stand up Attack Host with exploit tools

## Using Attack Host to exploit

## Destroying the environment

## Introducing and Setting up Kubebot

## Introducing Target and Attack Surface

## Automation of multiple tools
* Show local data conversion for both wfuzz and reposupervisor and storing in bq
* Show local bruteforcer
* Show k8s deployment of both wfuzz and reposupervisor pods simultaneously manually using yaml files.
* Once they finish, show k8s deployment of bruteforcer using yaml file
* If more time, introduce Kubebot and show how it can all be automated - using slack to start the workflow, using channels/queues to communicate b/w the tools and bruteforcer & sending results back to slack

## Kubebot experience
### Sending a request from Kubebot for a target company from mobile
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