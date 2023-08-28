

## build docker image
``` 
docker build -t golang-hello:0.1 .

docker run -it -p 8888:8888 golang-hello:0.1
```

## load image to kind cluster
``` 
kind load docker-image golang-hello:0.1 --name ipv6-cluster 
```

## deploy
``` 
kustomize build k8s | kubectl apply -f -
```

