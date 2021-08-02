package dimensions

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// String A filter for the metrics string dimensions.
type String struct {
	BaseDimension
	TextFilter *Filter `json:"textFilter"` // A filter for a string value based on the given operator.
}

func (me *String) GetType() FilterType {
	return FilterTypes.String
}

func (me *String) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The dimensions key on the metric",
		},
		"filter": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "A filter for a string value based on the given operator",
			Elem:        &hcl.Resource{Schema: new(Filter).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *String) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Key != nil {
		result["key"] = *me.Key
	}
	if me.TextFilter != nil {
		if marshalled, err := me.TextFilter.MarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err == nil {
			result["filter"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *String) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "filterType")
		delete(me.Unknowns, "textFilter")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("filter.#"); ok {
		me.TextFilter = new(Filter)
		if err := me.TextFilter.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *String) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.GetType(),
		"key":        me.Key,
		"textFilter": me.TextFilter,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *String) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
		"key":        &me.Key,
		"textFilter": &me.TextFilter,
	}); err != nil {
		return err
	}
	return nil
}
