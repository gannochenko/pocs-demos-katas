~~~
export TF_LOG=1
export DO_PAT="your_personal_access_token"
cd infra/setup
terraform init

~~~


~~~
resource "digitalocean_project" "devops" {
  name        = "devops"
  description = "This is a temporary project where I am gonna give k8s a try"
  purpose     = "Educational purposes"
  environment = local.backend_environment
}
~~~