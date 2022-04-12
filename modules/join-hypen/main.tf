variable "input_one" {
  type = string
}

variable "input_two" {
  type = string
}

output "joined_output" {
  value = "${var.input_one}-${var.input_two}"
}
