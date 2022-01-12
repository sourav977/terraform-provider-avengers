terraform {
  required_providers {
    avengers = {
      version = "1.0.1"
      source  = "github.com/sourav977/avengers"
    }
  }
}

provider "avengers" {
  host = "http://localhost:8000"
}

resource "avengers_resource" "foo" {
  name = "sourav patnaik"
  alias = "sikan"
  weapon = "hammer"
}

output "created_avenger" {
  value = avengers_resource.foo
}


data "avengers_datasource" "all" {}

# Returns all avengers
output "all_avengers" {
  value = data.avengers_datasource.all.avengers
}
