package load

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// SpikeDetection The configuration of load spikes detection.
type SpikeDetection struct {
	Enabled          bool                       `json:"enabled"`                                     // The detection is enabled (`true`) or disabled (`false`).
	LoadSpikePercent *int32                     `json:"loadSpikePercent,omitempty"`                  // Alert if the observed load is more than *X* % of the expected value.
	AbnormalMinutes  *int32                     `json:"minAbnormalStateDurationInMinutes,omitempty"` // Alert if the service stays in abnormal state for at least *X* minutes.
	Unknowns         map[string]json.RawMessage `json:"-"`
}

func (me *SpikeDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"percent": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Alert if the observed load is more than *X* % of the expected value",
		},
		"minutes": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Alert if the service stays in abnormal state for at least *X* minutes",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *SpikeDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.LoadSpikePercent != nil {
		result["percent"] = int(*me.LoadSpikePercent)
	}
	if me.AbnormalMinutes != nil {
		result["minutes"] = int(*me.AbnormalMinutes)
	}
	return result, nil
}

func (me *SpikeDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "loadSpikePercent")
		delete(me.Unknowns, "minAbnormalStateDurationInMinutes")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	me.Enabled = true
	if value, ok := decoder.GetOk("percent"); ok {
		me.LoadSpikePercent = opt.NewInt32(int32(value.(int)))
	}
	if value, ok := decoder.GetOk("minutes"); ok {
		me.AbnormalMinutes = opt.NewInt32(int32(value.(int)))
	}
	return nil
}

func (me *SpikeDetection) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"enabled":                           me.Enabled,
		"loadSpikePercent":                  me.LoadSpikePercent,
		"minAbnormalStateDurationInMinutes": me.AbnormalMinutes,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *SpikeDetection) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"enabled":                           &me.Enabled,
		"loadSpikePercent":                  &me.LoadSpikePercent,
		"minAbnormalStateDurationInMinutes": &me.AbnormalMinutes,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
