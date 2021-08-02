package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/tag"
	"github.com/dtcookie/hcl"
)

// Tag Comparison for `TAG` attributes.
type Tag struct {
	BaseComparison
	Value    *tag.Info    `json:"value,omitempty"` // Tag of a Dynatrace entity.
	Operator tag.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

func (tc *Tag) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.Tag
}

func (tc *Tag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be TAG",
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
			Description: "Operator of the comparison. Possible values are EQUALS and TAG_KEY_EQUALS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "Tag of a Dynatrace entity",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(tag.Info).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (tc *Tag) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(tc.Unknowns) > 0 {
		data, err := json.Marshal(tc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = tc.Negate
	result["operator"] = string(tc.Operator)
	if tc.Value != nil {
		if marshalled, err := tc.Value.MarshalHCL(); err == nil {
			result["value"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (tc *Tag) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), tc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &tc.Unknowns); err != nil {
			return err
		}
		delete(tc.Unknowns, "type")
		delete(tc.Unknowns, "negate")
		delete(tc.Unknowns, "operator")
		delete(tc.Unknowns, "value")
		if len(tc.Unknowns) == 0 {
			tc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		tc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		tc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		tc.Operator = tag.Operator(value.(string))
	}
	if _, ok := decoder.GetOk("value.#"); ok {
		tc.Value = new(tag.Info)
		if err := tc.Value.UnmarshalHCL(hcl.NewDecoder(decoder, "value", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (tc *Tag) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(tc.Unknowns) > 0 {
		for k, v := range tc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(tc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ComparisonBasicTypes.Tag)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&tc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if tc.Value != nil {
		rawMessage, err := json.Marshal(tc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (tc *Tag) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	tc.Type = tc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &tc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &tc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &tc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		tc.Unknowns = m
	}
	return nil
}
