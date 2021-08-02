package dimensions

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// Entity A filter for the metrics entity dimensions.
type Entity struct {
	BaseDimension
	NameFilter *Filter `json:"nameFilter"` // A filter for a string value based on the given operator.
}

func (me *Entity) GetType() FilterType {
	return FilterTypes.Entity
}

func (me *Entity) Schema() map[string]*hcl.Schema {
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

func (me *Entity) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	if me.NameFilter != nil {
		if marshalled, err := me.NameFilter.MarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err == nil {
			result["filter"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Entity) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "filterType")
		delete(me.Unknowns, "nameFilter")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("filter.#"); ok {
		me.NameFilter = new(Filter)
		if err := me.NameFilter.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *Entity) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.GetType(),
		"key":        me.Key,
		"nameFilter": me.NameFilter,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Entity) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
		"key":        &me.Key,
		"nameFilter": &me.NameFilter,
	}); err != nil {
		return err
	}
	return nil
}
