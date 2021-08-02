package load

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// DropDetection The configuration of load drops detection.
type DropDetection struct {
	Enabled         bool   `json:"enabled"`                                     // The detection is enabled (`true`) or disabled (`false`).
	LoadDropPercent *int32 `json:"loadDropPercent,omitempty"`                   // Alert if the observed load is less than *X* % of the expected value.
	AbnormalMinutes *int32 `json:"minAbnormalStateDurationInMinutes,omitempty"` // Alert if the service stays in abnormal state for at least *X* minutes.
}

func (me *DropDetection) Schema() map[string]*hcl.Schema {
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
	}
}

func (me *DropDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.LoadDropPercent != nil {
		result["percent"] = int(*me.LoadDropPercent)
	}
	if me.AbnormalMinutes != nil {
		result["minutes"] = int(*me.AbnormalMinutes)
	}
	return result, nil
}

func (me *DropDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Enabled = true
	if value, ok := decoder.GetOk("percent"); ok {
		me.LoadDropPercent = opt.NewInt32(int32(value.(int)))
	}
	if value, ok := decoder.GetOk("minutes"); ok {
		me.AbnormalMinutes = opt.NewInt32(int32(value.(int)))
	}
	return nil
}

func (me *DropDetection) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"enabled":                           me.Enabled,
		"loadDropPercent":                   me.LoadDropPercent,
		"minAbnormalStateDurationInMinutes": me.AbnormalMinutes,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *DropDetection) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"enabled":                           &me.Enabled,
		"loadDropPercent":                   &me.LoadDropPercent,
		"minAbnormalStateDurationInMinutes": &me.AbnormalMinutes,
	}); err != nil {
		return err
	}
	return nil
}
