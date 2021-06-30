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

func flattenAppRolesData(appRoles *[]client.AppRole) []map[string]interface{} {
	if appRoles != nil {
		ois := make([]map[string]interface{}, len(*appRoles), len(*appRoles))

		for i, appRole := range *appRoles {
			oi := make(map[string]interface{})

			oi["id"] = appRole.ID
			oi["display_name"] = appRole.DisplayName

			ois[i] = oi
		}

		return ois
	}

	return make([]map[string]interface{}, 0)
}

func dataSourceAppRolesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	appRoles, err := c.GetAppRoles()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("app_roles", flattenAppRolesData(&appRoles)); err != nil {
		return diag.FromErr(err)
	}
	log.Println(fmt.Sprintf("d %#v", d))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceAppRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAppRolesRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"app_roles": {
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
