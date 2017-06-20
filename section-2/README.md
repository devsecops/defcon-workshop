# Section-2

## Deploying a K8S cluster locally on minikube

* `minikube start`
* `eval $(minikube docker-env)`
* `docker ps -a` - Verify you are inside minikube's docker environment
* `kubectl apply -f local-deployment.yaml` - Deploys the local K8S cluster on Minikube
* `minikube dashboard` - Dashboard to view the deployment
* `kubectl get deployments --namespace=local-server` - Retrieve all the deployments in the namespace
* `kubectl get pods --namespace=local-server` - Retrieve all the pods in the namespace
* `kubectl scale deployment nginx-deployment --namespace=local-server --replicas 10` - Scales the deployment from 3 to 10
* `kubectl autoscale deployment nginx-deployment --namespace=local-server --min=10 --max=15 --cpu-percent=80` - Autoscale

References:
* https://kubernetes.io/docs/concepts/workloads/controllers/deployment/

## Deploying a K8S cluster remotely on GCP

*
* `kubectl apply -f remote-deployment.yaml`