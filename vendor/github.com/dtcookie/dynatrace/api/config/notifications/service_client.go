package notifications

import (
	"encoding/json"
	"errors"
	"fmt"

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
func (cs *ServiceClient) Create(config *NotificationRecord) (*Stub, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.NotificationConfig.GetID())) > 0 {
		return nil, errors.New("you MUST NOT provide an ID within the Notification payload upon creation")
	}

	if bytes, err = cs.client.POST("/notifications", config, 201); err != nil {
		return nil, err
	}
	var stub Stub
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(config *NotificationRecord) error {
	if len(opt.String(config.NotificationConfig.GetID())) == 0 {
		return errors.New("the Notification doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/notifications/%s", opt.String(config.NotificationConfig.GetID())), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the Notification to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/notifications/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*NotificationRecord, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the Notification to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/notifications/%s", id), 200); err != nil {
		return nil, err
	}
	var record NotificationRecord
	if err = json.Unmarshal(bytes, &record); err != nil {
		return nil, err
	}
	return &record, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/notifications", 200); err != nil {
		return nil, err
	}
	var stubList StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
