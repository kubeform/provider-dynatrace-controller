package monitors

import "github.com/dtcookie/hcl"

type LoadingTimeThresholdsPolicy struct {
	Enabled    bool                  `json:"enabled"`    // Performance threshold is enabled (`true`) or disabled (`false`)
	Thresholds LoadingTimeThresholds `json:"thresholds"` // The list of performance threshold rules
}

func (me *LoadingTimeThresholdsPolicy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Performance threshold is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"thresholds": {
			Type:        hcl.TypeList,
			Description: "The list of performance threshold rules",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(LoadingTimeThresholds).Schema(),
			},
		},
	}
}

func (me *LoadingTimeThresholdsPolicy) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["enabled"] = me.Enabled
	if len(me.Thresholds) > 0 {
		if marshalled, err := me.Thresholds.MarshalHCL(); err == nil {
			result["thresholds"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *LoadingTimeThresholdsPolicy) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("enabled", &me.Enabled); err != nil {
		return err
	}
	if err := decoder.Decode("thresholds", &me.Thresholds); err != nil {
		return err
	}
	if me.Thresholds == nil {
		me.Thresholds = LoadingTimeThresholds{}
	}
	return nil
}
