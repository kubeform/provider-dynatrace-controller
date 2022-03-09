package web

import "github.com/dtcookie/hcl"

// JavaScriptFrameworkSupport configures support of various JavaScript frameworks
type JavaScriptFrameworkSupport struct {
	Angular       bool `json:"angular"`       // AngularJS and Angular support enabled/disabled
	Dojo          bool `json:"dojo"`          // Dojo support enabled/disabled
	ExtJS         bool `json:"extJS"`         // ExtJS, Sencha Touch support enabled/disabled
	ICEfaces      bool `json:"icefaces"`      // ICEfaces support enabled/disabled
	JQuery        bool `json:"jQuery"`        // jQuery, Backbone.js support enabled/disabled
	MooTools      bool `json:"mooTools"`      // MooTools support enabled/disabled
	Prototype     bool `json:"prototype"`     // Prototype support enabled/disabled
	ActiveXObject bool `json:"activeXObject"` // ActiveXObject support enabled/disabled
}

func (me *JavaScriptFrameworkSupport) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"angular": {
			Type:        hcl.TypeBool,
			Description: "AngularJS and Angular support enabled/disabled",
			Optional:    true,
		},
		"dojo": {
			Type:        hcl.TypeBool,
			Description: "Dojo support enabled/disabled",
			Optional:    true,
		},
		"extjs": {
			Type:        hcl.TypeBool,
			Description: "ExtJS, Sencha Touch support enabled/disabled",
			Optional:    true,
		},
		"icefaces": {
			Type:        hcl.TypeBool,
			Description: "ICEfaces support enabled/disabled",
			Optional:    true,
		},
		"jquery": {
			Type:        hcl.TypeBool,
			Description: "jQuery, Backbone.js support enabled/disabled",
			Optional:    true,
		},
		"moo_tools": {
			Type:        hcl.TypeBool,
			Description: "MooTools support enabled/disabled",
			Optional:    true,
		},
		"prototype": {
			Type:        hcl.TypeBool,
			Description: "Prototype support enabled/disabled",
			Optional:    true,
		},
		"active_x_object": {
			Type:        hcl.TypeBool,
			Description: "ActiveXObject support enabled/disabled",
			Optional:    true,
		},
	}
}

func (me *JavaScriptFrameworkSupport) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"angular":         me.Angular,
		"dojo":            me.Dojo,
		"extjs":           me.ExtJS,
		"icefaces":        me.ICEfaces,
		"jquery":          me.JQuery,
		"moo_tools":       me.MooTools,
		"prototype":       me.Prototype,
		"active_x_object": me.ActiveXObject,
	})
}

func (me *JavaScriptFrameworkSupport) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"angular":         &me.Angular,
		"dojo":            &me.Dojo,
		"extjs":           &me.ExtJS,
		"icefaces":        &me.ICEfaces,
		"jquery":          &me.JQuery,
		"moo_tools":       &me.MooTools,
		"prototype":       &me.Prototype,
		"active_x_object": &me.ActiveXObject,
	})
}
