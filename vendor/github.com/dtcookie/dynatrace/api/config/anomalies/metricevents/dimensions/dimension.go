package dimensions

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// MetricEventDimension A single filter for the metrics dimensions.
// This is the base version of the filter, depending on the type,
// the actual JSON may contain additional fields.
type Dimension interface {
	GetType() FilterType
}

// BaseDimension A single filter for the metrics dimensions.
// This is the base version of the filter, depending on the type,
// the actual JSON may contain additional fields.
type BaseDimension struct {
	FilterType FilterType                 `json:"filterType"`    // Defines the actual set of fields depending on the value. See one of the following objects:  * `ENTITY` -> MetricEventEntityDimensions  * `STRING` -> MetricEventStringDimensions
	Key        *string                    `json:"key,omitempty"` // The dimensions key on the metric.
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (me *BaseDimension) GetType() FilterType {
	return me.FilterType
}

func (me *BaseDimension) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The dimensions key on the metric",
		},
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

func (me *BaseDimension) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.Key != nil {
		result["key"] = *me.Key
	}
	result["type"] = string(me.FilterType)
	return result, nil
}

func (me *BaseDimension) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "filterType")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("type"); ok {
		me.FilterType = FilterType(value.(string))
	}
	return nil
}

func (me *BaseDimension) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.FilterType,
		"key":        me.Key,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *BaseDimension) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
		"key":        &me.Key,
	}); err != nil {
		return err
	}
	return nil
}

// DimensionFilterType Defines the actual set of fields depending on the value. See one of the following objects:
// * `ENTITY` -> MetricEventEntityDimensions
// * `STRING` -> MetricEventStringDimensions
type FilterType string

// FilterTypes offers the known enum values
var FilterTypes = struct {
	Entity FilterType
	String FilterType
}{
	"ENTITY",
	"STRING",
}
