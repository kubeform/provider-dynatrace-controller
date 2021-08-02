package traffic

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// SpikeDetection The configuration of traffic spikes detection.
type SpikeDetection struct {
	Enabled             bool   `json:"enabled"`                       // The detection is enabled (`true`) or disabled (`false`).
	TrafficSpikePercent *int32 `json:"trafficSpikePercent,omitempty"` // Alert if the observed traffic is more than *X* % of the expected value.
}

func (me *SpikeDetection) Schema() map[string]*hcl.Schema {
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

func (me *SpikeDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["enabled"] = me.Enabled
	if me.TrafficSpikePercent != nil {
		result["percent"] = int(*me.TrafficSpikePercent)
	}
	return result, nil
}

func (me *SpikeDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("enabled"); ok {
		me.Enabled = value.(bool)
	}
	if value, ok := decoder.GetOk("percent"); ok {
		me.TrafficSpikePercent = opt.NewInt32(int32(value.(int)))
	}

	return nil
}
