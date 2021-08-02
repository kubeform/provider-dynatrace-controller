package dashboards

import "github.com/dtcookie/hcl"

type FilterMatch struct {
	Key    string
	Values []string
}

func (me *FilterMatch) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Description: "The entity type (e.g. HOST, SERVICE, ...)",
			Required:    true,
		},
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "the tiles this Dashboard consist of",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *FilterMatch) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["key"] = me.Key
	if len(me.Values) > 0 {
		result["values"] = me.Values
	}

	return result, nil
}

func (me *FilterMatch) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = value.(string)
	}
	me.Values = decoder.GetStringSet("values")
	return nil
}
