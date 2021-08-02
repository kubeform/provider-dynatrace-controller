package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/azure_sku"
	"github.com/dtcookie/hcl"
)

// AzureSku Comparison for `AZURE_SKU` attributes.
type AzureSku struct {
	BaseComparison
	Value    *azure_sku.Value   `json:"value,omitempty"` // The value to compare to.
	Operator azure_sku.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

func (asc *AzureSku) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.AzureSku
}

func (asc *AzureSku) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be AZURE_SKU",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"negate": {
			Type:        hcl.TypeBool,
			Description: "Reverses the operator. For example it turns EQUALS into DOES NOT EQUAL",
			Optional:    true,
		},
		"operator": {
			Type:        hcl.TypeString,
			Description: "Either EQUALS or EXISTS. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to. Possible values are BASIC, DYNAMIC, FREE, PREMIUM, SHARED and STANDARD.",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (asc *AzureSku) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(asc.Unknowns) > 0 {
		data, err := json.Marshal(asc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = asc.Negate
	result["operator"] = string(asc.Operator)
	if asc.Value != nil {
		result["value"] = asc.Value.String()
	}
	return result, nil
}

func (asc *AzureSku) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), asc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &asc.Unknowns); err != nil {
			return err
		}
		delete(asc.Unknowns, "type")
		delete(asc.Unknowns, "negate")
		delete(asc.Unknowns, "operator")
		delete(asc.Unknowns, "value")
		if len(asc.Unknowns) == 0 {
			asc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		asc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		asc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		asc.Operator = azure_sku.Operator(value.(string))
	}
	if value, ok := decoder.GetOk("value"); ok {
		asc.Value = azure_sku.Value(value.(string)).Ref()
	}
	return nil
}

func (asc *AzureSku) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(asc.Unknowns) > 0 {
		for k, v := range asc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(asc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(asc.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&asc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if asc.Value != nil {
		rawMessage, err := json.Marshal(asc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (asc *AzureSku) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	asc.Type = asc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &asc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &asc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &asc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		asc.Unknowns = m
	}
	return nil
}
