package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Client -
type Client struct {
	HostURL     string
	HTTPClient  *http.Client
	AccessToken string
}

// AuthResponse -
type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// NewClient -
func NewClient(host, clientId, clientSecret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	// if host != nil {
	// 	c.HostURL = *host
	// }

	if (clientId != nil) && (clientSecret != nil) {
		// form request body
		// rb, err := json.Marshal(AuthStruct{
		// 	Username: *username,
		// 	Password: *password,
		// })
		// if err != nil {
		// 	return nil, err
		// }

		data := url.Values{}
		data.Set("grant_type", "client_credentials")
		data.Set("scope", "urn:opc:idm:__myscopes__")
		encodedData := data.Encode()
		log.Println(encodedData)

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth2/v1/token", c.HostURL), strings.NewReader(encodedData))
		if err != nil {
			return nil, err
		}

		req.SetBasicAuth(*clientId, *clientSecret)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))

		body, err := c.doRequest(req)

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		c.AccessToken = ar.AccessToken
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.AccessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	}

	log.Println(req)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Println((string(body)))
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
