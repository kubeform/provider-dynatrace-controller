package alerting

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/common"
	"github.com/dtcookie/hcl"
)

// ProfileTagFilter Configuration of the tag filtering of the alerting profile.
type ProfileTagFilter struct {
	IncludeMode IncludeMode                `json:"includeMode"`          // The filtering mode:  * `INCLUDE_ANY`: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.  * `INCLUDE_ALL`: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.  * `NONE`: The rule applies to all monitored entities.
	TagFilters  []*common.TagFilter        `json:"tagFilters,omitempty"` // A list of required tags.
	Unknowns    map[string]json.RawMessage `json:"-"`
}

func (me *ProfileTagFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"include_mode": {
			Type:        hcl.TypeString,
			Description: "The filtering mode:  * `INCLUDE_ANY`: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.  * `INCLUDE_ALL`: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.  * `NONE`: The rule applies to all monitored entities",
			Required:    true,
		},
		"tag_filters": {
			Type:        hcl.TypeList,
			Description: "A list of required tags",
			Optional:    true,
			MinItems:    1,
			Elem: &hcl.Resource{
				Schema: new(common.TagFilter).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ProfileTagFilter) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["include_mode"] = string(me.IncludeMode)
	if me.TagFilters != nil {
		entries := []interface{}{}
		for _, entry := range me.TagFilters {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["tag_filters"] = entries
	}
	return result, nil
}

func (me *ProfileTagFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "include_mode")
		delete(me.Unknowns, "tag_filters")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("include_mode"); ok {
		me.IncludeMode = IncludeMode(value.(string))
	}
	if result, ok := decoder.GetOk("tag_filters.#"); ok {
		me.TagFilters = []*common.TagFilter{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(common.TagFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "tag_filters", idx)); err != nil {
				return err
			}
			me.TagFilters = append(me.TagFilters, entry)
		}
	}
	return nil
}

func (me *ProfileTagFilter) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.IncludeMode)
		if err != nil {
			return nil, err
		}
		m["includeMode"] = rawMessage
	}
	if len(me.TagFilters) > 0 {
		rawMessage, err := json.Marshal(me.TagFilters)
		if err != nil {
			return nil, err
		}
		m["tagFilters"] = rawMessage
	}

	return json.Marshal(m)
}

func (me *ProfileTagFilter) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["includeMode"]; found {
		if err := json.Unmarshal(v, &me.IncludeMode); err != nil {
			return err
		}
	}
	if v, found := m["tagFilters"]; found {
		if err := json.Unmarshal(v, &me.TagFilters); err != nil {
			return err
		}
	}

	delete(m, "includeMode")
	delete(m, "tagFilters")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// IncludeMode The filtering mode:
// * `INCLUDE_ANY`: The rule applies to monitored entities that have at least one of the specified tags. You can specify up to 100 tags.
// * `INCLUDE_ALL`: The rule applies to monitored entities that have **all** of the specified tags. You can specify up to 10 tags.
// * `NONE`: The rule applies to all monitored entities.
type IncludeMode string

// IncludeModes offers the known enum values
var IncludeModes = struct {
	IncludeAll IncludeMode
	IncludeAny IncludeMode
	None       IncludeMode
}{
	"INCLUDE_ALL",
	"INCLUDE_ANY",
	"NONE",
}
