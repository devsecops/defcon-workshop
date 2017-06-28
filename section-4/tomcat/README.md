# Building the vulnerable Tomcat Manager

## Building and Running the image
In a new terminal i.e. outside minikube's docker environment, type
* docker build -t test .
* docker run --rm -it -p 8080:8080 test

Start NGROK by typing `./ngrok http 8080`. You can now use the NGROK's URL as the target URL