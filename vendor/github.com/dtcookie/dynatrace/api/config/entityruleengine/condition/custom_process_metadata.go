package condition

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// CustomProcessMetadata The key for dynamic attributes of the `PROCESS_CUSTOM_METADATA_KEY` type.
type CustomProcessMetadata struct {
	BaseConditionKey
	DynamicKey *CustomProcessMetadataKey  `json:"dynamicKey"` // The key of the attribute, which need dynamic keys.  Not applicable otherwise, as the attibute itself acts as a key.
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (cpmck *CustomProcessMetadata) GetType() *ConditionKeyType {
	return &ConditionKeyTypes.ProcessCustomMetadataKey
}

func (cpmck *CustomProcessMetadata) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute": {
			Type:        hcl.TypeString,
			Description: "The attribute to be used for comparision",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be PROCESS_CUSTOM_METADATA_KEY",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"dynamic_key": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "The key of the attribute, which need dynamic keys. Not applicable otherwise, as the attibute itself acts as a key",
			Required:    true,
			Elem: &hcl.Resource{
				Schema: new(CustomProcessMetadataKey).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (cpmck *CustomProcessMetadata) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(cpmck.Unknowns) > 0 {
		data, err := json.Marshal(cpmck.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["attribute"] = string(cpmck.Attribute)
	if marshalled, err := cpmck.DynamicKey.MarshalHCL(); err == nil {
		result["dynamic_key"] = []interface{}{marshalled}
	} else {
		return nil, err
	}
	return result, nil
}

func (cpmck *CustomProcessMetadata) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), cpmck); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &cpmck.Unknowns); err != nil {
			return err
		}
		delete(cpmck.Unknowns, "attribute")
		delete(cpmck.Unknowns, "dynamic_key")
		delete(cpmck.Unknowns, "type")
		if len(cpmck.Unknowns) == 0 {
			cpmck.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("attribute"); ok {
		cpmck.Attribute = Attribute(value.(string))
	}
	if _, ok := decoder.GetOk("dynamic_key.#"); ok {
		cpmck.DynamicKey = new(CustomProcessMetadataKey)
		if err := cpmck.DynamicKey.UnmarshalHCL(hcl.NewDecoder(decoder, "dynamic_key", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (cpmck *CustomProcessMetadata) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(cpmck.Unknowns) > 0 {
		for k, v := range cpmck.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(cpmck.Attribute)
		if err != nil {
			return nil, err
		}
		m["attribute"] = rawMessage
	}
	if cpmck.GetType() != nil {
		rawMessage, err := json.Marshal(ConditionKeyTypes.ProcessCustomMetadataKey)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	if cpmck.DynamicKey != nil {
		rawMessage, err := json.Marshal(cpmck.DynamicKey)
		if err != nil {
			return nil, err
		}
		m["dynamicKey"] = rawMessage
	}
	return json.Marshal(m)
}

func (cpmck *CustomProcessMetadata) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	cpmck.Type = cpmck.GetType()
	if v, found := m["attribute"]; found {
		if err := json.Unmarshal(v, &cpmck.Attribute); err != nil {
			return err
		}
	}
	if v, found := m["dynamicKey"]; found {
		if err := json.Unmarshal(v, &cpmck.DynamicKey); err != nil {
			return err
		}
	}
	delete(m, "attribute")
	delete(m, "dynamicKey")
	delete(m, "type")
	if len(m) > 0 {
		cpmck.Unknowns = m
	}
	return nil
}
