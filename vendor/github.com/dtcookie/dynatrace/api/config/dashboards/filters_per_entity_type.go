package dashboards

import (
	"github.com/dtcookie/hcl"
)

type FiltersPerEntityType struct {
	Filters []*FilterForEntityType
}

func (me *FiltersPerEntityType) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "the tiles this Dashboard consist of",
			Elem: &hcl.Resource{
				Schema: new(FilterForEntityType).Schema(),
			},
		},
	}
}

func (me *FiltersPerEntityType) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.Filters) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Filters {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["filter"] = entries
	}
	return result, nil
}

func (me *FiltersPerEntityType) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("filter.#"); ok {
		me.Filters = []*FilterForEntityType{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(FilterForEntityType)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", idx)); err != nil {
				return err
			}
			me.Filters = append(me.Filters, entry)
		}
	}
	return nil
}
