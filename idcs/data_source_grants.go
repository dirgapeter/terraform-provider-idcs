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

func flattenGrantsData(grants *[]client.Grant) []map[string]interface{} {
	if grants != nil {
		ois := make([]map[string]interface{}, len(*grants), len(*grants))

		for i, grant := range *grants {
			oi := make(map[string]interface{})

			oi["id"] = grant.ID
			oi["composite_key"] = grant.CompositeKey

			ois[i] = oi
		}

		return ois
	}

	return make([]map[string]interface{}, 0)
}

func dataSourceGrantsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	grants, err := c.GetGrants()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("grants", flattenGrantsData(&grants)); err != nil {
		return diag.FromErr(err)
	}
	log.Println(fmt.Sprintf("d %#v", d))

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceGrants() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGrantsRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"grants": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"composite_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
