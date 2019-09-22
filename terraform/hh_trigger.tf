resource "aws_iam_role" "hh_trigger_iam_role" {
  name = "health-hacker-prototype-001-trigger"

  assume_role_policy = file("./assume_role.json")
}

resource "aws_lambda_function" "hh_trigger_lambda" {
  filename = "${path.module}/../build/hh_trigger/function.zip"
  function_name = "hh-trigger"
  role = aws_iam_role.hh_trigger_iam_role.arn
  handler = "main"
  runtime = "go1.x"
  source_code_hash = filebase64sha256("${path.module}/../build/hh_trigger/function.zip")


  environment {
    variables = {
      SLACK_API_KEY = var.slack_api_key
    }
  }
}
