package common

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// TagFilter A tag-based filter of monitored entities.
type TagFilter struct {
	Context Context `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.  Custom tags use the `CONTEXTLESS` value.
	Key     string  `json:"key"`             // The key of the tag. Custom tags have the tag value here.
	Value   *string `json:"value,omitempty"` // The value of the tag. Not applicable to custom tags.
}

func (me *TagFilter) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"context": {
			Type:        hcl.TypeString,
			Description: "The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "The key of the tag. Custom tags have the tag value here",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the tag. Not applicable to custom tags",
			Optional:    true,
		},
	}
}

func (me *TagFilter) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["context"] = string(me.Context)
	result["key"] = me.Key
	if me.Value != nil {
		result["value"] = *me.Value
	}
	return result, nil
}

func (me *TagFilter) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("context"); ok {
		me.Context = Context(value.(string))
	}
	if value, ok := decoder.GetOk("key"); ok {
		me.Key = value.(string)
	}
	if value, ok := decoder.GetOk("value"); ok {
		me.Value = opt.NewString(value.(string))
	}
	return nil
}

// Context The origin of the tag, such as AWS or Cloud Foundry.
// Custom tags use the `CONTEXTLESS` value.
type Context string

// Contexts offers the known enum values
var Contexts = struct {
	AWS          Context
	AWSGeneric   Context
	Azure        Context
	CloudFoundry Context
	Contextless  Context
	Environment  Context
	GoogleCloud  Context
	Kubernetes   Context
}{
	"AWS",
	"AWS_GENERIC",
	"AZURE",
	"CLOUD_FOUNDRY",
	"CONTEXTLESS",
	"ENVIRONMENT",
	"GOOGLE_CLOUD",
	"KUBERNETES",
}
