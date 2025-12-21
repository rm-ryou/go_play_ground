module "network" {
  source = "../../modules/network"

  project    = "simple-webapp"
  env        = "test"
  aws_region = "ap-northeast-1"
}

