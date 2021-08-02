package diskevents

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
func (cs *Service) Create(config *AnomalyDetection) (*api.EntityRef, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.POST("/anomalyDetection/diskEvents", config, 201); err != nil {
		return nil, err
	}
	var stub api.EntityRef
	if err = json.Unmarshal(bytes, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// Update TODO: documentation
func (cs *Service) Update(config *AnomalyDetection) error {
	if _, err := cs.client.PUT(fmt.Sprintf("/anomalyDetection/diskEvents/%s", *config.ID), config, 204); err != nil {
		return err
	}
	return nil
}

// Validate TODO: documentation
func (cs *Service) Validate(config *AnomalyDetection) error {
	if config.ID != nil {
		if _, err := cs.client.PUT(fmt.Sprintf("/anomalyDetection/diskEvents/%s/validator", *config.ID), config, 204); err != nil {
			return err
		}
	} else {
		if _, err := cs.client.PUT("/anomalyDetection/diskEvents/validator", config, 204); err != nil {
			return err
		}
	}
	return nil
}

// Delete TODO: documentation
func (cs *Service) Delete(id string) error {
	if _, err := cs.client.DELETE(fmt.Sprintf("/anomalyDetection/diskEvents/%s", id), 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *Service) Get(id string) (*AnomalyDetection, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/anomalyDetection/diskEvents/%s", id), 200); err != nil {
		return nil, err
	}
	var config AnomalyDetection
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// List TODO: documentation
func (cs *Service) List() (*api.EntityRefs, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/anomalyDetection/diskEvents", 200); err != nil {
		return nil, err
	}
	var stubList api.EntityRefs
	if err = json.Unmarshal(bytes, &stubList); err != nil {
		return nil, err
	}
	return &stubList, nil
}
