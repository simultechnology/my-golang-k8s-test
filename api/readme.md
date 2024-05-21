

## build docker image
``` 
docker build -t golang-api:0.2 .

docker run -it -p 8080:8080 golang-api:0.2
```

## load image to kind cluster
``` 
kind load docker-image golang-api:0.2 --name ipv6-cluster 
```

## deploy
``` 
kustomize build k8s | kubectl apply -f -
```

## Verification Steps

1. **Check the Service**:
   Verify that the service is correctly created after deployment.

   ```sh
   kubectl get svc -n api
   ```
2. **Check the EXTERNAL-IP of the Service**:
   Verify the EXTERNAL-IP assigned to the created service. For example, for api-service:

   ```sh
   kubectl get svc api-service -n api
   ```
3. **Access the EXTERNAL-IP**:
   Access the assigned EXTERNAL-IP on the specified port to ensure it is functioning. For example, for api-service:

   first, log in to a container(node)

   ```sh
   docker exec -it ipv6-cluster-control-plane bash
   ```

   ```sh 
   curl -g [<EXTERNAL-IP>]/get-pods-info
   ```
   
   example
   
   ```sh
   curl -g [fc00:f853:ccd:e793::c]/get-pods-info
   ```
4. **Check MetalLB Logs**:
   Verify MetalLB controller and speaker pod logs for any error messages.

   ``` 
   kubectl logs -n metallb-system -l component=controller
   kubectl logs -n metallb-system -l component=speaker
   ```

