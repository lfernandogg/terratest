terraform {
  cloud {
    organization = "Personal2"

    workspaces {
      name = "dealer"
    }
  }
}

provider "aws"{ 
  region = "us-east-2"
  access_key = var.AWS_ACCESS_KEY_ID
  secret_key = var.AWS_SECRET_ACCESS_KEY
}

  


variable "AWS_ACCESS_KEY_ID" {}
variable "AWS_SECRET_ACCESS_KEY" {}
variable bucket { 
    type = string
    default = "testbucket"
}
variable content{ 
    type = string
    default = "Esto es una simple liea de texto"
}





/*------------bucket------------*/

resource "aws_s3_bucket" "s3_bucket" {
  bucket = "projectjav-test-bucket"

  tags = {
    Name        = "projectjav-test-bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.s3_bucket.id
  acl    = "private"
}

/*------------content------------*/

resource "aws_s3_bucket_object" "object1" {
  bucket = aws_s3_bucket.s3_bucket.bucket
  key    = "index.html"
  acl = "public-read"
  content = var.content
  content_type = "text/html"
}


