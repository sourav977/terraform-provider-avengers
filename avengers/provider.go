package avengers

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"my_avengers": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"my_avengers": dataSourceAvengers(),
		},
	}
}
