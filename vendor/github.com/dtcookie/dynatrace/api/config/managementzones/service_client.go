package managementzones

import (
	"encoding/json"
	"errors"
	"fmt"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
	"github.com/dtcookie/opt"
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
func (cs *ServiceClient) Create(managementzone *ManagementZone) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if len(opt.String(managementzone.ID)) > 0 {
		return nil, errors.New("you MUST NOT provide an ID within the ManagementZone payload upon creation")
	}

	if bytes, err = cs.client.POST("/managementZones", managementzone, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(managementzone *ManagementZone) error {
	if len(opt.String(managementzone.ID)) == 0 {
		return errors.New("the configuration doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/managementZones/%s", opt.String(managementzone.ID)), managementzone, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the configuration to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/managementZones/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string, includeProcessGroupRefs bool) (*ManagementZone, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the configuration to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/managementZones/%s?includeProcessGroupRefs=%v", id, includeProcessGroupRefs), 200); err != nil {
		return nil, err
	}
	var managementzone ManagementZone
	if err = json.Unmarshal(bytes, &managementzone); err != nil {
		return nil, err
	}
	return &managementzone, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() ([]*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/managementZones", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return stubList.Values, nil
}
