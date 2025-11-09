variable "aws_region" {
  description = "AWS region for resources"
  type        = string
  default     = "us-east-1"
}

variable "project_name" {
  description = "Project name for resource naming"
  type        = string
  default     = "packager"
}

variable "environment" {
  description = "Environment name (dev, staging, prod)"
  type        = string
  default     = "demo"
}
