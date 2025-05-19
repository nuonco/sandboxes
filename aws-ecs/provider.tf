provider "aws" {
  region = {{.nuon.install_stack.region}}

  default_tags {
    tags = local.tags
  }
}
