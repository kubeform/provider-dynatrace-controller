package monitors

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

// CreateBrowser TODO: documentation
func (cs *ServiceClient) CreateBrowser(config *BrowserSyntheticMonitorUpdate) (*string, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.ID)) > 0 {
		return nil, errors.New("you must not provide an ID within the configuration payload upon creation")
	}

	if bytes, err = cs.client.POST("/synthetic/monitors", config, 200); err != nil {
		return nil, err
	}
	stub := struct {
		ID string `json:"entityId"`
	}{}
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	successfulAttempts := 0
	for successfulAttempts < 5 {
		attempts := 0
		for attempts < 50 {
			if _, err := cs.GetBrowser(stub.ID); err == nil {
				attempts = 50
			} else {
				attempts++
				time.Sleep(2 * time.Second)
			}
		}
		successfulAttempts++
	}

	return &stub.ID, nil
}

// CreateHTTP TODO: documentation
func (cs *ServiceClient) CreateHTTP(config *HTTPSyntheticMonitorUpdate) (*string, error) {
	var err error
	var bytes []byte

	if len(opt.String(config.ID)) > 0 {
		return nil, errors.New("you must not provide an ID within the configuration payload upon creation")
	}

	if bytes, err = cs.client.POST("/synthetic/monitors", config, 200); err != nil {
		return nil, err
	}
	stub := struct {
		ID string `json:"entityId"`
	}{}
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	successfulAttempts := 0
	for successfulAttempts < 5 {
		attempts := 0
		for attempts < 50 {
			if _, err := cs.GetHTTP(stub.ID); err == nil {
				attempts = 50
			} else {
				attempts++
				time.Sleep(2 * time.Second)
			}
		}
		successfulAttempts++
	}

	return &stub.ID, nil
}

// UpdateBrowser TODO: documentation
func (cs *ServiceClient) UpdateBrowser(config *BrowserSyntheticMonitorUpdate) error {
	if len(opt.String(config.ID)) == 0 {
		return errors.New("the configuration doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/synthetic/monitors/%s", opt.String(config.ID)), config, 204); err != nil {
		return err
	}
	return nil
}

// UpdateHTTP TODO: documentation
func (cs *ServiceClient) UpdateHTTP(config *HTTPSyntheticMonitorUpdate) error {
	if len(opt.String(config.ID)) == 0 {
		return errors.New("the configuration doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/synthetic/monitors/%s", opt.String(config.ID)), config, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the configuration to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/synthetic/monitors/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// GetBrowser TODO: documentation
func (cs *ServiceClient) GetBrowser(id string) (*BrowserSyntheticMonitorUpdate, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the configuration to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/synthetic/monitors/%s", id), 200); err != nil {
		return nil, err
	}
	var autoTag BrowserSyntheticMonitorUpdate
	if err = json.Unmarshal(bytes, &autoTag); err != nil {
		return nil, err
	}
	return &autoTag, nil
}

// GetHTTP TODO: documentation
func (cs *ServiceClient) GetHTTP(id string) (*HTTPSyntheticMonitorUpdate, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the configuration to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/synthetic/monitors/%s", id), 200); err != nil {
		return nil, err
	}
	var autoTag HTTPSyntheticMonitorUpdate
	if err = json.Unmarshal(bytes, &autoTag); err != nil {
		return nil, err
	}
	return &autoTag, nil
}

// ListBrowser TODO: documentation
func (cs *ServiceClient) ListBrowser() (*Monitors, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/synthetic/monitors?type=BROWSER", 200); err != nil {
		return nil, err
	}
	var monitors Monitors
	if err = json.Unmarshal(bytes, &monitors); err != nil {
		return nil, err
	}
	return &monitors, nil
}

// ListHTTP TODO: documentation
func (cs *ServiceClient) ListHTTP() (*Monitors, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/synthetic/monitors?type=HTTP", 200); err != nil {
		return nil, err
	}
	var monitors Monitors
	if err = json.Unmarshal(bytes, &monitors); err != nil {
		return nil, err
	}
	return &monitors, nil
}
