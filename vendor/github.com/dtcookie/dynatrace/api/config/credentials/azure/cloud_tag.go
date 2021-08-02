package azure

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// CloudTag A cloud tag.
type CloudTag struct {
	Value    *string                    `json:"value,omitempty"` // The value of the tag. If set to `null`, then resources with any value of the tag are monitored.
	Name     *string                    `json:"name,omitempty"`  // The name of the tag.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (ct *CloudTag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the tag.   If set to `null`, then resources with any value of the tag are monitored.",
			Optional:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the tag.",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (ct *CloudTag) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(ct.Unknowns) > 0 {
		for k, v := range ct.Unknowns {
			m[k] = v
		}
	}
	if ct.Value != nil {
		rawMessage, err := json.Marshal(ct.Value)
		if err != nil {
			return nil, err
		}
		m["value"] = rawMessage
	}
	if ct.Name != nil {
		rawMessage, err := json.Marshal(ct.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	return json.Marshal(m)
}

func (ct *CloudTag) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["value"]; found {
		if err := json.Unmarshal(v, &ct.Value); err != nil {
			return err
		}
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &ct.Name); err != nil {
			return err
		}
	}
	delete(m, "value")
	delete(m, "name")

	if len(m) > 0 {
		ct.Unknowns = m
	}
	return nil
}

func (ct *CloudTag) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ct.Unknowns) > 0 {
		data, err := json.Marshal(ct.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}

	if ct.Value != nil {
		result["value"] = *ct.Value
	}
	if ct.Name != nil {
		result["name"] = *ct.Name
	}
	return result, nil
}

func (ct *CloudTag) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ct); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ct.Unknowns); err != nil {
			return err
		}
		delete(ct.Unknowns, "value")
		delete(ct.Unknowns, "name")
		if len(ct.Unknowns) == 0 {
			ct.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("value"); ok {
		ct.Value = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("name"); ok {
		ct.Name = opt.NewString(value.(string))
	}
	return nil
}
