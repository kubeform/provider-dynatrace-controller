package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/stringc"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// String Comparison for `STRING` attributes.
type String struct {
	BaseComparison
	CaseSensitive bool             `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or insensitive (`false`).
	Operator      stringc.Operator `json:"operator"`                // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value         *string          `json:"value,omitempty"`         // The value to compare to.
}

func (sc *String) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.String
}

func (sc *String) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be STRING",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns the **begins with** into **does not begin with**",
			Optional:    true,
		},
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "The comparison is case-sensitive (`true`) or insensitive (`false`)",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator of the comparison. Possible values are BEGINS_WITH, CONTAINS, ENDS_WITH, EQUALS, EXISTS and REGEX_MATCHES. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
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

func (sc *String) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(sc.Unknowns) > 0 {
		data, err := json.Marshal(sc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = sc.Negate
	result["operator"] = string(sc.Operator)
	result["case_sensitive"] = sc.CaseSensitive
	if sc.Value != nil {
		result["value"] = *sc.Value
	}
	return result, nil
}

func (sc *String) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), sc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &sc.Unknowns); err != nil {
			return err
		}
		delete(sc.Unknowns, "type")
		delete(sc.Unknowns, "negate")
		delete(sc.Unknowns, "operator")
		delete(sc.Unknowns, "value")
		if len(sc.Unknowns) == 0 {
			sc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		sc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		sc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		sc.Operator = stringc.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		sc.Value = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("case_sensitive"); ok {
		sc.CaseSensitive = value.(bool)
	}

	return nil
}

func (sc *String) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(sc.Unknowns) > 0 {
		for k, v := range sc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(sc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ComparisonBasicTypes.String)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&sc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if sc.Value != nil {
		rawMessage, err := json.Marshal(sc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	if sc.Operator != stringc.Operators.Exists {
		rawMessage, err := json.Marshal(sc.CaseSensitive)
		if err != nil {
			return nil, err
		}
		m["caseSensitive"] = rawMessage
	}
	return json.Marshal(m)
}

func (sc *String) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	sc.Type = sc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &sc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &sc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &sc.Value); err != nil {
			return err
		}
	}
	if v, found := m["caseSensitive"]; found {
		if err := json.Unmarshal(v, &sc.CaseSensitive); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	delete(m, "caseSensitive")
	if len(m) > 0 {
		sc.Unknowns = m
	}
	return nil
}
