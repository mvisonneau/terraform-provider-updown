# Output ipv4 and ipv6 nodes addresses list
data "updown_nodes" "global" {}

output "updown_nodes_ipv4" {
  value = data.updown_nodes.global.ipv4
}

output "updown_nodes_ipv6" {
  value = data.updown_nodes.global.ipv6
}
