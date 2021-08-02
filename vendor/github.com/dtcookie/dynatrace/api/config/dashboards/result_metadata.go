package dashboards

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/dtcookie/hcl"
)

type ResultMetadata struct {
	Entries []*ResultMetadataEntry
}

func (me *ResultMetadata) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"config": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Additional metadata for charted metric",
			Elem: &hcl.Resource{
				Schema: new(ResultMetadataEntry).Schema(),
			},
		},
	}
}

func (me *ResultMetadata) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.Entries) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Entries {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		sort.Slice(entries, func(i, j int) bool {
			d1, _ := json.Marshal(entries[i])
			d2, _ := json.Marshal(entries[j])
			cmp := strings.Compare(string(d1), string(d2))
			return (cmp == -1)
		})

		result["config"] = entries
	}
	return result, nil
}

func (me *ResultMetadata) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("config.#"); ok {
		me.Entries = []*ResultMetadataEntry{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(ResultMetadataEntry)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "config", idx)); err != nil {
				return err
			}
			me.Entries = append(me.Entries, entry)
		}
	}
	return nil
}
