package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// FlawState Comparison for `FLAW_STATE` attributes.
type FlawState struct {
	BaseComparisonInfo
	Comparison FlawStateComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *FlawStateValue     `json:"value,omitempty"`  // The value to compare to.
	Values     []FlawStateValue    `json:"values,omitempty"` // The values to compare to.
}

func (me *FlawState) GetType() Type {
	return Types.FlawState
}

func (me *FlawState) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to. Possible values are `FLAWED` and `NOT_FLAWED`",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare to. Possible values are `FLAWED` and `NOT_FLAWED`",
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

func (me *FlawState) MarshalHCL() (map[string]interface{}, error) {
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

func (me *FlawState) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *FlawState) MarshalJSON() ([]byte, error) {
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

func (me *FlawState) UnmarshalJSON(data []byte) error {
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

// FlawStateComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FlawStateComparison string

// FlawStateComparisons offers the known enum values
var FlawStateComparisons = struct {
	Equals      FlawStateComparison
	EqualsAnyOf FlawStateComparison
	Exists      FlawStateComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// FlawStateValue The value to compare to.
type FlawStateValue string

// FlawStateValues offers the known enum values
var FlawStateValues = struct {
	Flawed    FlawStateValue
	NotFlawed FlawStateValue
}{
	"FLAWED",
	"NOT_FLAWED",
}
