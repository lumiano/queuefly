resource "aws_alb" "queuefly_alb" {
  name            = "${var.app_name}-${var.app_environment}-alb"
  subnets         = aws_subnet.public.*.id
  security_groups = [aws_security_group.queuefly_security_group.id]

  tags = {
    Name        = "${var.app_name}-alb"
    Environment = var.app_environment
  }
}

resource "aws_security_group" "queuefly_security_group" {
  vpc_id = aws_vpc.aws_vpc_queuefly.id

  ingress {
    protocol    = "tcp"
    from_port   = 80
    to_port     = 80
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name        = "${var.app_name}-security-group"
    Environment = var.app_environment
  }
}



resource "aws_lb_target_group" "queuefly_target_group" {
  name        = "${var.app_name}-${var.app_environment}-tg"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.aws_vpc_queuefly.id

  health_check {
    healthy_threshold   = "3"
    interval            = "300"
    protocol            = "HTTP"
    matcher             = "200"
    timeout             = "3"
    path                = "/v1/status"
    unhealthy_threshold = "2"
  }

  tags = {
    Name        = "${var.app_name}-loadbalancer-target-group"
    Environment = var.app_environment
  }
}


resource "aws_lb_listener" "queuefly_listener" {
  load_balancer_arn = aws_alb.queuefly_alb.id
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.queuefly_target_group.id
  }







}







