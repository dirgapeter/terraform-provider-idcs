terraform {
  required_providers {
    idcs = {
      version = "0.1"
      source  = "hashicorp.com/dirgapeter/idcs"
    }
  }
}

provider "idcs" {}

# module "psl" {
#   source = "./groups"

#   coffee_name = "Packer Spiced Latte"
# }

# output "psl" {
#   value = module.psl.coffee
# }
