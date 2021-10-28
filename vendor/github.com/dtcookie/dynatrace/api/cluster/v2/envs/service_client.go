package envs

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
	config := rest.Config{Insecure: true}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(environment *Environment) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if len(opt.String(environment.ID)) > 0 {
		return nil, errors.New("you MUST NOT provide an ID within the Dashboard payload upon creation")
	}

	if bytes, err = cs.client.POST("/environments", environment, 201); err != nil {
		retry := false
		switch rerr := err.(type) {
		case *rest.Error:
			if len(rerr.ConstraintViolations) > 0 {
				for _, violation := range rerr.ConstraintViolations {
					if violation.Message == "Cannot set Synthetic monitors quota as Synthetic monitors are not allowed for this license." {
						environment.Quotas.Synthetic = nil
						retry = true
					} else if violation.Message == "Cannot set DEM units quota as DEM units are not allowed for this license." {
						environment.Quotas.DEMUnits = nil
						retry = true
					} else if violation.Message == "Cannot set Log Monitoring quota as Log Monitoring is not allowed for this license." {
						environment.Quotas.LogMonitoring = nil
						retry = true
					}
				}
			}
			if retry {
				return cs.Create(environment)
			}
		default:
			return nil, err
		}
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(environment *Environment) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/environments/%s", opt.String(environment.ID)), environment, 204); err != nil {
		retry := false
		switch rerr := err.(type) {
		case *rest.Error:
			if len(rerr.ConstraintViolations) > 0 {
				for _, violation := range rerr.ConstraintViolations {
					if violation.Message == "Cannot set Synthetic monitors quota as Synthetic monitors are not allowed for this license." {
						environment.Quotas.Synthetic = nil
						retry = true
					} else if violation.Message == "Cannot set DEM units quota as DEM units are not allowed for this license." {
						environment.Quotas.DEMUnits = nil
						retry = true
					} else if violation.Message == "Cannot set Log Monitoring quota as Log Monitoring is not allowed for this license." {
						environment.Quotas.LogMonitoring = nil
						retry = true
					}
				}
			}
			if retry {
				return cs.Update(environment)
			}
		default:
			return err
		}
		return err
	}

	return nil
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("empty ID provided for the Dashboard to delete")
	}
	env, err := cs.Get(id)
	if err != nil {
		return err
	}
	if env.State == States.Enabled {
		env.State = States.Disabled
		if err = cs.Update(env); err != nil {
			return err
		}
	}
	if _, err := cs.client.DELETE(fmt.Sprintf("/environments/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*Environment, error) {
	if len(id) == 0 {
		return nil, errors.New("empty ID provided for the Dashboard to fetch")
	}

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/environments/%s?includeConsumptionInfo=true&includeStorageInfo=true", id), 200); err != nil {
		return nil, err
	}
	var environment Environment
	if err = json.Unmarshal(bytes, &environment); err != nil {
		return nil, err
	}
	return &environment, nil
}

// ListAll TODO: documentation
func (cs *ServiceClient) ListAll() (*EnvironmentList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/environments", 200); err != nil {
		return nil, err
	}
	var environmentList EnvironmentList
	if err = json.Unmarshal(bytes, &environmentList); err != nil {
		return nil, err
	}
	return &environmentList, nil
}

type EnvironmentList struct {
	Environments []*Environment `json:"environments"`
}
