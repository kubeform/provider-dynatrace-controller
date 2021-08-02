package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// String Comparison for `STRING` attributes.
type String struct {
	BaseComparisonInfo
	CaseSensitive *bool            `json:"caseSensitive,omitempty"` // The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	Comparison    StringComparison `json:"comparison"`              // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value         *string          `json:"value,omitempty"`         // The value to compare to.
	Values        []string         `json:"values,omitempty"`        // The values to compare to.
}

func (me *String) GetType() Type {
	return Types.String
}

func (me *String) Schema() map[string]*hcl.Schema {
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
			Description: "Operator of the comparison. You can reverse it by setting `negate` to `true`. Possible values are `BEGINS_WITH`, `BEGINS_WITH_ANY_OF`, `CONTAINS`, `ENDS_WITH`, `ENDS_WITH_ANY_OF`, `EQUALS`, `EQUALS_ANY_OF`, `EXISTS` and `REGEX_MATCHES`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *String) MarshalHCL() (map[string]interface{}, error) {
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

func (me *String) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":         &me.Values,
		"value":          &me.Value,
		"operator":       &me.Comparison,
		"case_sensitive": &me.CaseSensitive,
		"unknowns":       &me.Unknowns,
	})
}

func (me *String) MarshalJSON() ([]byte, error) {
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

func (me *String) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	err := properties.UnmarshalAll(map[string]interface{}{
		"negate":        &me.Negate,
		"values":        &me.Values,
		"value":         &me.Value,
		"comparison":    &me.Comparison,
		"caseSensitive": &me.CaseSensitive,
	})
	return err
}

// StringComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type StringComparison string

// StringComparisons offers the known enum values
var StringComparisons = struct {
	BeginsWith      StringComparison
	BeginsWithAnyOf StringComparison
	Contains        StringComparison
	EndsWith        StringComparison
	EndsWithAnyOf   StringComparison
	Equals          StringComparison
	EqualsAnyOf     StringComparison
	Exists          StringComparison
	RegexMatches    StringComparison
}{
	"BEGINS_WITH",
	"BEGINS_WITH_ANY_OF",
	"CONTAINS",
	"ENDS_WITH",
	"ENDS_WITH_ANY_OF",
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
	"REGEX_MATCHES",
}
