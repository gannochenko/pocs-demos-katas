//resource "google_project_iam_member" "project" {
//  project = local.project_id
//  role    = "roles/owner"
//  member  = "serviceAccount:test-564@go-app-390716.iam.gserviceaccount.com"
//}

resource "random_id" "bucket_prefix" {
  byte_length = 8
}

resource "google_storage_bucket" "default" {
  project       = local.project_id
  name          = "${random_id.bucket_prefix.hex}-bucket-tfstate"
  force_destroy = false
  location      = "us-east1"
  storage_class = "STANDARD"
  versioning {
    enabled = false
  }
//  depends_on = [
//    google_project_iam_member.project
//  ]
}
