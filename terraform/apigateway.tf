resource "aws_api_gateway_rest_api" "webhooks" {
  name = "hh-webhooks"
  policy = file("${path.module}/api_gateway.json")
}

resource "aws_api_gateway_deployment" "test" {
  depends_on  = ["aws_api_gateway_integration.test"]

  rest_api_id = aws_api_gateway_rest_api.webhooks.id
  stage_name  = "test"
}

resource "aws_api_gateway_stage" "test" {
  stage_name    = "prod"
  rest_api_id   = aws_api_gateway_rest_api.webhooks.id
  deployment_id = aws_api_gateway_deployment.test.id
}

resource "aws_api_gateway_resource" "test" {
  rest_api_id = aws_api_gateway_rest_api.webhooks.id
  parent_id   = aws_api_gateway_rest_api.webhooks.root_resource_id
  path_part   = "mytestresource"
}

resource "aws_api_gateway_method" "test" {
  rest_api_id   = aws_api_gateway_rest_api.webhooks.id
  resource_id   = aws_api_gateway_resource.test.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_method_settings" "s" {
  rest_api_id = aws_api_gateway_rest_api.webhooks.id
  stage_name  = aws_api_gateway_stage.test.stage_name
  method_path = "${aws_api_gateway_resource.test.path_part}/${aws_api_gateway_method.test.http_method}"

  settings {
    metrics_enabled = true
    logging_level   = "OFF"
  }
}

resource "aws_api_gateway_integration" "test" {
  rest_api_id = "${aws_api_gateway_rest_api.webhooks.id}"
  resource_id = "${aws_api_gateway_resource.test.id}"
  http_method = "${aws_api_gateway_method.test.http_method}"
  integration_http_method = "POST"

  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${aws_lambda_function.hh_webhook_lambda.arn}/invocations"
}

resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "AllowMyDemoAPIInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.hh_webhook_lambda.function_name}"
  principal     = "apigateway.amazonaws.com"

  # The /*/*/* part allows invocation from any stage, method and resource path
  # within API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.webhooks.execution_arn}/*/*/*"
}

resource "aws_api_gateway_account" "demo" {
  cloudwatch_role_arn = "${aws_iam_role.cloudwatch.arn}"
}

resource "aws_iam_role" "cloudwatch" {
  name = "api_gateway_cloudwatch_hh"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "cloudwatch" {
  name = "default"
  role = "${aws_iam_role.cloudwatch.id}"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams",
                "logs:PutLogEvents",
                "logs:GetLogEvents",
                "logs:FilterLogEvents"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}