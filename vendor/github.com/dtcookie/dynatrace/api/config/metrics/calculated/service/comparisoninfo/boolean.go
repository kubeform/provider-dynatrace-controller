package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Boolean Comparison for `BOOLEAN` attributes.
type Boolean struct {
	BaseComparisonInfo
	Values     []bool            `json:"values,omitempty"` // The values to compare to.
	Comparison BooleanComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *bool             `json:"value,omitempty"`  // The value to compare to.
}

func (me *Boolean) GetType() Type {
	return Types.Boolean
}

func (me *Boolean) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to",
			Elem:        &hcl.Schema{Type: hcl.TypeBool},
		},
		"value": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The value to compare to",
		},
		"operator": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Operator of the comparison. You can reverse it by setting `negate` to `true`. Possible values are `EQUALS`, `EQUALS_ANY_OF` and `EXISTS`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Boolean) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"values":   me.Values,
		"value":    me.Value,
		"operator": me.Comparison,
		"unknowns": me.Unknowns,
	})
}

func (me *Boolean) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *Boolean) MarshalJSON() ([]byte, error) {
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

func (me *Boolean) UnmarshalJSON(data []byte) error {
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

// BooleanComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type BooleanComparison string

// BooleanComparisons offers the known enum values
var BooleanComparisons = struct {
	Equals      BooleanComparison
	EqualsAnyOf BooleanComparison
	Exists      BooleanComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}
