package scope

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/common"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// TagFilter A scope filter for tags on entities.
type TagFilter struct {
	BaseAlertingScope
	TagFilter *common.TagFilter `json:"tagFilter"` // A tag-based filter of monitored entities.
}

func (me *TagFilter) GetType() FilterType {
	return FilterTypes.Tag
}

func (me *TagFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "A filter for a string value based on the given operator",
			Elem:        &hcl.Resource{Schema: new(common.TagFilter).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *TagFilter) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.TagFilter != nil {
		if marshalled, err := me.TagFilter.MarshalHCL(); err == nil {
			result["filter"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *TagFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "filterType")
		delete(me.Unknowns, "tagFilter")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, ok := decoder.GetOk("filter.#"); ok {
		me.TagFilter = new(common.TagFilter)
		if err := me.TagFilter.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *TagFilter) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.GetType(),
		"tagFilter":  me.TagFilter,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *TagFilter) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
		"tagFilter":  &me.TagFilter,
	}); err != nil {
		return err
	}
	return nil
}
