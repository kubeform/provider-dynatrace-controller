package customservices

import (
	"encoding/json"
	"fmt"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// Service TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService TODO: documentation
// "https://#######.live.dynatrace.com/api/config/v1", "###########"
func NewService(baseURL string, token string) *ServiceClient {
	rest.Verbose = false
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(customService *CustomService, technology Technology) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST(fmt.Sprintf("/service/customServices/%s", technology), customService, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(customService *CustomService, technology Technology) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/service/customServices/%s/%s", technology, *customService.ID), customService, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string, technology Technology) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/service/customServices/%s/%s", technology, id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string, technology Technology, includeProcessGroupReferences bool) (*CustomService, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/service/customServices/%s/%s?includeProcessGroupReferences=%v", technology, id, includeProcessGroupReferences), 200); err != nil {
		return nil, err
	}
	var customService CustomService
	if err = json.Unmarshal(bytes, &customService); err != nil {
		return nil, err
	}
	return &customService, nil
}

// List TODO: documentation
func (cs *ServiceClient) List(technology Technology) (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/service/customServices/%s", technology), 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}

// Technology has no documentation
type Technology string

// Technologies offers the known enum values
var Technologies = struct {
	DotNet Technology
	Go     Technology
	Java   Technology
	NodeJS Technology
	PHP    Technology
}{
	"dotNet",
	"go",
	"java",
	"nodeJS",
	"php",
}
