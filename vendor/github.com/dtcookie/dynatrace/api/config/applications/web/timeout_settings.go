package web

import "github.com/dtcookie/hcl"

// TimeoutSettings configures timed action capture
type TimeoutSettings struct {
	TimedActionSupport          bool  `json:"timedActionSupport"`          // Timed action support enabled/disabled. \n\nEnable to detect actions that trigger sending of XHRs via *setTimout* methods
	TemporaryActionLimit        int32 `json:"temporaryActionLimit"`        // Defines how deep temporary actions may cascade. 0 disables temporary actions completely. Recommended value if enabled is 3
	TemporaryActionTotalTimeout int32 `json:"temporaryActionTotalTimeout"` // The total timeout of all cascaded timeouts that should still be able to create a temporary action
}

func (me *TimeoutSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"timed_action_support": {
			Type:        hcl.TypeBool,
			Description: "Timed action support enabled/disabled. \n\nEnable to detect actions that trigger sending of XHRs via `setTimout` methods",
			Optional:    true,
		},
		"temporary_action_limit": {
			Type:        hcl.TypeInt,
			Description: "Defines how deep temporary actions may cascade. 0 disables temporary actions completely. Recommended value if enabled is 3",
			Required:    true,
		},
		"temporary_action_total_timeout": {
			Type:        hcl.TypeInt,
			Description: "The total timeout of all cascaded timeouts that should still be able to create a temporary action",
			Required:    true,
		},
	}
}

func (me *TimeoutSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"timed_action_support":           me.TimedActionSupport,
		"temporary_action_limit":         me.TemporaryActionLimit,
		"temporary_action_total_timeout": me.TemporaryActionTotalTimeout,
	})
}

func (me *TimeoutSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"timed_action_support":           &me.TimedActionSupport,
		"temporary_action_limit":         &me.TemporaryActionLimit,
		"temporary_action_total_timeout": &me.TemporaryActionTotalTimeout,
	})
}
