package failurerate

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/anomalies/common/detection"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Detection Configuration of failure rate increase detection.
type Detection struct {
	Thresholds         *Thresholds    `json:"thresholds,omitempty"`         // Fixed thresholds for failure rate increase detection.   Required if **detectionMode** is `DETECT_USING_FIXED_THRESHOLDS`. Not applicable otherwise.
	AutomaticDetection *Autodetection `json:"automaticDetection,omitempty"` // Parameters of failure rate increase auto-detection. Required if **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.  The absolute and relative thresholds **both** must exceed to trigger an alert.  Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	DetectionMode      detection.Mode `json:"detectionMode"`                // How to detect failure rate increase: automatically, or based on fixed thresholds, or do not detect.
}

func (me *Detection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"auto": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%",
			Elem:        &hcl.Resource{Schema: new(Autodetection).Schema()},
		},
		"thresholds": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Fixed thresholds for failure rate increase detection",
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
	} else if me.Thresholds != nil {
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
		me.DetectionMode = detection.Modes.DetectAutomatically
		me.AutomaticDetection = new(Autodetection)
		if err := me.AutomaticDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "auto", 0)); err != nil {
			return err
		}
	} else if _, ok := decoder.GetOk("thresholds.#"); ok {
		me.DetectionMode = detection.Modes.DetectUsingFixedThresholds
		me.Thresholds = new(Thresholds)
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
		"thresholds":         me.Thresholds,
		"automaticDetection": me.AutomaticDetection,
		"detectionMode":      me.DetectionMode,
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
		"thresholds":         &me.Thresholds,
		"automaticDetection": &me.AutomaticDetection,
		"detectionMode":      &me.DetectionMode,
	}); err != nil {
		return err
	}
	return nil
}
