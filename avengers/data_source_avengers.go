package avengers

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var HostURL string = "http://localhost:8000"

//dataSourceAvengers is the Avengers data source which will pull information on all Avengers served by avengers-backend.
func dataSourceAvengers() *schema.Resource {
	return &schema.Resource{
		//to read All Avengers, we can directly call resourceAvengersRead()
		//which is implemented in resource_avengers.go file.
		//But watch the Schema, here KEYs are 'Computed: true' not 'Required: true'
		//because we don't want to provide these values while read.
		ReadContext: resourceAvengersRead,
		Schema: map[string]*schema.Schema{
			"avengers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "the _id value returned from mongodb",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of avenger",
						},
						"alias": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "any alias of avenger",
						},
						"weapon": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "his/her special weapons",
						},
					},
				},
			},
		},
	}
}
