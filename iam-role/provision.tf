data "http" "sandbox_provision_policy" {
  url = "${local.artifact_base_url}/${chomp(data.http.sandbox_version.response_body)}/provision.json"
}

resource "aws_iam_policy" "provision" {
  name   = "${var.prefix}-nuon-${var.sandbox}-provision"
  policy = data.http.sandbox_provision_policy.response_body
}
