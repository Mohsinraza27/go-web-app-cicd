# DevOps Implementation of go web application

## Overview
This project demonstrates the full lifecycle of a DevOps implementation, from containerizing an application using Docker to deploying it on Kubernetes with a CI/CD pipeline using GitHub Actions. The project is designed to automate and streamline development, testing, and deployment processes, ensuring a reliable and scalable environment for software delivery.

The main goal of this project is to implement DevOps practices in the Go web application. The project is a simple website written in Golang. It uses the `net/http` package to serve HTTP requests.


## Table of Contents
1. [Project Description](#project-description)
2. [Technologies Used](#technologies-used)
3. [Setup and Installation](#setup-and-installation)
4. [CI/CD Pipeline](#cicd-pipeline)
5. [Kubernetes Deployment](#kubernetes-deployment)
6. [Challenges and Solutions](#challenges-and-solutions)
7. [Special Thanks](#special-thanks)

## Project Description
The goal of this project was to implement a robust DevOps pipeline that includes the following key components:

- *Containerization:* Application containerized using Docker for consistent environment across development, testing, and production.
- *Orchestration:* Managed containerized applications using Kubernetes, enabling automated deployment, scaling, and operations.
- *CI/CD Pipeline:* Automated building, testing, and deployment using GitHub Actions.
- *Environment Management:* Managed different environments (development, staging, production) using Helm charts.

## Technologies Used
- *Docker:* For containerizing the application.
- *Kubernetes:* For orchestrating and managing containers.
- *Helm:* For managing Kubernetes applications.
- *GitHub Actions:* For CI/CD pipeline automation.
- *Go (Golang):* The primary programming language used in the application.
- *Nginx:* Used as a reverse proxy and load balancer.

## Setup and Installation
### Prerequisites
- *Docker:* Ensure Docker is installed on your machine.
- *Kubernetes:* Set up a Kubernetes cluster (can be minikube for local development).
- *Helm:* Install Helm for Kubernetes management.
- *GitHub Account:* For repository management and setting up GitHub Actions.

## Creating Dockerfile (Multi-stage build)
The Dockerfile is used to build a Docker image. The Docker image contains the Go web application and its dependencies. The Docker image is then used to create a Docker container.

We will use a Multi-stage build to create the Docker image. The Multi-stage build is a feature of Docker that allows you to use multiple build stages in a single Dockerfile. This will reduce the final Docker image's size and secure the image by removing unnecessary files and packages.

## Containerization
Containerization is the process of packaging an application and its dependencies into a container. The container is then run on a container platform such as Docker. Containerization allows you to run the application in a consistent environment, regardless of the underlying infrastructure.

We will use Docker to containerize the Go web application. Docker is a container platform that allows you to build, ship, and run containers.

### Installation Steps
1. *Clone the Repository:*
    bash
    git clone https://github.com/Mohsinraza27/go-web-app-cicd.git
    cd go-web-app-cicd
    

2. *Build Docker Images:*
    bash
    docker build -t <your-dockerhub-username>/go-web-app .
    

3. *Run Docker Containers:*
    bash
    docker run -d -p 8080:8080 <your-dockerhub-username>/go-web-app
    

4. *Setup Kubernetes Resources:*
    - Deploy the application using Helm:
    bash
    helm install go-web-app ./helm-chart/
    

5. *Access the Application:*
    - If using minikube, access the application using:
    bash
    minikube service go-web-app-service
    

## CI/CD Pipeline

# Continuous Integration (CI)
Continuous Integration (CI) is the practice of automating the integration of code changes into a shared repository. CI helps to catch bugs early in the development process and ensures that the code is always in a deployable state.

We will use GitHub Actions to implement CI for the Go web application. GitHub Actions is a feature of GitHub that allows you to automate workflows, such as building, testing, and deploying code.

The CI/CD pipeline is fully automated using GitHub Actions. Every push to the main branch triggers the following steps:

1. *Build:* Docker images are built and tagged.
2. *Test:* Automated tests are run to ensure code quality.
3. *Deploy:* Successful builds are deployed to the Kubernetes cluster.

# Continuous Deployment (CD)
Continuous Deployment (CD) is the practice of automatically deploying code changes to a production environment. CD helps to reduce the time between code changes and deployment, allowing you to deliver new features and fixes to users faster.

We will use Argo CD to implement CD for the Go web application. Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes. It allows you to deploy applications to Kubernetes clusters using Git as the source of truth.

The Argo CD application will deploy the Go web application to a Kubernetes cluster. The application will be automatically synced with the Git repository, ensuring that the application is always up to date.

### Pipeline Configuration
The pipeline is configured in the .github/workflows/ci-cd.yml file. Key stages include:

- *Build Stage:* Uses docker build to create images.
- *Test Stage:* Runs unit and integration tests.
- *Deploy Stage:* Deploys the application to the specified Kubernetes environment using kubectl and Helm.

## Kubernetes Deployment
### Deployment Files
The deployment configuration is defined in the ./helm-chart/ directory, which includes:

- *values.yaml:* Configuration file for Helm charts.
- *deployment.yaml:* Kubernetes deployment file for defining pod replicas, containers, and environment variables.
- *service.yaml:* Defines the service type, ports, and load balancer settings.

### Deploying to Production
To deploy the application to production, update the values.yaml file with production configurations and run:

```bash
helm install go-web-app ./helm-chart/ 
