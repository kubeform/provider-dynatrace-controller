package maintenance

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// TagInfo Tag of a Dynatrace entity.
type TagInfo struct {
	Context  Context                    `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value.
	Key      string                     `json:"key"`             // The key of the tag. Custom tags have the tag value here.
	Value    *string                    `json:"value,omitempty"` // The value of the tag. Not applicable to custom tags.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *TagInfo) Schema() map[string]*hcl.Schema {
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
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *TagInfo) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["context"] = string(me.Context)
	result["key"] = string(me.Key)
	if me.Value != nil {
		result["value"] = *me.Value
	}

	return result, nil
}

func (me *TagInfo) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "context")
		delete(me.Unknowns, "key")
		delete(me.Unknowns, "value")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
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

func (me *TagInfo) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("context", me.Context); err != nil {
		return nil, err
	}
	if err := m.Marshal("key", me.Key); err != nil {
		return nil, err
	}
	if err := m.Marshal("value", me.Value); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *TagInfo) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("context", &me.Context); err != nil {
		return err
	}
	if err := m.Unmarshal("key", &me.Key); err != nil {
		return err
	}
	if err := m.Unmarshal("value", &me.Value); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
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
