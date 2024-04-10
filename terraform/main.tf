terraform {
  backend "local" {}
}

provider "aws" {
  region = "ap-northeast-1"
  s3_use_path_style = true
}

resource "aws_s3_bucket" "tf-test-bucket" {
  bucket = "tf-test-bucket"
}
