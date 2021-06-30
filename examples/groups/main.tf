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

data "idcs_groups" "all" {}

# Returns all coffees
output "all_groups" {
  value = data.idcs_groups.all.groups
}

# Only returns packer spiced latte
# output "coffee" {
#   value = {
#     for coffee in data.hashicups_coffees.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }
