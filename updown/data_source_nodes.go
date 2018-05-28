package updown

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mvisonneau/updown"
)

func nodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: nodesList,

		Schema: map[string]*schema.Schema{
      "ipv4_addresses": {
				Type:        schema.TypeList,
				Computed: true,
				Description: "Ipv4 addresses list of the nodes.",
        Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ipv6_addresses": {
				Type:        schema.TypeList,
				Computed: true,
				Description: "Ipv6 addresses list or the nodes.",
        Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
    },
	}
}

func nodesList(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*updown.Client)

  ipv4_addresses, _, err := client.Node.ListIPv4()
  if err != nil {
		return fmt.Errorf("Error reading ipv4 addresses from API")
	}

	ipv6_addresses, _, err := client.Node.ListIPv6()
  if err != nil {
		return fmt.Errorf("Error reading ipv6 addresses from API")
	}

	d.SetId("updown.io/nodes")
  d.Set("ipv4_addresses", ipv4_addresses)
	d.Set("ipv6_addresses", ipv6_addresses)

	return nil
}
