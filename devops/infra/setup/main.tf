terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "digitalocean_token" {}
variable "digitalocean_spaces_access_key_id" {}
variable "digitalocean_spaces_secret_access_key" {}

locals {
  project_name = "devops"
  backend_region = "fra1"
  backend_environment = "production"
}

provider "digitalocean" {
  token = var.digitalocean_token
  spaces_access_id = var.digitalocean_spaces_access_key_id
  spaces_secret_key = var.digitalocean_spaces_secret_access_key
}

resource "random_id" "bucket_prefix" {
  byte_length = 8
}

resource "digitalocean_spaces_bucket" "tfstate" {
  name   = "${random_id.bucket_prefix.hex}-${local.project_name}-tfstate"
  region = local.backend_region

  acl             = "private"
  force_destroy   = false
  versioning {
    enabled = false
  }
}
