# Microsoft Assignment EKS cluster

Pre-requisites needed:
- aws-cli can be downloaded from: https://formulae.brew.sh/formula/awscli
or following the steps from the official documentation:
https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

- terraform
https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli

- kubelet
https://kubernetes.io/docs/tasks/tools/#kubectl


Before starting connect into your AWS console and generate a new access key for your user:
https://docs.aws.amazon.com/cli/latest/userguide/cli-services-iam-create-creds.html

Download the csv file created add User Name into header like in the following example:
```
User name,Access key ID,Secret access key
power_usr,XXXXXXXX,XXXXXXXX
```

and import it using the following command
```bash
aws configure --profile power_usr import --csv file://<path_to_credentials_downloaded.csv>
```

# Cluster Launch
In an EKS cluster the RBAC is enabled by default and users are authenticated through IAM.

## Steps to launch the EKS cluster

- Go to terraform_eks folder and perform the following commands
```bash
terraform init
terraform plan
terraform apply
```

After applying the recipe defined in terraform files a new EKS cluster will be created in your account
and the following objects:

- VPC
- NAT Gateway
- Security Groups
- Route Table
- Public Subnets
- Private Subnets

To be able to connect to the new cluster created a kubeconfig should be generated via:
```bash
aws eks update-kubeconfig --name <cluster-name> --region <region>
```

# Ingress Controller Deployment

Before deploying the 2 services we need to deploy the ingress controller via

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.9.6/deploy/static/provider/cloud/deploy.yaml
```

Deployment can be verified via:
```bash
kubectl -n ingress-nginx get all
```

and the public address is the EXTERNAL-IP of service/ingress-nginx-controller and it will be used for making https calls: 


```bash
kubectl -n ingress-nginx get service

NAME                                         TYPE           CLUSTER-IP       EXTERNAL-IP                                                               PORT(S)                      AGE
service/ingress-nginx-controller             LoadBalancer   172.20.4.14      aa5872b9dd43f4ca7ab0870449c82290-1006774448.us-east-2.elb.amazonaws.com   80:32669/TCP,443:30549/TCP        67s
service/ingress-nginx-controller-admission   ClusterIP      172.20.156.175   <none>                                                                    443/TCP                      67s
```

# Service_A Deployment
The source code of the application is available in /src folder from service_A folder.
The image was built using the Dockerfile and pushed into razvantiu repository from DockerHub
In order to perform the deployment of the service_A the script deploy.sh from deploy folder
should be applied which contains the following commands:

Create a new namespace for service_A
```bash
kubectl apply -f bitcoin-server-namespace.yaml
```

Deploy the secret used to connect to razvantiu repository into DockerHub
```bash
kubectl apply -f bitcoin-server-razvantiu-imagepull-secret.yaml
```

Deploy into the new namespace the deployment associated with service_A
```bash
kubectl apply -f bitcoin-server-deployment.yaml
```

Create a new service type ClusterIP for service_A
```bash
kubectl apply -f bitcoin-server-service.yaml
```

Create a new nginx ingress in order to be accessible via public Cloud
```bash
kubectl apply -f bitcoin-server-razvantiu-ingress
```

# Service_B Deployment
The code of the API is available in the /src folder and it contains a simple Flask-API
with the following routes:
- index route /
- liveness route
- readiness route

The image was built using the Dockerfile associated and pushed into razvantiu repository
from DockerHub and in order to deploy the service the same deploy.sh script should be
executed.

## Test connection to service_A
```bash
curl -i -H "Host: bitcoin-server.com" <service/ingress-nginx-controller EXTERNAL_IP>
curl -i -H "Host: bitcoin-server.com" <service/ingress-nginx-controller EXTERNAL_IP>/price
curl -i -H "Host: bitcoin-server.com" <service/ingress-nginx-controller EXTERNAL_IP>/average
```

## Test connection to service_B
```bash
curl -i -H "Host: api.com" <service/ingress-nginx-controller EXTERNAL_IP>
curl -i -H "Host: api.com" <service/ingress-nginx-controller EXTERNAL_IP>/livez
curl -i -H "Host: api.com" <service/ingress-nginx-controller EXTERNAL_IP>/readyz
```

# Cleanup EKS cluster

In order to delete the infrastructure launched at the end of the process you should run

```bash
terraform destroy
```
