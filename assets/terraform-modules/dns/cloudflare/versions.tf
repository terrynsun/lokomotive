terraform {
  required_version = ">= 0.13"

  required_providers {
    cloudflare = {
      source  = "terraform-providers/cloudflare"
      version = "2.9.0"
    }
  }
}
