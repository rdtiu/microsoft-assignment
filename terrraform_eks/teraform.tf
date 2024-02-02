terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 5.34.0"
    }

    random = {
      source = "hashicorp/random"
      version = "~> 3.6.0"
    }

    tls = {
      source = "hashicorp/tls"
      version = "~> 4.0.5"
    }

    cloudinit = {
      source = "hashicorp/cloudinit"
      version = "~> 2.3.3"
    }

    kubernetes = {
      source = "hashicorp/kubernetes"
      version = "~> 2.25.2"
    }
  }

  required_version = "~> 1.7"
}
