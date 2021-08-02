package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// ValueCondition IBM integration bus label node name condition for which the value is captured.
type ValueCondition struct {
	Negate   *bool                      `json:"negate"`   // Negate the comparison.
	Operator Operator                   `json:"operator"` // Operator comparing the extracted value to the comparison value.
	Value    string                     `json:"value"`    // The value to compare to.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *ValueCondition) IsZero() bool {
	if opt.Bool(me.Negate) {
		return false
	}
	if len(me.Operator) > 0 {
		return false
	}
	if len(me.Value) > 0 {
		return false
	}
	if len(me.Unknowns) > 0 {
		return false
	}
	return true
}

func (me *ValueCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Negate the comparison",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator comparing the extracted value to the comparison value",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ValueCondition) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Negate != nil {
		result["negate"] = opt.Bool(me.Negate)
	}
	result["operator"] = string(me.Operator)
	result["value"] = me.Value
	return result, nil
}

func (me *ValueCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "negate")
		delete(me.Unknowns, "operator")
		delete(me.Unknowns, "value")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		me.Negate = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("operator"); ok {
		me.Operator = Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		me.Value = value.(string)
	}
	return nil
}

func (me *ValueCondition) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("negate", me.Negate); err != nil {
		return nil, err
	}
	if err := m.Marshal("operator", me.Operator); err != nil {
		return nil, err
	}
	if err := m.Marshal("value", me.Value); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *ValueCondition) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("negate", &me.Negate); err != nil {
		return err
	}
	if err := m.Unmarshal("operator", &me.Operator); err != nil {
		return err
	}
	if err := m.Unmarshal("value", &me.Value); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// Operator Operator comparing the extracted value to the comparison value.
type Operator string

// Operators offers the known enum values
var Operators = struct {
	BeginsWith Operator
	Contains   Operator
	EndsWith   Operator
	Equals     Operator
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
}
