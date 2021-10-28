package sharing

import (
	"encoding/json"
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
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &ServiceClient{client: client}
}

// Create TODO: documentation
func (cs *ServiceClient) Create(settings *DashboardSharing) (string, error) {
	if err := cs.Update(settings); err != nil {
		return "", err
	}
	// return settings.DashboardID + "-sharing", nil
	return settings.DashboardID, nil
}

// Update TODO: documentation
func (cs *ServiceClient) Update(settings *DashboardSharing) error {
	_, err := cs.client.PUT(fmt.Sprintf("/dashboards/%s/shareSettings", settings.DashboardID), settings, 201)
	if err != nil && strings.HasPrefix(err.Error(), "No Content (PUT)") {
		return nil
	}
	return err
}

// Delete TODO: documentation
func (cs *ServiceClient) Delete(id string) error {
	settings := DashboardSharing{
		DashboardID: id,
		Enabled:     false,
		Preset:      false,
		Permissions: []*SharePermission{
			{
				Type:       PermissionTypes.All,
				Permission: Permissions.View,
			},
		},
		PublicAccess: &AnonymousAccess{
			ManagementZoneIDs: []string{},
			URLs:              map[string]string{},
		},
	}
	return cs.Update(&settings)
}

// Get TODO: documentation
func (cs *ServiceClient) Get(id string) (*DashboardSharing, error) {
	// id = strings.TrimSuffix(id, "-sharing")

	var err error
	var bytes []byte

	if bytes, err = cs.client.GET(fmt.Sprintf("/dashboards/%s/shareSettings", id), 200); err != nil {
		return nil, err
	}
	var settings DashboardSharing
	if err = json.Unmarshal(bytes, &settings); err != nil {
		return nil, err
	}
	return &settings, nil
}
