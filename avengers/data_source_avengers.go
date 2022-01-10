package avengers

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	aclient "github.com/sourav977/avengers-client"
)

var HostURL string = "http://localhost:8000"

//dataSourceAvengers is the Avengers data source which will pull information on all Avengers served by avengers-backend.
func dataSourceAvengers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAvengersRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
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

func dataSourceAvengersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	avengersList, err := client.GetAllAvengers()
	if err != nil {
		return diag.FromErr(err)
	}

	mapavengerslist := flattenAvengersList(&avengersList)
	if err := d.Set("avengers", mapavengerslist); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenAvengersList(avengersList *[]aclient.Avenger) []interface{} {
	if avengersList != nil {
		avengers := make([]interface{}, len(*avengersList))

		for i, avenger := range *avengersList {
			al := make(map[string]interface{})

			al["_id"] = avenger.ID
			al["avenger_name"] = avenger.Name
			al["avenger_alias"] = avenger.Alias
			al["avenger_weapon"] = avenger.Weapon

			avengers[i] = al
		}
		return avengers
	}
	return make([]interface{}, 0)
}
