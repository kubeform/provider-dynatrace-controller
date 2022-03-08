package web

import "github.com/dtcookie/hcl"

// VisuallyCompleteSettings Settings for VisuallyComplete
type VisuallyCompleteSettings struct {
	ExcludeURLRegex      *string `json:"excludeUrlRegex"`             // A RegularExpression used to exclude images and iframes from being detected by the VC module
	IgnoredMutationsList *string `json:"ignoredMutationsList"`        // Query selector for mutation nodes to ignore in VC and SI calculation
	MutationTimeout      *int32  `json:"mutationTimeout,omitempty"`   // Determines the time in ms VC waits after an action closes to start calculation. Defaults to 50. Valid values range from 0 to 5000.
	InactivityTimeout    *int32  `json:"inactivityTimeout,omitempty"` // The time in ms the VC module waits for no mutations happening on the page after the load action. Defaults to 1000. Valid values range from 0 to 30000.
	Threshold            *int32  `json:"threshold,omitempty"`         // Minimum visible area in pixels of elements to be counted towards VC and SI. Defaults to 50. Valid values range from 0 to 10000.
}

func (me *VisuallyCompleteSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"exclude_url_regex": {
			Type:        hcl.TypeString,
			Description: "A RegularExpression used to exclude images and iframes from being detected by the VC module",
			Optional:    true,
		},
		"ignored_mutations_list": {
			Type:        hcl.TypeString,
			Description: "Query selector for mutation nodes to ignore in VC and SI calculation",
			Optional:    true,
		},
		"mutation_timeout": {
			Type:        hcl.TypeInt,
			Description: "Determines the time in ms VC waits after an action closes to start calculation. Defaults to 50. Valid values range from 0 to 5000.",
			Optional:    true,
		},
		"inactivity_timeout": {
			Type:        hcl.TypeInt,
			Description: "The time in ms the VC module waits for no mutations happening on the page after the load action. Defaults to 1000. Valid values range from 0 to 30000.",
			Optional:    true,
		},
		"threshold": {
			Type:        hcl.TypeInt,
			Description: "Minimum visible area in pixels of elements to be counted towards VC and SI. Defaults to 50. Valid values range from 0 to 10000.",
			Optional:    true,
		},
	}
}

func (me *VisuallyCompleteSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"exclude_url_regex":      me.ExcludeURLRegex,
		"ignored_mutations_list": me.IgnoredMutationsList,
		"mutation_timeout":       me.MutationTimeout,
		"inactivity_timeout":     me.InactivityTimeout,
		"threshold":              me.Threshold,
	})
}

func (me *VisuallyCompleteSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"exclude_url_regex":      &me.ExcludeURLRegex,
		"ignored_mutations_list": &me.IgnoredMutationsList,
		"mutation_timeout":       &me.MutationTimeout,
		"inactivity_timeout":     &me.InactivityTimeout,
		"threshold":              &me.Threshold,
	})
}
