package metricevents

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalies/metricevents/dimensions"
	"github.com/dtcookie/dynatrace/api/config/anomalies/metricevents/scope"
	"github.com/dtcookie/dynatrace/api/config/anomalies/metricevents/strategy"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// MetricEvent The configuration of the metric event.
type MetricEvent struct {
	ID                  *string                     `json:"id,omitempty"`                  // The ID of the metric event.
	MetricID            *string                     `json:"metricId"`                      // The ID of the metric evaluated by the metric event.
	AggregationType     *AggregationType            `json:"aggregationType,omitempty"`     // How the metric data points are aggregated for the evaluation.   The timeseries must support this aggregation.
	Description         string                      `json:"description"`                   // The description of the metric event.
	Name                string                      `json:"name"`                          // The name of the metric event displayed in the UI.
	WarningReason       *WarningReason              `json:"warningReason,omitempty"`       // The reason of a warning set on the config.  The `NONE` means config has no warnings.
	MetricDimensions    dimensions.Dimensions       `json:"metricDimensions,omitempty"`    // Defines the dimensions of the metric to alert on. The filters are combined by conjunction.
	DisabledReason      *DisabledReason             `json:"disabledReason,omitempty"`      // The reason of automatic disabling.  The `NONE` means config was not disabled automatically.
	Enabled             bool                        `json:"enabled"`                       // The metric event is enabled (`true`) or disabled (`false`).
	AlertingScope       scope.AlertingScopes        `json:"alertingScope,omitempty"`       // Defines the scope of the metric event. Only one filter is allowed per filter type, except for tags, where up to 3 are allowed. The filters are combined by conjunction.
	MonitoringStrategy  strategy.MonitoringStrategy `json:"monitoringStrategy"`            // A monitoring strategy for a metric event config. This is the base version of the monitoring strategy, depending on the type,  the actual JSON may contain additional fields.
	PrimaryDimensionKey *string                     `json:"primaryDimensionKey,omitempty"` // Defines which dimension key should be used for the **alertingScope**.
	Severity            *Severity                   `json:"severity,omitempty"`            // The type of the event to trigger on the threshold violation.  The `CUSTOM_ALERT` type is not correlated with other alerts. The `INFO` type does not open a problem.
	MetricSelector      *string                     `json:"metricSelector,omitempty"`      // The metric selector that should be executed
	Metadata            *api.ConfigMetadata         `json:"metadata,omitempty"`            // Metadata useful for debugging
	Unknowns            map[string]json.RawMessage  `json:"-"`
}

func (me *MetricEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"metric_id": {
			Type:          hcl.TypeString,
			Optional:      true,
			ConflictsWith: []string{"metric_selector"},
			Description:   "The ID of the metric evaluated by the metric event",
		},
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the metric event displayed in the UI",
		},
		"description": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The description of the metric event",
		},
		"aggregation_type": {
			Type:          hcl.TypeString,
			Optional:      true,
			ConflictsWith: []string{"metric_selector"},
			Description:   "How the metric data points are aggregated for the evaluation. The timeseries must support this aggregation",
		},
		"metric_selector": {
			Type:          hcl.TypeString,
			Optional:      true,
			ConflictsWith: []string{"metric_id", "scopes", "aggregation_type"},
			Description:   "The metric selector that should be executed",
		},
		"warning_reason": {
			Type:     hcl.TypeString,
			Optional: true,

			Deprecated:  "This property is not meant to be configured from the outside. It will get removed completely in future versions",
			Description: "The reason of a warning set on the config. The `NONE` means config has no warnings. The other supported value is `TOO_MANY_DIMS`",
		},
		"dimensions": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "Defines the dimensions of the metric to alert on. The filters are combined by conjunction",
			Elem:        &hcl.Resource{Schema: new(dimensions.Dimensions).Schema()},
		},
		"disabled_reason": {
			Type:        hcl.TypeString,
			Optional:    true,
			Deprecated:  "This property is not meant to be configured from the outside. It will get removed completely in future versions",
			Description: "The reason of automatic disabling.  The `NONE` means config was not disabled automatically. Possible values are `METRIC_DEFINITION_INCONSISTENCY`, `NONE`, `TOO_MANY_DIMS` and `TOPX_FORCIBLY_DEACTIVATED`",
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Required:    true,
			Description: "The metric event is enabled (`true`) or disabled (`false`)",
		},
		"scopes": {
			Type:          hcl.TypeList,
			ConflictsWith: []string{"metric_selector"},
			Optional:      true,
			Description:   "Defines the scope of the metric event. Only one filter is allowed per filter type, except for tags, where up to 3 are allowed. The filters are combined by conjunction",
			Elem:          &hcl.Resource{Schema: new(scope.AlertingScopes).Schema()},
		},
		"strategy": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "A monitoring strategy for a metric event config. This is the base version of the monitoring strategy, depending on the type,  the actual JSON may contain additional fields",
			Elem:        &hcl.Resource{Schema: new(strategy.Wrapper).Schema()},
		},
		"primary_dimension_key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Defines which dimension key should be used for the **alertingScope**",
		},
		"severity": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The type of the event to trigger on the threshold violation.  The `CUSTOM_ALERT` type is not correlated with other alerts. The `INFO` type does not open a problem",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *MetricEvent) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.MetricSelector != nil {
		result["metric_selector"] = me.MetricSelector
	}
	if me.MetricID != nil {
		result["metric_id"] = *me.MetricID
	}
	result["name"] = me.Name
	result["description"] = me.Description
	if me.AggregationType != nil {
		result["aggregation_type"] = string(*me.AggregationType)
	}
	// if me.WarningReason != nil {
	// 	result["warning_reason"] = string(*me.WarningReason)
	// }
	// if me.DisabledReason != nil {
	// 	result["disabled_reason"] = string(*me.DisabledReason)
	// }
	result["enabled"] = me.Enabled
	if me.PrimaryDimensionKey != nil {
		result["primary_dimension_key"] = *me.PrimaryDimensionKey
	}
	if me.Severity != nil {
		result["severity"] = string(*me.Severity)
	}
	if me.MetricDimensions != nil {
		if marshalled, err := me.MetricDimensions.MarshalHCL(hcl.NewDecoder(decoder, "dimensions", 0)); err == nil {
			result["dimensions"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.AlertingScope != nil {
		if marshalled, err := me.AlertingScope.MarshalHCL(hcl.NewDecoder(decoder, "scopes", 0)); err == nil {
			result["scopes"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.MonitoringStrategy != nil {
		wrapper := &strategy.Wrapper{Strategy: me.MonitoringStrategy}
		if marshalled, err := wrapper.MarshalHCL(hcl.NewDecoder(decoder, "strategy", 0)); err == nil {
			result["strategy"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *MetricEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "id")
		delete(me.Unknowns, "metricId")
		delete(me.Unknowns, "aggregationType")
		delete(me.Unknowns, "description")
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "warningReason")
		delete(me.Unknowns, "metricDimensions")
		delete(me.Unknowns, "disabledReason")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "alertingScope")
		delete(me.Unknowns, "monitoringStrategy")
		delete(me.Unknowns, "primaryDimensionKey")
		delete(me.Unknowns, "severity")
		delete(me.Unknowns, "metadata")
		delete(me.Unknowns, "metric_selector")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}

	if value, ok := decoder.GetOk("metric_selector"); ok {
		me.MetricSelector = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("metric_id"); ok {
		me.MetricID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("description"); ok {
		me.Description = value.(string)
	}
	if value, ok := decoder.GetOk("aggregation_type"); ok {
		me.AggregationType = AggregationType(value.(string)).Ref()
	}
	// if value, ok := decoder.GetOk("warning_reason"); ok {
	// 	me.WarningReason = WarningReason(value.(string)).Ref()
	// }
	if value, ok := decoder.GetOk("enabled"); ok {
		me.Enabled = value.(bool)
	}
	if value, ok := decoder.GetOk("primary_dimension_key"); ok {
		me.PrimaryDimensionKey = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("severity"); ok {
		me.Severity = Severity(value.(string)).Ref()
	}
	// if value, ok := decoder.GetOk("disabled_reason"); ok {
	// 	me.DisabledReason = DisabledReason(value.(string)).Ref()
	// }
	if _, ok := decoder.GetOk("strategy.#"); ok {
		cfg := new(strategy.Wrapper)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "strategy", 0)); err != nil {
			return err
		}
		me.MonitoringStrategy = cfg.Strategy
	}
	if _, ok := decoder.GetOk("dimensions.#"); ok {
		me.MetricDimensions = dimensions.Dimensions{}
		if err := me.MetricDimensions.UnmarshalHCL(hcl.NewDecoder(decoder, "dimensions", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("scopes.#"); ok {
		me.AlertingScope = scope.AlertingScopes{}
		if err := me.AlertingScope.UnmarshalHCL(hcl.NewDecoder(decoder, "scopes", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *MetricEvent) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":                  me.ID,
		"metricId":            me.MetricID,
		"aggregationType":     me.AggregationType,
		"description":         me.Description,
		"name":                me.Name,
		"warningReason":       me.WarningReason,
		"primaryDimensionKey": me.PrimaryDimensionKey,
		"severity":            me.Severity,
		"disabledReason":      me.DisabledReason,
		"enabled":             me.Enabled,
		"metricDimensions":    me.MetricDimensions,
		"alertingScope":       me.AlertingScope,
		"monitoringStrategy":  me.MonitoringStrategy,
		"metadata":            me.Metadata,
		"metricSelector":      me.MetricSelector,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *MetricEvent) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	wrapper := strategy.Wrapper{}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"id":                  &me.ID,
		"metricId":            &me.MetricID,
		"aggregationType":     &me.AggregationType,
		"description":         &me.Description,
		"name":                &me.Name,
		"warningReason":       &me.WarningReason,
		"primaryDimensionKey": &me.PrimaryDimensionKey,
		"severity":            &me.Severity,
		"disabledReason":      &me.DisabledReason,
		"enabled":             &me.Enabled,
		"metricDimensions":    &me.MetricDimensions,
		"alertingScope":       &me.AlertingScope,
		"monitoringStrategy":  &wrapper,
		"metadata":            &me.Metadata,
		"metricSelector":      &me.MetricSelector,
	}); err != nil {
		return err
	}
	me.MonitoringStrategy = wrapper.Strategy
	return nil
}

// AggregationType How the metric data points are aggregated for the evaluation.
//  The timeseries must support this aggregation.
type AggregationType string

func (me AggregationType) Ref() *AggregationType {
	return &me
}

// AggregationTypes offers the known enum values
var AggregationTypes = struct {
	Avg    AggregationType
	Count  AggregationType
	Max    AggregationType
	Median AggregationType
	Min    AggregationType
	P90    AggregationType
	Sum    AggregationType
	Value  AggregationType
}{
	"AVG",
	"COUNT",
	"MAX",
	"MEDIAN",
	"MIN",
	"P90",
	"SUM",
	"VALUE",
}

// WarningReason The reason of a warning set on the config.
// The `NONE` means config has no warnings.
type WarningReason string

func (me WarningReason) Ref() *WarningReason {
	return &me
}

// WarningReasons offers the known enum values
var WarningReasons = struct {
	None        WarningReason
	TooManyDims WarningReason
}{
	"NONE",
	"TOO_MANY_DIMS",
}

// DisabledReason The reason of automatic disabling.
// The `NONE` means config was not disabled automatically.
type DisabledReason string

func (me DisabledReason) Ref() *DisabledReason {
	return &me
}

// DisabledReasons offers the known enum values
var DisabledReasons = struct {
	MetricDefinitionInconsistency DisabledReason
	None                          DisabledReason
	TooManyDims                   DisabledReason
	TopxForciblyDeactivated       DisabledReason
}{
	"METRIC_DEFINITION_INCONSISTENCY",
	"NONE",
	"TOO_MANY_DIMS",
	"TOPX_FORCIBLY_DEACTIVATED",
}

// Severity The type of the event to trigger on the threshold violation.
// The `CUSTOM_ALERT` type is not correlated with other alerts.
// The `INFO` type does not open a problem.
type Severity string

func (me Severity) Ref() *Severity {
	return &me
}

// Severitys offers the known enum values
var Severitys = struct {
	Availability       Severity
	CustomAlert        Severity
	Error              Severity
	Info               Severity
	Performance        Severity
	ResourceContention Severity
}{
	"AVAILABILITY",
	"CUSTOM_ALERT",
	"ERROR",
	"INFO",
	"PERFORMANCE",
	"RESOURCE_CONTENTION",
}
