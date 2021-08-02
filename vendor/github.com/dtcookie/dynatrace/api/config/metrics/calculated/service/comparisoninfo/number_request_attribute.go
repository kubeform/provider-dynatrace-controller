package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/metrics/calculated/service/propagation"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// NumberRequestAttribute Comparison for `NUMBER_REQUEST_ATTRIBUTE` attributes, specifically of the generic **Number** type.
type NumberRequestAttribute struct {
	BaseComparisonInfo
	MatchOnChildCalls *bool                            `json:"matchOnChildCalls,omitempty"` // If `true`, the request attribute is matched on child service calls.    Default is `false`.
	RequestAttribute  string                           `json:"requestAttribute"`            // has no documentation
	Source            *propagation.Source              `json:"source,omitempty"`            // Defines valid sources of request attributes for conditions or placeholders.
	Value             *float64                         `json:"value,omitempty"`             // The value to compare to.
	Values            []float64                        `json:"values,omitempty"`            // The values to compare to.
	Comparison        NumberRequestAttributeComparison `json:"comparison"`                  // Operator of the comparision. You can reverse it by setting **negate** to `true`.
}

func (me *NumberRequestAttribute) GetType() Type {
	return Types.NumberRequestAttribute
}

func (me *NumberRequestAttribute) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
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

func (me *NumberRequestAttribute) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"values":               me.Values,
		"value":                me.Value,
		"operator":             me.Comparison,
		"match_on_child_calls": me.MatchOnChildCalls,
		"request_attribute":    me.RequestAttribute,
		"source":               me.Source,
		"unknowns":             me.Unknowns,
	})
}

func (me *NumberRequestAttribute) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":               &me.Values,
		"value":                &me.Value,
		"operator":             &me.Comparison,
		"match_on_child_calls": &me.MatchOnChildCalls,
		"request_attribute":    &me.RequestAttribute,
		"source":               &me.Source,
		"unknowns":             &me.Unknowns,
	})
}

func (me *NumberRequestAttribute) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":              me.GetType(),
		"negate":            me.Negate,
		"values":            me.Values,
		"value":             me.Value,
		"comparison":        me.Comparison,
		"matchOnChildCalls": me.MatchOnChildCalls,
		"requestAttribute":  me.RequestAttribute,
		"source":            me.Source,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *NumberRequestAttribute) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"negate":            &me.Negate,
		"values":            &me.Values,
		"value":             &me.Value,
		"comparison":        &me.Comparison,
		"matchOnChildCalls": &me.MatchOnChildCalls,
		"requestAttribute":  &me.RequestAttribute,
		"source":            &me.Source,
	})
}

// NumberRequestAttributeComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type NumberRequestAttributeComparison string

// NumberRequestAttributeComparisons offers the known enum values
var NumberRequestAttributeComparisons = struct {
	Equals             NumberRequestAttributeComparison
	EqualsAnyOf        NumberRequestAttributeComparison
	Exists             NumberRequestAttributeComparison
	GreaterThan        NumberRequestAttributeComparison
	GreaterThanOrEqual NumberRequestAttributeComparison
	LowerThan          NumberRequestAttributeComparison
	LowerThanOrEqual   NumberRequestAttributeComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"EXISTS",
	"GREATER_THAN",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN",
	"LOWER_THAN_OR_EQUAL",
}
