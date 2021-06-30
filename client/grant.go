package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetGrants - Returns list of apps
func (c *Client) GetGrants() ([]Grant, error) {
	search := map[string]interface{}{
		"schemas":    []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"},
		"sortBy":     "id",
		"count":      1000,
		"startIndex": 1,
	}

	rb, err := json.Marshal(search)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/v1/Grants/.search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	grants := Grants{}
	err = json.Unmarshal(body, &grants)
	if err != nil {
		return nil, err
	}

	return grants.Resources, nil
}

// CreateGrant -
func (c *Client) CreateGrant() (Grant, error) {
	search := map[string]interface{}{
		"schemas":    []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"},
		"sortBy":     "id",
		"count":      1000,
		"startIndex": 1,
	}

	rb, err := json.Marshal(search)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/v1/Grants/.search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	grants := Grants{}
	err = json.Unmarshal(body, &grants)
	if err != nil {
		return nil, err
	}

	return grants.Resources, nil
}
