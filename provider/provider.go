package provider

import (
	"context"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

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
		// Data source name has to starts with the provider
		DataSourcesMap: map[string]*schema.Resource {
			"demo_users": dataSourceService(),
		},
		// ResourceMap : map[string]*schema.Resource{},
		ConfigureContextFunc: setupServiceContext,
	}
}

func setupServiceContext(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := rd.Get("host").(string)
	port := rd.Get("port").(string)

	var d diag.Diagnostics
	var hostPort = fmt.Sprintf("%s:%s", host, port)
	fmt.Println(hostPort)

	client := http.Client{}
	return client, d
}


func dataAllUsers() string {
	var usersURL = fmt.Sprintf("%s", "http://localhost:3000/users")
	resp, err := http.Get(usersURL)
	if err != nil {
		fmt.Println("Unable to get response")
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	return string(body)
}

type User struct {
	ID        string
	FirstName string
	LastName  string
}

func dataUserById(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	host := rd.Get("host").(string)
	port := rd.Get("port").(string)
	userId := rd.Get("id")
	var userIdURL = fmt.Sprintf("http://%s:%s/users/%s", host, port, userId)
	
	// Get response of the request
	resp, err := http.Get(userIdURL)
	if err != nil {
		fmt.Println(err)
	}

	// Get response body of the request
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	user := User{}
	jsonErr := json.Unmarshal(body, &user)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	var ds diag.Diagnostics

	// single value of Diagnostic (not Diagnostics)
	ds = append(ds, diag.Diagnostic{
		Severity:	diag.Warning,
		Detail:		fmt.Sprintf("%+v", rd),
		Summary:	fmt.Sprintf("%+v", rd),
	})

	// receivedUsers := map[string]string{
	// 	"firstname": rd.FirstName,
	// 	"lastname": rd.LastName,
	// }

	// for key, val := range receivedUsers {
	// 	if err := d.Set(key, val); err != nil {
	// 		return diag.FromErr(err)
	// 	}
	// }
	
	rd.SetId(user.ID)

	return ds
}

func dataSourceService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataUserById,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:		schema.TypeString,
				Required:	true,
			},
			"firstname": {
				Type:		schema.TypeString,
				Computed:	true,
			},
			"lastname": {
				Type:		schema.TypeString,
				Computed:	true,
			},
			// "users": {
			// 	Type: schema.TypeMap,
			// 	Computed: true,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 		Sensitive: true,
			// 	},
			// },
		},
	}
}
