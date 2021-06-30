package idcs

import (
	"context"
	"log"

	"terraform-provider-idcs/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	idcsUrl := d.Get("idcs_url").(string)
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	log.Println("[DEBUG] Something happened!")

	if (idcsUrl != "") && (clientId != "") && (clientSecret != "") {
		c, err := client.NewClient(&idcsUrl, &clientId, &clientSecret)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := client.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"idcs_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("IDCS_URL", nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("IDCS_CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("IDCS_CLIENT_SECRET", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"idcs_grant": resourceGrant(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"idcs_groups":    dataSourceGroups(),
			"idcs_apps":      dataSourceApps(),
			"idcs_app_roles": dataSourceAppRoles(),
			"idcs_grants":    dataSourceGrants(),
			// "group":  dataSourceGroup(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
