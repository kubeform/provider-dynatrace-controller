package web

import "github.com/dtcookie/hcl"

type JavaScriptInjectionRules []*JavaScriptInjectionRule

func (me *JavaScriptInjectionRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Description: "Java script injection rule",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(JavaScriptInjectionRule).Schema()},
		},
	}
}

func (me JavaScriptInjectionRules) MarshalHCL() (map[string]interface{}, error) {
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

func (me *JavaScriptInjectionRules) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("rule", me); err != nil {
		return err
	}
	return nil
}

type JavaScriptInjectionRule struct {
	Enabled     bool             `json:"enabled"`               // The enable or disable rule of the java script injection
	URLOperator URLOperator      `json:"urlOperator"`           // The url operator of the java script injection. Possible values are `ALL_PAGES`, `CONTAINS`, `ENDS_WITH`, `EQUALS` and `STARTS_WITH`.
	URLPattern  *string          `json:"urlPattern,omitempty"`  // The url pattern of the java script injection
	Rule        JSInjectionRule  `json:"rule"`                  // The url rule of the java script injection. Possible values are `AFTER_SPECIFIC_HTML`, `AUTOMATIC_INJECTION`, `BEFORE_SPECIFIC_HTML` and `DO_NOT_INJECT`.
	HTMLPattern *string          `json:"htmlPattern,omitempty"` // The HTML pattern of the java script injection
	Target      *InjectionTarget `json:"target,omitempty"`      // The target against which the rule of the java script injection should be matched. Possible values are `PAGE_QUERY` and `URL`.
}

func (me *JavaScriptInjectionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "`fetch()` request capture enabled/disabled",
			Optional:    true,
		},
		"url_operator": {
			Type:        hcl.TypeString,
			Description: "The url operator of the java script injection. Possible values are `ALL_PAGES`, `CONTAINS`, `ENDS_WITH`, `EQUALS` and `STARTS_WITH`.",
			Required:    true,
		},
		"url_pattern": {
			Type:        hcl.TypeString,
			Description: "The url pattern of the java script injection",
			Optional:    true,
		},
		"rule": {
			Type:        hcl.TypeString,
			Description: "The url rule of the java script injection. Possible values are `AFTER_SPECIFIC_HTML`, `AUTOMATIC_INJECTION`, `BEFORE_SPECIFIC_HTML` and `DO_NOT_INJECT`.",
			Required:    true,
		},
		"html_pattern": {
			Type:        hcl.TypeString,
			Description: "The HTML pattern of the java script injection",
			Optional:    true,
		},
		"target": {
			Type:        hcl.TypeString,
			Description: "The target against which the rule of the java script injection should be matched. Possible values are `PAGE_QUERY` and `URL`.",
			Optional:    true,
		},
	}
}

func (me *JavaScriptInjectionRule) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"enabled":      me.Enabled,
		"url_operator": me.URLOperator,
		"url_pattern":  me.URLPattern,
		"rule":         me.Rule,
		"html_pattern": me.HTMLPattern,
		"target":       me.Target,
	})
}

func (me *JavaScriptInjectionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"enabled":      &me.Enabled,
		"url_operator": &me.URLOperator,
		"url_pattern":  &me.URLPattern,
		"rule":         &me.Rule,
		"html_pattern": &me.HTMLPattern,
		"target":       &me.Target,
	})
}
