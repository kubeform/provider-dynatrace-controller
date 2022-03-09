package mobile

import "github.com/dtcookie/hcl"

// NewAppConfig represents configuration of a mobile or custom application to be created
type NewAppConfig struct {
	ID                               string                         `json:"-"`
	Name                             string                         `json:"name"`                                       // The name of the application
	ApplicationType                  *ApplicationType               `json:"applicationType,omitempty"`                  // The type of the application
	ApplicationID                    *string                        `json:"applicationId,omitempty"`                    // The UUID of the application.\n\nIt is used only by OneAgent to send data to Dynatrace
	CostControlUserSessionPercentage *int32                         `json:"costControlUserSessionPercentage,omitempty"` // The percentage of user sessions to be analyzed
	ApdexSettings                    *MobileCustomApdex             `json:"apdexSettings,omitempty"`
	OptInModeEnabled                 bool                           `json:"optInModeEnabled"`            // The opt-in mode is enabled (`true`) or disabled (`false`).\n\nThis value is only applicable to mobile and not to custom apps
	SessionReplayEnabled             bool                           `json:"sessionReplayEnabled"`        // The session replay is enabled (`true`) or disabled (`false`).\nThis value is only applicable to mobile and not to custom apps
	SessionReplayOnCrashEnabled      bool                           `json:"sessionReplayOnCrashEnabled"` // The session replay on crash is enabled (`true`) or disabled (`false`). \n\nEnabling requires both **sessionReplayEnabled** and **optInModeEnabled** values set to `true`.\nAlso, this value is only applicable to mobile and not to custom apps
	BeaconEndpointType               BeaconEndpointType             `json:"beaconEndpointType"`          // The type of the beacon endpoint
	BeaconEndpointUrl                *string                        `json:"beaconEndpointUrl,omitempty"` // The URL of the beacon endpoint.\n\nOnly applicable when the **beaconEndpointType** is set to `ENVIRONMENT_ACTIVE_GATE` or `INSTRUMENTED_WEB_SERVER`
	KeyUserActions                   []string                       `json:"-"`
	Properties                       UserActionAndSessionProperties `json:"-"`
}

func (me *NewAppConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("name", &me.Name); err != nil {
		return err
	}
	me.ApplicationType = &ApplicationTypes.MobileApplication
	if err := decoder.Decode("application_id", &me.ApplicationID); err != nil {
		return err
	}
	if err := decoder.Decode("user_session_percentage", &me.CostControlUserSessionPercentage); err != nil {
		return err
	}
	if err := decoder.Decode("opt_in_mode", &me.OptInModeEnabled); err != nil {
		return err
	}
	if err := decoder.Decode("session_replay", &me.SessionReplayEnabled); err != nil {
		return err
	}
	if err := decoder.Decode("session_replay_on_crash", &me.SessionReplayOnCrashEnabled); err != nil {
		return err
	}
	if err := decoder.Decode("beacon_endpoint_type", &me.BeaconEndpointType); err != nil {
		return err
	}
	if err := decoder.Decode("beacon_endpoint_url", &me.BeaconEndpointUrl); err != nil {
		return err
	}
	if err := decoder.Decode("apdex", &me.ApdexSettings); err != nil {
		return err
	}
	if err := decoder.Decode("key_user_actions", &me.KeyUserActions); err != nil {
		return err
	}
	if err := decoder.Decode("properties", &me.Properties); err != nil {
		return err
	}
	return nil
}

func (me *NewAppConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if err := properties.Encode("name", me.Name); err != nil {
		return nil, err
	}
	if me.ApplicationID != nil {
		if err := properties.Encode("application_id", me.ApplicationID); err != nil {
			return nil, err
		}
	}
	if me.CostControlUserSessionPercentage != nil {
		if err := properties.Encode("user_session_percentage", me.CostControlUserSessionPercentage); err != nil {
			return nil, err
		}
	}
	if me.OptInModeEnabled {
		if err := properties.Encode("opt_in_mode", me.OptInModeEnabled); err != nil {
			return nil, err
		}
	}
	if me.SessionReplayEnabled {
		if err := properties.Encode("session_replay", me.SessionReplayEnabled); err != nil {
			return nil, err
		}
	}
	if me.SessionReplayOnCrashEnabled {
		if err := properties.Encode("session_replay_on_crash", me.SessionReplayOnCrashEnabled); err != nil {
			return nil, err
		}
	}
	if err := properties.Encode("beacon_endpoint_type", me.BeaconEndpointType); err != nil {
		return nil, err
	}
	if me.BeaconEndpointUrl != nil {
		if err := properties.Encode("beacon_endpoint_url", me.BeaconEndpointUrl); err != nil {
			return nil, err
		}
	}
	if me.ApdexSettings != nil {
		if err := properties.Encode("apdex", me.ApdexSettings); err != nil {
			return nil, err
		}
	}
	if len(me.KeyUserActions) > 0 {
		if err := properties.Encode("key_user_actions", me.KeyUserActions); err != nil {
			return nil, err
		}
	}
	if len(me.Properties) > 0 {
		if err := properties.Encode("properties", me.Properties); err != nil {
			return nil, err
		}
	}
	return properties, nil
}

func (me *NewAppConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the application",
			Required:    true,
		},
		"application_id": {
			Type:        hcl.TypeString,
			Description: "The UUID of the application.\n\nIt is used only by OneAgent to send data to Dynatrace. If not specified it will get generated.",
			Optional:    true,
		},
		"user_session_percentage": {
			Type:        hcl.TypeInt,
			Description: "The percentage of user sessions to be analyzed",
			Optional:    true,
		},
		"opt_in_mode": {
			Type:        hcl.TypeBool,
			Description: "The opt-in mode is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"session_replay": {
			Type:        hcl.TypeBool,
			Description: "The session replay is enabled (`true`) or disabled (`false`).",
			Optional:    true,
		},
		"session_replay_on_crash": {
			Type:        hcl.TypeBool,
			Description: "The session replay on crash is enabled (`true`) or disabled (`false`). \n\nEnabling requires both **sessionReplayEnabled** and **optInModeEnabled** values set to `true`.",
			Optional:    true,
		},
		"beacon_endpoint_type": {
			Type:        hcl.TypeString,
			Description: "The type of the beacon endpoint. Possible values are `CLUSTER_ACTIVE_GATE`, `ENVIRONMENT_ACTIVE_GATE` and `INSTRUMENTED_WEB_SERVER`.",
			Required:    true,
		},
		"beacon_endpoint_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the beacon endpoint.\n\nOnly applicable when the **beacon_endpoint_type** is set to `ENVIRONMENT_ACTIVE_GATE` or `INSTRUMENTED_WEB_SERVER`",
			Optional:    true,
		},
		"key_user_actions": {
			Type:        hcl.TypeSet,
			Description: "User Action names to be flagged as Key User Actions",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Optional:    true,
		},
		"apdex": {
			Type:        hcl.TypeList,
			Description: "Apdex configuration of a mobile application. A duration less than the **tolerable** threshold is considered satisfied",
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(MobileCustomApdex).Schema()},
		},
		"properties": {
			Type:        hcl.TypeList,
			Description: "User Action and Session Properties",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionAndSessionProperties).Schema()},
		},
	}
}

type ApplicationType string

var ApplicationTypes = struct {
	CustomApplication ApplicationType
	MobileApplication ApplicationType
}{
	CustomApplication: ApplicationType("CUSTOM_APPLICATION"),
	MobileApplication: ApplicationType("MOBILE_APPLICATION"),
}

type BeaconEndpointType string

var BeaconEndpointTypes = struct {
	ClusterActiveGate     BeaconEndpointType
	EnvironmentActiveGate BeaconEndpointType
	InstrumentedWebServer BeaconEndpointType
}{
	ClusterActiveGate:     BeaconEndpointType("CLUSTER_ACTIVE_GATE"),
	EnvironmentActiveGate: BeaconEndpointType("ENVIRONMENT_ACTIVE_GATE"),
	InstrumentedWebServer: BeaconEndpointType("INSTRUMENTED_WEB_SERVER"),
}
