package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// ExtractSubstring Preprocess by extracting a substring from the original value.
type ExtractSubstring struct {
	EndDelimiter *string                    `json:"endDelimiter,omitempty"` // The end-delimiter string.   Required if the **position** value is `BETWEEN`. Otherwise not allowed.
	Position     Position                   `json:"position"`               // The position of the extracted string relative to delimiters.
	Delimiter    string                     `json:"delimiter"`              // The delimiter string.
	Unknowns     map[string]json.RawMessage `json:"-"`
}

func (me *ExtractSubstring) IsZero() bool {
	if me.EndDelimiter != nil && len(*me.EndDelimiter) > 0 {
		return false
	}
	if len(me.Position) > 0 {
		return false
	}
	if len(me.Delimiter) > 0 {
		return false
	}
	if len(me.Unknowns) > 0 {
		return false
	}
	return true
}

func (me *ExtractSubstring) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"end_delimiter": {
			Type:        hcl.TypeString,
			Description: "The end-delimiter string.   Required if the **position** value is `BETWEEN`. Otherwise not allowed",
			Optional:    true,
		},
		"position": {
			Type:        hcl.TypeString,
			Description: "The position of the extracted string relative to delimiters",
			Required:    true,
		},
		"delimiter": {
			Type:        hcl.TypeString,
			Description: "The delimiter string",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ExtractSubstring) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.EndDelimiter != nil {
		result["end_delimiter"] = *me.EndDelimiter
	}
	result["position"] = string(me.Position)
	result["delimiter"] = me.Delimiter
	return result, nil
}

func (me *ExtractSubstring) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "end_delimiter")
		delete(me.Unknowns, "position")
		delete(me.Unknowns, "delimiter")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("end_delimiter"); ok {
		me.EndDelimiter = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("position"); ok {
		me.Position = Position(value.(string))
	}
	if value, ok := decoder.GetOk("delimiter"); ok {
		me.Delimiter = value.(string)
	}
	return nil
}

func (me *ExtractSubstring) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("endDelimiter", me.EndDelimiter); err != nil {
		return nil, err
	}
	if err := m.Marshal("position", me.Position); err != nil {
		return nil, err
	}
	if err := m.Marshal("delimiter", me.Delimiter); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *ExtractSubstring) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("endDelimiter", &me.EndDelimiter); err != nil {
		return err
	}
	if err := m.Unmarshal("position", &me.Position); err != nil {
		return err
	}
	if err := m.Unmarshal("delimiter", &me.Delimiter); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// Position The position of the extracted string relative to delimiters.
type Position string

// Positions offers the known enum values
var Positions = struct {
	After   Position
	Before  Position
	Between Position
}{
	"AFTER",
	"BEFORE",
	"BETWEEN",
}
