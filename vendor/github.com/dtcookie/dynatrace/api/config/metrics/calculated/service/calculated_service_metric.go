package service

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// CalculatedServiceMetric Descriptor of a calculated service metric.
type CalculatedServiceMetric struct {
	EntityID            *string                     `json:"entityId,omitempty" hcl:"entity_id"` // Restricts the metric usage to the specified service. This field is mutually exclusive with the **managementZones** field.
	Name                string                      `json:"name"`                               // The displayed name of the metric.
	Unit                Unit                        `json:"unit"`                               // The unit of the metric.
	DimensionDefinition *DimensionDefinition        `json:"dimensionDefinition,omitempty"`      // Parameters of a definition of a calculated service metric.
	Enabled             bool                        `json:"enabled"`                            // The metric is enabled (`true`) or disabled (`false`).
	ManagementZones     []string                    `json:"managementZones,omitempty"`          // Restricts the metric usage to specified management zones. This field is mutually exclusive with the **entityId** field.
	MetricDefinition    *CalculatedMetricDefinition `json:"metricDefinition"`                   // The definition of a calculated service metric.
	TsmMetricKey        string                      `json:"tsmMetricKey"`                       // The key of the calculated service metric.
	UnitDisplayName     *string                     `json:"unitDisplayName,omitempty"`          // The display name of the metric's unit.   Only applicable when the **unit** parameter is set to `UNSPECIFIED`.
	Conditions          Conditions                  `json:"conditions,omitempty"`               // The set of conditions for the metric usage.   **All** the specified conditions must be fulfilled to use the metric.
	Metadata            *api.ConfigMetadata         `json:"metadata,omitempty"`                 // Metadata useful for debugging
	Unknowns            map[string]json.RawMessage  `json:"-"`
}

func (me *CalculatedServiceMetric) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entity_id": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Restricts the metric usage to the specified service. This field is mutually exclusive with the `management_zones` field",
		},
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The displayed name of the metric",
		},
		"unit": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The unit of the metric. Possible values are `BIT`, `BIT_PER_HOUR`, `BIT_PER_MINUTE`, `BIT_PER_SECOND`, `BYTE`, `BYTE_PER_HOUR`, `BYTE_PER_MINUTE`, `BYTE_PER_SECOND`, `CORES`, `COUNT`, `DAY`, `DECIBEL_MILLI_WATT`, `GIBI_BYTE`, `GIGA`, `GIGA_BYTE`, `HOUR`, `KIBI_BYTE`, `KIBI_BYTE_PER_HOUR`, `KIBI_BYTE_PER_MINUTE`, `KIBI_BYTE_PER_SECOND`, `KILO`, `KILO_BYTE`, `KILO_BYTE_PER_HOUR`, `KILO_BYTE_PER_MINUTE`, `KILO_BYTE_PER_SECOND`, `MEBI_BYTE`, `MEBI_BYTE_PER_HOUR`, `MEBI_BYTE_PER_MINUTE`, `MEBI_BYTE_PER_SECOND`, `MEGA`, `MEGA_BYTE`, `MEGA_BYTE_PER_HOUR`, `MEGA_BYTE_PER_MINUTE`, `MEGA_BYTE_PER_SECOND`, `MICRO_SECOND`, `MILLI_CORES`, `MILLI_SECOND`, `MILLI_SECOND_PER_MINUTE`, `MINUTE`, `MONTH`, `MSU`, `NANO_SECOND`, `NANO_SECOND_PER_MINUTE`, `NOT_APPLICABLE`, `PERCENT`, `PER_HOUR`, `PER_MINUTE`, `PER_SECOND`, `PIXEL`, `PROMILLE`, `RATIO`, `SECOND`, `STATE`, `UNSPECIFIED`, `WEEK` and `YEAR`",
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The metric is enabled (`true`) or disabled (`false`)",
		},
		"metric_key": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The key of the calculated service metric",
		},
		"unit_display_name": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The display name of the metric's unit. Only applicable when the **unit** parameter is set to `UNSPECIFIED`",
		},
		"management_zones": {
			Type:        hcl.TypeSet,
			Optional:    true,
			Description: "Restricts the metric usage to specified management zones. This field is mutually exclusive with the `entity_id` field",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"metric_definition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The definition of a calculated service metric",
			Elem:        &hcl.Resource{Schema: new(CalculatedMetricDefinition).Schema()},
		},
		"dimension_definition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Parameters of a definition of a calculated service metric",
			Elem:        &hcl.Resource{Schema: new(DimensionDefinition).Schema()},
		},
		"conditions": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "The set of conditions for the metric usage. **All** the specified conditions must be fulfilled to use the metric",
			Elem:        &hcl.Resource{Schema: new(Conditions).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *CalculatedServiceMetric) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	if me.MetricDefinition != nil && me.MetricDefinition.Metric == nil && me.MetricDefinition.RequestAttribute == nil {
		me.MetricDefinition = nil
	}
	return properties.EncodeAll(map[string]interface{}{
		"entity_id":            me.EntityID,
		"name":                 me.Name,
		"unit":                 me.Unit,
		"enabled":              me.Enabled,
		"metric_key":           me.TsmMetricKey,
		"unit_display_name":    me.UnitDisplayName,
		"management_zones":     me.ManagementZones,
		"metric_definition":    me.MetricDefinition,
		"dimension_definition": me.DimensionDefinition,
		"conditions":           me.Conditions,
		"unknowns":             me.Unknowns,
	})
}

func (me *CalculatedServiceMetric) UnmarshalHCL(decoder hcl.Decoder) error {
	err := decoder.DecodeAll(map[string]interface{}{
		"entity_id":            &me.EntityID,
		"name":                 &me.Name,
		"unit":                 &me.Unit,
		"enabled":              &me.Enabled,
		"metric_key":           &me.TsmMetricKey,
		"unit_display_name":    &me.UnitDisplayName,
		"management_zones":     &me.ManagementZones,
		"metric_definition":    &me.MetricDefinition,
		"dimension_definition": &me.DimensionDefinition,
		"conditions":           &me.Conditions,
		"unknowns":             &me.Unknowns,
	})
	if me.UnitDisplayName == nil || len(*me.UnitDisplayName) == 0 {
		me.UnitDisplayName = nil
	}
	if me.MetricDefinition == nil {
		me.MetricDefinition = new(CalculatedMetricDefinition)
	}
	return err
}

func (me *CalculatedServiceMetric) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"entityId":            me.EntityID,
		"name":                me.Name,
		"unit":                me.Unit,
		"enabled":             me.Enabled,
		"tsmMetricKey":        me.TsmMetricKey,
		"unitDisplayName":     me.UnitDisplayName,
		"managementZones":     me.ManagementZones,
		"metricDefinition":    me.MetricDefinition,
		"dimensionDefinition": me.DimensionDefinition,
		"conditions":          me.Conditions,
		"metadata":            me.Metadata,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *CalculatedServiceMetric) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	err := properties.UnmarshalAll(map[string]interface{}{
		"entityId":            &me.EntityID,
		"name":                &me.Name,
		"unit":                &me.Unit,
		"enabled":             &me.Enabled,
		"tsmMetricKey":        &me.TsmMetricKey,
		"unitDisplayName":     &me.UnitDisplayName,
		"managementZones":     &me.ManagementZones,
		"metricDefinition":    &me.MetricDefinition,
		"dimensionDefinition": &me.DimensionDefinition,
		"conditions":          &me.Conditions,
		"metadata":            &me.Metadata,
	})
	if me.UnitDisplayName != nil && len(*me.UnitDisplayName) == 0 {
		me.UnitDisplayName = nil
	}
	return err
}

// Unit The unit of the metric.
type Unit string

// Units offers the known enum values
var Units = struct {
	Bit                  Unit
	BitPerHour           Unit
	BitPerMinute         Unit
	BitPerSecond         Unit
	Byte                 Unit
	BytePerHour          Unit
	BytePerMinute        Unit
	BytePerSecond        Unit
	Cores                Unit
	Count                Unit
	Day                  Unit
	DecibelMilliWatt     Unit
	GibiByte             Unit
	Giga                 Unit
	GigaByte             Unit
	Hour                 Unit
	KibiByte             Unit
	KibiBytePerHour      Unit
	KibiBytePerMinute    Unit
	KibiBytePerSecond    Unit
	Kilo                 Unit
	KiloByte             Unit
	KiloBytePerHour      Unit
	KiloBytePerMinute    Unit
	KiloBytePerSecond    Unit
	MebiByte             Unit
	MebiBytePerHour      Unit
	MebiBytePerMinute    Unit
	MebiBytePerSecond    Unit
	Mega                 Unit
	MegaByte             Unit
	MegaBytePerHour      Unit
	MegaBytePerMinute    Unit
	MegaBytePerSecond    Unit
	MicroSecond          Unit
	MilliCores           Unit
	MilliSecond          Unit
	MilliSecondPerMinute Unit
	Minute               Unit
	Month                Unit
	Msu                  Unit
	NanoSecond           Unit
	NanoSecondPerMinute  Unit
	NotApplicable        Unit
	Percent              Unit
	PerHour              Unit
	PerMinute            Unit
	PerSecond            Unit
	Pixel                Unit
	Promille             Unit
	Ratio                Unit
	Second               Unit
	State                Unit
	Unspecified          Unit
	Week                 Unit
	Year                 Unit
}{
	"BIT",
	"BIT_PER_HOUR",
	"BIT_PER_MINUTE",
	"BIT_PER_SECOND",
	"BYTE",
	"BYTE_PER_HOUR",
	"BYTE_PER_MINUTE",
	"BYTE_PER_SECOND",
	"CORES",
	"COUNT",
	"DAY",
	"DECIBEL_MILLI_WATT",
	"GIBI_BYTE",
	"GIGA",
	"GIGA_BYTE",
	"HOUR",
	"KIBI_BYTE",
	"KIBI_BYTE_PER_HOUR",
	"KIBI_BYTE_PER_MINUTE",
	"KIBI_BYTE_PER_SECOND",
	"KILO",
	"KILO_BYTE",
	"KILO_BYTE_PER_HOUR",
	"KILO_BYTE_PER_MINUTE",
	"KILO_BYTE_PER_SECOND",
	"MEBI_BYTE",
	"MEBI_BYTE_PER_HOUR",
	"MEBI_BYTE_PER_MINUTE",
	"MEBI_BYTE_PER_SECOND",
	"MEGA",
	"MEGA_BYTE",
	"MEGA_BYTE_PER_HOUR",
	"MEGA_BYTE_PER_MINUTE",
	"MEGA_BYTE_PER_SECOND",
	"MICRO_SECOND",
	"MILLI_CORES",
	"MILLI_SECOND",
	"MILLI_SECOND_PER_MINUTE",
	"MINUTE",
	"MONTH",
	"MSU",
	"NANO_SECOND",
	"NANO_SECOND_PER_MINUTE",
	"NOT_APPLICABLE",
	"PERCENT",
	"PER_HOUR",
	"PER_MINUTE",
	"PER_SECOND",
	"PIXEL",
	"PROMILLE",
	"RATIO",
	"SECOND",
	"STATE",
	"UNSPECIFIED",
	"WEEK",
	"YEAR",
}
