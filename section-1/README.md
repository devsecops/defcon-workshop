# Section-1

## Table of Contents
<!-- TOC -->

- [Section-1](#section-1)
    - [Table of Contents](#table-of-contents)
    - [Overview](#overview)
    - [Installation / Setting up](#installation--setting-up)
        - [Requirements](#requirements)
        - [Google Cloud Platform (GCP) account](#google-cloud-platform-gcp-account)
        - [MacOS/Linux](#macoslinux)
        - [Windows](#windows)
    - [Building the first Docker image](#building-the-first-docker-image)
        - [Linux/Mac](#linuxmac)
        - [Windows](#windows-1)
    - [Pushing the Docker image to Google Container Registry (GCR)](#pushing-the-docker-image-to-google-container-registry-gcr)
        - [Linux/Mac](#linuxmac-1)
        - [Windows](#windows-2)
    - [References:](#references)

<!-- /TOC -->

## Overview
In this section, we will:
1. Install and set up our environment for the workshop.
2. Build our first Docker image.
3. Push our Docker image to Google Container Registry (GCR).

## Installation / Setting up

### Requirements
1. Google Cloud Platform (GCP) account
2. Docker
3. Minikube
4. kubectl 
5. VirtualBox
6. Virtualenv
7. GoLang
8. Google Cloud SDK

### Google Cloud Platform (GCP) account 
* You can use the GCP Free Tier to get one
    https://cloud.google.com/free/


### MacOS/Linux

* **Install Xcode from app store (this will install Git so you can clone the repo)**

* **Install Docker**

    https://www.docker.com/products/docker-toolbox

    https://docs.docker.com/docker-for-mac/install/#download-docker-for-mac

    * if you're using Kali, use the install script [kali-install-docker.sh](./kali-install-docker.sh). 
    * Unfortunately, this [script] only works for x64 Kali builds. 

* **Install Google Cloud SDK**

    https://cloud.google.com/sdk/docs/quickstart-mac-os-x (download tar, extract, and run ./install.sh)

    * Initialize Google SDK: Open a new terminal and run `gcloud init` (set project to defcon-workshop and default region to us-west1-a)
    
    Google Cloud SDK 163.0.0

* **Install Minikube:**

    https://github.com/kubernetes/minikube

    minikube version: v0.20.0

* **Install Kubectl:**

Install kubectl via the GCLOUD SDK - `gcloud components install kubectl`

Client Version: version.Info{Major:"1", Minor:"7", GitVersion:"v1.7.0", GitCommit:"d3ada0119e776222f11ec7945e6d860061339aad", GitTreeState:"clean", BuildDate:"2017-06-29T23:15:59Z", GoVersion:"go1.8.3", Compiler:"gc", Platform:"darwin/amd64"}
Server Version: version.Info{Major:"1", Minor:"6", GitVersion:"v1.6.4", GitCommit:"d6f433224538d4f9ca2f7ae19b252e6fcb66a3ae", GitTreeState:"dirty", BuildDate:"2017-06-22T04:31:09Z", GoVersion:"go1.7.5", Compiler:"gc", Platform:"linux/amd64"}

* **Install Virtualbox**

    https://www.virtualbox.org/wiki/Downloads


* **Install Virtualenv**

    https://virtualenv.pypa.io/en/stable/installation/


* **Install the GO programming language**

    https://golang.org/dl/ (download and install .pkg)

    go version go1.8.3 darwin/amd64

### Windows

* **Install Docker Toolbox (installs virtualbox, Git, etc.)**

    https://download.docker.com/win/stable/DockerToolbox.exe

* **Install Minikube**
    * Download this file [https:storage.googleapis.com/minikube/releases/latest/minikube-windows-amd64.exe](https://storage.googleapis.com/minikube/releases/latest/minikube-windows-amd64.exe) and rename it to minikube.exe
    * Place it in your Program Files directory under Minikube
    * Add the folder to your PATH by right clicking on the Windows icon in the lower left hand side, clicking System, then clicking on Advanced system settings in the left hand pane
    * Click on Environment Variables at the bottom of that screen and double clicking on ‘Path’ under System variables
    * Add the following to the end: `;C:\Program Files\Minikube`
    * Now when you open a command prompt and type minikube it should run using the file in that directory

* **Install Google SDK and KubeCtl**

    * Download and run the SDK from here: https://dl.google.com/dl/cloudsdk/channels/rapid/GoogleCloudSDKInstaller.exe
    * Uncheck all of the boxes and click finish
    * Restart your command prompt and then type: `gcloud components install kubectl`

* **Install Virtualenv**

    https://virtualenv.pypa.io/en/stable/installation/

* **Install the GO programming language**

    https://storage.googleapis.com/golang/go1.8.3.windows-amd64.msi

Note: If any of the tools above don’t seem like they’re working in a command prompt after installation, try closing your prompt and opening a new one.  The tool should then work.


## Building the first Docker image
* Open a shell/command prompt and clone this repository with the following command:
    * `git clone https://github.com/devsecops/defcon-workshop.git`
*  Change into the section-1 directory of the defcon-workshop repo

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
* Use `gcloud init` to make sure your gcloud configuration is for the right account and project if you run into this error: `denied: Unable to create the repository, please check that you have access to do so.`

### Linux/Mac
* `gcloud docker -- push us.gcr.io/$PROJECT_ID/test:v1`
### Windows
* `gcloud docker -- push us.gcr.io/%project_id%/test:v1`

## References:
- https://gist.github.com/nikallass/e5124756d0e2bdcf8981827f3ed40bcc
- https://gist.github.com/apolloclark/f0e3974601346883c731
