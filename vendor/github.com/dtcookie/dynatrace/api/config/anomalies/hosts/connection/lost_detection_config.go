package connection

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// LostDetectionConfig Configuration of lost connection detection.
type LostDetectionConfig struct {
	EnabledOnGracefulShutdowns bool `json:"enabledOnGracefulShutdowns"` // Alert (`true`) on graceful host shutdowns.
	Enabled                    bool `json:"enabled"`                    // The detection is enabled (`true`) or disabled (`false`).
}

func (me *LostDetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "The detection is enabled (`true`) or disabled (`false`)",
		},
		"enabled_on_graceful_shutdowns": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "Alert (`true`) on graceful host shutdowns",
		},
	}
}

func (me *LostDetectionConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["enabled"] = me.Enabled
	result["enabled_on_graceful_shutdowns"] = me.EnabledOnGracefulShutdowns
	return result, nil
}

func (me *LostDetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	adapter := hcl.Adapt(decoder)
	me.Enabled = opt.Bool(adapter.GetBool("enabled"))
	me.EnabledOnGracefulShutdowns = opt.Bool(adapter.GetBool("enabled_on_graceful_shutdowns"))
	return nil
}
