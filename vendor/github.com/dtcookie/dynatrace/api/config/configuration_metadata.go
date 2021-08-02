package api

import "encoding/json"

// ConfigurationMetadata is Metadata useful for debugging
type ConfigurationMetadata struct {
	CurrentConfigurationVersions []string                   `json:"currentConfigurationVersions,omitempty"` // A Sorted list of string version numbers of the configuration
	ClusterVersion               string                     `json:"clusterVersion,omitempty"`               // Dynatrace server version
	ConfigurationVersions        []int64                    `json:"configurationVersions,omitempty"`        // A Sorted list of the version numbers of the configuration
	Unknowns                     map[string]json.RawMessage `json:"-"`
}

func (cm *ConfigurationMetadata) UnmarshalJSON(data []byte) error {
	cm.Unknowns = map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &cm.Unknowns); err != nil {
		return err
	}
	if v, found := cm.Unknowns["currentConfigurationVersions"]; found {
		if err := json.Unmarshal(v, &cm.CurrentConfigurationVersions); err != nil {
			return err
		}
		delete(cm.Unknowns, "currentConfigurationVersions")
	}
	if v, found := cm.Unknowns["clusterVersion"]; found {
		if err := json.Unmarshal(v, &cm.ClusterVersion); err != nil {
			return err
		}
		delete(cm.Unknowns, "clusterVersion")
	}
	if v, found := cm.Unknowns["configurationVersions"]; found {
		if err := json.Unmarshal(v, &cm.ConfigurationVersions); err != nil {
			return err
		}
		delete(cm.Unknowns, "configurationVersions")
	}
	return nil
}

func (cm *ConfigurationMetadata) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if cm.Unknowns != nil {
		for k, v := range cm.Unknowns {
			m[k] = v
		}
	}
	if cm.CurrentConfigurationVersions != nil {
		rawMessage, err := json.Marshal(cm.CurrentConfigurationVersions)
		if err != nil {
			return nil, err
		}
		m["currentConfigurationVersions"] = rawMessage
	}
	if cm.ClusterVersion != "" {
		rawMessage, err := json.Marshal(cm.ClusterVersion)
		if err != nil {
			return nil, err
		}
		m["clusterVersion"] = rawMessage
	}
	if cm.ConfigurationVersions != nil {
		rawMessage, err := json.Marshal(cm.ConfigurationVersions)
		if err != nil {
			return nil, err
		}
		m["configurationVersions"] = rawMessage
	}
	return json.Marshal(m)
}

// String provides a string representation in JSON format for debugging purposes
func (cm ConfigurationMetadata) String() string {
	return MarshalJSON(&cm)
}
