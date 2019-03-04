package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-fortios/fortios"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortios.Provider})
}
