package responsetime

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/anomalies/common/load"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Autodetection Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.
// Violation of **any** criterion triggers an alert.
type Autodetection struct {
	LoadThreshold       load.Threshold             `json:"loadThreshold"`                              // Minimal service load to detect response time degradation.   Response time degradation of services with smaller load won't trigger alerts.
	Milliseconds        int32                      `json:"responseTimeDegradationMilliseconds"`        // Alert if the response time degrades by more than *X* milliseconds.
	Percent             int32                      `json:"responseTimeDegradationPercent"`             // Alert if the response time degrades by more than *X* %.
	SlowestMilliseconds int32                      `json:"slowestResponseTimeDegradationMilliseconds"` // Alert if the response time of the slowest 10% degrades by more than *X* milliseconds.
	SlowestPercent      int32                      `json:"slowestResponseTimeDegradationPercent"`      // Alert if the response time of the slowest 10% degrades by more than *X* %.
	Unknowns            map[string]json.RawMessage `json:"-"`
}

func (me *Autodetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"load": {
			Type:        hcl.TypeString,
			Description: "Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`",
			Required:    true,
		},
		"milliseconds": {
			Type:        hcl.TypeInt,
			Description: "Alert if the response time degrades by more than *X* milliseconds",
			Required:    true,
		},
		"percent": {
			Type:        hcl.TypeInt,
			Description: "Alert if the response time degrades by more than *X* %",
			Required:    true,
		},
		"slowest_milliseconds": {
			Type:        hcl.TypeInt,
			Description: "Alert if the response time of the slowest 10% degrades by more than *X* milliseconds",
			Required:    true,
		},
		"slowest_percent": {
			Type:        hcl.TypeInt,
			Description: "Alert if the response time of the slowest 10% degrades by more than *X* milliseconds",
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
	result["load"] = string(me.LoadThreshold)
	result["milliseconds"] = int(me.Milliseconds)
	result["percent"] = int(me.Percent)
	result["slowest_milliseconds"] = int(me.SlowestMilliseconds)
	result["slowest_percent"] = int(me.SlowestPercent)
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
		delete(me.Unknowns, "loadThreshold")
		delete(me.Unknowns, "responseTimeDegradationMilliseconds")
		delete(me.Unknowns, "responseTimeDegradationPercent")
		delete(me.Unknowns, "slowestResponseTimeDegradationMilliseconds")
		delete(me.Unknowns, "slowestResponseTimeDegradationPercent")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("load"); ok {
		me.LoadThreshold = load.Threshold(value.(string))
	}
	if value, ok := decoder.GetOk("milliseconds"); ok {
		me.Milliseconds = int32(value.(int))
	}
	if value, ok := decoder.GetOk("percent"); ok {
		me.Percent = int32(value.(int))
	}
	if value, ok := decoder.GetOk("slowest_milliseconds"); ok {
		me.SlowestMilliseconds = int32(value.(int))
	}
	if value, ok := decoder.GetOk("slowest_percent"); ok {
		me.SlowestPercent = int32(value.(int))
	}
	return nil
}

func (me *Autodetection) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"loadThreshold":                              me.LoadThreshold,
		"responseTimeDegradationMilliseconds":        me.Milliseconds,
		"responseTimeDegradationPercent":             me.Percent,
		"slowestResponseTimeDegradationMilliseconds": me.SlowestMilliseconds,
		"slowestResponseTimeDegradationPercent":      me.SlowestPercent,
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
		"loadThreshold":                              &me.LoadThreshold,
		"responseTimeDegradationMilliseconds":        &me.Milliseconds,
		"responseTimeDegradationPercent":             &me.Percent,
		"slowestResponseTimeDegradationMilliseconds": &me.SlowestMilliseconds,
		"slowestResponseTimeDegradationPercent":      &me.SlowestPercent,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
