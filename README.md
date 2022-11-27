CircleCI status badge: 
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/laihoangdung/devops-capston-k8s-project/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/laihoangdung/devops-capston-k8s-project/tree/master)

## Project Overview

In this project, I use the k8s to build to Golang APIs server application. This application uses the rolling update strategy. I use the ECR to save to docker image and EKS to deploy the application

### Files explanation
1. circleci/config.yml: config file to run circleci
2. main.go: server application
3. cloudformation/eks.yml: eks cloudformation configuration file
4. cloudformation/create-stack.yml: command to create the eks
5. cloudformation/update-stack.yml: command to update the eks
6. k8s/deployment.yml: k8s deployment file
7. k8s/service.yml: k8s service file
8. src/models: entity models use in the application
9. src/controllers: the api controllers
10. Dockerfile: define container to run the application
11. Makefile: define common commands help to install libraries, test, lint the application
12. go.mod: golang libraries to use in the web application
13. run_docker.sh: use to test the container image
15. upload_docker.sh: test to push docker image to docker hub


# GitHub repository
https://github.com/laihoangdung/devops-capston-k8s-project

# APIs
1. Check health: http://ab4bbb6dc72f845b2a4d98d74bfc5ae2-1054660999.us-east-1.elb.amazonaws.com:8080/check-health
2. Get list of users: http://ab4bbb6dc72f845b2a4d98d74bfc5ae2-1054660999.us-east-1.elb.amazonaws.com:8080/users
