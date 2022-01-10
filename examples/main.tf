terraform {
  required_providers {
    avengers = {
      version = "1.0.0"
      source  = "github.com/custom-terraform-provider/avengers"
    }
  }
}

provider "avengers" {}

module "psl" {
  source = "./avengers"

  name = "zido"
}

output "psl" {
  value = module.psl.avenger
}
