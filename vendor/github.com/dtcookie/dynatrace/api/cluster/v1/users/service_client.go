package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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
	config := rest.Config{Insecure: true}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(userConfig *UserConfig) (*UserConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/users", userConfig, 200); err != nil {
		return nil, err
	}
	var createdUserConfig UserConfig
	if err = json.Unmarshal(bytes, &createdUserConfig); err != nil {
		return nil, err
	}
	return &createdUserConfig, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(userConfig *UserConfig) error {
	if _, err := cs.client.PUT("/users", userConfig, 200); err != nil {
		return err
	}

	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the Dashboard to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/users/%s", id), 200); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*UserConfig, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the Dashboard to fetch")
	}

	var err error
	var bytes []byte
	if bytes, err = cs.client.GET(fmt.Sprintf("/users/%s", id), 200); err != nil {
		if strings.HasPrefix(err.Error(), "Not Found (GET) ") {
			return nil, fmt.Errorf("user '%s' doesn't exist", id)
		}
		return nil, err
	}
	var userConfig UserConfig
	if err = json.Unmarshal(bytes, &userConfig); err != nil {
		return nil, err
	}
	return &userConfig, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() ([]*UserConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/users", 200); err != nil {
		return nil, err
	}
	var users []*UserConfig
	if err = json.Unmarshal(bytes, &users); err != nil {
		return nil, err
	}
	return users, nil
}
