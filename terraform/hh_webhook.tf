resource "aws_iam_role" "hh_webhook_iam_role" {
  name = "health-hacker-prototype-001-webhook"

  assume_role_policy = file("./assume_role.json")
}

resource "aws_iam_role_policy" "hh_webhook_iam_policy" {
  name = "health-hacker-prototype-001-webhook"
  role = aws_iam_role.hh_webhook_iam_role.id
  policy = file("./webhook_policy.json")
}

resource "aws_lambda_function" "hh_webhook_lambda" {
  filename = "${path.module}/../build/hh_webhook/function.zip"
  function_name = "hh-webhook"
  role = aws_iam_role.hh_webhook_iam_role.arn
  handler = "main"
  runtime = "go1.x"
  source_code_hash = filebase64sha256("${path.module}/../build/hh_webhook/function.zip")


  environment {
    variables = {
      SLACK_API_KEY = var.slack_api_key
    }
  }
}

resource "aws_dynamodb_table" "hh_webhook_mode_db" {
  name              = "patientMode"
  read_capacity     = 1
  write_capacity    = 1
  hash_key          = "UserId"
  attribute {
    name = "UserId"
    type = "S"
  }
}
