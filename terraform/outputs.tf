output "http_api_url" {
  value = aws_apigatewayv2_api.http_tester.api_endpoint
}

output "http_api_id" {
  value = aws_apigatewayv2_api.http_tester.id
}