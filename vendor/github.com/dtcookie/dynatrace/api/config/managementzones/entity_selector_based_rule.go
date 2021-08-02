package managementzones

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// EntitySelectorBasedRule is an entity-selector-based rule for management zone usage. It allows adding entities to a management zone via an entity selector
type EntitySelectorBasedRule struct {
	Enabled  *bool                      `json:"enabled,omitempty"` // The rule is enabled (`true`) or disabled (`false`)
	Selector string                     `json:"entitySelector"`    // The entity selector string, by which the entities are selected
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *EntitySelectorBasedRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The rule is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"selector": {
			Type:        hcl.TypeString,
			Description: "The entity selector string, by which the entities are selected",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *EntitySelectorBasedRule) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["selector"] = me.Selector
	result["enabled"] = opt.Bool(me.Enabled)
	return result, nil
}

func (me *EntitySelectorBasedRule) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "selector")
		delete(me.Unknowns, "enabled")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("selector"); ok {
		me.Selector = value.(string)
	}
	if _, value := decoder.GetChange("enabled"); value != nil {
		me.Enabled = opt.NewBool(value.(bool))
	}
	return nil
}

func (me *EntitySelectorBasedRule) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.Selector)
		if err != nil {
			return nil, err
		}
		m["entitySelector"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(me.Enabled))
		if err != nil {
			return nil, err
		}
		m["enabled"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *EntitySelectorBasedRule) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["enabled"]; found {
		if err := json.Unmarshal(v, &me.Enabled); err != nil {
			return err
		}
	}
	if v, found := m["entitySelector"]; found {
		if err := json.Unmarshal(v, &me.Selector); err != nil {
			return err
		}
	}
	delete(m, "entitySelector")
	delete(m, "enabled")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
