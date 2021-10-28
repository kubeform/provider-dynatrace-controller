package mobile

import "github.com/dtcookie/hcl"

// MobileCustomApdex represents Apdex configuration of a mobile or custom application. \n\nA duration less than the **tolerable** threshold is considered satisfied
type MobileCustomApdex struct {
	ToleratedThreshold   int32 `json:"toleratedThreshold"`   // Apdex **tolerable** threshold, in milliseconds: a duration greater than or equal to this value is considered tolerable
	FrustratingThreshold int32 `json:"frustratingThreshold"` // Apdex **frustrated** threshold, in milliseconds: a duration greater than or equal to this value is considered frustrated
	FrustratedOnError    bool  `json:"frustratedOnError"`    // Apdex error condition: if `true` the user session is considered frustrated when an error is reported
}

func (me *MobileCustomApdex) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"tolerable": {
			Type:        hcl.TypeInt,
			Description: "Apdex **tolerable** threshold, in milliseconds: a duration greater than or equal to this value is considered tolerable",
			Required:    true,
		},
		"frustrated": {
			Type:        hcl.TypeInt,
			Description: "Apdex **frustrated** threshold, in milliseconds: a duration greater than or equal to this value is considered frustrated",
			Required:    true,
		},
		"frustrated_on_error": {
			Type:        hcl.TypeBool,
			Description: "Apdex error condition: if `true` the user session is considered frustrated when an error is reported",
			Optional:    true,
		},
	}
}

func (me *MobileCustomApdex) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("tolerable", &me.ToleratedThreshold); err != nil {
		return err
	}
	if err := decoder.Decode("frustrated", &me.FrustratingThreshold); err != nil {
		return err
	}
	if err := decoder.Decode("frustrated_on_error", &me.FrustratedOnError); err != nil {
		return err
	}
	return nil
}

func (me *MobileCustomApdex) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if err := properties.Encode("tolerable", me.ToleratedThreshold); err != nil {
		return nil, err
	}
	if err := properties.Encode("frustrated", me.FrustratingThreshold); err != nil {
		return nil, err
	}
	if me.FrustratedOnError {
		if err := properties.Encode("frustrated_on_error", me.FrustratedOnError); err != nil {
			return nil, err
		}
	}
	return properties, nil
}
