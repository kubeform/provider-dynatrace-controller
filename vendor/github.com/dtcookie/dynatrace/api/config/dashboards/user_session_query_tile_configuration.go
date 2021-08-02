package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// UserSessionQueryTileConfiguration Configuration of a User session query visualization tile
type UserSessionQueryTileConfiguration struct {
	HasAxisBucketing *bool                      `json:"hasAxisBucketing,omitempty"` // The axis bucketing when enabled groups similar series in the same virtual axis
	Unknowns         map[string]json.RawMessage `json:"-"`
}

func (me *UserSessionQueryTileConfiguration) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"has_axis_bucketing": {
			Type:        hcl.TypeBool,
			Description: "The axis bucketing when enabled groups similar series in the same virtual axis",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *UserSessionQueryTileConfiguration) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.HasAxisBucketing != nil {
		result["has_axis_bucketing"] = opt.Bool(me.HasAxisBucketing)
	}
	return result, nil
}

func (me *UserSessionQueryTileConfiguration) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "has_axis_bucketing")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, value := decoder.GetChange("has_axis_bucketing"); value != nil {
		me.HasAxisBucketing = opt.NewBool(value.(bool))
	}
	return nil
}

func (me *UserSessionQueryTileConfiguration) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	if me.HasAxisBucketing != nil {
		rawMessage, err := json.Marshal(me.HasAxisBucketing)
		if err != nil {
			return nil, err
		}
		m["hasAxisBucketing"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *UserSessionQueryTileConfiguration) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["hasAxisBucketing"]; found {
		if err := json.Unmarshal(v, &me.HasAxisBucketing); err != nil {
			return err
		}
	}
	delete(m, "hasAxisBucketing")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
