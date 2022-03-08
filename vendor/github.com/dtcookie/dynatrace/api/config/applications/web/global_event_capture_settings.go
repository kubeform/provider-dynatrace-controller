package web

import "github.com/dtcookie/hcl"

// GlobalEventCaptureSettings Global event capture settings
type GlobalEventCaptureSettings struct {
	MouseUp                            bool   `json:"mouseUp"`                            // MouseUp enabled/disabled
	MouseDown                          bool   `json:"mouseDown"`                          // MouseDown enabled/disabled
	Click                              bool   `json:"click"`                              // Click enabled/disabled
	DoubleClick                        bool   `json:"doubleClick"`                        // DoubleClick enabled/disabled
	KeyUp                              bool   `json:"keyUp"`                              // KeyUp enabled/disabled
	KeyDown                            bool   `json:"keyDown"`                            // KeyDown enabled/disabled
	Scroll                             bool   `json:"scroll"`                             // Scroll enabled/disabled
	AdditionalEventCapturedAsUserInput string `json:"additionalEventCapturedAsUserInput"` // Additional events to be captured globally as user input. \n\nFor example `DragStart` or `DragEnd`. Maximum 100 characters.
}

func (me *GlobalEventCaptureSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"mouseup": {
			Type:        hcl.TypeBool,
			Description: "MouseUp enabled/disabled",
			Optional:    true,
		},
		"mousedown": {
			Type:        hcl.TypeBool,
			Description: "MouseDown enabled/disabled",
			Optional:    true,
		},
		"click": {
			Type:        hcl.TypeBool,
			Description: "Click enabled/disabled",
			Optional:    true,
		},
		"doubleclick": {
			Type:        hcl.TypeBool,
			Description: "DoubleClick enabled/disabled",
			Optional:    true,
		},
		"keyup": {
			Type:        hcl.TypeBool,
			Description: "KeyUp enabled/disabled",
			Optional:    true,
		},
		"keydown": {
			Type:        hcl.TypeBool,
			Description: "KeyDown enabled/disabled",
			Optional:    true,
		},
		"scroll": {
			Type:        hcl.TypeBool,
			Description: "Scroll enabled/disabled",
			Optional:    true,
		},
		"additional_event_captured_as_user_input": {
			Type:        hcl.TypeString,
			Description: "Additional events to be captured globally as user input. \n\nFor example `DragStart` or `DragEnd`. Maximum 100 characters.",
			Optional:    true,
		},
	}
}

func (me *GlobalEventCaptureSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"mouseup":     me.MouseUp,
		"mousedown":   me.MouseDown,
		"click":       me.Click,
		"doubleclick": me.DoubleClick,
		"keyup":       me.KeyUp,
		"keydown":     me.KeyDown,
		"scroll":      me.Scroll,
		"additional_event_captured_as_user_input": me.AdditionalEventCapturedAsUserInput,
	})
}

func (me *GlobalEventCaptureSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"mouseup":     &me.MouseUp,
		"mousedown":   &me.MouseDown,
		"click":       &me.Click,
		"doubleclick": &me.DoubleClick,
		"keyup":       &me.KeyUp,
		"keydown":     &me.KeyDown,
		"scroll":      &me.Scroll,
		"additional_event_captured_as_user_input": &me.AdditionalEventCapturedAsUserInput,
	})
}
