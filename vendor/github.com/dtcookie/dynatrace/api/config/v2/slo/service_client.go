package slo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v2"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(item *SLO) (string, error) {
	var err error

	var id string
	post := cs.client.NewPOST("/slo", item)
	post = post.Customize(func(resp *http.Response) {
		location := resp.Header.Get("Location")
		// log.Println("Location: " + location)
		if len(location) > 0 {
			parts := strings.Split(location, "/")
			// log.Println(fmt.Sprintf("len(parts): %d", len(parts)))
			if len(parts) > 0 {
				id = parts[len(parts)-1]
				// log.Println("id: " + id)
			}
		}
	}).Expect(201)
	if _, err = post.Send(); err != nil {
		return id, err
	}
	length := 0
	var bytes []byte
	for length == 0 {
		if bytes, err = cs.client.GET(fmt.Sprintf("/slo?sloSelector=id(\"%s\")&pageSize=10000&sort=name&timeFrame=CURRENT&pageIdx=1&demo=false&evaluate=false", id), 200); err != nil {
			return id, err
		}
		var slos sloList
		if err = json.Unmarshal(bytes, &slos); err != nil {
			return id, err
		}
		length = len(slos.SLOs)
		if length == 0 {
			time.Sleep(time.Second * 2)
		}
		for _, stub := range slos.SLOs {
			item.Timeframe = stub.Timeframe
		}
	}

	return id, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, item *SLO) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/slo/%s", id), item, 200); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the item to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/slo/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*SLO, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the config to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/slo/%s", id), 200); err != nil {
		return nil, err
	}
	var item SLO
	if err = json.Unmarshal(bytes, &item); err != nil {
		return nil, err
	}
	length := 0

	for length == 0 {
		if bytes, err = cs.client.GET(fmt.Sprintf("/slo?sloSelector=id(\"%s\")&pageSize=10000&sort=name&timeFrame=CURRENT&pageIdx=1&demo=false&evaluate=false", id), 200); err != nil {
			return nil, err
		}
		var slos sloList
		if err = json.Unmarshal(bytes, &slos); err != nil {
			return nil, err
		}
		length = len(slos.SLOs)
		if length == 0 {
			time.Sleep(time.Second * 2)
		}
		for _, stub := range slos.SLOs {
			item.Timeframe = stub.Timeframe
		}
	}

	return &item, nil
}

// List TODO: documentation
func (cs *ServiceClient) List() ([]string, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/slo?pageSize=10000&sort=name&timeFrame=CURRENT&pageIdx=1&demo=false&evaluate=false", 200); err != nil {
		return nil, err
	}
	var slos sloList
	if err = json.Unmarshal(bytes, &slos); err != nil {
		return nil, err
	}
	ids := []string{}
	for _, stub := range slos.SLOs {
		ids = append(ids, stub.ID)
	}

	return ids, nil
}

type sloList struct {
	SLOs        []*sloListEntry `json:"slo"`
	PageSize    *int32          `json:"pageSize"`
	NextPageKey *string         `json:"nextPageKey,omitempty"`
	TotalCount  *int64          `json:"totalCount"`
}

type sloListEntry struct {
	ID        string `json:"id"`
	Timeframe string `json:"timeframe"`
}
