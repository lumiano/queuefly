resource "aws_vpc" "aws_vpc_queuefly" {
  cidr_block           = "10.10.0.0/16"
  instance_tenancy     = "default"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "${var.app_environment}-vpc"
    Environment = var.app_environment
  }
}

data "aws_availability_zones" "available_zones" {
  state = "available"
}