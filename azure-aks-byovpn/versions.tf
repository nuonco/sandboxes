terraform {
  backend "s3" {}

  required_providers {
    azapi = {
      source  = "Azure/azapi"
      version = "~> 1.12.1"
    }

    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.116.0"
    }
  }
}
