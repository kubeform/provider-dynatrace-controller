package dashboards

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Dashboard the configuration of a dashboard
type Dashboard struct {
	ID                    *string                    `json:"id,omitempty"`      // the ID of the dashboard
	Metadata              *DashboardMetadata         `json:"dashboardMetadata"` // contains parameters of a dashboard
	Tiles                 []*Tile                    `json:"tiles"`             // the tiles the dashboard consists of
	ConfigurationMetadata *api.ConfigMetadata        `json:"metadata,omitempty"`
	Unknowns              map[string]json.RawMessage `json:"-"`
}

func (me *Dashboard) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboard_metadata": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "contains parameters of a dashboard",
			Elem: &hcl.Resource{
				Schema: new(DashboardMetadata).Schema(),
			},
		},
		"tile": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "the tiles this Dashboard consist of",
			Elem: &hcl.Resource{
				Schema: new(Tile).Schema(),
			},
		},
		"metadata": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "`metadata` exists for backwards compatibility but shouldn't get specified anymore",
			Deprecated:  "`metadata` exists for backwards compatibility but shouldn't get specified anymore",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(api.ConfigMetadata).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Dashboard) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Metadata != nil {
		if marshalled, err := me.Metadata.MarshalHCL(); err == nil {
			result["dashboard_metadata"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Tiles) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Tiles {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["tile"] = entries
	}
	return result, nil
}

func (me *Dashboard) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "dashboard_metadata")
		delete(me.Unknowns, "tile")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, ok := decoder.GetOk("dashboard_metadata.#"); ok {
		me.Metadata = new(DashboardMetadata)
		if err := me.Metadata.UnmarshalHCL(hcl.NewDecoder(decoder, "dashboard_metadata", 0)); err != nil {
			return err
		}
	}
	if result, ok := decoder.GetOk("tile.#"); ok {
		me.Tiles = []*Tile{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Tile)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "tile", idx)); err != nil {
				return err
			}
			me.Tiles = append(me.Tiles, entry)
		}
	}
	return nil
}

func (me *Dashboard) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("id", me.ID); err != nil {
		return nil, err
	}
	if err := m.Marshal("dashboardMetadata", me.Metadata); err != nil {
		return nil, err
	}
	if err := m.Marshal("metadata", me.ConfigurationMetadata); err != nil {
		return nil, err
	}
	if err := m.Marshal("tiles", me.Tiles); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *Dashboard) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("id", &me.ID); err != nil {
		return err
	}
	if err := m.Unmarshal("dashboardMetadata", &me.Metadata); err != nil {
		return err
	}
	if err := m.Unmarshal("tiles", &me.Tiles); err != nil {
		return err
	}
	if err := m.Unmarshal("metadata", &me.ConfigurationMetadata); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
