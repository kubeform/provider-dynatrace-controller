package responsetime

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/anomalies/common/detection"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Detection Configuration of response time degradation detection.
type Detection struct {
	AutomaticDetection *Autodetection `json:"automaticDetection,omitempty"` // Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.  Violation of **any** criterion triggers an alert.
	DetectionMode      detection.Mode `json:"detectionMode"`                // How to detect response time degradation: automatically, or based on fixed thresholds, or do not detect.
	Thresholds         *Thresholds    `json:"thresholds,omitempty"`         // Fixed thresholds for response time degradation detection.   Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
}

func (me *Detection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert",
			Elem:        &hcl.Resource{Schema: new(Autodetection).Schema()},
		},
		"thresholds": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Fixed thresholds for response time degradation detection",
			Elem:        &hcl.Resource{Schema: new(Thresholds).Schema()},
		},
	}
}

func (me *Detection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.AutomaticDetection != nil {
		if marshalled, err := me.AutomaticDetection.MarshalHCL(hcl.NewDecoder(decoder, "auto", 0)); err == nil {
			result["auto"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Thresholds != nil {
		if marshalled, err := me.Thresholds.MarshalHCL(hcl.NewDecoder(decoder, "thresholds", 0)); err == nil {
			result["thresholds"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Detection) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("auto.#"); ok {
		me.AutomaticDetection = new(Autodetection)
		me.DetectionMode = detection.Modes.DetectAutomatically
		if err := me.AutomaticDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "auto", 0)); err != nil {
			return err
		}
	} else if _, ok := decoder.GetOk("thresholds.#"); ok {
		me.Thresholds = new(Thresholds)
		me.DetectionMode = detection.Modes.DetectUsingFixedThresholds
		if err := me.Thresholds.UnmarshalHCL(hcl.NewDecoder(decoder, "thresholds", 0)); err != nil {
			return err
		}
	} else {
		me.DetectionMode = detection.Modes.DontDetect
	}
	return nil
}

func (me *Detection) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"automaticDetection": me.AutomaticDetection,
		"detectionMode":      me.DetectionMode,
		"thresholds":         me.Thresholds,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Detection) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"automaticDetection": &me.AutomaticDetection,
		"detectionMode":      &me.DetectionMode,
		"thresholds":         &me.Thresholds,
	}); err != nil {
		return err
	}
	return nil
}
