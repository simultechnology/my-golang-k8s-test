
``` 
helm repo add metallb https://metallb.github.io/metallb
kubectl create namespace metallb-system
helm install metallb metallb/metallb -n metallb-system

kubectl apply -f ipaddresspool.yaml
```