#!/bin/bash
# source: https://gist.github.com/nikallass/e5124756d0e2bdcf8981827f3ed40bcc

# update apt-get
export DEBIAN_FRONTEND="noninteractive"
sudo apt-get update

# remove previously installed Docker
sudo apt-get remove docker docker-engine docker.io* lxc-docker*

# install dependencies 4 cert
sudo apt-get install apt-transport-https ca-certificates curl gnupg2 software-properties-common

# add Docker repo gpg key
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -

echo "deb https://download.docker.com/linux/debian stretch stable" >> /etc/apt/sources.list 

sudo apt-get update

# install Docker
sudo apt-get install docker-ce

# run Hello World image
sudo docker run hello-world

# manage Docker as a non-root user
sudo groupadd docker
sudo usermod -aG docker $USER

# configure Docker to start on boot
sudo systemctl enable docker