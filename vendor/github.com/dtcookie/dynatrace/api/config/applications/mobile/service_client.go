package mobile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

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

type nothing struct{}

func (n *nothing) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

// Create TODO: documentation
func (cs *ServiceClient) Create(applicationConfig *NewAppConfig) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/applications/mobile", applicationConfig, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	for i := 0; i < 40; i++ {
		if _, err = cs.Get(stub.ID); err == nil {
			break
		}
		time.Sleep(time.Second * 3)
	}

	if len(applicationConfig.KeyUserActions) > 0 {
		for _, keyUserAction := range applicationConfig.KeyUserActions {
			if _, err = cs.client.POST(fmt.Sprintf("/applications/mobile/%s/keyUserActions/%s", stub.ID, url.PathEscape(keyUserAction)), new(nothing), 201); err != nil {
				return nil, err
			}
		}
	}
	if len(applicationConfig.Properties) > 0 {
		for _, property := range applicationConfig.Properties {
			if _, err = cs.client.POST(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties", stub.ID), property, 201); err != nil {
				return nil, err
			}
		}
	}
	return &stub, nil
}

type keyUserActionsResponse struct {
	KeyUserActions []*keyUserActionsResponseItem `json:"keyUserActions"`
}

type keyUserActionsResponseItem struct {
	Name string `json:"name"`
}

// Update TODO: documentation
func (cs *ServiceClient) Update(applicationConfig *NewAppConfig) error {
	if len(applicationConfig.ID) == 0 {
		return errors.New("the config doesn't contain an ID")
	}
	applicationConfig.ApplicationType = nil
	applicationConfig.ApplicationID = nil
	if _, err := cs.client.PUT(fmt.Sprintf("/applications/mobile/%s", applicationConfig.ID), applicationConfig, 204); err != nil {
		return err
	}
	var err error
	var bytes []byte
	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s/keyUserActions", applicationConfig.ID), 200); err != nil {
		return err
	}
	remoteKeyUserActions := map[string]string{}
	var resp keyUserActionsResponse
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return err
	}
	for _, item := range resp.KeyUserActions {
		remoteKeyUserActions[item.Name] = item.Name
	}
	keyUserActionsToDelete := map[string]string{}
	for keyUserAction := range remoteKeyUserActions {
		keyUserActionsToDelete[keyUserAction] = keyUserAction
	}
	keyUserActionsToAdd := []string{}
	if len(applicationConfig.KeyUserActions) > 0 {
		for _, keyUserAction := range applicationConfig.KeyUserActions {
			delete(keyUserActionsToDelete, keyUserAction)
			if _, found := remoteKeyUserActions[keyUserAction]; !found {
				keyUserActionsToAdd = append(keyUserActionsToAdd, keyUserAction)
			}
		}
	}
	for keyUserAction := range keyUserActionsToDelete {
		if _, err := cs.client.DELETE(fmt.Sprintf("/applications/mobile/%s/keyUserActions/%s", applicationConfig.ID, url.PathEscape(keyUserAction)), 204); err != nil {
			return err
		}
	}
	for _, keyUserAction := range keyUserActionsToAdd {
		if _, err = cs.client.POST(fmt.Sprintf("/applications/mobile/%s/keyUserActions/%s", applicationConfig.ID, url.PathEscape(keyUserAction)), new(nothing), 201); err != nil {
			return err
		}
	}
	remoteProperties := map[string]*UserActionAndSessionProperty{}
	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties", applicationConfig.ID), 200); err != nil {
		return err
	}
	var presp userActionsAndSessionPropertiesResponse
	if err = json.Unmarshal(bytes, &presp); err != nil {
		return err
	}
	propKeys := map[string]string{}
	for _, v := range presp.SessionProperties {
		propKeys[v.Key] = v.Key
	}
	for _, v := range presp.UserActionProperties {
		propKeys[v.Key] = v.Key
	}
	for propKey := range propKeys {
		if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties/%s", applicationConfig.ID, url.PathEscape(propKey)), 200); err != nil {
			return err
		}
		var property UserActionAndSessionProperty
		if err = json.Unmarshal(bytes, &property); err != nil {
			return err
		}
		remoteProperties[propKey] = &property
	}
	propsToDelete := map[string]string{}
	for propKey := range remoteProperties {
		propsToDelete[propKey] = propKey
	}
	propsToUpdate := map[string]*UserActionAndSessionProperty{}
	propsToCreate := map[string]*UserActionAndSessionProperty{}
	if len(applicationConfig.Properties) > 0 {
		for _, property := range applicationConfig.Properties {
			propKey := property.Key
			delete(propsToDelete, propKey)
			if _, found := remoteProperties[propKey]; found {
				propsToUpdate[propKey] = property
			} else {
				propsToCreate[propKey] = property
			}
		}
	}
	for propKey := range propsToDelete {
		if _, err := cs.client.DELETE(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties/%s", applicationConfig.ID, url.PathEscape(propKey)), 204); err != nil {
			return err
		}
	}
	for _, property := range propsToCreate {
		if _, err = cs.client.POST(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties", applicationConfig.ID), property, 201); err != nil {
			return err
		}
	}
	for propKey, property := range propsToUpdate {
		if _, err = cs.client.PUT(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties/%s", applicationConfig.ID, url.PathEscape(propKey)), property, 201); err != nil {
			if !strings.Contains(err.Error(), "No Content (PUT)") {
				return err
			}
		}
	}

	return nil
}

type userActionsAndSessionPropertiesResponse struct {
	SessionProperties    []*keyName `json:"sessionProperties"`
	UserActionProperties []*keyName `json:"userActionProperties"`
}

type keyName struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the application config to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/applications/mobile/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*NewAppConfig, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the Dashboard to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s", id), 200); err != nil {
		return nil, err
	}
	var applicationConfig NewAppConfig
	if err = json.Unmarshal(bytes, &applicationConfig); err != nil {
		return nil, err
	}
	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s/keyUserActions", id), 200); err != nil {
		return nil, err
	}
	var resp keyUserActionsResponse
	if err = json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	names := []string{}
	for _, item := range resp.KeyUserActions {
		names = append(names, item.Name)
	}
	if len(names) > 0 {
		applicationConfig.KeyUserActions = names
	}

	remoteProperties := map[string]*UserActionAndSessionProperty{}
	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties", id), 200); err != nil {
		return nil, err
	}
	var presp userActionsAndSessionPropertiesResponse
	if err = json.Unmarshal(bytes, &presp); err != nil {
		return nil, err
	}
	propKeys := map[string]string{}
	for _, v := range presp.SessionProperties {
		propKeys[v.Key] = v.Key
	}
	for _, v := range presp.UserActionProperties {
		propKeys[v.Key] = v.Key
	}
	for propKey := range propKeys {
		if bytes, err = cs.client.GET(fmt.Sprintf("/applications/mobile/%s/userActionAndSessionProperties/%s", id, url.PathEscape(propKey)), 200); err != nil {
			return nil, err
		}
		var property UserActionAndSessionProperty
		if err = json.Unmarshal(bytes, &property); err != nil {
			return nil, err
		}
		remoteProperties[propKey] = &property
	}
	if len(remoteProperties) > 0 {
		applicationConfig.Properties = UserActionAndSessionProperties{}
		for _, property := range remoteProperties {
			applicationConfig.Properties = append(applicationConfig.Properties, property)
		}
	}

	return &applicationConfig, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) List() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/applications/mobile", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
