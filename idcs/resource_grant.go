package idcs

import (
	"context"
	"log"
	"strconv"
	"terraform-provider-idcs/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGrant() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGrantCreate,
		ReadContext:   resourceGrantRead,
		UpdateContext: resourceGrantUpdate,
		DeleteContext: resourceGrantDelete,
		Schema: map[string]*schema.Schema{
			"app": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"grantee": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"entitlement": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"attribute_value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"grant_mechanism": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGrantCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//items := d.Get("items").([]interface{})
	grant := client.Grant{}
	app := d.Get("app").([]interface{})[0].(map[string]interface{})
	grantee := d.Get("grantee").([]interface{})[0].(map[string]interface{})
	entitlement := d.Get("entitlement").([]interface{})[0].(map[string]interface{})

	grant.App.Value = app["value"].(string)
	grant.Grantee.Type = grantee["type"].(string)
	grant.Grantee.Value = grantee["value"].(string)
	grant.Entitlement.AttributeName = entitlement["attribute_name"].(string)
	grant.Entitlement.AttributeValue = entitlement["attribute_value"].(string)
	grant.GrantMechanism = d.Get("grant_mechanism").(string)

	log.Println(grant)
	// for _, item := range items {
	// 	i := item.(map[string]interface{})

	// 	co := i["coffee"].([]interface{})[0]
	// 	coffee := co.(map[string]interface{})

	// 	oi := hc.OrderItem{
	// 		Coffee: hc.Coffee{
	// 			ID: coffee["id"].(int),
	// 		},
	// 		Quantity: i["quantity"].(int),
	// 	}

	// 	ois = append(ois, oi)
	// }

	g, err := c.CreateGrant(grant)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(g.ID))

	return diags
}

func resourceGrantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceGrantUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceGrantRead(ctx, d, m)
}

func resourceGrantDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
