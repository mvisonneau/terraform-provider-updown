package provider

import (
	"github.com/antoineaugusti/updown"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New returns a Terraform provider resource
func New() func() *schema.Provider {
	return func() *schema.Provider {
		return &schema.Provider{
			Schema: map[string]*schema.Schema{
				"api_key": {
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
				"updown_check":     checkResource(),
				"updown_webhook":   webhookResource(),
				"updown_recipient": recipientResource(),
			},
		}
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return updown.NewClient(d.Get("api_key").(string), nil), nil
}
