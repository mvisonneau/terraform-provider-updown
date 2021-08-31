# terraform-provider-updown

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mvisonneau/gitlab-ci-pipelines-exporter)](https://pkg.go.dev/mod/github.com/mvisonneau/gitlab-ci-pipelines-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/gitlab-ci-pipelines-exporter)](https://goreportcard.com/report/github.com/mvisonneau/gitlab-ci-pipelines-exporter)
[![test](https://github.com/mvisonneau/gitlab-ci-pipelines-exporter/actions/workflows/test.yml/badge.svg)](https://github.com/mvisonneau/gitlab-ci-pipelines-exporter/actions/workflows/test.yml)
[![release](https://github.com/mvisonneau/gitlab-ci-pipelines-exporter/actions/workflows/release.yml/badge.svg)](https://github.com/mvisonneau/gitlab-ci-pipelines-exporter/actions/workflows/release.yml)

Terraform provider for [updown.io](https://updown.io)

## Resources

| TYPE | NAME | DESCRIPTION |
|---|---|---|
| **data** |`updown_nodes`| Returns the list of testing nodes ipv4 and ipv6 addresses |
| **resource** |`updown_check`| Creates a check |
| **resource** |`updown_webhook`| Creates a webhook |

## Example usage

```hcl
# Import the provider
terraform {
  required_providers {
    updown = {
      source = "mvisonneau/updown"
    }
  }
}

# Configure it
provider "updown" {
  # Can also be set using UPDOWN_API_KEY env variable.
  api_key = "<YOUR_UPDOWN_API_KEY>"
}

# Create a check
resource "updown_check" "mywebsite" {
  alias        = "https://example.com"
  apdex_t      = 1.0
  enabled      = true
  period       = 30
  published    = true
  url          = "https://test.example.com/healthz"
  string_match = "OK"
  mute_until   = "tomorrow"

  disabled_locations = [
    "mia",
  ]

  custom_headers = {
    "X-GREAT-HEADER" = "yay!"
  }
}

# Add a webhook
resource "updown_webhook" "mywebhook" {
  url = "https://my-nice-webhook.com"
}

# Output ipv4 and ipv6 nodes addresses list
data "updown_nodes" "global" {}

output "updown_nodes_ipv4" {
  value = data.updown_nodes.global.ipv4
}

output "updown_nodes_ipv6" {
  value = data.updown_nodes.global.ipv6
}
```

## Building the provider

```bash
~$ export PROVIDER_PATH=${GOPATH}/src/github.com/mvisonneau/terraform-provider-updown
~$ mkdir -p ${PROVIDER_PATH}; cd ${PROVIDER_PATH}
~$ git clone git@github.com:mvisonneau/terraform-provider-updown .
~$ make install
```

## TODO

- Add tests, need to figure out how to get a mocking endpoint
