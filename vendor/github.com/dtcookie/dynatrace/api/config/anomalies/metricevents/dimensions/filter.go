package dimensions

import (
	"github.com/dtcookie/hcl"
)

// Filter A filter for a string value based on the given operator.
type Filter struct {
	Operator Operator `json:"operator"` // The operator to match on.
	Value    string   `json:"value"`    // The value to match on.
}

func (me *Filter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"operator": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The operator to match on",
		},
		"value": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The value to match on",
		},
	}
}

func (me *Filter) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["operator"] = string(me.Operator)
	result["value"] = me.Value
	return result, nil
}

func (me *Filter) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("operator"); ok {
		me.Operator = Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		me.Value = value.(string)
	}
	return nil
}

// Operator The operator to match on.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	Equals Operator
}{
	"EQUALS",
}
