terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }

  # https://dev.to/jmarhee/digitalocean-spaces-as-a-terraform-backend-3lck
  backend "s3" {
    endpoint                    = "fra1.digitaloceanspaces.com/"
    key                         = "tfstate"
    bucket                      = "5797f00ba72993d1-devops-tfstate"
    region                      = "us-west-1"
    skip_credentials_validation = true
    skip_metadata_api_check     = true
  }
}

variable "digitalocean_token" {}
variable "digitalocean_spaces_access_key_id" {}
variable "digitalocean_spaces_secret_access_key" {}

provider "digitalocean" {
  token = var.digitalocean_token
  spaces_access_id = var.digitalocean_spaces_access_key_id
  spaces_secret_key = var.digitalocean_spaces_secret_access_key
}
