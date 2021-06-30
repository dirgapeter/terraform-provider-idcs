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

data "idcs_apps" "all" {
  filter {
    name   = "display_name"
    values = ["pdgoic1-frnzv4l8z4tx-fr"]
  }
}

# Returns all coffees
output "all_apps" {
  value = data.idcs_apps.all.apps
}

# Only returns packer spiced latte
# output "coffee" {
#   value = {
#     for coffee in data.hashicups_coffees.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }
