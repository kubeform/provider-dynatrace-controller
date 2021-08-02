package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// FailedState Comparison for `FAILED_STATE` attributes.
type FailedState struct {
	BaseComparisonInfo
	Comparison FailedStateComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *FailedStateValue     `json:"value,omitempty"`  // The value to compare to.
	Values     []FailedStateValue    `json:"values,omitempty"` // The values to compare to.
}

func (me *FailedState) GetType() Type {
	return Types.FailedState
}

func (me *FailedState) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to. Possible values are `FAILED` and `FAILED`",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare to. Possible values are `FAILED` and `FAILED`",
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

func (me *FailedState) MarshalHCL() (map[string]interface{}, error) {
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

func (me *FailedState) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *FailedState) MarshalJSON() ([]byte, error) {
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

func (me *FailedState) UnmarshalJSON(data []byte) error {
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

// FailedStateValue The value to compare to.
type FailedStateValue string

// FailedStateValues offers the known enum values
var FailedStateValues = struct {
	Failed     FailedStateValue
	Successful FailedStateValue
}{
	"FAILED",
	"SUCCESSFUL",
}

// FailedStateComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FailedStateComparison string

// FailedStateComparisons offers the known enum values
var FailedStateComparisons = struct {
	Equals      FailedStateComparison
	EqualsAnyOf FailedStateComparison
	Exists      FailedStateComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}
