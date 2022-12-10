provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

resource "aws_s3_bucket" "terraform_bucket" {
  bucket = "tfstate-backend-development"
  tags = {
    Name        = "${var.app_name}-s3"
    Environment = var.app_environment
  }
}


resource "aws_s3_bucket_acl" "terraform_state" {
  bucket = aws_s3_bucket.terraform_bucket.id
  acl    = "private"
}


resource "aws_s3_bucket_server_side_encryption_configuration" "terraform_configuration" {
  bucket = "tfstate-backend-development"

  rule {
    apply_server_side_encryption_by_default {
      kms_master_key_id = aws_kms_key.terraform_bucket_key.arn
      sse_algorithm     = "aws:kms"
    }
  }

}

resource "aws_s3_bucket_public_access_block" "terraform_block" {
  bucket = "tfstate-backend-development"

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_dynamodb_table" "terraform_state" {
  name           = "terraform_state"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}

output "load_balancer_ip" {
  value = aws_alb.queuefly_alb.dns_name
}
