

## build docker image
``` 
docker build -t golang-api:0.1 .

docker run -it -p 8080:8080 golang-api:0.1
```

## load image to kind cluster
``` 
kind load docker-image golang-api:0.1 --name ipv6-cluster 
```

## deploy
``` 
kustomize build k8s | kubectl apply -f -
```

