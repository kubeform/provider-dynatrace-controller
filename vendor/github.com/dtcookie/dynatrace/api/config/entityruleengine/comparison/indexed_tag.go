package comparison

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/indexed_tag"
	"github.com/dtcookie/dynatrace/api/config/entityruleengine/comparison/tag"
	"github.com/dtcookie/hcl"
)

// IndexedTag Comparison for `INDEXED_TAG` attributes.
type IndexedTag struct {
	BaseComparison
	Operator indexed_tag.Operator `json:"operator"`        // Operator of the comparison. You can reverse it by setting **negate** to `true`.  Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Value    *tag.Info            `json:"value,omitempty"` // Tag of a Dynatrace entity.
}

func (itc *IndexedTag) GetType() ComparisonBasicType {
	return ComparisonBasicTypes.IndexedTag
}

func (itc *IndexedTag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be INDEXED_TAG",
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
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (itc *IndexedTag) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(itc.Unknowns) > 0 {
		data, err := json.Marshal(itc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["negate"] = itc.Negate
	result["operator"] = string(itc.Operator)
	if itc.Value != nil {
		if marshalled, err := itc.Value.MarshalHCL(); err == nil {
			result["value"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (itc *IndexedTag) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), itc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &itc.Unknowns); err != nil {
			return err
		}
		delete(itc.Unknowns, "type")
		delete(itc.Unknowns, "negate")
		delete(itc.Unknowns, "operator")
		delete(itc.Unknowns, "value")
		if len(itc.Unknowns) == 0 {
			itc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		itc.Type = ComparisonBasicType(value.(string))
	}
	if _, value := decoder.GetChange("negate"); value != nil {
		itc.Negate = value.(bool)
	}
	if value, ok := decoder.GetOk("operator"); ok {
		itc.Operator = indexed_tag.Operator(value.(string))
	}
	if _, ok := decoder.GetOk("value.#"); ok {
		itc.Value = new(tag.Info)
		if err := itc.Value.UnmarshalHCL(hcl.NewDecoder(decoder, "value", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (itc *IndexedTag) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(itc.Unknowns) > 0 {
		for k, v := range itc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(itc.Negate)
		if err != nil {
			return nil, err
		}
		m["negate"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ComparisonBasicTypes.IndexedTag)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(&itc.Operator)
		if err != nil {
			return nil, err
		}
		m["operator"] = rawMessage
	}
	if itc.Value != nil {
		rawMessage, err := json.Marshal(itc.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (itc *IndexedTag) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	itc.Type = itc.GetType()
	if v, found := m["negate"]; found {
		if err := json.Unmarshal(v, &itc.Negate); err != nil {
			return err
		}
	}
	if v, found := m["operator"]; found {
		if err := json.Unmarshal(v, &itc.Operator); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &itc.Value); err != nil {
			return err
		}
	}
	delete(m, "negate")
	delete(m, "operator")
	delete(m, "value")
	delete(m, "type")
	if len(m) > 0 {
		itc.Unknowns = m
	}
	return nil
}
