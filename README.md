# terraform-provider-updown

[![GoDoc](https://godoc.org/github.com/mvisonneau/terraform-provider-updown?status.svg)](https://godoc.org/github.com/mvisonneau/terraform-provider-updown/app)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvisonneau/terraform-provider-updown)](https://goreportcard.com/report/github.com/mvisonneau/terraform-provider-updown)
[![Build Status](https://cloud.drone.io/api/badges/mvisonneau/terraform-provider-updown/status.svg)](https://cloud.drone.io/mvisonneau/terraform-provider-updown)
[![Coverage Status](https://coveralls.io/repos/github/mvisonneau/terraform-provider-updown/badge.svg?branch=main)](https://coveralls.io/github/mvisonneau/terraform-provider-updown?branch=main)

Terraform provider for [updown.io](https://updown.io)

## Resources

| TYPE | NAME | DESCRIPTION |
|---|---|---|
| **data** |`updown_nodes`| Returns the list of testing nodes ipv4 and ipv6 addresses |
| **resource** |`updown_check`| Creates a check |

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

# Output ipv4 and ipv6 nodes addresses list
data "updown_nodes" "global" {}

output "updown_nodes_ipv4" {
  value = data.updown_nodes.global.ipv4
}

output "updown_nodes_ipv6" {
  value = data.updown_nodes.global.ipv6
}
```

## Import

### updown_check

```
terraform import updown_check.my_website <check_id>
```

The check_id is basically whatever 4 characters you have after https://updown.io/
It looks like the following regexp : ^https:\/\/updown.io\/([a-z0-9]{4})$

## Building the provider

```bash
~$ export PROVIDER_PATH=${GOPATH}/src/github.com/mvisonneau/terraform-provider-updown
~$ mkdir -p ${PROVIDER_PATH}; cd ${PROVIDER_PATH}
~$ git clone git@github.com:mvisonneau/terraform-provider-updown .
~$ make build-local
```

## TODO

- Add tests, need to figure out how to get a mocking endpoint
- Documentation
