terraform {
  required_providers {
    idcs = {
      version = "0.1"
      source  = "hashicorp.com/dirgapeter/idcs"
    }
  }
}

# variable "coffee_name" {
#   type    = string
#   default = "Vagrante espresso"
# }

data "idcs_app_roles" "all" {
  filter {
    name   = "display_name"
    values = ["ServiceInvoker"]
    regex  = true
  }
}

# Returns all coffees
output "all_app_roles" {
  value = data.idcs_app_roles.all.app_roles
}

# Only returns packer spiced latte
# output "coffee" {
#   value = {
#     for coffee in data.hashicups_coffees.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }
