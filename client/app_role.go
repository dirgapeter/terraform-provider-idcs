package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAppRoles - Returns list of apps
func (c *Client) GetAppRoles() ([]AppRole, error) {
	search := map[string]interface{}{
		"schemas":    []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"},
		"sortBy":     "displayName",
		"count":      1000,
		"startIndex": 1,
	}

	rb, err := json.Marshal(search)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/v1/AppRoles/.search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	appRoles := AppRoles{}
	err = json.Unmarshal(body, &appRoles)
	if err != nil {
		return nil, err
	}

	return appRoles.Resources, nil
}
