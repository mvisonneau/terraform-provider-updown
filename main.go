package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mvisonneau/terraform-provider-updown/updown"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: updown.Provider,
	})
}
