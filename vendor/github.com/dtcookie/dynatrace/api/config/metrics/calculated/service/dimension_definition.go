package service

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// DimensionDefinition Parameters of a definition of a calculated service metric.
type DimensionDefinition struct {
	TopXDirection   TopXDirection              `json:"topXDirection"`          // How to calculate the **topX** values.
	Dimension       string                     `json:"dimension"`              // The dimension value pattern.   You can define custom placeholders in the **placeholders** field and use them here.
	Name            string                     `json:"name"`                   // The name of the dimension.
	Placeholders    Placeholders               `json:"placeholders,omitempty"` // The list of custom placeholders to be used in a dimension value pattern.
	TopX            int32                      `json:"topX"`                   // The number of top values to be calculated.
	TopXAggregation TopXAggregation            `json:"topXAggregation"`        // The aggregation of the dimension.
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *DimensionDefinition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the dimension",
		},
		"dimension": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The dimension value pattern. You can define custom placeholders in the `placeholders` field and use them here",
		},
		"top_x_direction": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "How to calculate the **topX** values. Possible values are `ASCENDING` and `DESCENDING`",
		},
		"top_x": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "The number of top values to be calculated",
		},
		"top_x_aggregation": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The aggregation of the dimension. Possible values are `AVERAGE`, `COUNT`, `MAX`, `MIN`, `OF_INTEREST_RATIO`, `OTHER_RATIO`, `SINGLE_VALUE` and `SUM`",
		},
		"placeholders": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "The list of custom placeholders to be used in a dimension value pattern",
			Elem:        &hcl.Resource{Schema: new(Placeholders).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *DimensionDefinition) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"name":              me.Name,
		"dimension":         me.Dimension,
		"top_x_direction":   me.TopXDirection,
		"top_x":             me.TopX,
		"top_x_aggregation": me.TopXAggregation,
		"placeholders":      me.Placeholders,
		"unknowns":          me.Unknowns,
	})
}

func (me *DimensionDefinition) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":              &me.Name,
		"dimension":         &me.Dimension,
		"top_x_direction":   &me.TopXDirection,
		"top_x":             &me.TopX,
		"top_x_aggregation": &me.TopXAggregation,
		"placeholders":      &me.Placeholders,
		"unknowns":          &me.Unknowns,
	})
}

func (me *DimensionDefinition) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"name":            me.Name,
		"dimension":       me.Dimension,
		"placeholders":    me.Placeholders,
		"topXDirection":   me.TopXDirection,
		"topX":            me.TopX,
		"topXAggregation": me.TopXAggregation,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *DimensionDefinition) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"name":            &me.Name,
		"dimension":       &me.Dimension,
		"placeholders":    &me.Placeholders,
		"topXDirection":   &me.TopXDirection,
		"topX":            &me.TopX,
		"topXAggregation": &me.TopXAggregation,
	})
}

// TopXDirection How to calculate the **topX** values.
type TopXDirection string

// TopXDirections offers the known enum values
var TopXDirections = struct {
	Ascending  TopXDirection
	Descending TopXDirection
}{
	"ASCENDING",
	"DESCENDING",
}

// TopXAggregation The aggregation of the dimension.
type TopXAggregation string

// TopXAggregations offers the known enum values
var TopXAggregations = struct {
	Average         TopXAggregation
	Count           TopXAggregation
	Max             TopXAggregation
	Min             TopXAggregation
	OfInterestRatio TopXAggregation
	OtherRatio      TopXAggregation
	SingleValue     TopXAggregation
	Sum             TopXAggregation
}{
	"AVERAGE",
	"COUNT",
	"MAX",
	"MIN",
	"OF_INTEREST_RATIO",
	"OTHER_RATIO",
	"SINGLE_VALUE",
	"SUM",
}
