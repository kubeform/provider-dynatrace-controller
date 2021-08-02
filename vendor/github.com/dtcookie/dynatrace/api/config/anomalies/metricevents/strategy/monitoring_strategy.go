package strategy

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// MonitoringStrategy A monitoring strategy for a metric event config.
// This is the base version of the monitoring strategy, depending on the type,
// the actual JSON may contain additional fields.
type MonitoringStrategy interface {
	GetType() Type
}

// BaseMetricEventMonitoringStrategy A monitoring strategy for a metric event config.
// This is the base version of the monitoring strategy, depending on the type,
// the actual JSON may contain additional fields.
type BaseMonitoringStrategy struct {
	Type     Type                       `json:"type"` // Defines the actual set of fields depending on the value. See one of the following objects:  * `STATIC_THRESHOLD` -> MetricEventStaticThresholdMonitoringStrategy  * `AUTO_ADAPTIVE_BASELINE` -> MetricEventAutoAdaptiveBaselineMonitoringStrategy
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *BaseMonitoringStrategy) GetType() Type {
	return me.Type
}

func (me *BaseMonitoringStrategy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Defines the actual set of fields depending on the value",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *BaseMonitoringStrategy) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["type"] = string(me.Type)
	return result, nil
}

func (me *BaseMonitoringStrategy) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "type")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		me.Type = Type(value.(string))
	}
	return nil
}

func (me *BaseMonitoringStrategy) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type": me.Type,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *BaseMonitoringStrategy) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"type": &me.Type,
	}); err != nil {
		return err
	}
	return nil
}

// Type Defines the actual set of fields depending on the value. See one of the following objects:
// * `STATIC_THRESHOLD` -> MetricEventStaticThresholdMonitoringStrategy
// * `AUTO_ADAPTIVE_BASELINE` -> MetricEventAutoAdaptiveBaselineMonitoringStrategy
type Type string

// Types offers the known enum values
var Types = struct {
	AutoAdaptiveBaseline Type
	StaticThreshold      Type
}{
	"AUTO_ADAPTIVE_BASELINE",
	"STATIC_THRESHOLD",
}

// AlertCondition The condition for the **threshold** value check: above or below.
type AlertCondition string

// AlertConditions offers the known enum values
var AlertConditions = struct {
	Above AlertCondition
	Below AlertCondition
}{
	"ABOVE",
	"BELOW",
}
