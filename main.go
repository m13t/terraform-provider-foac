package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/m13t/terraform-provider-foac/foac"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: foac.Provider,
	})
}
