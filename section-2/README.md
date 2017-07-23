# Section-2

## Table of Contents
<!-- TOC -->

- [Section-2](#section-2)
    - [Table of Contents](#table-of-contents)
    - [Overview](#overview)
    - [GCP Web Services overview](#gcp-web-services-overview)
    - [Kubernetes aka K8s review](#kubernetes-aka-k8s-review)
        - [K8s Core](#k8s-core)
        - [K8s Overview](#k8s-overview)
    - [Deploying a K8S cluster locally on minikube](#deploying-a-k8s-cluster-locally-on-minikube)
    - [Deploying a K8S cluster remotely on GCP](#deploying-a-k8s-cluster-remotely-on-gcp)

<!-- /TOC -->


## Overview
In this section, we will
1. Briefly discuss the various Google Cloud Platform (GCP) web services.
2. Cover the basic concepts of [Kubernetes](https://kubernetes.io/).
3. Deploy our first Kubernetes (K8S) cluster locally using [Minikube](https://github.com/kubernetes/minikube).
4. Deploy our first Kubernetes (K8S) cluster remotely on GCP.


## GCP Web Services overview (only the ones we will be using)
* Google Container Engine (GKE) - Runs containers on a cluster. Entire backend is based on Kubernetes.
* Google BigQuery - Query large amounts of data real quick. Append only.
* Google Container Registry (GCR) - Registry/Repository to store Docker images used in a GCP account
* Google PubSub - Messaging system. Publisher/Subscription model.
    * Publish messages to a topic
    * Create a subscription to that topic
    * Listen for messages on that subscription
    * Acknowledge messages to remove from further consumption


## Kubernetes aka K8s review

### K8s Core
![k8s core](imgs/k8s.png)

Reference: [link](https://blog.heptio.com/core-kubernetes-jazz-improv-over-orchestration-a7903ea92ca)

### K8s Overview
![k8s overview](imgs/k8s4.png)

Reference: [link](https://www.redhat.com/en/containers/what-is-kubernetes)

## Deploying a K8S cluster locally on minikube

1. `minikube start`
2. `eval $(minikube docker-env)`
3. `docker ps -a` - Verify you are inside minikube's docker environment
4. `kubectl apply -f local-deployment.yaml` - Deploys the local K8S cluster on Minikube
5. `minikube dashboard` - Dashboard to view the deployment
6. `kubectl get deployments --namespace=local-server` - Retrieve all the deployments in the namespace
7. `kubectl get pods --namespace=local-server` - Retrieve all the pods in the namespace
8. `kubectl scale deployment nginx-deployment --namespace=local-server --replicas 10` - Scales the deployment from 3 to 10
9. `kubectl autoscale deployment nginx-deployment --namespace=local-server --min=10 --max=15 --cpu-percent=80` - Autoscale
10. `kubectl delete deployments --namespace=local-server --all` - Deletes the local deployments in the namespace
11. `kubectl delete namespace local-server` - Deletes the namespace


## Deploying a K8S cluster remotely on GCP

1. `gcloud components install alpha` - Install google alpha components
2. `gcloud alpha container clusters create remote-cluster --enable-kubernetes-alpha --scopes bigquery,storage-rw,compute-ro,https://www.googleapis.com/auth/pubsub` - Creates an alpha K8S cluster with scopes
3. `gcloud container clusters get-credentials remote-cluster --zone us-west1-a --project $PROJECT_ID` - Connecting to the remote K8S cluster and generating an entry in the `~/.kube/config` file for it
4. `kubectl get nodes` - Verify you are talking to the remote K8S cluster
5. `kubectl proxy` - Starts a proxy locally to view the remote K8S dashboard. You can then view the Minikube dashboard by navigating to the URL in the browser. This is the same as typing `minikube dashboard` in the above usecase
6. `kubectl apply -f remote-deployment.yaml` - Deploys the remote K8S cluster on GCP. Similar commands as above apply here as well
7. `kubectl delete deployments --namespace=remote-server --all` - Deletes the remote deployments in the namespace
8. `kubectl delete namespace remote-server` - Deletes the namespace
9. You can kill the `kubectl proxy` as well now.

References:
* https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
