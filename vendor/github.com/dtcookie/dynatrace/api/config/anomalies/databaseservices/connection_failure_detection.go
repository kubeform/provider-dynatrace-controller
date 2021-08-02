package databaseservices

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// ConnectionFailureDetection Parameters of the failed database connections detection.
// The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period.
type ConnectionFailureDetection struct {
	ConnectionFailsCount *int32 `json:"connectionFailsCount,omitempty"` // Number of failed database connections during any **timePeriodMinutes** minutes period to trigger an alert.
	Enabled              bool   `json:"enabled"`                        // The detection is enabled (`true`) or disabled (`false`).
	TimePeriodMinutes    *int32 `json:"timePeriodMinutes,omitempty"`    // The *X* minutes time period during which the **connectionFailsCount** is evaluated.
}

func (me *ConnectionFailureDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"connection_fails_count": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Number of failed database connections during any **eval_period** minutes period to trigger an alert",
		},
		"eval_period": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "The *X* minutes time period during which the **connection_fails_count** is evaluated",
		},
	}
}

func (me *ConnectionFailureDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["connection_fails_count"] = int(*me.ConnectionFailsCount)
	result["eval_period"] = int(*me.TimePeriodMinutes)

	return result, nil
}

func (me *ConnectionFailureDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("connection_fails_count"); ok {
		me.ConnectionFailsCount = opt.NewInt32(int32(value.(int)))
	}
	if value, ok := decoder.GetOk("eval_period"); ok {
		me.TimePeriodMinutes = opt.NewInt32(int32(value.(int)))
	}
	me.Enabled = true
	return nil
}
