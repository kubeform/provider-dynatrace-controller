package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// FailureReason Comparison for `FAILURE_REASON` attributes.
type FailureReason struct {
	BaseComparisonInfo
	Comparison FailureReasonComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *FailureReasonValue     `json:"value,omitempty"`  // The value to compare to.
	Values     []FailureReasonValue    `json:"values,omitempty"` // The values to compare to.
}

func (me *FailureReason) GetType() Type {
	return Types.FailureReason
}

func (me *FailureReason) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to. Possible values are `EXCEPTION_AT_ENTRY_NODE`, `EXCEPTION_ON_ANY_NODE`, `HTTP_CODE` and `REQUEST_ATTRIBUTE`",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare to. Possible values are `EXCEPTION_AT_ENTRY_NODE`, `EXCEPTION_ON_ANY_NODE`, `HTTP_CODE` and `REQUEST_ATTRIBUTE`",
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

func (me *FailureReason) MarshalHCL() (map[string]interface{}, error) {
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

func (me *FailureReason) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *FailureReason) MarshalJSON() ([]byte, error) {
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

func (me *FailureReason) UnmarshalJSON(data []byte) error {
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

// FailureReasonComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FailureReasonComparison string

// FailureReasonComparisons offers the known enum values
var FailureReasonComparisons = struct {
	Equals      FailureReasonComparison
	EqualsAnyOf FailureReasonComparison
	Exists      FailureReasonComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
}

// FailureReasonValue The value to compare to.
type FailureReasonValue string

// FailureReasonValues offers the known enum values
var FailureReasonValues = struct {
	ExceptionAtEntryNode FailureReasonValue
	ExceptionOnAnyNode   FailureReasonValue
	HTTPCode             FailureReasonValue
	RequestAttribute     FailureReasonValue
}{
	"EXCEPTION_AT_ENTRY_NODE",
	"EXCEPTION_ON_ANY_NODE",
	"HTTP_CODE",
	"REQUEST_ATTRIBUTE",
}
