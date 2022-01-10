terraform {
  required_providers {
    avengers = {
      version = "1.0.0"
      source  = "sourav.com/custom-terraform-provider/avengers"
    }
  }
}

provider "avengers" {}

resource "my_avengers" "first" {
    name= "Sourav"
    alias= "sikan"
    weapon= "gun"
}

output "first_avenger" {
  value = my_avengers.first
}