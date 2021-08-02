package common

import (
	"sort"
	"strings"

	"github.com/dtcookie/hcl"
)

type TagFilters []*TagFilter

func (me TagFilters) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A Tag Filter",
			Elem:        &hcl.Resource{Schema: new(TagFilter).Schema()},
		},
	}
}

func (me TagFilters) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	tagFilters := TagFilters{}
	tagFilters = append(tagFilters, me...)
	sort.Slice(tagFilters, func(i int, j int) bool {
		a := tagFilters[i]
		b := tagFilters[j]
		return strings.Compare(a.Key, b.Key) > 0
	})
	filters := []interface{}{}
	for _, filter := range tagFilters {
		if marshalled, err := filter.MarshalHCL(); err == nil {
			filters = append(filters, marshalled)
		} else {
			return nil, err
		}
	}
	if len(filters) > 0 {
		result["filter"] = filters
	}
	return result, nil
}

func (me *TagFilters) UnmarshalHCL(decoder hcl.Decoder) error {
	nme := TagFilters{}
	if result, ok := decoder.GetOk("filter.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(TagFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	*me = nme
	return nil
}
