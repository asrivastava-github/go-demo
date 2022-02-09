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

resource "demo_user_service" "user_asdd" {
  firstname = "Avinash"
  lastname = "Srivastava"
}