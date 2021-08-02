package applications

import (
	"encoding/json"

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

// Update TODO: documentation
func (cs *Service) Update(config *AnomalyDetection) error {
	if _, err := cs.client.PUT("/anomalyDetection/applications", config, 204); err != nil {
		return err
	}
	return nil
}

// Validate TODO: documentation
func (cs *Service) Validate(config *AnomalyDetection) error {
	if _, err := cs.client.POST("/anomalyDetection/applications/validator", config, 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *Service) Get() (*AnomalyDetection, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/anomalyDetection/applications", 200); err != nil {
		return nil, err
	}
	var response AnomalyDetection
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
