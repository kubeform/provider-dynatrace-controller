package web

import "github.com/dtcookie/hcl"

// ResourceTimingSettings configures resource timings capture
type ResourceTimingSettings struct {
	W3CResourceTimings                        bool                      `json:"w3cResourceTimings"`                        // W3C resource timings for third party/CDN enabled/disabled
	NonW3CResourceTimings                     bool                      `json:"nonW3cResourceTimings"`                     // Timing for JavaScript files and images on non-W3C supported browsers enabled/disabled
	NonW3CResourceTimingsInstrumentationDelay int32                     `json:"nonW3cResourceTimingsInstrumentationDelay"` // Instrumentation delay for monitoring resource and image resource impact in browsers that don't offer W3C resource timings. \n\nValid values range from 0 to 9999.\n\nOnly effective if **nonW3cResourceTimings** is enabled
	ResourceTimingCaptureType                 ResourceTimingCaptureType `json:"resourceTimingCaptureType"`                 // Defines how detailed resource timings are captured.\n\nOnly effective if **w3cResourceTimings** or **nonW3cResourceTimings** is enabled. Possible values are `CAPTURE_ALL_SUMMARIES`, `CAPTURE_FULL_DETAILS` and `CAPTURE_LIMITED_SUMMARIES`
	ResourceTimingsDomainLimit                int32                     `json:"resourceTimingsDomainLimit"`                // Limits the number of domains for which W3C resource timings are captured.\n\nOnly effective if **resourceTimingCaptureType** is `CAPTURE_LIMITED_SUMMARIES`. Valid values range from 0 to 50.
}

func (me *ResourceTimingSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"w3c_resource_timings": {
			Type:        hcl.TypeBool,
			Description: "W3C resource timings for third party/CDN enabled/disabled",
			Optional:    true,
		},
		"non_w3c_resource_timings": {
			Type:        hcl.TypeBool,
			Description: "Timing for JavaScript files and images on non-W3C supported browsers enabled/disabled",
			Optional:    true,
		},
		"instrumentation_delay": {
			Type:        hcl.TypeInt,
			Description: "Instrumentation delay for monitoring resource and image resource impact in browsers that don't offer W3C resource timings. \n\nValid values range from 0 to 9999.\n\nOnly effective if `nonW3cResourceTimings` is enabled",
			Required:    true,
		},
		"resource_timing_capture_type": {
			Type:        hcl.TypeString,
			Description: "Defines how detailed resource timings are captured.\n\nOnly effective if **w3cResourceTimings** or **nonW3cResourceTimings** is enabled. Possible values are `CAPTURE_ALL_SUMMARIES`, `CAPTURE_FULL_DETAILS` and `CAPTURE_LIMITED_SUMMARIES`",
			Required:    true,
		},
		"resource_timings_domain_limit": {
			Type:        hcl.TypeInt,
			Description: "Limits the number of domains for which W3C resource timings are captured.\n\nOnly effective if **resourceTimingCaptureType** is `CAPTURE_LIMITED_SUMMARIES`. Valid values range from 0 to 50.",
			Required:    true,
		},
	}
}

func (me *ResourceTimingSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"w3c_resource_timings":          me.W3CResourceTimings,
		"non_w3c_resource_timings":      me.NonW3CResourceTimings,
		"instrumentation_delay":         me.NonW3CResourceTimingsInstrumentationDelay,
		"resource_timing_capture_type":  me.ResourceTimingCaptureType,
		"resource_timings_domain_limit": me.ResourceTimingsDomainLimit,
	})
}

func (me *ResourceTimingSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"w3c_resource_timings":          &me.W3CResourceTimings,
		"non_w3c_resource_timings":      &me.NonW3CResourceTimings,
		"instrumentation_delay":         &me.NonW3CResourceTimingsInstrumentationDelay,
		"resource_timing_capture_type":  &me.ResourceTimingCaptureType,
		"resource_timings_domain_limit": &me.ResourceTimingsDomainLimit,
	})
}
