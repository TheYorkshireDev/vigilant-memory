variable "example_input_one" {
  type = string
}

variable "example_input_two" {
  type = string
}

module "join_up" {
  source = "../../../modules/join-hypen"

  input_one = var.example_input_one
  input_two = var.example_input_two
}

output "joined_name" {
  value = module.join_up.joined_output
}
