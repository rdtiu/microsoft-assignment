#!/bin/bash

kubectl apply -f api-namespace.yaml
kubectl apply -f api-razvantiu-imagepull-secret.yaml
kubectl apply -f api-deployment.yaml
kubectl apply -f api-service.yaml
kubectl apply -f api-ingress.yaml
kubectl apply -f api-np.yaml
