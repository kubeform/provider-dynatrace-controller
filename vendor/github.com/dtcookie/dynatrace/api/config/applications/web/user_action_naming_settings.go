package web

import "github.com/dtcookie/hcl"

// UserActionNamingSettings The settings of user action naming
type UserActionNamingSettings struct {
	Placeholders               UserActionNamingPlaceholders `json:"placeholders"`                     // User action placeholders
	LoadActionNamingRules      UserActionNamingRules        `json:"loadActionNamingRules"`            // User action naming rules for loading actions
	XHRActionNamingRules       UserActionNamingRules        `json:"xhrActionNamingRules"`             // User action naming rules for XHR actions
	CustomActionNamingRules    UserActionNamingRules        `json:"customActionNamingRules"`          // User action naming rules for custom actions
	IgnoreCase                 bool                         `json:"ignoreCase"`                       // Case insensitive naming
	UseFirstDetectedLoadAction bool                         `json:"useFirstDetectedLoadAction"`       // First load action found under an XHR action should be used when true. Else the deepest one under the xhr action is used
	SplitUserActionsByDomain   bool                         `json:"splitUserActionsByDomain"`         // Deactivate this setting if different domains should not result in separate user actions
	QueryParameterCleanups     []string                     `json:"queryParameterCleanups,omitempty"` // List of parameters that should be removed from the query before using the query in the user action name
}

func (me *UserActionNamingSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"placeholders": {
			Type:        hcl.TypeList,
			Description: "User action placeholders",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingPlaceholders).Schema()},
		},
		"load_action_naming_rules": {
			Type:        hcl.TypeList,
			Description: "User action naming rules for loading actions",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingRules).Schema()},
		},
		"xhr_action_naming_rules": {
			Type:        hcl.TypeList,
			Description: "User action naming rules for XHR actions",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingRules).Schema()},
		},
		"custom_action_naming_rules": {
			Type:        hcl.TypeList,
			Description: "User action naming rules for custom actions",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingRules).Schema()},
		},
		"ignore_case": {
			Type:        hcl.TypeBool,
			Description: "Case insensitive naming",
			Optional:    true,
		},
		"use_first_detected_load_action": {
			Type:        hcl.TypeBool,
			Description: "First load action found under an XHR action should be used when true. Else the deepest one under the xhr action is used",
			Optional:    true,
		},
		"split_user_actions_by_domain": {
			Type:        hcl.TypeBool,
			Description: "Deactivate this setting if different domains should not result in separate user actions",
			Optional:    true,
		},
		"query_parameter_cleanups": {
			Type:        hcl.TypeSet,
			Description: "User action naming rules for custom actions",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *UserActionNamingSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"placeholders":                   me.Placeholders,
		"load_action_naming_rules":       me.LoadActionNamingRules,
		"xhr_action_naming_rules":        me.XHRActionNamingRules,
		"custom_action_naming_rules":     me.CustomActionNamingRules,
		"ignore_case":                    me.IgnoreCase,
		"use_first_detected_load_action": me.UseFirstDetectedLoadAction,
		"split_user_actions_by_domain":   me.SplitUserActionsByDomain,
		"query_parameter_cleanups":       me.QueryParameterCleanups,
	})
}

func (me *UserActionNamingSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeAll(map[string]interface{}{
		"placeholders":                   &me.Placeholders,
		"load_action_naming_rules":       &me.LoadActionNamingRules,
		"xhr_action_naming_rules":        &me.XHRActionNamingRules,
		"custom_action_naming_rules":     &me.CustomActionNamingRules,
		"ignore_case":                    &me.IgnoreCase,
		"use_first_detected_load_action": &me.UseFirstDetectedLoadAction,
		"split_user_actions_by_domain":   &me.SplitUserActionsByDomain,
		"query_parameter_cleanups":       &me.QueryParameterCleanups,
	}); err != nil {
		return err
	}

	if me.Placeholders == nil {
		me.Placeholders = UserActionNamingPlaceholders{}
	}
	if me.LoadActionNamingRules == nil {
		me.LoadActionNamingRules = UserActionNamingRules{}
	}
	if me.XHRActionNamingRules == nil {
		me.XHRActionNamingRules = UserActionNamingRules{}
	}
	if me.CustomActionNamingRules == nil {
		me.CustomActionNamingRules = UserActionNamingRules{}
	}
	return nil
}
