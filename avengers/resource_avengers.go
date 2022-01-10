package avengers

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	aclient "github.com/sourav977/avengers-client"
)

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAvengersCreate,
		ReadContext:   resourceAvengersRead,
		UpdateContext: resourceAvengersUpdate,
		DeleteContext: resourceAvengersDelete,
		Schema: map[string]*schema.Schema{
			"_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the _id value returned from mongodb",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of avenger",
			},
			"alias": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "any alias of avenger",
			},
			"weapon": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "his/her special weapons",
			},
			"DeletedCount": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "deleted item count",
			},
			"MatchedCount": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "total matched item found",
			},
			"ModifiedCount": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "total item modified",
			},
			"UpsertedCount": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "total item upserted",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAvengersCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(aclient.Client)

	name := d.Get("name")
	alias := d.Get("alias")
	weapon := d.Get("weapon")

	a := aclient.Avenger{
		Name:   name.(string),
		Alias:  alias.(string),
		Weapon: weapon.(string),
	}

	res, err := c.CreateAvenger(a)
	if err != nil {
		return diag.FromErr(err)
	}

	resItem := flattenAvenger(res)
	r := resItem.(aclient.Avenger)
	if err := d.Set("_id", r.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", r.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("alias", r.Alias); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("weapon", r.Weapon); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(res.ID))

	return diags
}

func resourceAvengersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(aclient.Client)
	res, err := c.GetAllAvengers()
	if err != nil {
		return diag.FromErr(err)
	}

	resItems := flattenAvengers(&res)
	if err := d.Set("avengers", resItems); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceAvengersUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(aclient.Client)

	name := d.Get("name")
	alias := d.Get("alias")
	weapon := d.Get("weapon")

	a := aclient.Avenger{
		Name:   name.(string),
		Alias:  alias.(string),
		Weapon: weapon.(string),
	}
	res, err := c.UpdateAvengerByName(a)
	if err != nil {
		return diag.FromErr(err)
	}

	resItem := flattenUpdateItem(res)
	r := resItem.(aclient.UpdateResult)
	if err := d.Set("MatchedCount", r.MatchedCount); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("ModifiedCount", r.ModifiedCount); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("UpsertedCount", r.UpsertedCount); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
func resourceAvengersDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	c := m.(aclient.Client)
	name := d.Get("name")
	del, err := c.DeleteAvengerByName(name.(string))
	if err != nil {
		return diag.FromErr(err)
	}
	resItems := flattenDeleteItem(del)
	if err := d.Set("DeletedCount", resItems); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func flattenAvenger(avenger *aclient.Avenger) interface{} {
	al := make(map[string]interface{})
	al["_id"] = avenger.ID
	al["avenger_name"] = avenger.Name
	al["avenger_alias"] = avenger.Alias
	al["avenger_weapon"] = avenger.Weapon
	return al
}

func flattenUpdateItem(update *aclient.UpdateResult) interface{} {
	up := make(map[string]interface{})
	up["MatchedCount"] = update.MatchedCount
	up["ModifiedCount"] = update.ModifiedCount
	up["UpsertedCount"] = update.UpsertedCount
	return up
}

func flattenAvengers(avengersList *[]aclient.Avenger) []interface{} {
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

func flattenDeleteItem(res *aclient.DeleteResult) []interface{} {
	a := make(map[string]interface{})
	a["DeletedCount"] = res.DeletedCount
	return []interface{}{a}
}
