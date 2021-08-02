package alerting

import (
	"encoding/json"
	"fmt"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// Service TODO: documentation
type Service struct {
	client *rest.Client
}

// NewService TODO: documentation
// "https://#######.live.dynatrace.com/api/config/v1", "###########"
func NewService(baseURL string, token string) *Service {
	rest.Verbose = false
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &Service{client: client}
}

// Create TODO: documentation
func (cs *Service) Create(alertingProfile *Profile) (*api.EntityShortRepresentation, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/alertingProfiles", alertingProfile, 201); err != nil {
		return nil, err
	}
	var stub api.EntityShortRepresentation
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *Service) Update(alertingProfile *Profile) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/alertingProfiles/%s", *alertingProfile.ID), alertingProfile, 204); err != nil {
		return err
	}
	return nil
}

// Delete TODO: documentation
func (cs *Service) Delete(id string) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/alertingProfiles/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *Service) Get(id string) (*Profile, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/alertingProfiles/%s", id), 200); err != nil {
		return nil, err
	}
	var alertingProfile Profile
	if err = json.Unmarshal(bytes, &alertingProfile); err != nil {
		return nil, err
	}
	return &alertingProfile, nil
}

// List TODO: documentation
func (cs *Service) List() (*api.StubList, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/alertingProfiles", 200); err != nil {
		return nil, err
	}
	var stubList api.StubList
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
