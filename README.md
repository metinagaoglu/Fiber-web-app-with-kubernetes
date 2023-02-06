Golang Fiber Web Application with Docker and Kubernetes Deployment

This project is a demonstration of a simple web application built with the Fiber library in Golang, Dockerized, and deployed with Kubernetes.
Prerequisites

    Docker
    Kubernetes (minikube)
    Fiber

Running the Application

Clone the repository:

    $ git clone https://github.com/metinagaoglu/basic-go-app-with-kubernetes.git

Build the Docker image:

    $ cd basic-go-app-with-kubernetes
    $ docker build -t <image-name>:<tag> .

Deploy the application to Kubernetes:

    $ kubectl apply -f manifests/

Verify the deployment:

    $ kubectl get pods

Open your web browser and navigate to your cluster ip address. You should see the web page served by your Golang Fiber application.

### Conclusion

This project shows how to build a simple web application with Golang Fiber, package it in a Docker image, and deploy it to a Kubernetes cluster. It can be used as a starting point for developing more complex web applications.
