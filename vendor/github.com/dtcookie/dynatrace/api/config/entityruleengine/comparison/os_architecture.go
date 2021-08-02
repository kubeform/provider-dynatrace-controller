package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/osarch"
	"github.com/dtcookie/hcl"
)

// OSArchitecture Comparison for `OS_ARCHITECTURE` attributes.
type OSArchitecture struct {
	BaseComparison
	Value    *osarch.Value   `json:"value,omitempty"` // The value to compare to.
	Operator osarch.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

func (oac *OSArchitecture) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.OSArchitecture
}

func (oac *OSArchitecture) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be OS_ARCHITECTURE",
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

			Type:        hcl.TypeString,
			Description: "The value to compare to. Possible values are ARM, IA64, PARISC, PPC, PPCLE, S390, SPARC, X86 and ZOS.",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (oac *OSArchitecture) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(oac.Unknowns) > 0 {
		data, err := json.Marshal(oac.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = oac.Negate
	result["operator"] = string(oac.Operator)
	if oac.Value != nil {
		result["value"] = oac.Value.String()
	}
	return result, nil
}

func (oac *OSArchitecture) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), oac); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &oac.Unknowns); err != nil {
			return err
		}
		delete(oac.Unknowns, "type")
		delete(oac.Unknowns, "negate")
		delete(oac.Unknowns, "operator")
		delete(oac.Unknowns, "value")
		if len(oac.Unknowns) == 0 {
			oac.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		oac.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		oac.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		oac.Operator = osarch.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		oac.Value = osarch.Value(value.(string)).Ref()
	}
	return nil
}

func (oac *OSArchitecture) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(oac.Unknowns) > 0 {
		for k, v := range oac.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(oac.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ComparisonBasicTypes.PaasType)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&oac.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if oac.Value != nil {
		rawMessage, err := json.Marshal(oac.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (oac *OSArchitecture) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	oac.Type = oac.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &oac.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &oac.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &oac.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		oac.Unknowns = m
	}
	return nil
}
