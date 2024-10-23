terraform {
  backend "s3" {
    endpoint                    = "sfo2.digitaloceanspaces.com"
    key                         = "terraform.tfstate"
    bucket                      = "5797f00ba72993d1-devops-tfstate"
    region                      = "fra1"
    skip_requesting_account_id  = true
    skip_credentials_validation = true
    skip_get_ec2_platforms      = true
    skip_metadata_api_check     = true
  }
}
