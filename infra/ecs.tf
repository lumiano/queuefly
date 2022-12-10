resource "aws_ecs_cluster" "queuefly_cluster" {
  name = "${var.app_name}-${var.app_environment}-ecs-cluster"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }

  tags = {
    Name        = "${var.app_name}-ecs-cluster"
    Environment = var.app_environment
  }
}

resource "aws_ecs_task_definition" "queuefly_task_definition" {
  family                   = "${var.app_name}-task"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  execution_role_arn       = aws_iam_role.ecsTaskExecutionRole.arn
  task_role_arn            = aws_iam_role.ecsTaskExecutionRole.arn


  container_definitions = <<DEFINITION
  [
    {
      "name": "${var.app_name}-${var.app_environment}-container",
      "image": "${aws_ecr_repository.queuefly_ecr_repository.repository_url}:latest",
      "entryPoint": [],
      "essential": true,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group.queuefly_cloudwatch.id}",
          "awslogs-region": "${var.aws_region}",
          "awslogs-stream-prefix": "${var.app_name}-${var.app_environment}"
        }
      },
      "portMappings": [
        {
          "containerPort": 3000,
          "hostPort": 3000
        }
      ],
    "cpu": 1024,
    "memory": 2048,
      "networkMode": "awsvpc"
    }
  ]
  DEFINITION


  tags = {
    Name        = "${var.app_name}-queuefly-ecs"
    Environment = var.app_environment
  }
}

data "aws_ecs_task_definition" "main" {
  task_definition = aws_ecs_task_definition.queuefly_task_definition.family
}


resource "aws_ecs_service" "queuefly_service" {
  name                 = "${var.app_name}-${var.app_environment}-ecs-service"
  cluster              = aws_ecs_cluster.queuefly_cluster.id
  task_definition      = "${aws_ecs_task_definition.queuefly_task_definition.family}:${max(aws_ecs_task_definition.queuefly_task_definition.revision, data.aws_ecs_task_definition.main.revision)}"
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = var.app_count
  force_new_deployment = true

  network_configuration {
    subnets = aws_subnet.private.*.id
    security_groups = [
      aws_security_group.queuefly_security.id,
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.queuefly_target_group.id
    container_name   = "${var.app_name}-${var.app_environment}-container"
    container_port   = 3000
  }


  tags = {
    Name        = "${var.app_name}-service"
    Environment = var.app_environment
  }

  depends_on = [aws_lb_listener.queuefly_listener]

}


resource "aws_security_group" "queuefly_security" {
  vpc_id = aws_vpc.aws_vpc_queuefly.id

  ingress {
    protocol        = "tcp"
    from_port       = 3000
    to_port         = 3000
    security_groups = [aws_security_group.queuefly_security_group.id]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.app_name}-service-sg"
    Environment = var.app_environment
  }
}

