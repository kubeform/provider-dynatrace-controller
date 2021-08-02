package space

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for low disk space. If not set, automatic mode is used.
type Thresholds struct {
	FreeSpacePercentage int32 `json:"freeSpacePercentage"` // Alert if free disk space is lower than *X*% in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"percentage": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Alert if free disk space is lower than *X*% in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["percentage"] = int(me.FreeSpacePercentage)
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("percentage"); ok {
		me.FreeSpacePercentage = int32(value.(int))
	}
	return nil
}
