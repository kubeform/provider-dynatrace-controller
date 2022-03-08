package comparison

import (
	"encoding/json"
	"log"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/entity_id"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// EntityID Comparison for `ENTITY_ID` attributes.
type EntityID struct {
	BaseComparison
	Value    *string            `json:"value,omitempty"` // The value to compare to.
	Operator entity_id.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
}

func (eic *EntityID) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.EntityID
}

func (eic *EntityID) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be ENTITY_ID",
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
			Description: "Currently only EQUALS is supported. You can reverse it by setting **negate** to `true`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to compare to",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (eic *EntityID) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(eic.Unknowns) > 0 {
		data, err := json.Marshal(eic.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = eic.Negate
	result["operator"] = string(eic.Operator)
	if eic.Value != nil {
		result["value"] = *eic.Value
	}
	return result, nil
}

func (eic *EntityID) UnmarshalHCL(decoder hcl.Decoder) error {
	log.Println("UnmarshalHCL")
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), eic); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &eic.Unknowns); err != nil {
			return err
		}
		delete(eic.Unknowns, "type")
		delete(eic.Unknowns, "negate")
		delete(eic.Unknowns, "operator")
		delete(eic.Unknowns, "value")
		if len(eic.Unknowns) == 0 {
			eic.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("value"); ok {
		eic.Value = opt.NewString(value.(string))
	}
	eic.Type = ComparisonBasicTypes.EntityID
	if _, value := decoder.GetChange("negate"); value != nil {
		eic.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		eic.Operator = entity_id.Operator(value.(string))
	}
	return nil
}

func (eic *EntityID) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(eic.Unknowns) > 0 {
		for k, v := range eic.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(eic.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(eic.GetType())
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&eic.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if eic.Value != nil {
		rawMessage, err := json.Marshal(eic.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (eic *EntityID) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	eic.Type = eic.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &eic.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &eic.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &eic.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		eic.Unknowns = m
	}
	return nil
}
