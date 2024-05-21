# My Golang K8s Test Project

## Overview
This project is a sample application using Golang and Kubernetes.

## Setup

### Prerequisites
- Go
- Kubernetes
- Helm
- istioctl
- kustomize

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
3. set up istio
   
   ```
   istioctl install --set profile=demo -y
   ```

4. deploy common Kubernetes configurations and manifests

   ```
   kustomize build general-k8s/ | kubectl apply -f - 
   ```

5. Build and deploy the application:
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


## Verification Steps

1. **Check the Service**:
   Verify that the service is correctly created after deployment.

   ```sh
   kubectl get svc -n api
   kubectl get svc -n hello
   ```

2. **Check the EXTERNAL-IP of the Service**:
   Verify the EXTERNAL-IP assigned to the created service. For example, for api-service:

   ```sh
   kubectl get svc api-service -n api
   kubectl get svc api-service -n hello
   ```
   
3. **Check the Service**:
   log in to the network tools container

   ```sh
   kubectl exec -it multitool-****** -- bash
   ```

4. **Check access**:
 
 	```sh
 	curl -g [fc00:f853:ccd:e793::c]/get-pods-info
	curl -g [fc00:f853:ccd:e793::d]/hello
	
	curl -g myapp.example.com/hello
	```
	

## Technologies Used

- Golang
- Kubernetes
- Helm
- MetalLB

