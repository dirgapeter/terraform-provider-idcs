terraform {
  required_providers {
    idcs = {
      version = "0.1"
      source  = "hashicorp.com/dirgapeter/idcs"
    }
  }
}

provider "idcs" {
  idcs_url      = "https://idcs-725b1697f54d4336b245078b0d3f1d8b.identity.oraclecloud.com"
  client_id     = "bc110484f81141f9b748f5c3a2e682ac"
  client_secret = "f53cdf85-4377-4639-a8d2-15c90fdcc2b9"
}

module "groups" {
  source = "./groups"

  # coffee_name = "Packer Spiced Latte"
}

output "groups" {
  value = module.groups.all_groups
}


module "apps" {
  source = "./apps"

  # coffee_name = "Packer Spiced Latte"
}

output "apps" {
  value = module.apps.all_apps
}


module "app_roles" {
  source = "./app_roles"

  # coffee_name = "Packer Spiced Latte"
}

output "app_roles" {
  value = module.app_roles.all_app_roles
}

module "grants" {
  source = "./grants"

  # coffee_name = "Packer Spiced Latte"
}

output "grants" {
  value = module.grants.all_grants
}

resource "idcs_grant" "grant" {
  app {
    value = "0983f72cbdd1404883515b9ef0d05d85"
  }
  grantee {
    type  = "Group"
    value = "cbce7a3c94174127a5525e9162969337"
  }
  entitlement {
    attribute_name  = "appRoles"
    attribute_value = "013b4c4a17a741c892cd8248bbe04242"
  }
  grant_mechanism = "ADMINISTRATOR_TO_GROUP"
}
