package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/tech"
	"github.com/dtcookie/hcl"
)

// SimpleHostTech Comparison for `SIMPLE_HOST_TECH` attributes.
type SimpleHostTech struct {
	BaseComparison
	Value    *tech.Host        `json:"value,omitempty"` // The value to compare to.
	Operator tech.HostOperator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

func (shtc *SimpleHostTech) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.SimpleHostTech
}

func (shtc *SimpleHostTech) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be SIMPLE_HOST_TECH",
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
				Schema: new(tech.Host).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (shtc *SimpleHostTech) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(shtc.Unknowns) > 0 {
		data, err := json.Marshal(shtc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = shtc.Negate
	result["operator"] = string(shtc.Operator)
	if shtc.Value != nil {
		if marshalled, err := shtc.Value.MarshalHCL(); err == nil {
			result["value"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (shtc *SimpleHostTech) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), shtc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &shtc.Unknowns); err != nil {
			return err
		}
		delete(shtc.Unknowns, "type")
		delete(shtc.Unknowns, "negate")
		delete(shtc.Unknowns, "operator")
		delete(shtc.Unknowns, "value")
		if len(shtc.Unknowns) == 0 {
			shtc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		shtc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		shtc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		shtc.Operator = tech.HostOperator(value.(string))
	}
	if _, ok := decoder.GetOk("value.#"); ok {
		shtc.Value = new(tech.Host)
		if err := shtc.Value.UnmarshalHCL(hcl.NewDecoder(decoder, "value", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (shtc *SimpleHostTech) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(shtc.Unknowns) > 0 {
		for k, v := range shtc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(shtc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(shtc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&shtc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if shtc.Value != nil {
		rawMessage, err := json.Marshal(shtc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (shtc *SimpleHostTech) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	shtc.Type = shtc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &shtc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &shtc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &shtc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		shtc.Unknowns = m
	}
	return nil
}
