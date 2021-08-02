package inodes

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for low disk inodes number. If not set, automatic mode is used.
type Thresholds struct {
	FreeInodesPercentage int32 `json:"freeInodesPercentage"` // Alert if percentage of available inodes is lower than *X*% in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"percentage": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Alert if percentage of available inodes is lower than *X*% in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["percentage"] = int(me.FreeInodesPercentage)
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("percentage"); ok {
		me.FreeInodesPercentage = int32(value.(int))
	}
	return nil
}
