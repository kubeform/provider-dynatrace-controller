package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AbstractTile the configuration of a tile.
// The actual set of fields depends on the type of the tile. See the description of the **tileType** field
type Tile struct {
	Name                      string                             `json:"name"` // the name of the tile
	TileType                  TileType                           `json:"tileType"`
	Configured                *bool                              `json:"configured,omitempty"`                // The tile is configured and ready to use (`true`) or just placed on the dashboard (`false`)
	Bounds                    *TileBounds                        `json:"bounds"`                              // Bounds the position and size of a tile
	Filter                    *TileFilter                        `json:"tileFilter,omitempty"`                // is filter applied to a tile. It overrides dashboard's filter
	AssignedEntities          []string                           `json:"assignedEntities"`                    // The list of Dynatrace entities, assigned to the tile
	Metric                    *string                            `json:"metric,omitempty"`                    // The metric assigned to the tile
	CustomName                *string                            `json:"customName"`                          // The name of the tile, set by user
	Query                     *string                            `json:"query"`                               // A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile
	Visualization             *UserSessionQueryTileType          `json:"type"`                                // The visualization of the tile
	TimeFrameShift            *string                            `json:"timeFrameShift,omitempty"`            // The comparison timeframe of the query. If specified, you additionally get the results of the same query with the specified time shift
	VisualizationConfig       *UserSessionQueryTileConfiguration `json:"visualizationConfig,omitempty"`       // Configuration of a User session query visualization tile
	Limit                     *int32                             `json:"limit,omitempty"`                     // The limit of the results, if not set will use the default value of the system
	FilterConfig              *CustomFilterConfig                `json:"filterConfig,omitempty"`              // Configuration of the custom filter of a tile
	Markdown                  *string                            `json:"markdown"`                            // The markdown-formatted content of the tile
	ExcludeMaintenanceWindows *bool                              `json:"excludeMaintenanceWindows,omitempty"` // Include (`false') or exclude (`true`) maintenance windows from availability calculations
	ChartVisible              *bool                              `json:"chartVisible,omitempty"`

	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *Tile) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "the name of the tile",
			Required:    true,
		},
		"tile_type": {
			Type:        hcl.TypeString,
			Description: "the type of the tile. Must be either `APPLICATION_WORLDMAP`, `RESOURCES`, `THIRD_PARTY_MOST_ACTIVE`, `UEM_CONVERSIONS_PER_GOAL`, `PROCESS_GROUPS_ONE` or `HOST` .",
			Required:    true,
		},
		"configured": {
			Type:        hcl.TypeBool,
			Description: "The tile is configured and ready to use (`true`) or just placed on the dashboard (`false`)",
			Optional:    true,
		},
		"bounds": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "the position and size of a tile",
			Elem: &hcl.Resource{
				Schema: new(TileBounds).Schema(),
			},
		},
		"filter": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "is filter applied to a tile. It overrides dashboard's filter",
			Elem: &hcl.Resource{
				Schema: new(TileFilter).Schema(),
			},
		},
		"assigned_entities": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The list of Dynatrace entities, assigned to the tile",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"metric": {
			Type:        hcl.TypeString,
			Description: "The metric assigned to the tile",
			Optional:    true,
		},
		"custom_name": {
			Type:        hcl.TypeString,
			Description: "The name of the tile, set by user",
			Optional:    true,
		},
		"query": {
			Type:        hcl.TypeString,
			Description: "A [user session query](https://www.dynatrace.com/support/help/shortlink/usql-info) executed by the tile",
			Optional:    true,
		},
		"visualization": {
			Type:        hcl.TypeString,
			Description: "The visualization of the tile. Possible values are: `COLUMN_CHART`, `FUNNEL`, `LINE_CHART`, `PIE_CHART`, `SINGLE_VALUE`, `TABLE`",
			Optional:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The attribute `type` exists for backwards compatibilty. Usage is discouraged. You should use `visualization` instead.",
			Optional:    true,
		},
		"time_frame_shift": {
			Type:        hcl.TypeString,
			Description: "The comparison timeframe of the query. If specified, you additionally get the results of the same query with the specified time shift",
			Optional:    true,
		},
		"visualization_config": {
			Type:        hcl.TypeList,
			Description: "Configuration of a User session query visualization tile",
			MaxItems:    1,
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(UserSessionQueryTileConfiguration).Schema(),
			},
		},
		"limit": {
			Type:        hcl.TypeInt,
			Description: "The limit of the results, if not set will use the default value of the system",
			Optional:    true,
		},
		"filter_config": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "the position and size of a tile",
			Elem: &hcl.Resource{
				Schema: new(CustomFilterConfig).Schema(),
			},
		},
		"markdown": {
			Type:        hcl.TypeString,
			Description: "The markdown-formatted content of the tile",
			Optional:    true,
		},
		"exclude_maintenance_windows": {
			Type:        hcl.TypeBool,
			Description: "Include (`false') or exclude (`true`) maintenance windows from availability calculations",
			Optional:    true,
		},
		"chart_visible": {
			Type:        hcl.TypeBool,
			Description: "",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Tile) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	result["tile_type"] = string(me.TileType)
	if me.Configured != nil {
		result["configured"] = opt.Bool(me.Configured)
	}
	if me.Bounds != nil {
		if marshalled, err := me.Bounds.MarshalHCL(); err == nil {
			result["bounds"] = []interface{}{marshalled}
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
	if len(me.AssignedEntities) > 0 {
		result["assigned_entities"] = me.AssignedEntities
	}
	if me.Metric != nil {
		result["metric"] = opt.String(me.Metric)
	}
	if me.CustomName != nil {
		result["custom_name"] = opt.String(me.CustomName)
	}
	if me.Query != nil {
		result["query"] = opt.String(me.Query)
	}
	if me.Visualization != nil {
		result["visualization"] = string(*me.Visualization)
	}
	if me.TimeFrameShift != nil {
		result["time_frame_shift"] = string(*me.TimeFrameShift)
	}
	if me.VisualizationConfig != nil {
		if marshalled, err := me.VisualizationConfig.MarshalHCL(); err == nil {
			result["visualization_config"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Limit != nil {
		result["limit"] = int(opt.Int32(me.Limit))
	}

	if me.FilterConfig != nil {
		if marshalled, err := me.FilterConfig.MarshalHCL(); err == nil {
			result["filter_config"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Markdown != nil {
		result["markdown"] = string(*me.Markdown)
	}
	if me.ExcludeMaintenanceWindows != nil {
		result["exclude_maintenance_windows"] = opt.Bool(me.ExcludeMaintenanceWindows)
	}
	if me.ChartVisible != nil {
		result["chart_visible"] = opt.Bool(me.ChartVisible)
	}
	return result, nil
}

func (me *Tile) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "tile_type")
		delete(me.Unknowns, "configured")
		delete(me.Unknowns, "bounds")
		delete(me.Unknowns, "filter")
		delete(me.Unknowns, "assigned_entities")
		delete(me.Unknowns, "metric")
		delete(me.Unknowns, "custom_name")
		delete(me.Unknowns, "query")
		delete(me.Unknowns, "visualization")
		delete(me.Unknowns, "type")
		delete(me.Unknowns, "time_frame_shift")
		delete(me.Unknowns, "visualization_config")
		delete(me.Unknowns, "limit")
		delete(me.Unknowns, "filter_config")
		delete(me.Unknowns, "markdown")
		delete(me.Unknowns, "exclude_maintenance_windows")
		delete(me.Unknowns, "chart_visible")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("tile_type"); ok {
		me.TileType = TileType(value.(string))
	}
	if _, value := decoder.GetChange("configured"); value != nil {
		me.Configured = opt.NewBool(value.(bool))
	}
	if _, ok := decoder.GetOk("bounds.#"); ok {
		me.Bounds = new(TileBounds)
		if err := me.Bounds.UnmarshalHCL(hcl.NewDecoder(decoder, "bounds", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("filter.#"); ok {
		me.Filter = new(TileFilter)
		if err := me.Filter.UnmarshalHCL(hcl.NewDecoder(decoder, "filter", 0)); err != nil {
			return err
		}
	}
	me.AssignedEntities = decoder.GetStringSet("assigned_entities")
	if value, ok := decoder.GetOk("metric"); ok {
		me.Metric = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("custom_name"); ok {
		me.CustomName = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("query"); ok {
		me.Query = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("visualization"); ok {
		me.Visualization = UserSessionQueryTileType(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("time_frame_shift"); ok {
		me.TimeFrameShift = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("visualization_config.#"); ok {
		me.VisualizationConfig = new(UserSessionQueryTileConfiguration)
		if err := me.VisualizationConfig.UnmarshalHCL(hcl.NewDecoder(decoder, "visualization_config", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("limit"); ok {
		me.Limit = opt.NewInt32(int32(value.(int)))
	}
	if _, ok := decoder.GetOk("filter_config.#"); ok {
		me.FilterConfig = new(CustomFilterConfig)
		if err := me.FilterConfig.UnmarshalHCL(hcl.NewDecoder(decoder, "filter_config", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("markdown"); ok {
		me.Markdown = opt.NewString(value.(string))
	}
	if _, value := decoder.GetChange("exclude_maintenance_windows"); value != nil {
		me.ExcludeMaintenanceWindows = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("chart_visible"); value != nil {
		me.ChartVisible = opt.NewBool(value.(bool))
	}
	return nil
}

func (me *Tile) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.TileType)
		if err != nil {
			return nil, err
		}
		m["tileType"] = rawMessage
	}
	if me.Configured != nil {
		rawMessage, err := json.Marshal(me.Configured)
		if err != nil {
			return nil, err
		}
		m["configured"] = rawMessage
	}
	if me.Bounds != nil {
		rawMessage, err := json.Marshal(me.Bounds)
		if err != nil {
			return nil, err
		}
		m["bounds"] = rawMessage
	}
	if me.Filter != nil {
		rawMessage, err := json.Marshal(me.Filter)
		if err != nil {
			return nil, err
		}
		m["tileFilter"] = rawMessage
	}
	if len(me.AssignedEntities) > 0 {
		rawMessage, err := json.Marshal(me.AssignedEntities)
		if err != nil {
			return nil, err
		}
		m["assignedEntities"] = rawMessage
	} else {
		rawMessage, err := json.Marshal([]string{})
		if err != nil {
			return nil, err
		}
		m["assignedEntities"] = rawMessage
	}
	if me.Metric != nil {
		rawMessage, err := json.Marshal(me.Metric)
		if err != nil {
			return nil, err
		}
		m["metric"] = rawMessage
	}
	if me.CustomName != nil {
		rawMessage, err := json.Marshal(me.CustomName)
		if err != nil {
			return nil, err
		}
		m["customName"] = rawMessage
	}
	if me.Query != nil {
		rawMessage, err := json.Marshal(me.Query)
		if err != nil {
			return nil, err
		}
		m["query"] = rawMessage
	}
	if me.Visualization != nil {
		rawMessage, err := json.Marshal(me.Visualization)
		if err != nil {
			return nil, err
		}
		m["type"] = rawMessage
	}
	if me.TimeFrameShift != nil {
		rawMessage, err := json.Marshal(me.TimeFrameShift)
		if err != nil {
			return nil, err
		}
		m["timeFrameShift"] = rawMessage
	}
	if me.VisualizationConfig != nil {
		rawMessage, err := json.Marshal(me.VisualizationConfig)
		if err != nil {
			return nil, err
		}
		m["visualizationConfig"] = rawMessage
	}
	if me.Limit != nil {
		rawMessage, err := json.Marshal(me.Limit)
		if err != nil {
			return nil, err
		}
		m["limit"] = rawMessage
	}
	if me.FilterConfig != nil {
		rawMessage, err := json.Marshal(me.FilterConfig)
		if err != nil {
			return nil, err
		}
		m["filterConfig"] = rawMessage
	}
	if me.Markdown != nil {
		rawMessage, err := json.Marshal(me.Markdown)
		if err != nil {
			return nil, err
		}
		m["markdown"] = rawMessage
	}
	if me.ExcludeMaintenanceWindows != nil {
		rawMessage, err := json.Marshal(me.ExcludeMaintenanceWindows)
		if err != nil {
			return nil, err
		}
		m["excludeMaintenanceWindows"] = rawMessage
	}
	if me.ChartVisible != nil {
		rawMessage, err := json.Marshal(me.ChartVisible)
		if err != nil {
			return nil, err
		}
		m["chartVisible"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *Tile) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &me.Name); err != nil {
			return err
		}
	}
	if v, found := m["tileType"]; found {
		if err := json.Unmarshal(v, &me.TileType); err != nil {
			return err
		}
	}
	if v, found := m["configured"]; found {
		if err := json.Unmarshal(v, &me.Configured); err != nil {
			return err
		}
	}
	if v, found := m["bounds"]; found {
		if err := json.Unmarshal(v, &me.Bounds); err != nil {
			return err
		}
	}
	if v, found := m["tileFilter"]; found {
		if err := json.Unmarshal(v, &me.Filter); err != nil {
			return err
		}
	}
	if me.Filter != nil && me.Filter.IsZero() {
		me.Filter = nil
	}
	if v, found := m["assignedEntities"]; found {
		if err := json.Unmarshal(v, &me.AssignedEntities); err != nil {
			return err
		}
	}
	if v, found := m["metric"]; found {
		if err := json.Unmarshal(v, &me.Metric); err != nil {
			return err
		}
	}
	if v, found := m["customName"]; found {
		if err := json.Unmarshal(v, &me.CustomName); err != nil {
			return err
		}
	}
	if v, found := m["query"]; found {
		if err := json.Unmarshal(v, &me.Query); err != nil {
			return err
		}
	}
	if v, found := m["type"]; found {
		if err := json.Unmarshal(v, &me.Visualization); err != nil {
			return err
		}
	}
	if v, found := m["timeFrameShift"]; found {
		if err := json.Unmarshal(v, &me.TimeFrameShift); err != nil {
			return err
		}
	}
	if v, found := m["visualizationConfig"]; found {
		if err := json.Unmarshal(v, &me.VisualizationConfig); err != nil {
			return err
		}
	}
	if v, found := m["limit"]; found {
		if err := json.Unmarshal(v, &me.Limit); err != nil {
			return err
		}
	}
	if v, found := m["filterConfig"]; found {
		if err := json.Unmarshal(v, &me.FilterConfig); err != nil {
			return err
		}
	}
	if v, found := m["markdown"]; found {
		if err := json.Unmarshal(v, &me.Markdown); err != nil {
			return err
		}
	}
	if v, found := m["excludeMaintenanceWindows"]; found {
		if err := json.Unmarshal(v, &me.ExcludeMaintenanceWindows); err != nil {
			return err
		}
	}
	if v, found := m["chartVisible"]; found {
		if err := json.Unmarshal(v, &me.ChartVisible); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "tileType")
	delete(m, "configured")
	delete(m, "bounds")
	delete(m, "tileFilter")
	delete(m, "assignedEntities")
	delete(m, "metric")
	delete(m, "customName")
	delete(m, "query")
	delete(m, "type")
	delete(m, "timeFrameShift")
	delete(m, "visualizationConfig")
	delete(m, "limit")
	delete(m, "filterConfig")
	delete(m, "markdown")
	delete(m, "excludeMaintenanceWindows")
	delete(m, "chartVisible")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// UserSessionQueryTileType has no documentation
type UserSessionQueryTileType string

func (me UserSessionQueryTileType) Ref() *UserSessionQueryTileType {
	return &me
}

// UserSessionQueryTileTypes offers the known enum values
var UserSessionQueryTileTypes = struct {
	ColumnChart UserSessionQueryTileType
	Funnel      UserSessionQueryTileType
	LineChart   UserSessionQueryTileType
	PieChart    UserSessionQueryTileType
	SingleValue UserSessionQueryTileType
	Table       UserSessionQueryTileType
}{
	"COLUMN_CHART",
	"FUNNEL",
	"LINE_CHART",
	"PIE_CHART",
	"SINGLE_VALUE",
	"TABLE",
}
