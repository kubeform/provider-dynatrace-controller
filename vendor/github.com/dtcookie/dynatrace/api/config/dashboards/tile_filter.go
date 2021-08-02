package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// TileFilter is filter applied to a tile. It overrides dashboard's filter
type TileFilter struct {
	Timeframe      *string                    `json:"timeframe,omitempty"` // the default timeframe of the dashboard
	ManagementZone *EntityRef                 `json:"managementZone,omitempty"`
	Unknowns       map[string]json.RawMessage `json:"-"`
}

func (me *TileFilter) IsZero() bool {
	return me.Timeframe == nil && me.ManagementZone == nil
}

func (me *TileFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"timeframe": {
			Type:        hcl.TypeString,
			Description: "the default timeframe of the tile",
			Optional:    true,
		},
		"management_zone": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "the management zone this tile applies to",
			Elem: &hcl.Resource{
				Schema: new(EntityRef).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *TileFilter) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Timeframe != nil {
		result["timeframe"] = opt.String(me.Timeframe)
	}
	if me.ManagementZone != nil {
		if marshalled, err := me.ManagementZone.MarshalHCL(); err == nil {
			result["management_zone"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *TileFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "timeframe")
		delete(me.Unknowns, "management_zone")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}

	if value, ok := decoder.GetOk("timeframe"); ok {
		me.Timeframe = opt.NewString(value.(string))
	}

	if _, ok := decoder.GetOk("management_zone.#"); ok {
		me.ManagementZone = new(EntityRef)
		if err := me.ManagementZone.UnmarshalHCL(hcl.NewDecoder(decoder, "management_zone", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *TileFilter) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	if me.Timeframe != nil {
		rawMessage, err := json.Marshal(me.Timeframe)
		if err != nil {
			return nil, err
		}
		m["timeframe"] = rawMessage
	}
	if me.ManagementZone != nil {
		rawMessage, err := json.Marshal(me.ManagementZone)
		if err != nil {
			return nil, err
		}
		m["managementZone"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *TileFilter) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("timeframe", &me.Timeframe); err != nil {
		return err
	}
	if err := m.Unmarshal("managementZone", &me.ManagementZone); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
