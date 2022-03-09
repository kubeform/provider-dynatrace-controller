package web

import "github.com/dtcookie/hcl"

type UserActionNamingRules []*UserActionNamingRule

func (me *UserActionNamingRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Description: "The settings of naming rule",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingRule).Schema()},
		},
	}
}

func (me UserActionNamingRules) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rule"] = entries
	}
	return result, nil
}

func (me *UserActionNamingRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("rule", me); err != nil {
		return err
	}
	return nil
}

// UserActionNamingRule The settings of naming rule
type UserActionNamingRule struct {
	Template        string                         `json:"template"`             // Naming pattern. Use Curly brackets `{}` to select placeholders
	Conditions      UserActionNamingRuleConditions `json:"conditions,omitempty"` // Defines the conditions when the naming rule should apply
	UseOrConditions bool                           `json:"useOrConditions"`      // If set to `true` the conditions will be connected by logical OR instead of logical AND
}

func (me *UserActionNamingRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"template": {
			Type:        hcl.TypeString,
			Description: "Naming pattern. Use Curly brackets `{}` to select placeholders",
			Required:    true,
		},
		"conditions": {
			Type:        hcl.TypeList,
			Description: "Defines the conditions when the naming rule should apply",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingRuleConditions).Schema()},
		},
		"use_or_conditions": {
			Type:        hcl.TypeBool,
			Description: "If set to `true` the conditions will be connected by logical OR instead of logical AND",
			Optional:    true,
		},
	}
}

func (me *UserActionNamingRule) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"template":          me.Template,
		"conditions":        me.Conditions,
		"use_or_conditions": me.UseOrConditions,
	})
}

func (me *UserActionNamingRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"template":          &me.Template,
		"conditions":        &me.Conditions,
		"use_or_conditions": &me.UseOrConditions,
	})
}
