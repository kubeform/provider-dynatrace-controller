package errors

import "github.com/dtcookie/hcl"

// Thresholds Custom thresholds for network errors. If not set, automatic mode is used.
//  **All** of these conditions must be met to trigger an alert.
type Thresholds struct {
	TotalPacketsRate int32 `json:"totalPacketsRate"` // Total receive/transmit packets rate is higher than *X* packets per second in 3 out of 5 samples.
	ErrorsPercentage int32 `json:"errorsPercentage"` // Receive/transmit error packet percentage is higher than *X*% in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"total_packets_rate": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Total receive/transmit packets rate is higher than *X* packets per second in 3 out of 5 samples",
		},
		"errors_percentage": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Receive/transmit error packet percentage is higher than *X*% in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["errors_percentage"] = int(me.ErrorsPercentage)
	result["total_packets_rate"] = int(me.TotalPacketsRate)
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("errors_percentage"); ok {
		me.ErrorsPercentage = int32(value.(int))
	}
	if value, ok := decoder.GetOk("total_packets_rate"); ok {
		me.TotalPacketsRate = int32(value.(int))
	}
	return nil
}
