terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"

  backend "s3" {
    bucket         = "tfstate-backend-development"
    key            = "tfstate/terraform_state.tfstate"
    region         = "us-east-1"
    kms_key_id     = "alias/terraform_bucket_key"
    encrypt        = true
    dynamodb_table = "terraform_state"
  }
}

