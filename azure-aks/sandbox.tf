module "sandbox" {
  # source = "nuonco/aks-sandbox/azure"
  # version = "1.3.12"
  source = "github.com/nuonco/terraform-azure-aks-sandbox?ref=ja%2F6033-add-tag-policy"

  location             = var.location
  nuon_id              = var.nuon_id
  internal_root_domain = var.internal_root_domain
  public_root_domain   = var.public_root_domain
  cluster_version      = var.cluster_version
  vm_size              = var.vm_size
  node_count           = var.node_count
}
