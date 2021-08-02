package dashboards

import "github.com/dtcookie/hcl"

type FilterForEntityType struct {
	EntityType string
	Filters    []*FilterMatch
}

func (me *FilterForEntityType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entity_type": {
			Type:        hcl.TypeString,
			Description: "The entity type (e.g. HOST, SERVICE, ...)",
			Required:    true,
		},
		"match": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "the tiles this Dashboard consist of",
			Elem: &hcl.Resource{
				Schema: new(FilterMatch).Schema(),
			},
		},
	}
}

func (me *FilterForEntityType) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["entity_type"] = me.EntityType
	if len(me.Filters) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Filters {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["match"] = entries
	}
	return result, nil
}

func (me *FilterForEntityType) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("entity_type"); ok {
		me.EntityType = value.(string)
	}
	if result, ok := decoder.GetOk("match.#"); ok {
		me.Filters = []*FilterMatch{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(FilterMatch)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "match", idx)); err != nil {
				return err
			}
			me.Filters = append(me.Filters, entry)
		}
	}
	return nil
}
