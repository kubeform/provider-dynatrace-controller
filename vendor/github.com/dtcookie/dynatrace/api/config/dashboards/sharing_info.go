package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// SharingInfo represents sharing configuration of a dashboard
type SharingInfo struct {
	LinkShared *bool                      `json:"linkShared,omitempty"` // If `true`, the dashboard is shared via link and authenticated users with the link can view
	Published  *bool                      `json:"published,omitempty"`  // If `true`, the dashboard is published to anyone on this environment
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (me *SharingInfo) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"link_shared": {
			Type:        hcl.TypeBool,
			Description: "If `true`, the dashboard is shared via link and authenticated users with the link can view",
			Optional:    true,
		},
		"published": {
			Type:        hcl.TypeBool,
			Description: "If `true`, the dashboard is published to anyone on this environment",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *SharingInfo) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.LinkShared != nil {
		result["link_shared"] = opt.Bool(me.LinkShared)
	}
	if me.Published != nil {
		result["published"] = opt.Bool(me.Published)
	}
	return result, nil
}

func (me *SharingInfo) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "link_shared")
		delete(me.Unknowns, "published")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, value := decoder.GetChange("link_shared"); value != nil {
		me.LinkShared = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("published"); value != nil {
		me.Published = opt.NewBool(value.(bool))
	}
	return nil
}

func (me *SharingInfo) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	if me.Published != nil {
		rawMessage, err := json.Marshal(me.Published)
		if err != nil {
			return nil, err
		}
		m["published"] = rawMessage
	}

	if me.LinkShared != nil {
		rawMessage, err := json.Marshal(me.LinkShared)
		if err != nil {
			return nil, err
		}
		m["linkShared"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *SharingInfo) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("linkShared", &me.LinkShared); err != nil {
		return err
	}
	if err := m.Unmarshal("published", &me.Published); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
