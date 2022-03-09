package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// DashboardMetadata contains parameters of a dashboard
type DashboardMetadata struct {
	Name            string                     `json:"name"`                     // the name of the dashboard
	Shared          *bool                      `json:"shared,omitempty"`         // the dashboard is shared (`true`) or private (`false`)
	Owner           *string                    `json:"owner,omitempty"`          // the owner of the dashboard
	SharingDetails  *SharingInfo               `json:"sharingDetails,omitempty"` // represents sharing configuration of a dashboard
	Filter          *DashboardFilter           `json:"dashboardFilter,omitempty"`
	Tags            []string                   `json:"tags,omitempty"`            // a set of tags assigned to the dashboard
	Preset          *bool                      `json:"preset,omitempty"`          // the dashboard is a preset (`true`)
	ValidFilterKeys []string                   `json:"validFilterKeys,omitempty"` // a set of all possible global dashboard filters that can be applied to dashboard
	DynamicFilters  *DynamicFilters            `json:"dynamicFilters,omitempty"`  // Dashboard filter configuration of a dashboard
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *DashboardMetadata) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "the name of the dashboard",
			Required:    true,
		},
		"shared": {
			Type:        hcl.TypeBool,
			Description: "the dashboard is shared (`true`) or private (`false`)",
			Optional:    true,
		},
		"owner": {
			Type:        hcl.TypeString,
			Description: "the owner of the dashboard",
			Required:    true,
		},
		"sharing_details": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "represents sharing configuration of a dashboard",
			Elem: &hcl.Resource{
				Schema: new(SharingInfo).Schema(),
			},
		},
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Global filter Settings for the Dashboard",
			Elem: &hcl.Resource{
				Schema: new(DashboardFilter).Schema(),
			},
		},
		"dynamic_filters": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Dashboard filter configuration of a dashboard",
			Elem: &hcl.Resource{
				Schema: new(DynamicFilters).Schema(),
			},
		},
		"tags": {
			Type:        hcl.TypeSet,
			Description: "a set of tags assigned to the dashboard",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		// PRESET IS READONLY
		// "preset": {
		// 	Type:        hcl.TypeBool,
		// 	Description: "the dashboard is a preset (`true`)",
		// 	Optional:    true,
		// },
		"valid_filter_keys": {
			Type:        hcl.TypeSet,
			Description: "a set of all possible global dashboard filters that can be applied to dashboard",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *DashboardMetadata) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	if me.Shared != nil {
		result["shared"] = opt.Bool(me.Shared)
	}
	if me.Owner != nil {
		result["owner"] = opt.String(me.Owner)
	}
	if len(me.Tags) > 0 {
		result["tags"] = me.Tags
	}
	if len(me.ValidFilterKeys) > 0 {
		result["valid_filter_keys"] = me.Tags
	}
	if me.SharingDetails != nil {
		if marshalled, err := me.SharingDetails.MarshalHCL(); err == nil {
			result["name"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Filter != nil {
		if marshalled, err := me.Filter.MarshalHCL(); err == nil {
			result["filter"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.DynamicFilters != nil {
		if marshalled, err := me.DynamicFilters.MarshalHCL(); err == nil {
			result["dynamic_filters"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *DashboardMetadata) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "shared")
		delete(me.Unknowns, "owner")
		delete(me.Unknowns, "sharing_details")
		delete(me.Unknowns, "dashboard_filter")
		delete(me.Unknowns, "dynamic_filters")
		delete(me.Unknowns, "tags")
		delete(me.Unknowns, "valid_filter_keys")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if _, value := decoder.GetChange("shared"); value != nil {
		me.Shared = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("owner"); ok {
		me.Owner = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("sharing_details.#"); ok {
		me.SharingDetails = new(SharingInfo)
		if err := me.SharingDetails.UnmarshalHCL(hcl.NewDecoder(decoder, "sharing_details", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("filter.#"); ok {
		me.Filter = new(DashboardFilter)
		if err := me.Filter.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("dynamic_filters.#"); ok {
		me.DynamicFilters = new(DynamicFilters)
		if err := me.DynamicFilters.UnmarshalHCL(hcl.NewDecoder(decoder, "dynamic_filters", 0)); err != nil {
			return err
		}
	}
	me.Tags = decoder.GetStringSet("tags")
	if _, value := decoder.GetChange("preset"); value != nil {
		me.Preset = opt.NewBool(value.(bool))
	}
	me.ValidFilterKeys = decoder.GetStringSet("valid_filter_keys")
	return nil
}

func (me *DashboardMetadata) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("name", me.Name); err != nil {
		return nil, err
	}
	if err := m.Marshal("shared", me.Shared); err != nil {
		return nil, err
	}
	if err := m.Marshal("owner", me.Owner); err != nil {
		return nil, err
	}
	if err := m.Marshal("sharingDetails", me.SharingDetails); err != nil {
		return nil, err
	}
	if err := m.Marshal("dashboardFilter", me.Filter); err != nil {
		return nil, err
	}
	if err := m.Marshal("dynamicFilters", me.DynamicFilters); err != nil {
		return nil, err
	}
	if err := m.Marshal("tags", me.Tags); err != nil {
		return nil, err
	}
	if err := m.Marshal("preset", me.Preset); err != nil {
		return nil, err
	}
	if err := m.Marshal("validFilterKeys", me.ValidFilterKeys); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *DashboardMetadata) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("name", &me.Name); err != nil {
		return nil
	}
	if err := m.Unmarshal("shared", &me.Shared); err != nil {
		return nil
	}
	if err := m.Unmarshal("owner", &me.Owner); err != nil {
		return nil
	}
	if err := m.Unmarshal("sharingDetails", &me.SharingDetails); err != nil {
		return nil
	}
	if err := m.Unmarshal("dashboardFilter", &me.Filter); err != nil {
		return nil
	}
	if err := m.Unmarshal("dynamicFilters", &me.DynamicFilters); err != nil {
		return nil
	}
	if err := m.Unmarshal("tags", &me.Tags); err != nil {
		return nil
	}
	if err := m.Unmarshal("preset", &me.Preset); err != nil {
		return nil
	}
	if err := m.Unmarshal("validFilterKeys", &me.ValidFilterKeys); err != nil {
		return nil
	}
	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
