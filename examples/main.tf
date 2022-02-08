terraform {
  required_providers {
    demo = {
      version = "0.1.0"
      source  = "example.com/learn/demo"
    }
  }
}

provider "demo" {
  host = "localhost"
  port = "5000"
}

data "demo_user_by_id" "user" {
    id = 1
}

output "name" {
  value = data.demo_user_by_id.user
}

# data "demo_users" "users" {
# }

# output "users" {
#   value = data.demo_users.users
# }

resource "demo_user_service" "user_asdd" {
  firstname = "Avinash"
  lastname = "Srivastava"
}