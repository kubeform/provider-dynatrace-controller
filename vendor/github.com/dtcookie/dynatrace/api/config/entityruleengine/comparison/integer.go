package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/integer"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// Integer Comparison for `INTEGER` attributes.
type Integer struct {
	BaseComparison
	Operator integer.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *int32           `json:"value,omitempty"` // The value to compare to.
}

func (ic *Integer) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.Integer
}

func (ic *Integer) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be INTEGER",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns the **begins with** into **does not begin with**",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator of the comparison. Possible values are EQUALS, EXISTS, GREATER_THAN, GREATER_THAN_OR_EQUAL, LOWER_THAN and LOWER_THAN_OR_EQUAL. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeInt,
			Description: "The value to compare to",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (ic *Integer) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ic.Unknowns) > 0 {
		data, err := json.Marshal(ic.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = ic.Negate
	result["operator"] = string(ic.Operator)
	if ic.Value != nil {
		result["value"] = int(opt.Int32(ic.Value))
	}
	return result, nil
}

func (ic *Integer) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ic); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ic.Unknowns); err != nil {
			return err
		}
		delete(ic.Unknowns, "type")
		delete(ic.Unknowns, "negate")
		delete(ic.Unknowns, "operator")
		delete(ic.Unknowns, "value")
		if len(ic.Unknowns) == 0 {
			ic.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		ic.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		ic.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		ic.Operator = integer.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		ic.Value = opt.NewInt32(int32(value.(int)))
	}
	return nil
}

func (ic *Integer) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(ic.Unknowns) > 0 {
		for k, v := range ic.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(ic.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ic.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&ic.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if ic.Value != nil {
		rawMessage, err := json.Marshal(ic.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (ic *Integer) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	ic.Type = ic.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &ic.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &ic.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &ic.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		ic.Unknowns = m
	}
	return nil
}
