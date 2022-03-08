package web

import "github.com/dtcookie/hcl"

// ContentCapture contains settings for content capture
type ContentCapture struct {
	ResourceTimingSettings        *ResourceTimingSettings   `json:"resourceTimingSettings"`        // Settings for resource timings capture
	JavaScriptErrors              bool                      `json:"javaScriptErrors"`              // JavaScript errors monitoring enabled/disabled
	TimeoutSettings               *TimeoutSettings          `json:"timeoutSettings"`               // Settings for timed action capture
	VisuallyCompleteAndSpeedIndex bool                      `json:"visuallyCompleteAndSpeedIndex"` // Visually complete and Speed index support enabled/disabled
	VisuallyCompleteSettings      *VisuallyCompleteSettings `json:"visuallyComplete2Settings"`     // Settings for VisuallyComplete
}

func (me *ContentCapture) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"resource_timing_settings": {
			Type:        hcl.TypeList,
			Description: "Settings for resource timings capture",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ResourceTimingSettings).Schema()},
		},
		"javascript_errors": {
			Type:        hcl.TypeBool,
			Description: "JavaScript errors monitoring enabled/disabled",
			Optional:    true,
		},
		"timeout_settings": {
			Type:        hcl.TypeList,
			Description: "Settings for timed action capture",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(TimeoutSettings).Schema()},
		},
		"visually_complete_and_speed_index": {
			Type:        hcl.TypeBool,
			Description: "Visually complete and Speed index support enabled/disabled",
			Optional:    true,
		},
		"visually_complete_settings": {
			Type:        hcl.TypeList,
			Description: "Settings for VisuallyComplete",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(VisuallyCompleteSettings).Schema()},
		},
	}
}

func (me *ContentCapture) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"resource_timing_settings":          me.ResourceTimingSettings,
		"javascript_errors":                 me.JavaScriptErrors,
		"timeout_settings":                  me.TimeoutSettings,
		"visually_complete_and_speed_index": me.VisuallyCompleteAndSpeedIndex,
		"visually_complete_settings":        me.VisuallyCompleteSettings,
	})
}

func (me *ContentCapture) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"resource_timing_settings":          &me.ResourceTimingSettings,
		"javascript_errors":                 &me.JavaScriptErrors,
		"timeout_settings":                  &me.TimeoutSettings,
		"visually_complete_and_speed_index": &me.VisuallyCompleteAndSpeedIndex,
		"visually_complete_settings":        &me.VisuallyCompleteSettings,
	})
}
