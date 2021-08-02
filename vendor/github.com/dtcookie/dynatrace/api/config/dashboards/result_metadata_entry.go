package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type ResultMetadataEntry struct {
	Key      string
	Config   *CustomChartingItemMetadataConfig
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *ResultMetadataEntry) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "A generated key by the Dynatrace Server",
		},
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

func (me *ResultMetadataEntry) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Config.LastModified != nil {
		result["last_modified"] = int(opt.Int64(me.Config.LastModified))
	}
	result["custom_color"] = me.Config.CustomColor
	result["key"] = me.Key
	return result, nil
}

func (me *ResultMetadataEntry) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "last_modified")
		delete(me.Unknowns, "custom_color")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = value.(string)
	}
	if value, ok := decoder.GetOk("last_modified"); ok {
		if me.Config == nil {
			me.Config = new(CustomChartingItemMetadataConfig)
		}
		me.Config.LastModified = opt.NewInt64(int64(value.(int)))
	}
	if value, ok := decoder.GetOk("custom_color"); ok {
		if me.Config == nil {
			me.Config = new(CustomChartingItemMetadataConfig)
		}
		me.Config.CustomColor = value.(string)
	}
	return nil
}
