data "aws_caller_identity" "current" {}
data "aws_availability_zones" "available" {
  state = "available"
}

locals {
  name_prefix    = "${var.project}-${var.env}"
  azs            = slice(data.aws_availability_zones.available.names, 0, 2)
  log_group_name = "/ecs/${local.name_prefix}-api"
}
