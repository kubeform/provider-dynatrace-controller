package utilization

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for high network utilization. If not set, automatic mode is used.
type Thresholds struct {
	UtilizationPercentage int32 `json:"utilizationPercentage"` // Alert if sent/received traffic utilization is higher than *X*% in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"utilization": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Alert if sent/received traffic utilization is higher than *X*% in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["utilization"] = int(me.UtilizationPercentage)
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("utilization"); ok {
		me.UtilizationPercentage = int32(value.(int))
	}
	return nil
}
