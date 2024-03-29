---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "updown_nodes Data Source - terraform-provider-updown"
subcategory: ""
description: |-
  updown_nodes data source can be used to retrieve the IP addresses of their servers.
---

# updown_nodes (Data Source)

`updown_nodes` data source can be used to retrieve the IP addresses of their servers.

## Example Usage

```terraform
# Output ipv4 and ipv6 nodes addresses list
data "updown_nodes" "global" {}

output "updown_nodes_ipv4" {
  value = data.updown_nodes.global.ipv4
}

output "updown_nodes_ipv6" {
  value = data.updown_nodes.global.ipv6
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **ipv4** (List of String) Ipv4 addresses list of the nodes.
- **ipv6** (List of String) Ipv6 addresses list or the nodes.


