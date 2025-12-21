terraform {
  backend "s3" {
    bucket  = "tfstate-simple-webapp-test-ap-northeast-1"
    key     = "simple-webapp/test/state"
    region  = "ap-northeast-1"
    encrypt = true
  }
}
