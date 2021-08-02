package tag

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// Info Tag of a Dynatrace entity.
type Info struct {
	Context  Context                    `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value.
	Key      string                     `json:"key"`             // The key of the tag. Custom tags have the tag value here.
	Value    *string                    `json:"value,omitempty"` // The value of the tag. Not applicable to custom tags.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (ti *Info) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"context": {
			Type:        hcl.TypeString,
			Description: "The origin of the tag, such as AWS or Cloud Foundry. Possible values are AWS, AWS_GENERIC, AZURE, CLOUD_FOUNDRY, CONTEXTLESS, ENVIRONMENT, GOOGLE_CLOUD and KUBERNETES. Custom tags use the `CONTEXTLESS` value",
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
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider but have meanwhile gotten introduced by a newer version of the Dynatrace REST API",
			Optional:    true,
		},
	}
}

func (ti *Info) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ti.Unknowns) > 0 {
		data, err := json.Marshal(ti.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["context"] = string(ti.Context)
	result["key"] = ti.Key
	if ti.Value != nil {
		result["value"] = opt.String(ti.Value)
	}
	return result, nil
}

func (ti *Info) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ti); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ti.Unknowns); err != nil {
			return err
		}
		delete(ti.Unknowns, "context")
		delete(ti.Unknowns, "key")
		delete(ti.Unknowns, "value")
		if len(ti.Unknowns) == 0 {
			ti.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("context"); ok {
		ti.Context = Context(value.(string))
	}
	if _, value := decoder.GetChange("key"); value != nil {
		ti.Key = value.(string)
	}
	if value, ok := decoder.GetOk("value"); ok {
		ti.Value = opt.NewString(value.(string))
	}
	return nil
}

// Context The origin of the tag, such as AWS or Cloud Foundry.
//  Custom tags use the `CONTEXTLESS` value.
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
