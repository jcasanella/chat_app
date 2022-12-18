terraform {
  backend "s3" {
    bucket         = "<your_bucket_name>"
    key            = "terraform.tfstate"
    region         = var.region
    dynamodb_table = "<your_dynamo_dbtable_name>"
  }
}