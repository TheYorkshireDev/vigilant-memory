variable "product" {
  type = string
}

variable "environment" {
  type = string
}

module "repo_name" {
  source = "../../"

  product     = var.product
  environment = var.environment
}

output "repository_name" {
  value = module.repo_name.repository_name
}
