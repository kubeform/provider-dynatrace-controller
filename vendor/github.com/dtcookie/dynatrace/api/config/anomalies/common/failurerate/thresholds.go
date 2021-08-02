package failurerate

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/anomalies/common"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Thresholds Fixed thresholds for failure rate increase detection.
//  Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type Thresholds struct {
	Sensitivity common.Sensitivity         `json:"sensitivity"` // Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert.
	Threshold   int32                      `json:"threshold"`   // Failure rate during any 5-minute period to trigger an alert, %.
	Unknowns    map[string]json.RawMessage `json:"-"`
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert",
			Required:    true,
		},
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "Failure rate during any 5-minute period to trigger an alert, %",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["sensitivity"] = string(me.Sensitivity)
	result["threshold"] = int(me.Threshold)
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "sensitivity")
		delete(me.Unknowns, "threshold")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("sensitivity"); ok {
		me.Sensitivity = common.Sensitivity(value.(string))
	}
	if value, ok := decoder.GetOk("threshold"); ok {
		me.Threshold = int32(value.(int))
	}
	return nil
}

func (me *Thresholds) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"sensitivity": me.Sensitivity,
		"threshold":   me.Threshold,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Thresholds) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"sensitivity": &me.Sensitivity,
		"threshold":   &me.Threshold,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
