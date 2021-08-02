package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/ip_address"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// IPAddress Comparison for `IP_ADDRESS` attributes.
type IPAddress struct {
	BaseComparison
	CaseSensitive *bool               `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or insensitive (`false`).
	Operator      ip_address.Operator `json:"operator"`                // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value         *string             `json:"value,omitempty"`         // The value to compare to.
}

func (iac *IPAddress) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.IPAddress
}

func (iac *IPAddress) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be IP_ADDRESS",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: " The comparison is case-sensitive (`true`) or insensitive (`false`)",
			Optional:    true,
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns the **begins with** into **does not begin with**",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Operator of the comparison. Possible values are BEGINS_WITH, CONTAINS, ENDS_WITH, EQUALS, EXISTS, IS_IP_IN_RANGE and REGEX_MATCHES. You can reverse it by setting **negate** to `true`",
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

func (iac *IPAddress) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(iac.Unknowns) > 0 {
		data, err := json.Marshal(iac.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = iac.Negate
	result["operator"] = string(iac.Operator)
	if iac.Value != nil {
		result["value"] = *iac.Value
	}
	result["case_sensitive"] = opt.Bool(iac.CaseSensitive)
	return result, nil
}

func (iac *IPAddress) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), iac); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &iac.Unknowns); err != nil {
			return err
		}
		delete(iac.Unknowns, "type")
		delete(iac.Unknowns, "negate")
		delete(iac.Unknowns, "operator")
		delete(iac.Unknowns, "value")
		delete(iac.Unknowns, "case_sensitive")
		if len(iac.Unknowns) == 0 {
			iac.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		iac.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		iac.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		iac.Operator = ip_address.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		iac.Value = opt.NewString(value.(string))
	}
	if _, value := decoder.GetChange("case_sensitive"); value != nil {
		iac.CaseSensitive = opt.NewBool(value.(bool))
	}

	return nil
}

func (iac *IPAddress) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(iac.Unknowns) > 0 {
		for k, v := range iac.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(iac.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(iac.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&iac.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if iac.Value != nil {
		rawMessage, err := json.Marshal(iac.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(iac.CaseSensitive))
		if err != nil {
			return nil, err
		}
		m["caseSensitive"] = rawMessage
	}
	return json.Marshal(m)
}

func (iac *IPAddress) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	iac.Type = iac.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &iac.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &iac.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &iac.Value); err != nil {
			return err
		}
	}
	if v, found := m["caseSensitive"]; found {
		if err := json.Unmarshal(v, &iac.CaseSensitive); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	delete(m, "caseSensitive")
	if len(m) > 0 {
		iac.Unknowns = m
	}
	return nil
}
