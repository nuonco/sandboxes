resource "aws_iam_policy" "deprovision" {
  name   = "${var.prefix}-nuon-${var.sandbox}-deprovision"
  policy = file("../${var.sandbox}/artifacts/deprovision.json")
}
