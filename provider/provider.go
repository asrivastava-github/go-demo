package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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
			"demo_user_by_id": dataSourceUserbyID(),
			"demo_users": dataSourceUsers(),
		},
		ResourcesMap : map[string]*schema.Resource{
			"demo_user_service": resourceService(),
		},
		// This triggers to initialise the configuration. Connection etc.
		ConfigureContextFunc: setupServiceContext,
	}
}

type User struct {
	ID        int
	FirstName string
	LastName  string
}

const HostURL string = "http://localhost:5000"

func catchErr(err error) {
	fmt.Println(err)
}

type Client struct {
	HostURL    string
	HTTPClient *http.Client
}

func NewClient(host, port *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: fmt.Sprintf("http://%s:%s/users", *host, *port),
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	
	res, err := c.HTTPClient.Do(req)
	catchErr(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	catchErr(err)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

func setupServiceContext(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := rd.Get("host").(string)
	port := rd.Get("port").(string)

	client, err := NewClient(&host, &port)
	catchErr(err)

	var d diag.Diagnostics
	return client, d
}

func dataUserById(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// userId := rd.Get("id")
	userId := rd.Id()
	fmt.Println(userId)
	client := meta.(*Client)
	
	// Get response of the request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", client.HostURL, userId), nil)
	body, err := client.doRequest(req)
	catchErr(err)

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

	receivedUsers := map[string]string{
		"firstname": 	user.FirstName,
		"lastname": 	user.LastName,
	}

	for key, val := range receivedUsers {
		if err := rd.Set(key, val); err != nil {
			return diag.FromErr(err)
		}
	}
	rd.Set("id", user.ID)
	
	rd.SetId(userId)

	return nil
}

func dataUsers(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Client)
	
	// Get response of the request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s", client.HostURL), nil)
	body, err := client.doRequest(req)
	catchErr(err)

	users := []User{}
	jsonErr := json.Unmarshal(body, &users)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	numberOfUsers := fmt.Sprint(len(users))

	// for i := range users {
	// 	receivedUsers := map[string]string{
	// 		"firstname": 	users[i].FirstName,
	// 		"lastname": 	users[i].LastName,
	// 	};
	// 	for key, val := range receivedUsers {
	// 		if err := rd.Set(key, val); err != nil {
	// 			return diag.FromErr(err)
	// 		}
	// 	}
	// }

	if usrerr := rd.Set("numberofusers", numberOfUsers); usrerr != nil {
		return diag.FromErr(usrerr)
	}
	
	rd.SetId(fmt.Sprint(numberOfUsers))
	return nil
}

func dataSourceUserbyID() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataUserById,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:		schema.TypeInt,
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
		},
	}
}

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataUsers,
		Schema: map[string]*schema.Schema{
			"numberofusers": {
				Type:		schema.TypeInt,
				Computed:	true,
			},
			// "users": {
			// 	Type: schema.TypeList,
			// 	Computed: true,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeMap,
			// 	},
			// },
		},
	}
}

func resourceService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserService,
		UpdateContext: updateResourceUserService,
		ReadContext: dataUserById,
		DeleteContext: dataUserById,
		Schema: map[string]*schema.Schema{
			"firstname": {
				Type:		schema.TypeString,
				Required: 	true,
			},
			"lastname": {
				Type:		schema.TypeString,
				Required: 	true,
			},
		},
	}
}

func resourceUserService(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	firstname := rd.Get("firstname").(string)
	lastname := rd.Get("lastname").(string)
	client := meta.(*Client)

	userData, _ := json.Marshal(map[string]string{
		"firstname":	firstname,
		"lastname":		lastname,
	 })
	responseBody := bytes.NewBuffer(userData)
	fmt.Println(responseBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s", client.HostURL), responseBody)
	body, err := client.doRequest(req)
	catchErr(err)
	defer req.Body.Close()
	users := []User{}
	jsonErr := json.Unmarshal(body, &users)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	rd.SetId(fmt.Sprint(len(users)))

	return nil
}

func updateResourceUserService(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	firstname := rd.Get("firstname").(string)
	lastname := rd.Get("lastname").(string)
	client := meta.(*Client)

	userData, _ := json.Marshal(map[string]string{
		"firstname":	firstname,
		"lastname":		lastname,
	 })
	responseBody := bytes.NewBuffer(userData)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s", client.HostURL), responseBody)
	body, err := client.doRequest(req)
	catchErr(err)
	defer req.Body.Close()
	users := []User{}
	jsonErr := json.Unmarshal(body, &users)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	rd.SetId(string(len(users)))

	return nil
}
