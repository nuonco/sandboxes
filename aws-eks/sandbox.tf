module "sandbox" {
  source  = "nuonco/eks-sandbox/aws"
  version = "1.3.0"

  install_name          = var.install_name
  cluster_name          = var.cluster_name
  cluster_version       = var.cluster_version
  min_size              = var.min_size
  max_size              = var.max_size
  desired_size          = var.desired_size
  default_instance_type = var.default_instance_type
  additional_tags       = var.additional_tags

  nuon_id                           = var.nuon_id
  region                            = var.region
  waypoint_odr_namespace            = var.waypoint_odr_namespace
  waypoint_odr_service_account_name = var.waypoint_odr_service_account_name
  public_root_domain                = var.public_root_domain
  internal_root_domain              = var.internal_root_domain
  tags                              = var.tags
  enable_nginx_ingress_controller   = var.enable_nginx_ingress_controller
  runner_install_role               = var.runner_install_role

  # Need to explicitly pass in providers, because of no_tags provider.
  providers = {
    aws         = aws
    aws.no_tags = aws.no_tags
    helm        = helm
    kubectl     = kubectl
    kubernetes  = kubernetes
  }
}
