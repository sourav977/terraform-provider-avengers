terraform {
  required_providers {
    avengers = {
      version = "1.0.0"
      source  = "github.com/custom-terraform-provider/avengers"
    }
  }
}

variable "name" {
  type    = string
  default = "zido"
}

data "my_avengers" "all" {}

# Returns all Avengers
output "all_avengers" {
  value = data.my_avengers.all.avengers
}

# Only returns packer spiced latte
output "avenger" {
  value = {
    for avenger in data.my_avengers.all.avengers :
    avenger.id => avg
    if avg.name == var.name
  }
}
