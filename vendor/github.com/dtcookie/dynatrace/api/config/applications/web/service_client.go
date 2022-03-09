package web

import (
	"encoding/json"
	"errors"
	"fmt"
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

// Create TODO: documentation
func (cs *ServiceClient) Create(applicationConfig *ApplicationConfig) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/applications/web", applicationConfig, 201); err != nil {
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
			if _, err = cs.client.POST(fmt.Sprintf("/applications/web/%s/keyUserActions", stub.ID), &keyUserAction, 201); err != nil {
				return nil, err
			}
		}
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(applicationConfig *ApplicationConfig) error {
	if applicationConfig.ID == nil {
		return errors.New("the config doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/applications/web/%s", *applicationConfig.ID), applicationConfig, 204); err != nil {
		return err
	}
	var err error
	var bytes []byte
	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/web/%s/keyUserActions", *applicationConfig.ID), 200); err != nil {
		return err
	}
	remoteKeyUserActions := map[string]*KeyUserAction{}
	var kual KeyUserActionList
	if err = json.Unmarshal(bytes, &kual); err != nil {
		return err
	}
	for _, item := range kual.KeyUserActions {
		remoteKeyUserActions[item.String()] = item
	}
	keyUserActionsToDelete := map[string]*KeyUserAction{}
	for _, keyUserAction := range remoteKeyUserActions {
		keyUserActionsToDelete[keyUserAction.String()] = keyUserAction
	}
	keyUserActionsToAdd := []*KeyUserAction{}
	if len(applicationConfig.KeyUserActions) > 0 {
		for _, keyUserAction := range applicationConfig.KeyUserActions {
			delete(keyUserActionsToDelete, keyUserAction.String())
			if _, found := remoteKeyUserActions[keyUserAction.String()]; !found {
				keyUserActionsToAdd = append(keyUserActionsToAdd, keyUserAction)
			}
		}
	}
	for _, keyUserAction := range keyUserActionsToDelete {
		if _, err := cs.client.DELETE(fmt.Sprintf("/applications/web/%s/keyUserActions/%s", *applicationConfig.ID, *keyUserAction.ID), 204); err != nil {
			return err
		}
	}
	for _, keyUserAction := range keyUserActionsToAdd {
		tmp := struct {
			Name   string            `json:"name"`
			Type   KeyUserActionType `json:"actionType"`
			Domain *string           `json:"domain,omitempty"`
		}{
			keyUserAction.Name,
			keyUserAction.Type,
			keyUserAction.Domain,
		}
		if _, err = cs.client.POST(fmt.Sprintf("/applications/web/%s/keyUserActions", *applicationConfig.ID), tmp, 201); err != nil {
			return err
		}
	}
	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the application config to delete")
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/applications/web/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*ApplicationConfig, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the config to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/web/%s", id), 200); err != nil {
		return nil, err
	}
	var applicationConfig ApplicationConfig
	if err = json.Unmarshal(bytes, &applicationConfig); err != nil {
		return nil, err
	}
	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/web/%s/keyUserActions", id), 200); err != nil {
		return nil, err
	}
	var kual KeyUserActionList
	if err = json.Unmarshal(bytes, &kual); err != nil {
		return nil, err
	}
	actions := []*KeyUserAction{}
	for _, action := range kual.KeyUserActions {
		actions = append(actions, action)
	}
	if len(actions) > 0 {
		applicationConfig.KeyUserActions = actions
	}

	return &applicationConfig, nil
}

func (cs *ServiceClient) GetAppDataPrivacy(id string) (*ApplicationDataPrivacy, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the config to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/web/%s/dataPrivacy", id), 200); err != nil {
		return nil, err
	}
	var config ApplicationDataPrivacy
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (cs *ServiceClient) GetErrorRules(id string) (*ApplicationErrorRules, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the config to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/applications/web/%s/errorRules", id), 200); err != nil {
		return nil, err
	}
	var config ApplicationErrorRules
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	config.WebApplicationID = id

	return &config, nil
}

func (cs *ServiceClient) StoreAppDataPrivacy(config *ApplicationDataPrivacy) error {
	if config.WebApplicationID == nil {
		return errors.New("the config doesn't contain an ID")
	}
	copy := *config
	copy.WebApplicationID = nil
	if _, err := cs.client.PUT(fmt.Sprintf("/applications/web/%s/dataPrivacy", *config.WebApplicationID), &copy, 204); err != nil {
		return err
	}
	return nil
}

func (cs *ServiceClient) StoreErrorRules(config *ApplicationErrorRules) error {
	if len(config.WebApplicationID) == 0 {
		return errors.New("the config doesn't contain an ID")
	}
	if _, err := cs.client.PUT(fmt.Sprintf("/applications/web/%s/errorRules", config.WebApplicationID), config, 204); err != nil {
		return err
	}
	return nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) List() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/applications/web", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
