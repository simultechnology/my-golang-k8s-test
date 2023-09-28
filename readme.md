# My Golang K8s Test Project

## Overview
This project is a sample application using Golang and Kubernetes.

## Setup

### Prerequisites
- Go
- Kubernetes
- Helm

### Installation Steps

1. Install dependencies:

   ```
   go mod download
   ```

2. Set up MetalLB:

   ```
   helm repo add metallb https://metallb.github.io/metallb
   kubectl create namespace metallb-system
   helm install metallb metallb/metallb -n metallb-system
   ```

3. deploy common Kubernetes configurations and manifests

   ```
   kustomize build general-k8s/ | kubectl apply -f - 
   ```

3. Build and deploy the application:
   (Please add specific steps based on relevant Kubernetes manifests or scripts)
   
   - **api** Golang app [here](api/readme.md)
   - **hello** Golang app [here](hello/readme.md)


## Directory and File Descriptions

- **general-k8s**: Contains Kubernetes configurations and manifests.
- **api** and **api-updated**: Contains code and configurations related to the API.
- **hello**: Contains sample or test-related code.
- **main.go**: The main Go program.
- **go.mod** and **go.sum**: Files related to Go module dependencies.
- **Makefile**: A Makefile for automating build, test, and other tasks.

## Technologies Used

- Golang
- Kubernetes
- Helm
- MetalLB

