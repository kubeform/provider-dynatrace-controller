package aws

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// AWSConfigTag An AWS tag of the resource to be monitored.
type AWSConfigTag struct {
	Name     string                     `json:"name"`  // The key of the AWS tag.
	Value    string                     `json:"value"` // The value of the AWS tag.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (act *AWSConfigTag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "the key of the AWS tag.",
			Optional:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "the value of the AWS tag",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (act *AWSConfigTag) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &act.Name); err != nil {
			return err
		}
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &act.Value); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "value")
	if len(m) > 0 {
		act.Unknowns = m
	}
	return nil
}

// MarshalJSON provides custom JSON serialization
func (act *AWSConfigTag) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(act.Unknowns) > 0 {
		for k, v := range act.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(act.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(act.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	return json.Marshal(m)
}

func (act *AWSConfigTag) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(act.Unknowns) > 0 {
		data, err := json.Marshal(act.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = act.Name
	result["value"] = act.Value
	return result, nil
}

func (act *AWSConfigTag) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), act); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &act.Unknowns); err != nil {
			return err
		}
		delete(act.Unknowns, "name")
		delete(act.Unknowns, "value")
		if len(act.Unknowns) == 0 {
			act.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		act.Name = value.(string)
	}
	if value, ok := decoder.GetOk("value"); ok {
		act.Value = value.(string)
	}
	return nil
}
