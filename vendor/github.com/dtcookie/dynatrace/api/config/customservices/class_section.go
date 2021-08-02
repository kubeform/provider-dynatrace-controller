package customservices

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type ClassSection struct {
	Name  *string
	Match *ClassNameMatcher
}

func (me *ClassSection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The full name of the class / the name to match the class name with",
			Required:    true,
		},
		"match": {
			Type:        hcl.TypeString,
			Description: "Matcher applying to the class name (ENDS_WITH, EQUALS or STARTS_WITH). STARTS_WITH can only be used if there is at least one annotation defined. Default value is EQUALS",
			Optional:    true,
			Default:     "EQUALS",
		},
	}
}

func (me *ClassSection) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Name != nil {
		result["name"] = opt.String(me.Name)
	}
	if me.Name != nil {
		result["match"] = string(*me.Match)
	}
	return result, nil
}

func (me *ClassSection) UnmarshalHCL(decoder hcl.Decoder) error {
	adapter := hcl.Adapt(decoder)
	me.Name = adapter.GetString("name")
	if value := adapter.GetString("match"); value != nil {
		me.Match = ClassNameMatcher(*value).Ref()
	}
	return nil
}
