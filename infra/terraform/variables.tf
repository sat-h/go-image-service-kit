variable "AWS_ACCESS_KEY_ID" {
  type        = string
  description = "AWS Access Key ID"
}

variable "AWS_SECRET_ACCESS_KEY" {
  type        = string
  description = "AWS Secret Access Key"
}

variable "AWS_REGION" {
  type        = string
  description = "AWS Region"
}

variable "s3_bucket_name" {
  type        = string
  description = "Name of the S3 bucket"
}

variable "dynamodb_table_name" {
  type        = string
  description = "Name of the DynamoDB table"
}
