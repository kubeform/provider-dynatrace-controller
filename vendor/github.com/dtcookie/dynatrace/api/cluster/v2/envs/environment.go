package envs

import (
	"github.com/dtcookie/dynatrace/api/cluster/v2/envs/quota"
	"github.com/dtcookie/dynatrace/api/cluster/v2/envs/storage"
	"github.com/dtcookie/hcl"
)

// Environment representas basic configuration for an environment
type Environment struct {
	ID      *string           `json:"id,omitempty"`   // The ID of the environment
	Name    string            `json:"name"`           // The display name of the environment
	Trial   *bool             `json:"trial"`          // Specifies whether the environment is a trial environment or a non-trial environment. Creating a trial environment is only possible if your license allows that. The default value is false (non-trial)
	State   State             `json:"state"`          // Indicates whether the environment is enabled or disabled. The default value is ENABLED
	Tags    []string          `json:"tags,omitempty"` // A set of tags that are assigned to this environment. Every tag can have a maximum length of 100 characters
	Quotas  *quota.Settings   `json:"quotas"`         // Environment level consumption and quotas information. Only returned if includeConsumptionInfo or includeUncachedConsumptionInfo param is true
	Storage *storage.Settings `json:"storage"`        // Environment level storage usage and limit information. Not returned if includeStorageInfo param is not true. If skipped when editing via PUT method then already set limits will remain
}

func (me *Environment) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The display name of the environment",
		},
		"trial": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "Specifies whether the environment is a trial environment or a non-trial environment. Creating a trial environment is only possible if your license allows that. The default value is false (non-trial)",
		},
		"state": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Indicates whether the environment is enabled or disabled. Possible values are `ENABLED` and `DISABLED`. The default value is ENABLED",
		},
		"tags": {
			Type:        hcl.TypeSet,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Optional:    true,
			Description: "A set of tags that are assigned to this environment. Every tag can have a maximum length of 100 characters",
		},
		"quotas": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(quota.Settings).Schema()},
			Description: "Environment level consumption and quotas information",
		},
		"storage": {
			Type:        hcl.TypeList,
			Required:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(storage.Settings).Schema()},
			Description: "Environment level storage usage and limit information",
		},
	}
}

func (me *Environment) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if _, err := properties.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"trial":   me.Trial,
		"state":   me.State,
		"tags":    me.Tags,
		"quotas":  me.Quotas,
		"storage": me.Storage,
	}); err != nil {
		return nil, err
	}
	if me.Trial == nil || !*me.Trial {
		delete(properties, "trial")
	}
	if me.Quotas.IsEmpty() {
		delete(properties, "quotas")
	}
	return properties, nil
}

func (me *Environment) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"trial":   &me.Trial,
		"state":   &me.State,
		"tags":    &me.Tags,
		"quotas":  &me.Quotas,
		"storage": &me.Storage,
	}); err != nil {
		return err
	}
	if me.Quotas == nil {
		me.Quotas = &quota.Settings{
			HostUnits:     &quota.HostUnits{MaxLimit: nil},
			DEMUnits:      &quota.DEMUnits{MonthlyLimit: nil, AnnualLimit: nil},
			UserSessions:  &quota.UserSessions{TotalMonthlyLimit: nil, TotalAnnualLimit: nil},
			Synthetic:     &quota.Synthetic{MonthlyLimit: nil, AnnualLimit: nil},
			DDUs:          &quota.DavisDataUnits{MonthlyLimit: nil, AnnualLimit: nil},
			LogMonitoring: &quota.LogMonitoring{MonthlyLimit: nil, AnnualLimit: nil},
		}
	}
	return nil
}

type State string

var States = struct {
	Enabled  State
	Disabled State
}{
	Enabled:  State("ENABLED"),
	Disabled: State("DISABLED"),
}
