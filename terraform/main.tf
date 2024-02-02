terraform {
  required_providers {
    dotenv = {
      source  = "jrhouston/dotenv"
      version = "~> 1.0"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "= 5.31.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  access_key                  = "mock_access_key"
  secret_key                  = "mock_secret_key"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

data dotenv envs {
  filename = "../.env"
}

resource "aws_iam_role" "lambda_role" {
  name = "lambda_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_lambda_function" "app" {
  function_name = "app"
  handler       = "main"
  runtime       = "go1.x"
  role          = aws_iam_role.lambda_role.arn

  filename = "../bin/main.zip"
}

resource "aws_apigatewayv2_api" "http_tester" {
  name          = "-http-connection-tester"
  protocol_type = "HTTP"

  cors_configuration {
    allow_credentials = true
    allow_headers = ["*"]
    allow_methods = ["GET", "POST", "PATCH", "PUT", "DELETE", "HEAD", "OPTIONS"]
    allow_origins = ["http://localhost:5713", "http://localhost:5173"]
    expose_headers = ["*"]
    max_age = 300
  }
}

resource "aws_apigatewayv2_stage" "http_tester_stage" {
  api_id      = aws_apigatewayv2_api.http_tester.id
  name        = "dev"
  auto_deploy = true
}

resource "aws_apigatewayv2_integration" "tester_integration" {
  api_id           = aws_apigatewayv2_api.http_tester.id
  integration_type = "AWS_PROXY"

  connection_type      = "INTERNET"
  description          = "Integration with Lambda"
  integration_uri      = aws_lambda_function.app.invoke_arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "_route" {
  api_id    = aws_apigatewayv2_api.http_tester.id
  route_key = "ANY /{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.tester_integration.id}"
}