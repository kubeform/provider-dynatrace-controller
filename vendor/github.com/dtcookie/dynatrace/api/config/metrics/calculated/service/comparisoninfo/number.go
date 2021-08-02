package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Number Comparison for `NUMBER` attributes.
type Number struct {
	BaseComparisonInfo
	Comparison NumberComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *float64         `json:"value,omitempty"`  // The value to compare to.
	Values     []float64        `json:"values,omitempty"` // The values to compare to.
}

func (me *Number) GetType() Type {
	return Types.Number
}

func (me *Number) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to",
			Elem:        &hcl.Schema{Type: hcl.TypeFloat},
		},
		"value": {
			Type:        hcl.TypeFloat,
			Optional:    true,
			Description: "The value to compare to",
		},
		"operator": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Operator of the comparison. You can reverse it by setting `negate` to `true`. Possible values are `EQUALS`, `EQUALS_ANY_OF`, `EXISTS`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LOWER_THAN` and `LOWER_THAN_OR_EQUAL`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Number) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	if _, err = properties.EncodeAll(map[string]interface{}{
		"values":   me.Values,
		"value":    me.Value,
		"operator": me.Comparison,
		"unknowns": me.Unknowns,
	}); err != nil {
		return nil, err
	}
	// if len(me.Values) > 0 {
	// 	properties["values"] = me.Values
	// }
	return properties, nil
}

func (me *Number) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *Number) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":       me.GetType(),
		"negate":     me.Negate,
		"values":     me.Values,
		"value":      me.Value,
		"comparison": me.Comparison,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Number) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"negate":     &me.Negate,
		"values":     &me.Values,
		"value":      &me.Value,
		"comparison": &me.Comparison,
	})
}

// NumberComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type NumberComparison string

// NumberComparisons offers the known enum values
var NumberComparisons = struct {
	Equals             NumberComparison
	EqualsAnyOf        NumberComparison
	Exists             NumberComparison
	GreaterThan        NumberComparison
	GreaterThanOrEqual NumberComparison
	LowerThan          NumberComparison
	LowerThanOrEqual   NumberComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
}
