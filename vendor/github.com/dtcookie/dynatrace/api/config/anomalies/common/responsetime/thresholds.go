package responsetime

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/anomalies/common"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/load"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Thresholds Fixed thresholds for response time degradation detection.
//  Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
type Thresholds struct {
	LoadThreshold       load.Threshold             `json:"loadThreshold"`                            // Minimal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	Sensitivity         common.Sensitivity         `json:"sensitivity"`                              // Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert.
	Milliseconds        int32                      `json:"responseTimeThresholdMilliseconds"`        // Response time during any 5-minute period to trigger an alert, in milliseconds.
	SlowestMilliseconds int32                      `json:"slowestResponseTimeThresholdMilliseconds"` // Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds.
	Unknowns            map[string]json.RawMessage `json:"-"`
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"load": {
			Type:        hcl.TypeString,
			Description: "Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`",
			Required:    true,
		},
		"sensitivity": {
			Type:        hcl.TypeString,
			Description: "Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert",
			Required:    true,
		},
		"milliseconds": {
			Type:        hcl.TypeInt,
			Description: "Response time during any 5-minute period to trigger an alert, in milliseconds",
			Required:    true,
		},
		"slowest_milliseconds": {
			Type:        hcl.TypeInt,
			Description: "Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds",
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
	result["load"] = string(me.LoadThreshold)
	result["sensitivity"] = string(me.Sensitivity)
	result["milliseconds"] = int(me.Milliseconds)
	result["slowest_milliseconds"] = int(me.SlowestMilliseconds)
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
		delete(me.Unknowns, "loadThreshold")
		delete(me.Unknowns, "sensitivity")
		delete(me.Unknowns, "responseTimeThresholdMilliseconds")
		delete(me.Unknowns, "slowestResponseTimeThresholdMilliseconds")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("load"); ok {
		me.LoadThreshold = load.Threshold(value.(string))
	}
	if value, ok := decoder.GetOk("sensitivity"); ok {
		me.Sensitivity = common.Sensitivity(value.(string))
	}
	if value, ok := decoder.GetOk("milliseconds"); ok {
		me.Milliseconds = int32(value.(int))
	}
	if value, ok := decoder.GetOk("slowest_milliseconds"); ok {
		me.SlowestMilliseconds = int32(value.(int))
	}
	return nil
}

func (me *Thresholds) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"loadThreshold":                            me.LoadThreshold,
		"sensitivity":                              me.Sensitivity,
		"responseTimeThresholdMilliseconds":        me.Milliseconds,
		"slowestResponseTimeThresholdMilliseconds": me.SlowestMilliseconds,
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
		"loadThreshold":                            &me.LoadThreshold,
		"sensitivity":                              &me.Sensitivity,
		"responseTimeThresholdMilliseconds":        &me.Milliseconds,
		"slowestResponseTimeThresholdMilliseconds": &me.SlowestMilliseconds,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
