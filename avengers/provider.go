package avengers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	aclient "github.com/sourav977/avengers-client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		//as this terraform provider is created for learning purpose, i kept it very simple.
		//we can access all APIs available at avengers-backend without any authentication and authorization.
		//so it will only contain ResourcesMap & DataSourcesMap, no schema
		ResourcesMap: map[string]*schema.Resource{
			"avengers_resource": resourceAvengers(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"avengers_datasource": dataSourceAvengers(),
		},
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVENGERS_BACKEND_HOST_URL", "http://localhost:8000"),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

type ApiClient struct {
	data           *schema.ResourceData
	avengersclient *aclient.Client
}

//NewApiClient will return a new instance of ApiClient using which we can communicate with Avengers-backend
func NewApiClient(d *schema.ResourceData) (*ApiClient, diag.Diagnostics) {
	c := &ApiClient{data: d}
	client, err := c.NewAvengersClient()
	if err != nil {
		return c, diag.FromErr(err)
	}
	c.avengersclient = client
	return c, nil

}

func (a *ApiClient) NewAvengersClient() (*aclient.Client, error) {
	host := a.data.Get("host").(string)
	c, err := aclient.NewClient(&host)
	if err != nil {
		return c, err
	}
	return c, nil
}
