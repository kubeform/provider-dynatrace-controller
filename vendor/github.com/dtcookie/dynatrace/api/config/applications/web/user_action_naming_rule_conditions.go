package web

import "github.com/dtcookie/hcl"

type UserActionNamingRuleConditions []*UserActionNamingRuleCondition

func (me *UserActionNamingRuleConditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Description: "Defines the conditions when the naming rule should apply",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingRuleCondition).Schema()},
		},
	}
}

func (me UserActionNamingRuleConditions) MarshalHCL() (map[string]interface{}, error) {
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
		result["condition"] = entries
	}
	return result, nil
}

func (me *UserActionNamingRuleConditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("condition", me); err != nil {
		return err
	}
	return nil
}

// UserActionNamingRuleCondition The settings of conditions for user action naming
type UserActionNamingRuleCondition struct {
	Operand1 string   `json:"operand1"`           // Must be a defined placeholder wrapped in curly braces
	Operand2 *string  `json:"operand2,omitempty"` // Must be null if operator is `IS_EMPTY`, a regex if operator is `MATCHES_REGULAR_ERPRESSION`. In all other cases the value can be a freetext or a placeholder wrapped in curly braces
	Operator Operator `json:"operator"`           // The operator of the condition. Possible values are `CONTAINS`, `ENDS_WITH`, `EQUALS`, `IS_EMPTY`, `IS_NOT_EMPTY`, `MATCHES_REGULAR_EXPRESSION`, `NOT_CONTAINS`, `NOT_ENDS_WITH`, `NOT_EQUALS`, `NOT_MATCHES_REGULAR_EXPRESSION`, `NOT_STARTS_WITH` and `STARTS_WITH`.
}

func (me *UserActionNamingRuleCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"operand1": {
			Type:        hcl.TypeString,
			Description: "Must be a defined placeholder wrapped in curly braces",
			Required:    true,
		},
		"operand2": {
			Type:        hcl.TypeString,
			Description: "Must be null if operator is `IS_EMPTY`, a regex if operator is `MATCHES_REGULAR_ERPRESSION`. In all other cases the value can be a freetext or a placeholder wrapped in curly braces",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "The operator of the condition. Possible values are `CONTAINS`, `ENDS_WITH`, `EQUALS`, `IS_EMPTY`, `IS_NOT_EMPTY`, `MATCHES_REGULAR_EXPRESSION`, `NOT_CONTAINS`, `NOT_ENDS_WITH`, `NOT_EQUALS`, `NOT_MATCHES_REGULAR_EXPRESSION`, `NOT_STARTS_WITH` and `STARTS_WITH`.",
			Required:    true,
		},
	}
}

func (me *UserActionNamingRuleCondition) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"operand1": me.Operand1,
		"operand2": me.Operand2,
		"operator": me.Operator,
	})
}

func (me *UserActionNamingRuleCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"operand1": &me.Operand1,
		"operand2": &me.Operand2,
		"operator": &me.Operator,
	})
}
