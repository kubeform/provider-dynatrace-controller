package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// DynamicFilters Dashboard filter configuration of a dashboard
type DynamicFilters struct {
	Filters            []string                   `json:"filters,omitempty"`            // A set of all possible global dashboard filters that can be applied to a dashboard \n\nCurrently supported values are: \n\n\tOS_TYPE,\n\tSERVICE_TYPE,\n\tDEPLOYMENT_TYPE,\n\tAPPLICATION_INJECTION_TYPE,\n\tPAAS_VENDOR_TYPE,\n\tDATABASE_VENDOR,\n\tHOST_VIRTUALIZATION_TYPE,\n\tHOST_MONITORING_MODE,\n\tKUBERNETES_CLUSTER,\n\tRELATED_CLOUD_APPLICATION,\n\tRELATED_NAMESPACE,\n\tTAG_KEY:<tagname>
	TagSuggestionTypes []string                   `json:"tagSuggestionTypes,omitempty"` // A set of entities applied for tag filter suggestions. You can fetch the list of possible values with the [GET all entity types](https://dt-url.net/dw03s7h)request. \n\nOnly applicable if the **filters** set includes `TAG_KEY:<tagname>`
	Unknowns           map[string]json.RawMessage `json:"-"`
}

func (me *DynamicFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filters": {
			Type:        hcl.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "A set of all possible global dashboard filters that can be applied to a dashboard \n\nCurrently supported values are: \n\n\tOS_TYPE,\n\tSERVICE_TYPE,\n\tDEPLOYMENT_TYPE,\n\tAPPLICATION_INJECTION_TYPE,\n\tPAAS_VENDOR_TYPE,\n\tDATABASE_VENDOR,\n\tHOST_VIRTUALIZATION_TYPE,\n\tHOST_MONITORING_MODE,\n\tKUBERNETES_CLUSTER,\n\tRELATED_CLOUD_APPLICATION,\n\tRELATED_NAMESPACE,\n\tTAG_KEY:<tagname>",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"tag_suggestion_types": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "A set of entities applied for tag filter suggestions. You can fetch the list of possible values with the [GET all entity types](https://dt-url.net/dw03s7h)request. \n\nOnly applicable if the **filters** set includes `TAG_KEY:<tagname>`",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *DynamicFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "filters")
		delete(me.Unknowns, "tag_suggestion_types")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	me.Filters = decoder.GetStringSet("filters")
	me.TagSuggestionTypes = decoder.GetStringSet("tag_suggestion_types")
	return nil
}

func (me *DynamicFilters) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if len(me.Filters) > 0 {
		result["filters"] = me.Filters
	}
	if len(me.TagSuggestionTypes) > 0 {
		result["tag_suggestion_types"] = me.TagSuggestionTypes
	}
	return result, nil
}

func (me *DynamicFilters) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	m.Marshal("filters", me.Filters)
	m.Marshal("tagSuggestionTypes", me.TagSuggestionTypes)
	return json.Marshal(m)
}

func (me *DynamicFilters) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("filters", &me.Filters); err != nil {
		return err
	}
	if err := m.Unmarshal("tagSuggestionTypes", &me.TagSuggestionTypes); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
