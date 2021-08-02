package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// TagInfo Tag of a Dynatrace entity.
type TagInfo struct {
	Context TagInfoContext `json:"context"`         // The origin of the tag, such as AWS or Cloud Foundry.   Custom tags use the `CONTEXTLESS` value.
	Key     string         `json:"key"`             // The key of the tag.   Custom tags have the tag value here.
	Value   *string        `json:"value,omitempty"` // The value of the tag.   Not applicable to custom tags.
}

func (me *TagInfo) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The key of the tag. Custom tags have the tag value here",
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value of the tag. Not applicable to custom tags",
		},
		"context": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The origin of the tag, such as AWS or Cloud Foundry. Custom tags use the `CONTEXTLESS` value. Possible values are `AWS`, `AWS_GENERIC`, `AZURE`, `CLOUD_FOUNDRY`, `CONTEXTLESS`, `ENVIRONMENT`, `GOOGLE_CLOUD` and `KUBERNETES`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *TagInfo) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"key":     me.Key,
		"value":   me.Value,
		"context": me.Context,
	})
}

func (me *TagInfo) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"key":     &me.Key,
		"value":   &me.Value,
		"context": &me.Context,
	})
}

func (me *TagInfo) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"key":     me.Key,
		"value":   me.Value,
		"context": me.Context,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *TagInfo) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"key":     &me.Key,
		"value":   &me.Value,
		"context": &me.Context,
	})
}

// TagInfoContext The origin of the tag, such as AWS or Cloud Foundry.
//  Custom tags use the `CONTEXTLESS` value.
type TagInfoContext string

// TagInfoContexts offers the known enum values
var TagInfoContexts = struct {
	AWS          TagInfoContext
	AWSGeneric   TagInfoContext
	Azure        TagInfoContext
	CloudFoundry TagInfoContext
	Contextless  TagInfoContext
	Environment  TagInfoContext
	GoogleCloud  TagInfoContext
	Kubernetes   TagInfoContext
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

type TagInfos []*TagInfo

func (me TagInfos) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"value": {
			Type:        hcl.TypeList,
			MinItems:    1,
			Optional:    true,
			Description: "The values to compare to",
			Elem:        &hcl.Resource{Schema: new(TagInfo).Schema()},
		},
	}
}

func (me TagInfos) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("value", me)
}

func (me *TagInfos) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("value", me); err != nil {
		return err
	}
	return nil
}
