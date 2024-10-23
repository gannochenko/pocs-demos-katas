// Create project
resource "digitalocean_project" "devops" {
  name        = "devops"
  description = "This is a temporary project where I am gonna give k8s a try"
  purpose     = "Educational purposes"
  environment = local.environment
}

// Create registry
