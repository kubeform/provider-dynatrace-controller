package requestnaming

import (
	"encoding/json"
	"errors"
	"fmt"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://siz65484.live.dynatrace.com/api/config/v1"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(mw *RequestNaming) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/service/requestNaming", mw, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(id string, mw *RequestNaming) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/service/requestNaming/%s", id), mw, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the MaintenanceWindow to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/service/requestNaming/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*RequestNaming, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the MaintenanceWindow to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/service/requestNaming/%s", id), 200); err != nil {
		return nil, err
	}
	var mw RequestNaming
	if err = json.Unmarshal(bytes, &mw); err != nil {
		return nil, err
	}
	return &mw, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/service/requestNaming", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

func (cs *ServiceClient) GetOrder() (*Order, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/service/requestNaming", 200); err != nil {
		return nil, err
	}
	var result Order
	if err = json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (cs *ServiceClient) UpdateOrder(order *Order) error {
	if _, err := cs.client.PUT("/service/requestNaming/order", order, 204); err != nil {
		return err
	}
	return nil
}
