package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// FastString Comparison for `FAST_STRING` attributes. Use it for all service property attributes.
type FastString struct {
	BaseComparisonInfo
	CaseSensitive *bool                `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	Comparison    FastStringComparison `json:"comparison"`              // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value         *string              `json:"value,omitempty"`         // The value to compare to.
	Values        []string             `json:"values,omitempty"`        // The values to compare to.
}

func (me *FastString) GetType() Type {
	return Types.FastString
}

func (me *FastString) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The comparison is case-sensitive (`true`) or not case-sensitive (`false`)",
		},
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The values to compare to",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare to",
		},
		"operator": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Operator of the comparison. You can reverse it by setting `negate` to `true`. Possible values are `EQUALS`, `EQUALS_ANY_OF` and `CONTAINS`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *FastString) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"values":         me.Values,
		"value":          me.Value,
		"operator":       me.Comparison,
		"case_sensitive": me.CaseSensitive,
		"unknowns":       me.Unknowns,
	})
}

func (me *FastString) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":         &me.Values,
		"value":          &me.Value,
		"operator":       &me.Comparison,
		"case_sensitive": &me.CaseSensitive,
		"unknowns":       &me.Unknowns,
	})
}

func (me *FastString) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":          me.GetType(),
		"negate":        me.Negate,
		"values":        me.Values,
		"value":         me.Value,
		"comparison":    me.Comparison,
		"caseSensitive": me.CaseSensitive,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *FastString) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"negate":        &me.Negate,
		"values":        &me.Values,
		"value":         &me.Value,
		"comparison":    &me.Comparison,
		"caseSensitive": &me.CaseSensitive,
	})
}

// FastStringComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type FastStringComparison string

// FastStringComparisons offers the known enum values
var FastStringComparisons = struct {
	Contains    FastStringComparison
	Equals      FastStringComparison
	EqualsAnyOf FastStringComparison
}{
	"CONTAINS",
	"EQUALS",
	"EQUALS_ANY_OF",
}
