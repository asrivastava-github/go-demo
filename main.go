package main

import (
	"github.com/learn/godemo/provider"

	// Third Party
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"fmt"
)

func main() {
	fmt.Println("I am build a terraform provider using GO!!")
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider()
		},
	})
}