package web

import (
	"github.com/dtcookie/hcl"
)

// Defines the Apdex settings of an application
type ApdexSettings struct {
	Threshold                    *int `json:"threshold,omitempty"`                    // no documentation available
	ToleratedThreshold           *int `json:"toleratedThreshold,omitempty"`           // Maximal value of apdex, which is considered as satisfied user experience. Values between 0 and 60000 are allowed.
	FrustratingThreshold         *int `json:"frustratingThreshold,omitempty"`         // Maximal value of apdex, which is considered as tolerable user experience. Values between 0 and 240000 are allowed.
	ToleratedFallbackThreshold   *int `json:"toleratedFallbackThreshold,omitempty"`   // Fallback threshold of an XHR action, defining a satisfied user experience, when the configured KPM is not available. Values between 0 and 60000 are allowed.
	FrustratingFallbackThreshold *int `json:"frustratingFallbackThreshold,omitempty"` // Fallback threshold of an XHR action, defining a tolerable user experience, when the configured KPM is not available. Values between 0 and 240000 are allowed.
}

func (me *ApdexSettings) IsEmpty() bool {
	return me.Threshold == nil && me.ToleratedThreshold == nil && me.FrustratingThreshold == nil && me.ToleratedFallbackThreshold == nil && me.FrustratingFallbackThreshold == nil
}

func (me *ApdexSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "no documentation available",
			Optional:    true,
		},
		"tolerated_threshold": {
			Type:        hcl.TypeInt,
			Description: "Maximal value of apdex, which is considered as satisfied user experience. Values between 0 and 60000 are allowed.",
			Optional:    true,
		},
		"frustrating_threshold": {
			Type:        hcl.TypeInt,
			Description: "Maximal value of apdex, which is considered as tolerable user experience. Values between 0 and 240000 are allowed.",
			Optional:    true,
		},
		"tolerated_fallback_threshold": {
			Type:        hcl.TypeInt,
			Description: "Fallback threshold of an XHR action, defining a satisfied user experience, when the configured KPM is not available. Values between 0 and 60000 are allowed.",
			Optional:    true,
		},
		"frustrating_fallback_threshold": {
			Type:        hcl.TypeInt,
			Description: "Fallback threshold of an XHR action, defining a tolerable user experience, when the configured KPM is not available. Values between 0 and 240000 are allowed.",
			Optional:    true,
		},
	}
}

func (me *ApdexSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"threshold":                      me.Threshold,
		"tolerated_threshold":            me.ToleratedThreshold,
		"frustrating_threshold":          me.FrustratingThreshold,
		"tolerated_fallback_threshold":   me.ToleratedFallbackThreshold,
		"frustrating_fallback_threshold": me.FrustratingFallbackThreshold,
	})
}

func (me *ApdexSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"threshold":                      &me.Threshold,
		"tolerated_threshold":            &me.ToleratedThreshold,
		"frustrating_threshold":          &me.FrustratingThreshold,
		"tolerated_fallback_threshold":   &me.ToleratedFallbackThreshold,
		"frustrating_fallback_threshold": &me.FrustratingFallbackThreshold,
	})
}
