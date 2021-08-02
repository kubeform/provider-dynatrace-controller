package traffic

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// DropDetection The configuration of traffic drops detection.
type DropDetection struct {
	Enabled            bool   `json:"enabled"`                      // The detection is enabled (`true`) or disabled (`false`).
	TrafficDropPercent *int32 `json:"trafficDropPercent,omitempty"` // Alert if the observed traffic is less than *X* % of the expected value.
}

func (me *DropDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "The detection is enabled (`true`) or disabled (`false`)",
		},
		"percent": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Alert if the observed traffic is less than *X* % of the expected value",
		},
	}
}

func (me *DropDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["enabled"] = me.Enabled
	if me.TrafficDropPercent != nil {
		result["percent"] = int(*me.TrafficDropPercent)
	}
	return result, nil
}

func (me *DropDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("enabled"); ok {
		me.Enabled = value.(bool)
	}
	if value, ok := decoder.GetOk("percent"); ok {
		me.TrafficDropPercent = opt.NewInt32(int32(value.(int)))
	}

	return nil
}
