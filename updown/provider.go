package updown

import (
	"github.com/antoineaugusti/updown"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPDOWN_API_KEY", ""),
				Description: "API key to use in order to authenticated against updown.io API.",
			},
		},

		ConfigureFunc: providerConfigure,

		DataSourcesMap: map[string]*schema.Resource{
			"updown_nodes": nodesDataSource(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"updown_check": checkResource(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return updown.NewClient(d.Get("api_key").(string), nil), nil
}
