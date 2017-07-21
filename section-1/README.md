# Section-1

## Overview
In this section,
1. We will install all the software & set everything up for the rest of the workshop.
2. We will build our first Docker image.
3. We will push that Docker image to Google Container Registry (GCR).

## Installation / Setting up

A Google Cloud Platform (GCP) account – You can use the GCP Free Tier to get one
 https://cloud.google.com/free/


### MacOS/Linux

* Install Xcode from app store (this will install Git so you can clone the repo)

* Docker
https://www.docker.com/products/docker-toolbox
https://docs.docker.com/docker-for-mac/install/#download-docker-for-mac

* Minikube installed on the laptop
https://kubernetes.io/docs/tasks/tools/install-minikube/

`curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.19.1/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/
Kubectl
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl && chmod +x ./kubectl && sudo mv ./kubectl /usr/local/bin/kubectl`

* Virtualbox
https://www.virtualbox.org/wiki/Downloads

* The GO programming language installed on the laptop
https://golang.org/dl/ (download and install .pkg)

* Google Cloud SDK installed on the laptop
https://cloud.google.com/sdk/docs/quickstart-mac-os-x (download tar, extract, and run ./install.sh)

* Initialize Google SDK: Open a new terminal and run gcloud init (set project to defcon-workshop and default region to us-west1-a)


### Windows

* Install Docker Toolbox (installs virtualbox, Git, etc.)
https://download.docker.com/win/stable/DockerToolbox.exe

* Install Minikube
Download this file [minikube-windows-amd64.exe](https://storage.googleapis.com/minikube/releases/latest/minikube-windows-amd64.exe) and rename it to minikube.exe
Place it in your Program Files directory under Minikube
Add the folder to your PATH by right clicking on the Windows icon in the lower left hand side, clicking System, then clicking on Advanced system settings in the left hand pane
Click on Environment Variables at the bottom of that screen and double clicking on ‘Path’ under System variables
Add the following to the end: `;C:\Program Files\Minikube`
Now when you open a command prompt and type minikube it should run using the file in that directory

* Install Google SDK and KubeCtl
Download and run the SDK from here: https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe
Uncheck all of the boxes and click finish
Restart your command prompt and then type: `gcloud components install kubectl`

* Install the GO programming language
https://storage.googleapis.com/golang/go1.8.3.windows-amd64.msi

Note: If any of the tools above don’t seem like they’re working in a command prompt after installation, try closing your prompt and opening a new one.  The tool should then work.




## Building the first Docker image
* Open a shell/command prompt and clone this repository with the following command:
* `git clone https://github.com/devsecops/defcon-workshop.git`
*  Change into the section-1 directory of the defcon-workshop repo
* Identify your GCP project ID under Project info in your GCP project console

### Linux/Mac
* `export PROJECT_ID=<GCP-Project-ID>`
* `docker build -t us.gcr.io/$PROJECT_ID/test:v1 .`
* `docker run --rm us.gcr.io/$PROJECT_ID/test:v1`

### Windows
* `set project_id=<GCP-Project-ID>`
* `docker build -t us.gcr.io/%project_id%/test:v1 .`
* `docker run --rm us.gcr.io/%project_id%/test:v1`


## Pushing the Docker image to Google Container Registry (GCR)

* Enable Google Container Registry API in GCP’s API Manager.
* Use `gcloud init` to make sure your gcloud configuration is for the right account and project if you run into this error: `denied: Unable to create the repository, please check that you have access to do so.` You might also need: `docker tag us.gcr.io/$PROJECT_ID/test:v1 us.gcr.io/$PROJECT_ID/test:v1`

### Linux/Mac
* `gcloud docker -- push us.gcr.io/$PROJECT_ID/test:v1`
### Windows
* `gcloud docker -- push us.gcr.io/%project_id%/test:v1`
