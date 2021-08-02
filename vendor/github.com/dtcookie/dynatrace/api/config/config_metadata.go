package api

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// ConfigMetadata Metadata useful for debugging
type ConfigMetadata struct {
	ClusterVersion               *string  `json:"clusterVersion,omitempty"`               // Dynatrace server version.
	ConfigurationVersions        []int64  `json:"configurationVersions,omitempty"`        // A Sorted list of the version numbers of the configuration.
	CurrentConfigurationVersions []string `json:"currentConfigurationVersions,omitempty"` // A Sorted list of string version numbers of the configuration.
}

func (me *ConfigMetadata) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cluster_version": {
			Type:        hcl.TypeString,
			Description: "Dynatrace server version",
			Optional:    true,
		},
		"configuration_versions": {
			Type:        hcl.TypeList,
			Description: "A Sorted list of the version numbers of the configuration",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeInt},
		},
		"current_configuration_versions": {
			Type:        hcl.TypeList,
			Description: "A Sorted list of the version numbers of the configuration",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *ConfigMetadata) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.ClusterVersion != nil && len(*me.ClusterVersion) > 0 {
		result["cluster_version"] = *me.ClusterVersion
	}
	if len(me.ConfigurationVersions) > 0 {
		result["configuration_versions"] = me.ConfigurationVersions
	}
	if len(me.CurrentConfigurationVersions) > 0 {
		result["current_configuration_versions"] = me.CurrentConfigurationVersions
	}
	return result, nil
}

func (me *ConfigMetadata) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("cluster_version"); ok {
		me.ClusterVersion = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("configuration_versions.#"); ok {
		me.ConfigurationVersions = []int64{}
		if entries, ok := decoder.GetOk("configuration_versions"); ok {
			for _, entry := range entries.([]interface{}) {
				me.ConfigurationVersions = append(me.ConfigurationVersions, int64(entry.(int)))
			}
		}
	}
	if _, ok := decoder.GetOk("current_configuration_versions.#"); ok {
		me.CurrentConfigurationVersions = []string{}
		if entries, ok := decoder.GetOk("current_configuration_versions"); ok {
			for _, entry := range entries.([]interface{}) {
				me.CurrentConfigurationVersions = append(me.CurrentConfigurationVersions, entry.(string))
			}
		}
	}
	return nil
}
