provider "aws" {
  region     = "eu-west-1"
}

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
  filename = "${path.module}/../build/telegram_webhook/function.zip"
  function_name = "${local.app_name}-telegram-webhook"
  role = aws_iam_role.lambda_telegram_webhook.arn
  handler = "main"
  runtime = "go1.x"
  source_code_hash = filebase64sha256("${path.module}/../build/telegram_webhook/function.zip")
}
