output "lambda_function_url" {
  description = "Lambda Function URL for the API"
  value       = aws_lambda_function_url.api.function_url
}

output "frontend_bucket_name" {
  description = "S3 bucket name for frontend"
  value       = aws_s3_bucket.frontend.bucket
}

output "frontend_website_url" {
  description = "S3 website URL for frontend"
  value       = "http://${aws_s3_bucket.frontend.bucket}.s3-website-${data.aws_region.current.name}.amazonaws.com"
}

output "api_endpoint" {
  description = "GraphQL API endpoint (for frontend PUBLIC_API_URL)"
  value       = "${aws_lambda_function_url.api.function_url}query"
}
