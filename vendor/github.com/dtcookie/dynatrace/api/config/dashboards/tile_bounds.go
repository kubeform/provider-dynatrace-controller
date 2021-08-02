package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// TileBounds the position and size of a tile
type TileBounds struct {
	Top      int32                      `json:"top"`    // the vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels
	Left     int32                      `json:"left"`   // the horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels
	Width    int32                      `json:"width"`  // the width of the tile, in pixels
	Height   int32                      `json:"height"` // the height of the tile, in pixels
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *TileBounds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"top": {
			Type:        hcl.TypeInt,
			Description: "the vertical distance from the top left corner of the dashboard to the top left corner of the tile, in pixels",
			Required:    true,
		},
		"left": {
			Type:        hcl.TypeInt,
			Description: "the horizontal distance from the top left corner of the dashboard to the top left corner of the tile, in pixels",
			Required:    true,
		},
		"width": {
			Type:        hcl.TypeInt,
			Description: "the width of the tile, in pixels",
			Required:    true,
		},
		"height": {
			Type:        hcl.TypeInt,
			Description: "the height of the tile, in pixels",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *TileBounds) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["left"] = int(me.Left)
	result["top"] = int(me.Top)
	result["width"] = int(me.Width)
	result["height"] = int(me.Height)
	return result, nil
}

func (me *TileBounds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "left")
		delete(me.Unknowns, "top")
		delete(me.Unknowns, "width")
		delete(me.Unknowns, "height")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("left"); ok {
		me.Left = int32(value.(int))
	}
	if value, ok := decoder.GetOk("top"); ok {
		me.Top = int32(value.(int))
	}
	if value, ok := decoder.GetOk("width"); ok {
		me.Width = int32(value.(int))
	}
	if value, ok := decoder.GetOk("height"); ok {
		me.Height = int32(value.(int))
	}
	return nil
}

func (me *TileBounds) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.Left)
		if err != nil {
			return nil, err
		}
		m["left"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.Width)
		if err != nil {
			return nil, err
		}
		m["width"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.Height)
		if err != nil {
			return nil, err
		}
		m["height"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.Top)
		if err != nil {
			return nil, err
		}
		m["top"] = rawMessage
	}

	return json.Marshal(m)
}

func (me *TileBounds) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("top", &me.Top); err != nil {
		return err
	}
	if err := m.Unmarshal("left", &me.Left); err != nil {
		return err
	}
	if err := m.Unmarshal("width", &me.Width); err != nil {
		return err
	}
	if err := m.Unmarshal("height", &me.Height); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
