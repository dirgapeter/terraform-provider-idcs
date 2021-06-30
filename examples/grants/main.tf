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

data "idcs_grants" "all" {
  # filter {
  #   name   = "display_name"
  #   values = ["ServiceInvoker"]
  #   regex  = true
  # }
}

# Returns all coffees
output "all_grants" {
  value = data.idcs_grants.all.grants
}

# Only returns packer spiced latte
# output "coffee" {
#   value = {
#     for coffee in data.hashicups_coffees.all.coffees :
#     coffee.id => coffee
#     if coffee.name == var.coffee_name
#   }
# }
