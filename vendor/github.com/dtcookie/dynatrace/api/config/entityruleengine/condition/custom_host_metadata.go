package condition

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// CustomHostMetadata The key for dynamic attributes of the `HOST_CUSTOM_METADATA_KEY` type.
type CustomHostMetadata struct {
	BaseConditionKey
	DynamicKey *CustomHostMetadataKey `json:"dynamicKey"` // The key of the attribute, which need dynamic keys.  Not applicable otherwise, as the attibute itself acts as a key.
}

func (chmck *CustomHostMetadata) GetType() *ConditionKeyType {
	return &ConditionKeyTypes.HostCustomMetadataKey
}

func (chmck *CustomHostMetadata) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"attribute": {
			Type:        hcl.TypeString,
			Description: "The attribute to be used for comparision",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "if specified, needs to be HOST_CUSTOM_METADATA_KEY",
			Optional:    true,
			Deprecated:  "The value of the attribute type is implicit, therefore shouldn't get specified",
		},
		"dynamic_key": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "The key of the attribute, which need dynamic keys. Not applicable otherwise, as the attibute itself acts as a key",
			Required:    true,
			Elem: &hcl.Resource{
				Schema: new(CustomHostMetadataKey).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (chmck *CustomHostMetadata) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(chmck.Unknowns) > 0 {
		data, err := json.Marshal(chmck.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["attribute"] = string(chmck.Attribute)
	if marshalled, err := chmck.DynamicKey.MarshalHCL(); err == nil {
		result["dynamic_key"] = []interface{}{marshalled}
	} else {
		return nil, err
	}
	return result, nil
}

func (chmck *CustomHostMetadata) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), chmck); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &chmck.Unknowns); err != nil {
			return err
		}
		delete(chmck.Unknowns, "attribute")
		delete(chmck.Unknowns, "dynamic_key")
		delete(chmck.Unknowns, "type")
		if len(chmck.Unknowns) == 0 {
			chmck.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("attribute"); ok {
		chmck.Attribute = Attribute(value.(string))
	}
	if _, ok := decoder.GetOk("dynamic_key.#"); ok {
		chmck.DynamicKey = new(CustomHostMetadataKey)
		if err := chmck.DynamicKey.UnmarshalHCL(hcl.NewDecoder(decoder, "dynamic_key", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (chmck *CustomHostMetadata) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(chmck.Unknowns) > 0 {
		for k, v := range chmck.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(chmck.Attribute)
		if err != nil {
			return nil, err
		}
		m["attribute"] = rawMessage
	}
	if chmck.GetType() != nil {
		rawMessage, err := json.Marshal(ConditionKeyTypes.HostCustomMetadataKey)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	if chmck.DynamicKey != nil {
		rawMessage, err := json.Marshal(chmck.DynamicKey)
		if err != nil {
			return nil, err
		}
		m["dynamicKey"] = rawMessage
	}
	return json.Marshal(m)
}

func (chmck *CustomHostMetadata) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	chmck.Type = chmck.GetType()
	if v, found := m["attribute"]; found {
		if err := json.Unmarshal(v, &chmck.Attribute); err != nil {
			return err
		}
	}
	if v, found := m["dynamicKey"]; found {
		if err := json.Unmarshal(v, &chmck.DynamicKey); err != nil {
			return err
		}
	}
	delete(m, "attribute")
	delete(m, "dynamicKey")
	delete(m, "type")
	if len(m) > 0 {
		chmck.Unknowns = m
	}
	return nil
}
