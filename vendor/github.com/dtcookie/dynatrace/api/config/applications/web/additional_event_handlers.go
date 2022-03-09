package web

import "github.com/dtcookie/hcl"

// AdditionalEventHandlers Additional event handlers and wrappers
type AdditionalEventHandlers struct {
	UseMouseUpEventForClicks bool  `json:"userMouseupEventForClicks"` // Use mouseup event for clicks enabled/disabled
	ClickEventHandler        bool  `json:"clickEventHandler"`         // Click event handler enabled/disabled
	MouseUpEventHandler      bool  `json:"mouseupEventHandler"`       // Mouseup event handler enabled/disabled
	BlurEventHandler         bool  `json:"blurEventHandler"`          // Blur event handler enabled/disabled
	ChangeEventHandler       bool  `json:"changeEventHandler"`        // Change event handler enabled/disabled
	ToStringMethod           bool  `json:"toStringMethod"`            // toString method enabled/disabled
	MaxDomNodesToInstrument  int32 `json:"maxDomNodesToInstrument"`   // Max. number of DOM nodes to instrument. Valid values range from 0 to 100000.
}

func (me *AdditionalEventHandlers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"use_mouse_up_event_for_clicks": {
			Type:        hcl.TypeBool,
			Description: "Use mouseup event for clicks enabled/disabled",
			Optional:    true,
		},
		"click": {
			Type:        hcl.TypeBool,
			Description: "Click event handler enabled/disabled",
			Optional:    true,
		},
		"mouseup": {
			Type:        hcl.TypeBool,
			Description: "Mouseup event handler enabled/disabled",
			Optional:    true,
		},
		"blur": {
			Type:        hcl.TypeBool,
			Description: "Blur event handler enabled/disabled",
			Optional:    true,
		},
		"change": {
			Type:        hcl.TypeBool,
			Description: "Change event handler enabled/disabled",
			Optional:    true,
		},
		"to_string_method": {
			Type:        hcl.TypeBool,
			Description: "toString method enabled/disabled",
			Optional:    true,
		},
		"max_dom_nodes": {
			Type:        hcl.TypeInt,
			Description: "Max. number of DOM nodes to instrument. Valid values range from 0 to 100000.",
			Required:    true,
		},
	}
}

func (me *AdditionalEventHandlers) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"use_mouse_up_event_for_clicks": me.UseMouseUpEventForClicks,
		"click":                         me.ClickEventHandler,
		"mouseup":                       me.MouseUpEventHandler,
		"blur":                          me.BlurEventHandler,
		"change":                        me.ChangeEventHandler,
		"to_string_method":              me.ToStringMethod,
		"max_dom_nodes":                 me.MaxDomNodesToInstrument,
	})
}

func (me *AdditionalEventHandlers) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"use_mouse_up_event_for_clicks": &me.UseMouseUpEventForClicks,
		"click":                         &me.ClickEventHandler,
		"mouseup":                       &me.MouseUpEventHandler,
		"blur":                          &me.BlurEventHandler,
		"change":                        &me.ChangeEventHandler,
		"to_string_method":              &me.ToStringMethod,
		"max_dom_nodes":                 &me.MaxDomNodesToInstrument,
	})
}
