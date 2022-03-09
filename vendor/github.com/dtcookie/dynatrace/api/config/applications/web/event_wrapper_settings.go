package web

import "github.com/dtcookie/hcl"

// EventWrapperSettings In addition to the event handlers, events called using `addEventListener` or `attachEvent` can be captured. Be careful with this option! Event wrappers can conflict with the JavaScript code on a web page
type EventWrapperSettings struct {
	Click      bool `json:"click"`      // Click enabled/disabled
	MouseUp    bool `json:"mouseUp"`    // MouseUp enabled/disabled
	Change     bool `json:"change"`     // Change enabled/disabled
	Blur       bool `json:"blur"`       // Blur enabled/disabled
	TouchStart bool `json:"touchStart"` // TouchStart enabled/disabled
	TouchEnd   bool `json:"touchEnd"`   // TouchEnd enabled/disabled
}

func (me *EventWrapperSettings) IsDefault() bool {
	return !(me.Click || me.MouseUp || me.Change || me.Blur || me.TouchStart || me.TouchEnd)
}

func (me *EventWrapperSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"click": {
			Type:        hcl.TypeBool,
			Description: "Click enabled/disabled",
			Optional:    true,
		},
		"mouseup": {
			Type:        hcl.TypeBool,
			Description: "MouseUp enabled/disabled",
			Optional:    true,
		},
		"change": {
			Type:        hcl.TypeBool,
			Description: "Change enabled/disabled",
			Optional:    true,
		},
		"blur": {
			Type:        hcl.TypeBool,
			Description: "Blur enabled/disabled",
			Optional:    true,
		},
		"touch_start": {
			Type:        hcl.TypeBool,
			Description: "TouchStart enabled/disabled",
			Optional:    true,
		},
		"touch_end": {
			Type:        hcl.TypeBool,
			Description: "TouchEnd enabled/disabled",
			Optional:    true,
		},
	}
}

func (me *EventWrapperSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"click":       me.Click,
		"mouseup":     me.MouseUp,
		"change":      me.Change,
		"blur":        me.Blur,
		"touch_start": me.TouchStart,
		"touch_end":   me.TouchEnd,
	})
}

func (me *EventWrapperSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"click":       &me.Click,
		"mouseup":     &me.MouseUp,
		"change":      &me.Change,
		"blur":        &me.Blur,
		"touch_start": &me.TouchStart,
		"touch_end":   &me.TouchEnd,
	})
}
