package failurerate

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Autodetection Parameters of failure rate increase auto-detection. Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.
// The absolute and relative thresholds **both** must exceed to trigger an alert.
// Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:
// Absolute: 1.5% + **1%** = 2.5%
// Relative: 1.5% + 1.5% * **50%** = 2.25%
type Autodetection struct {
	PercentAbsolute int32                      `json:"failingServiceCallPercentageIncreaseAbsolute"` // Absolute increase of failing service calls to trigger an alert, %.
	PercentRelative int32                      `json:"failingServiceCallPercentageIncreaseRelative"` // Relative increase of failing service calls to trigger an alert, %.
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *Autodetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"absolute": {
			Type:        hcl.TypeInt,
			Description: "Absolute increase of failing service calls to trigger an alert, %",
			Required:    true,
		},
		"relative": {
			Type:        hcl.TypeInt,
			Description: "Relative increase of failing service calls to trigger an alert, %",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Autodetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["absolute"] = int(me.PercentAbsolute)
	result["relative"] = int(me.PercentRelative)
	return result, nil
}

func (me *Autodetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "failingServiceCallPercentageIncreaseAbsolute")
		delete(me.Unknowns, "failingServiceCallPercentageIncreaseRelative")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("relative"); ok {
		me.PercentRelative = int32(value.(int))
	}
	if value, ok := decoder.GetOk("absolute"); ok {
		me.PercentAbsolute = int32(value.(int))
	}
	return nil
}

func (me *Autodetection) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"failingServiceCallPercentageIncreaseAbsolute": me.PercentAbsolute,
		"failingServiceCallPercentageIncreaseRelative": me.PercentRelative,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Autodetection) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"failingServiceCallPercentageIncreaseAbsolute": &me.PercentAbsolute,
		"failingServiceCallPercentageIncreaseRelative": &me.PercentRelative,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
