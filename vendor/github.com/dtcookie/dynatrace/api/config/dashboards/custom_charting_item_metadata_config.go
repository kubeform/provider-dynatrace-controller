package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// CustomChartingItemMetadataConfig Additional metadata for charted metric
type CustomChartingItemMetadataConfig struct {
	LastModified *int64                     `json:"lastModified,omitempty"` // The timestamp of the last metadata modification, in UTC milliseconds
	CustomColor  string                     `json:"customColor"`            // The color of the metric in the chart, hex format
	Unknowns     map[string]json.RawMessage `json:"-"`
}

func (me *CustomChartingItemMetadataConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"last_modified": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "The timestamp of the last metadata modification, in UTC milliseconds",
		},
		"custom_color": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The color of the metric in the chart, hex format",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *CustomChartingItemMetadataConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.LastModified != nil {
		result["last_modified"] = int(opt.Int64(me.LastModified))
	}
	result["custom_color"] = me.CustomColor
	return result, nil
}

func (me *CustomChartingItemMetadataConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "last_modified")
		delete(me.Unknowns, "custom_color")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("last_modified"); ok {
		me.LastModified = opt.NewInt64(int64(value.(int)))
	}
	if value, ok := decoder.GetOk("custom_color"); ok {
		me.CustomColor = value.(string)
	}
	return nil
}

func (me *CustomChartingItemMetadataConfig) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	if me.LastModified != nil {
		rawMessage, err := json.Marshal(me.LastModified)
		if err != nil {
			return nil, err
		}
		m["lastModified"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.CustomColor)
		if err != nil {
			return nil, err
		}
		m["customColor"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *CustomChartingItemMetadataConfig) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["lastModified"]; found {
		if err := json.Unmarshal(v, &me.LastModified); err != nil {
			return err
		}
	}
	if v, found := m["customColor"]; found {
		if err := json.Unmarshal(v, &me.CustomColor); err != nil {
			return err
		}
	}
	delete(m, "lastModified")
	delete(m, "customColor")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
