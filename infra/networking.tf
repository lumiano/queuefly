resource "aws_internet_gateway" "gateway" {
  vpc_id = aws_vpc.aws_vpc_queuefly.id

  tags = {
    Name        = "${var.app_name}-igw"
    Environment = var.app_environment
  }
}


resource "aws_subnet" "private" {
  vpc_id            = aws_vpc.aws_vpc_queuefly.id
  count             = length(var.private_subnets)
  cidr_block        = cidrsubnet(aws_vpc.aws_vpc_queuefly.cidr_block, 8, count.index)
  availability_zone = data.aws_availability_zones.available_zones.names[count.index]

  tags = {
    Name        = "${var.app_name}-private-subnet-${count.index + 1}"
    Environment = var.app_environment
  }
}

resource "aws_subnet" "public" {
  vpc_id                  = aws_vpc.aws_vpc_queuefly.id
  count                   = length(var.public_subnets)
  map_public_ip_on_launch = true
  cidr_block              = cidrsubnet(aws_vpc.aws_vpc_queuefly.cidr_block, 8, 2 + count.index)
  availability_zone       = data.aws_availability_zones.available_zones.names[count.index]

  tags = {
    Name        = "${var.app_name}-public-subnet-${count.index + 1}"
    Environment = var.app_environment
  }
}


resource "aws_route_table" "private" {
  count  = var.app_count
  vpc_id = aws_vpc.aws_vpc_queuefly.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = element(aws_nat_gateway.gateway.*.id, count.index)
  }

  tags = {
    Name        = "${var.app_name}-routing-table-public"
    Environment = var.app_environment
  }
}

resource "aws_route" "internet_access" {
  route_table_id         = aws_vpc.aws_vpc_queuefly.main_route_table_id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.gateway.id

}

resource "aws_route_table_association" "private" {
  count          = length(var.public_subnets)
  subnet_id      = element(aws_subnet.private.*.id, count.index)
  route_table_id = element(aws_route_table.private.*.id, count.index)
}

resource "aws_eip" "gateway" {
  count      = var.app_count
  vpc        = true
  depends_on = [aws_internet_gateway.gateway]
}


resource "aws_nat_gateway" "gateway" {
  count         = var.app_count
  subnet_id     = element(aws_subnet.public.*.id, count.index)
  allocation_id = element(aws_eip.gateway.*.id, count.index)
}





