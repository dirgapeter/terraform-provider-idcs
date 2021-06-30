package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"time"
)

// GetCoffees - Returns list of coffees (no auth required)
func (c *Client) GetGroups() ([]Group, error) {
	// Create trace struct.
	trace, debug := trace()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/v1/Groups", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// Print report.
	data, err := json.MarshalIndent(debug, "", "    ")
	log.Println(string(data))
	log.Println(string(body))

	groups := Groups{}
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return nil, err
	}

	return groups.Resources, nil
}

type Debug struct {
	DNS struct {
		Start   string       `json:"start"`
		End     string       `json:"end"`
		Host    string       `json:"host"`
		Address []net.IPAddr `json:"address"`
		Error   error        `json:"error"`
	} `json:"dns"`
	Dial struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"dial"`
	Connection struct {
		Time string `json:"time"`
	} `json:"connection"`
	WroteAllRequestHeaders struct {
		Time string `json:"time"`
	} `json:"wrote_all_request_header"`
	WroteAllRequest struct {
		Time string `json:"time"`
	} `json:"wrote_all_request"`
	FirstReceivedResponseByte struct {
		Time string `json:"time"`
	} `json:"first_received_response_byte"`
}

func trace() (*httptrace.ClientTrace, *Debug) {
	d := &Debug{}

	t := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			t := time.Now().UTC().String()
			log.Println(t, "dns start")
			d.DNS.Start = t
			d.DNS.Host = info.Host
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			t := time.Now().UTC().String()
			log.Println(t, "dns end")
			d.DNS.End = t
			d.DNS.Address = info.Addrs
			d.DNS.Error = info.Err
		},
		ConnectStart: func(network, addr string) {
			t := time.Now().UTC().String()
			log.Println(t, "dial start")
			d.Dial.Start = t
		},
		ConnectDone: func(network, addr string, err error) {
			t := time.Now().UTC().String()
			log.Println(t, "dial end")
			d.Dial.End = t
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			t := time.Now().UTC().String()
			log.Println(t, "conn time")
			d.Connection.Time = t
		},
		WroteHeaders: func() {
			t := time.Now().UTC().String()
			log.Println(t, "wrote all request headers")
			d.WroteAllRequestHeaders.Time = t
		},
		WroteRequest: func(wr httptrace.WroteRequestInfo) {
			t := time.Now().UTC().String()
			log.Println(t, "wrote all request")
			d.WroteAllRequest.Time = t
		},
		GotFirstResponseByte: func() {
			t := time.Now().UTC().String()
			log.Println(t, "first received response byte")
			d.FirstReceivedResponseByte.Time = t
		},
	}

	return t, d
}

// // GetOrder - Returns a specifc order
// func (c *Client) GetGroup(orderID string) (*Order, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	order := Order{}
// 	err = json.Unmarshal(body, &order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &order, nil
// }

// // CreateOrder - Create new order
// func (c *Client) CreateOrder(orderItems []OrderItem) (*Order, error) {
// 	rb, err := json.Marshal(orderItems)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/orders", c.HostURL), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	order := Order{}
// 	err = json.Unmarshal(body, &order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &order, nil
// }

// // UpdateOrder - Updates an order
// func (c *Client) UpdateOrder(orderID string, orderItems []OrderItem) (*Order, error) {
// 	rb, err := json.Marshal(orderItems)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	order := Order{}
// 	err = json.Unmarshal(body, &order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &order, nil
// }

// // DeleteOrder - Deletes an order
// func (c *Client) DeleteOrder(orderID string) error {
// 	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/orders/%s", c.HostURL, orderID), nil)
// 	if err != nil {
// 		return err
// 	}

// 	body, err := c.doRequest(req)
// 	if err != nil {
// 		return err
// 	}

// 	if string(body) != "Deleted order" {
// 		return errors.New(string(body))
// 	}

// 	return nil
// }
