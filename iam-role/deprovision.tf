data "http" "sandbox_deprovision_policy" {
  url = "${local.artifact_base_url}/${chomp(data.http.sandbox_version.response_body)}/deprovision.json"
}

resource "aws_iam_policy" "deprovision" {
  name   = "${var.prefix}-nuon-${var.sandbox}-deprovision"
  policy = data.http.sandbox_deprovision_policy.response_body
}
