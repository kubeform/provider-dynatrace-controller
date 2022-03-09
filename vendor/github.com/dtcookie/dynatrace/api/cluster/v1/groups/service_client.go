package groups

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// ServiceClient TODO: documentation
type ServiceClient struct {
	client *rest.Client
}

// NewService creates a new Service Client
// baseURL should look like this: "https://#######.live.dynatrace.com/api/config/v1"
// token is an API Token
func NewService(baseURL string, token string) *ServiceClient {
	credentials := credentials.New(token)
	config := rest.Config{Insecure: true}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(groupConfig *GroupConfig) (*GroupConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/groups", groupConfig, 200); err != nil {
		return nil, err
	}
	var createdGroupConfig GroupConfig
	if err = json.Unmarshal(bytes, &createdGroupConfig); err != nil {
		return nil, err
	}
	return &createdGroupConfig, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(groupConfig *GroupConfig) error {
	if _, err := cs.client.PUT("/groups", groupConfig, 200); err != nil {
		return err
	}

	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/groups/%s", id), 200); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*GroupConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/groups/%s", id), 200); err != nil {
		return nil, err
	}
	var groupConfig GroupConfig
	if err = json.Unmarshal(bytes, &groupConfig); err != nil {
		return nil, err
	}
	return &groupConfig, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() ([]*GroupConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/groups", 200); err != nil {
		return nil, err
	}
	var groups []*GroupConfig
	if err = json.Unmarshal(bytes, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}
