#!/bin/bash

kubectl apply -f bitcoin-server-namespace.yaml
kubectl apply -f bitcoin-server-razvantiu-imagepull-secret.yaml
kubectl apply -f bitcoin-server-deployment.yaml
kubectl apply -f bitcoin-server-service.yaml
kubectl apply -f bitcoin-server-ingress.yaml
kubectl apply -f bitcoin-server-np.yaml
