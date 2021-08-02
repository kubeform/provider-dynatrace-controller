package ctxprop

import (
	"github.com/dtcookie/hcl"
)

// PropagationSetting Context propagation enables you to connect PurePaths through OpenTelemetry/OpenTracing. Define rules to enable context propagation for certain spans within OneAgent
type PropagationSetting struct {
	PropagationRule *PropagationRule `json:"contextPropagationRule"`
}

func (me *PropagationSetting) Schema() map[string]*hcl.Schema {
	return new(PropagationRule).Schema()
}

func (me *PropagationSetting) MarshalHCL() (map[string]interface{}, error) {
	return me.PropagationRule.MarshalHCL()
}

func (me *PropagationSetting) UnmarshalHCL(decoder hcl.Decoder) error {
	me.PropagationRule = new(PropagationRule)
	return me.PropagationRule.UnmarshalHCL(decoder)
}
