package idcs

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"terraform-provider-idcs/client"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenAppsData(apps *[]client.App) []map[string]interface{} {
	if apps != nil {
		ois := make([]map[string]interface{}, len(*apps), len(*apps))

		for i, app := range *apps {
			oi := make(map[string]interface{})

			oi["id"] = app.ID
			oi["display_name"] = app.DisplayName

			ois[i] = oi
		}

		return ois
	}

	return make([]map[string]interface{}, 0)
}

func dataSourceAppsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	apps, err := c.GetApps()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("apps", flattenAppsData(&apps)); err != nil {
		return diag.FromErr(err)
	}
	log.Println(fmt.Sprintf("d %#v", d))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceApps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppsRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"apps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
