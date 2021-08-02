package oom

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for Java out of memory. If not set, automatic mode is used.
type Thresholds struct {
	ExceptionCount int32 `json:"outOfMemoryExceptionsNumber"` // Alert if the number of Java out of memory exceptions is *X* per minute or higher.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"exception_count": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Alert if the number of Java out of memory exceptions is *X* per minute or higher",
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["exception_count"] = int(me.ExceptionCount)
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("exception_count"); ok {
		me.ExceptionCount = int32(value.(int))
	}
	return nil
}
