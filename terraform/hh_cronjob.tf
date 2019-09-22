resource "aws_iam_role" "hh_cronjob_iam_role" {
  name = "health-hacker-prototype-001-cronjob"

  assume_role_policy = file("./assume_role.json")
}

resource "aws_lambda_function" "hh_cronjob_lambda" {
  filename = "${path.module}/../build/hh_cronjob/function.zip"
  function_name = "hh-cronjob"
  role = aws_iam_role.hh_cronjob_iam_role.arn
  handler = "main"
  runtime = "go1.x"
  source_code_hash = filebase64sha256("${path.module}/../build/hh_cronjob/function.zip")
  timeout = 900

  environment {
    variables = {
      SLACK_API_KEY = var.slack_api_key
    }
  }
}
