provider "aws" {
  region = "ap-northeast-1"

  default_tags {
    tags = {
      Project         = "simple-webapp"
      Env             = "test"
      HashiCorp-Learn = "aws-default-tags"
    }
  }
}
