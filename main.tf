variable "product" {
  type = string
}

variable "environment" {
  type = string
}

output "repository_name" {
  value = "${var.product}-${var.environment}"
}
