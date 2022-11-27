CircleCI status badge: 
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/laihoangdung/devops-capston-k8s-project/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/laihoangdung/devops-capston-k8s-project/tree/master)

## Project Overview

In this project, I use the k8s to build to Golang APIs server application. This application uses the rolling update strategy. I use the ECR to save to docker image and EKS to deploy the application

### Files explanation
1. circleci/config.yml: config file to run circleci
3. main.go: server application
4. cloudformation/eks.yml: eks cloudformation configuration file
4. cloudformation/create-stack.yml: command to create the eks
4. cloudformation/update-stack.yml: command to update the eks
4. k8s/deployment.yml: k8s deployment file
4. k8s/service.yml: k8s service file
4. src/models: entity models use in the application
4. src/controllers: the api controllers
4. Dockerfile: define container to run the application
6. Makefile: define common commands help to install libraries, test, lint the application
7. go.mod: golang libraries to use in the web application
8. run_docker.sh: use to test the container image
9. run_kubernetes.sh: deploy application to kubenetes cluster
10. upload_docker.sh: test to push docker image to docker hub