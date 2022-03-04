locals {
  list_of_objects  = ["a", "b", "c"]
  object_not_found = index(local.list_of_objects, "d")
}