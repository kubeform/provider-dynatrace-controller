package customservices

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type FileSection struct {
	Name  *string
	Match *FileNameMatcher
}

func (me *FileSection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The full name of the file / the name to match the file name with",
			Required:    true,
		},
		"match": {
			Type:        hcl.TypeString,
			Description: "Matcher applying to the file name (ENDS_WITH, EQUALS or STARTS_WITH). Default value is ENDS_WITH (if applicable)",
			Optional:    true,
		},
	}
}

func (me *FileSection) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Name != nil {
		result["name"] = opt.String(me.Name)
	}
	if me.Match != nil {
		result["match"] = string(*me.Match)
	}
	return result, nil
}

func (me *FileSection) UnmarshalHCL(decoder hcl.Decoder) error {
	adapter := hcl.Adapt(decoder)
	me.Name = adapter.GetString("name")
	if value := adapter.GetString("match"); value != nil {
		me.Match = FileNameMatcher(*value).Ref()
	}
	return nil
}
