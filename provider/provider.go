package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// const (
// 	host_name	= "host"
// 	port_num	= "port"
// )

// Provider returns a terraform provider for the Demo WebService
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is host to connect to",
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is port to connect to",
			},
		},
		ConfigureContextFunc: setupServiceContext,
	}
}

func setupServiceContext(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := rd.Get("host").(string)
	port := rd.Get("port").(string)

	var d diag.Diagnostics
	var hostPort = fmt.Sprintf("%s:%s", host, port)

	client := http.ListenAndServe(hostPort, nil)
	return client, d
}
