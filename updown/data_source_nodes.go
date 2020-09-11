package updown

import (
	"fmt"

	"github.com/antoineaugusti/updown"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func nodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: nodesList,

		Schema: map[string]*schema.Schema{
			"ipv4": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Ipv4 addresses list of the nodes.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ipv6": {
				Type:        schema.TypeList,
				Computed:    true,
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

	ipv4, _, err := client.Node.ListIPv4()
	if err != nil {
		return fmt.Errorf("Error reading ipv4 addresses from API")
	}

	ipv6, _, err := client.Node.ListIPv6()
	if err != nil {
		return fmt.Errorf("Error reading ipv6 addresses from API")
	}

	d.SetId("updown.io/nodes")

	for k, v := range map[string]interface{}{
		"ipv4": ipv4,
		"ipv6": ipv6,
	} {
		if err := d.Set(k, v); err != nil {
			return err
		}
	}

	return nil
}
