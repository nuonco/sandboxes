module "sandbox" {
  # NOTE(fd): example format for testing branches
  # source  = "nuonco/eks-sandbox/aws"
  # version = "1.4.7"
  #
  source = "github.com/nuonco/terraform-aws-eks-sandbox?ref=721f84f304038c02e28a89ed2afbea53ac7f7f02"

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
}
