package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/tech"
	"github.com/dtcookie/hcl"
)

// SimpleTech Comparison for `SIMPLE_TECH` attributes.
type SimpleTech struct {
	BaseComparison
	Operator tech.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *tech.Simple  `json:"value,omitempty"` // The value to compare to.
}

func (stc *SimpleTech) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.SimpleTech
}

func (stc *SimpleTech) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be SIMPLE_TECH",
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
			Description: "Operator of the comparison. Possible values are EQUALS and EXISTS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "The value to compare to",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(tech.Simple).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (stc *SimpleTech) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(stc.Unknowns) > 0 {
		data, err := json.Marshal(stc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = stc.Negate
	result["operator"] = string(stc.Operator)
	if stc.Value != nil {
		if marshalled, err := stc.Value.MarshalHCL(); err == nil {
			result["value"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (stc *SimpleTech) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), stc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &stc.Unknowns); err != nil {
			return err
		}
		delete(stc.Unknowns, "type")
		delete(stc.Unknowns, "negate")
		delete(stc.Unknowns, "operator")
		delete(stc.Unknowns, "value")
		if len(stc.Unknowns) == 0 {
			stc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		stc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		stc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		stc.Operator = tech.Operator(value.(string))
	}
	if _, ok := decoder.GetOk("value.#"); ok {
		stc.Value = new(tech.Simple)
		if err := stc.Value.UnmarshalHCL(hcl.NewDecoder(decoder, "value", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (stc *SimpleTech) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(stc.Unknowns) > 0 {
		for k, v := range stc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(stc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(stc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&stc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if stc.Value != nil {
		rawMessage, err := json.Marshal(stc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (stc *SimpleTech) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	stc.Type = stc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &stc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &stc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &stc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		stc.Unknowns = m
	}
	return nil
}
