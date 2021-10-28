package attributes

import (
	"github.com/dtcookie/hcl"
)

// SpanAttribute has no documentation
type SpanAttribute struct {
	Key string `json:"key"`
}

func (me *SpanAttribute) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "the key of the attribute to capture",
		},
	}
}

func (me *SpanAttribute) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{
		"key": me.Key,
	}, nil
}

func (me *SpanAttribute) UnmarshalHCL(decoder hcl.Decoder) error {
	if key, ok := decoder.GetOk("key"); ok {
		me.Key = key.(string)
	}
	return nil
}
