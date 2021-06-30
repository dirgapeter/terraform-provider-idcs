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

func flattenGroupsData(groups *[]client.Group) []map[string]interface{} {
	if groups != nil {
		ois := make([]map[string]interface{}, len(*groups), len(*groups))

		for i, group := range *groups {
			oi := make(map[string]interface{})

			oi["id"] = group.ID
			oi["display_name"] = group.DisplayName

			ois[i] = oi
		}

		return ois
	}

	return make([]map[string]interface{}, 0)
}

func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	groups, err := c.GetGroups()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("groups", flattenGroupsData(&groups)); err != nil {
		return diag.FromErr(err)
	}
	log.Println(fmt.Sprintf("d %#v", d))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Schema: map[string]*schema.Schema{
			"groups": {
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
