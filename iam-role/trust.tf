data "http" "sandbox_trust_policy" {
  url = "${local.artifact_base_url}/${chomp(data.http.sandbox_version.response_body)}/trust.json"
}
