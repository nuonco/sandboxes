module "sandbox" {
  # source  = "nuonco/ecs-byovpc-sandbox/aws"
  # version = "1.0.0"
  source = "github.com/nuonco/terraform-aws-ecs-byovpc-sandbox"

  prefix_override = var.prefix_override
  vpc_id          = var.vpc_id

  nuon_id                                = var.nuon_id
  region                                 = var.region
  assume_role_arn                        = var.assume_role_arn
  public_root_domain                     = var.public_root_domain
  internal_root_domain                   = var.internal_root_domain
  nuon_runner_install_trust_iam_role_arn = var.nuon_runner_install_trust_iam_role_arn
  tags                                   = var.tags
}
