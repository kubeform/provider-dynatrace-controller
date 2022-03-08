package web

import "github.com/dtcookie/hcl"

type CustomErrorRules []*CustomErrorRule

func (me *CustomErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Description: "Configuration of the custom error in the web application",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(CustomErrorRule).Schema()},
		},
	}
}

func (me CustomErrorRules) MarshalHCL() (map[string]interface{}, error) {
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

// CustomErrorRule represents configuration of the custom error in the web application
type CustomErrorRule struct {
	KeyPattern     *string                      `json:"keyPattern,omitempty"`   // The key of the error to look for
	KeyMatcher     *CustomErrorRuleKeyMatcher   `json:"keyMatcher,omitempty"`   // The matching operation for the **keyPattern**. Possible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH` and `EQUALS`.
	ValuePattern   *string                      `json:"valuePattern,omitempty"` // The value of the error to look for
	ValueMatcher   *CustomErrorRuleValueMatcher `json:"valueMatcher,omitempty"` // The matching operation for the **valuePattern**. Possible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH` and `EQUALS`.
	Capture        bool                         `json:"capture"`                // Capture (`true`) or ignore (`false`) the error
	ImpactApdex    bool                         `json:"impactApdex"`            // Include (`true`) or exclude (`false`) the error in Apdex calculation
	CustomAlerting bool                         `json:"customAlerting"`         // Include (`true`) or exclude (`false`) the error in Davis AI [problem detection and analysis](https://dt-url.net/a963kd2)
}

func (me *CustomErrorRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key_pattern": {
			Type:        hcl.TypeString,
			Description: "The key of the error to look for",
			Optional:    true,
		},
		"key_matcher": {
			Type:        hcl.TypeString,
			Description: "The matching operation for the **keyPattern**. Possible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH` and `EQUALS`",
			Optional:    true,
		},
		"value_pattern": {
			Type:        hcl.TypeString,
			Description: "The value of the error to look for",
			Optional:    true,
		},
		"value_matcher": {
			Type:        hcl.TypeString,
			Description: "The matching operation for the **valuePattern**. Possible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH` and `EQUALS`.",
			Optional:    true,
		},
		"capture": {
			Type:        hcl.TypeBool,
			Description: "Capture (`true`) or ignore (`false`) the error",
			Optional:    true,
		},
		"impact_apdex": {
			Type:        hcl.TypeBool,
			Description: "Include (`true`) or exclude (`false`) the error in Apdex calculation",
			Optional:    true,
		},
		"custom_alerting": {
			Type:        hcl.TypeBool,
			Description: "Include (`true`) or exclude (`false`) the error in Davis AI [problem detection and analysis](https://dt-url.net/a963kd2)",
			Optional:    true,
		},
	}
}

func (me *CustomErrorRule) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"key_pattern":     me.KeyPattern,
		"key_matcher":     me.KeyMatcher,
		"value_pattern":   me.ValuePattern,
		"value_matcher":   me.ValueMatcher,
		"capture":         me.Capture,
		"impact_apdex":    me.ImpactApdex,
		"custom_alerting": me.CustomAlerting,
	})
}

func (me *CustomErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"key_pattern":     &me.KeyPattern,
		"key_matcher":     &me.KeyMatcher,
		"value_pattern":   &me.ValuePattern,
		"value_matcher":   &me.ValueMatcher,
		"capture":         &me.Capture,
		"impact_apdex":    &me.ImpactApdex,
		"custom_alerting": &me.CustomAlerting,
	})
}
