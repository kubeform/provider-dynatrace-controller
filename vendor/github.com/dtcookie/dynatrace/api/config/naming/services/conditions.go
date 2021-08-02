package services

import (
	"github.com/dtcookie/dynatrace/api/config/entityruleengine"
	"github.com/dtcookie/hcl"
)

type Conditions []*entityruleengine.Condition

func (me *Conditions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"condition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A conditions for the metric usage",
			Elem:        &hcl.Resource{Schema: new(entityruleengine.Condition).Schema()},
		},
	}
}

func (me Conditions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("condition", me)
}

func (me *Conditions) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("condition", me); err != nil {
		return err
	}
	return nil
}
