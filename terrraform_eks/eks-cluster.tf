module "eks" {
  source = "terraform-aws-modules/eks/aws"
  version = "19.21.0"

  cluster_name = local.cluster_name
  cluster_version = "1.29"

  vpc_id = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnets
  cluster_endpoint_public_access = true

  eks_managed_node_group_defaults = {
    ami_type = "AL2_x86_64"

  }

  eks_managed_node_groups = {
    g1 = {
      name = "node-group-1"

      instance_types = ["t3.small"]

      min_size = 1
      max_size = 3
      desired_size = 2

      additional_security_group_ids = [aws_security_group.g1-sg.id]
    }

    g2 = {
      name = "node-group-2"

      instance_types = ["t3.small"]

      min_size = 1
      max_size = 2
      desired_size = 1

      additional_security_group_ids = [aws_security_group.g2-sg.id]
    }
  }
}
