package api

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// EntityRef is the short representation of a Dynatrace entity
type EntityRef struct {
	ID          string                     `json:"id"`                    // the ID of the Dynatrace entity
	Name        *string                    `json:"name,omitempty"`        // the name of the Dynatrace entity
	Description *string                    `json:"description,omitempty"` // a short description of the Dynatrace entity
	Unknowns    map[string]json.RawMessage `json:"-"`
}

func (me *EntityRef) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Description: "the ID of the Dynatrace entity",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "the name of the Dynatrace entity",
			Optional:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "a short description of the Dynatrace entity",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *EntityRef) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["id"] = me.ID
	if me.Name != nil {
		result["name"] = opt.String(me.Name)
	}
	if me.Description != nil {
		result["description"] = opt.String(me.Description)
	}
	return result, nil
}

func (me *EntityRef) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "id")
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "description")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}

	if value, ok := decoder.GetOk("id"); ok {
		me.ID = value.(string)
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("description"); ok {
		me.Description = opt.NewString(value.(string))
	}
	return nil
}

func (me *EntityRef) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	if me.Name != nil {
		rawMessage, err := json.Marshal(me.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	if me.Description != nil {
		rawMessage, err := json.Marshal(me.Description)
		if err != nil {
			return nil, err
		}
		m["description"] = rawMessage
	}
	return json.Marshal(m)
}
