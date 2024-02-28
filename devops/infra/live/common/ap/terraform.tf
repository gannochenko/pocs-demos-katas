terraform {
  backend "gcs" {
    bucket = "8108535439b56ce6-bucket-tfstate"
    prefix = "live/common/ap"
  }
}
