package mobile

// AppConfig represents configuration of an existing mobile or custom application
type AppConfig struct {
	Name                             string              `json:"name"`                             // The name of the application
	ApplicationType                  ApplicationType     `json:"applicationType"`                  // The type of the application
	ApplicationID                    *string             `json:"applicationId,omitempty"`          // The UUID of the application.\n\nIt is used only by OneAgent to send data to Dynatrace
	CostControlUserSessionPercentage *int32              `json:"costControlUserSessionPercentage"` // The percentage of user sessions to be analyzed
	ApdexSettings                    *MobileCustomApdex  `json:"apdexSettings"`                    // Apdex configuration of a mobile or custom application. \n\nA duration less than the **tolerable** threshold is considered satisfied
	OptInModeEnabled                 *bool               `json:"optInModeEnabled"`                 // The opt-in mode is enabled (`true`) or disabled (`false`).\n\nThis value is only applicable to mobile and not to custom apps
	SessionReplayEnabled             *bool               `json:"sessionReplayEnabled"`             // The session replay is enabled (`true`) or disabled (`false`).\nThis value is only applicable to mobile and not to custom apps
	SessionReplayOnCrashEnabled      *bool               `json:"sessionReplayOnCrashEnabled"`      // The session replay on crash is enabled (`true`) or disabled (`false`). \n\nEnabling requires both **sessionReplayEnabled** and **optInModeEnabled** values set to `true`.\nAlso, this value is only applicable to mobile and not to custom apps
	BeaconEndpointType               *BeaconEndpointType `json:"beaconEndpointType"`               // The type of the beacon endpoint
	BeaconEndpointUrl                *string             `json:"beaconEndpointUrl"`                // The URL of the beacon endpoint.\n\nOnly applicable when the **beaconEndpointType** is set to `ENVIRONMENT_ACTIVE_GATE` or `INSTRUMENTED_WEB_SERVER`
}
