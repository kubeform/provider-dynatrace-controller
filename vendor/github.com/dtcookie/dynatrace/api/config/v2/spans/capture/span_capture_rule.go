package capture

import (
	"github.com/dtcookie/dynatrace/api/config/v2/spans/match"
	"github.com/dtcookie/hcl"
)

// SpanCaptureRule has no documentation
type SpanCaptureRule struct {
	Name     string             `json:"ruleName"`
	Action   SpanCaptureAction  `json:"ruleAction"`
	Matchers match.SpanMatchers `json:"matchers" min:"1" max:"100"`
}

func (me *SpanCaptureRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name of the rule",
		},
		"action": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Whether to create an entry point or not",
		},
		"matches": {
			Type:        hcl.TypeList,
			MinItems:    1,
			MaxItems:    1,
			Required:    true,
			Description: "Matching strategies for the Span",
			Elem:        &hcl.Resource{Schema: new(match.SpanMatchers).Schema()},
		},
	}
}

func (me *SpanCaptureRule) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"name":    me.Name,
		"action":  me.Action,
		"matches": me.Matchers,
	})
}

func (me *SpanCaptureRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":    &me.Name,
		"action":  &me.Action,
		"matches": &me.Matchers,
	})
}
