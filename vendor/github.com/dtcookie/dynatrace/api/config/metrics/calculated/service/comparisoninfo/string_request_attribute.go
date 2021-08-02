package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/metrics/calculated/service/propagation"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// StringRequestAttribute Comparison for `STRING_REQUEST_ATTRIBUTE` attributes, specifically of the **String** type.
type StringRequestAttribute struct {
	BaseComparisonInfo
	Source            *propagation.Source              `json:"source,omitempty"`            // Defines valid sources of request attributes for conditions or placeholders.
	Value             *string                          `json:"value,omitempty"`             // The value to compare to.
	Values            []string                         `json:"values,omitempty"`            // The values to compare to.
	CaseSensitive     *bool                            `json:"caseSensitive,omitempty"`     // The comparison is case-sensitive (`true`) or not case-sensitive (`false`).
	Comparison        StringRequestAttributeComparison `json:"comparison"`                  // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	MatchOnChildCalls *bool                            `json:"matchOnChildCalls,omitempty"` // If `true`, the request attribute is matched on child service calls.   Default is `false`.
	RequestAttribute  string                           `json:"requestAttribute"`            // has no documentation
}

func (me *StringRequestAttribute) GetType() Type {
	return Types.StringRequestAttribute
}

func (me *StringRequestAttribute) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The comparison is case-sensitive (`true`) or not case-sensitive (`false`)",
		},
		"match_on_child_calls": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "If `true`, the request attribute is matched on child service calls. Default is `false`",
		},
		"source": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "Defines valid sources of request attributes for conditions or placeholders",
			Elem:        &hcl.Resource{Schema: new(propagation.Source).Schema()},
		},
		"request_attribute": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "No documentation available for this attribute",
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

func (me *StringRequestAttribute) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"values":               me.Values,
		"value":                me.Value,
		"operator":             me.Comparison,
		"case_sensitive":       me.CaseSensitive,
		"match_on_child_calls": me.MatchOnChildCalls,
		"request_attribute":    me.RequestAttribute,
		"source":               me.Source,
		"unknowns":             me.Unknowns,
	})
}

func (me *StringRequestAttribute) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":               &me.Values,
		"value":                &me.Value,
		"operator":             &me.Comparison,
		"case_sensitive":       &me.CaseSensitive,
		"match_on_child_calls": &me.MatchOnChildCalls,
		"request_attribute":    &me.RequestAttribute,
		"source":               &me.Source,
		"unknowns":             &me.Unknowns,
	})
}

func (me *StringRequestAttribute) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":              me.GetType(),
		"negate":            me.Negate,
		"values":            me.Values,
		"value":             me.Value,
		"comparison":        me.Comparison,
		"caseSensitive":     me.CaseSensitive,
		"matchOnChildCalls": me.MatchOnChildCalls,
		"requestAttribute":  me.RequestAttribute,
		"source":            me.Source,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *StringRequestAttribute) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"negate":            &me.Negate,
		"values":            &me.Values,
		"value":             &me.Value,
		"comparison":        &me.Comparison,
		"caseSensitive":     &me.CaseSensitive,
		"matchOnChildCalls": &me.MatchOnChildCalls,
		"requestAttribute":  &me.RequestAttribute,
		"source":            &me.Source,
	})
}

// StringRequestAttributeComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type StringRequestAttributeComparison string

// StringRequestAttributeComparisons offers the known enum values
var StringRequestAttributeComparisons = struct {
	BeginsWith      StringRequestAttributeComparison
	BeginsWithAnyOf StringRequestAttributeComparison
	Contains        StringRequestAttributeComparison
	EndsWith        StringRequestAttributeComparison
	EndsWithAnyOf   StringRequestAttributeComparison
	Equals          StringRequestAttributeComparison
	EqualsAnyOf     StringRequestAttributeComparison
	Exists          StringRequestAttributeComparison
	RegexMatches    StringRequestAttributeComparison
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
