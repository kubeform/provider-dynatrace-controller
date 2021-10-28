package monitors

import (
	"github.com/dtcookie/hcl"
)

// KeyPerformanceMetrics The key performance metrics configuration
type KeyPerformanceMetrics struct {
	LoadActionKPM LoadActionKPM `json:"loadActionKpm"` // Defines the key performance metric for load actions
	XHRActionKPM  XHRActionKPM  `json:"xhrActionKpm"`  // Defines the key performance metric for XHR actions
}

func (me *KeyPerformanceMetrics) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"load_action_kpm": {
			Type:        hcl.TypeString,
			Description: "Defines the key performance metric for load actions. Supported values are `VISUALLY_COMPLETE`, `SPEED_INDEX`, `USER_ACTION_DURATION`, `TIME_TO_FIRST_BYTE`, `HTML_DOWNLOADED`, `DOM_INTERACTIVE`, `LOAD_EVENT_START` and `LOAD_EVENT_END`.",
			Required:    true,
		},
		"xhr_action_kpm": {
			Type:        hcl.TypeString,
			Description: "Defines the key performance metric for XHR actions. Supported values are `VISUALLY_COMPLETE`, `USER_ACTION_DURATION`, `TIME_TO_FIRST_BYTE` and `RESPONSE_END`.",
			Required:    true,
		},
	}
}

func (me *KeyPerformanceMetrics) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["load_action_kpm"] = string(me.LoadActionKPM)
	result["xhr_action_kpm"] = string(me.XHRActionKPM)
	return result, nil
}

func (me *KeyPerformanceMetrics) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("load_action_kpm", &me.LoadActionKPM); err != nil {
		return err
	}
	if err := decoder.Decode("xhr_action_kpm", &me.XHRActionKPM); err != nil {
		return err
	}
	return nil
}

// LoadActionKPM Defines the key performance metric for load actions
type LoadActionKPM string

// LoadActionKPMs offers the known enum values
var LoadActionKPMs = struct {
	VisuallyComplete   LoadActionKPM
	SpeedIndex         LoadActionKPM
	UserActionDuration LoadActionKPM
	TimeToFirstByte    LoadActionKPM
	HTMLDownloaded     LoadActionKPM
	DOMInteractive     LoadActionKPM
	LoadEventStart     LoadActionKPM
	LoadEventEnd       LoadActionKPM
}{
	"VISUALLY_COMPLETE",
	"SPEED_INDEX",
	"USER_ACTION_DURATION",
	"TIME_TO_FIRST_BYTE",
	"HTML_DOWNLOADED",
	"DOM_INTERACTIVE",
	"LOAD_EVENT_START",
	"LOAD_EVENT_END",
}

// XHRActionKPM Defines the key performance metric for XHR actions
type XHRActionKPM string

// LoadActionKPMs offers the known enum values
var XHRActionKPMs = struct {
	VisuallyComplete   XHRActionKPM
	UserActionDuration XHRActionKPM
	TimeToFirstByte    XHRActionKPM
	ResponseEnd        XHRActionKPM
}{
	"VISUALLY_COMPLETE",
	"USER_ACTION_DURATION",
	"TIME_TO_FIRST_BYTE",
	"RESPONSE_END",
}
