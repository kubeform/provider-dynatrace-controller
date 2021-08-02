package propagation

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// UniversalTag Use only request attributes from services that have this tag. Use either this or `managementZone`
type UniversalTag struct {
	Key     string               `json:"key"`               // The key of the tag. For custom tags, put the tag value here. The key allows categorization of multiple tags. It is possible that there are multiple values for a single key which will all be represented as standalone tags. Therefore, the key does not have the semantic of a map key but is more like a key of a key-value tuple. In some cases, for example custom tags, the key represents the actual tag value and the value field is not set – those are called valueless tags.
	TagKey  *UniversalTagKey     `json:"tagKey,omitempty"`  // has no documentation
	Value   *string              `json:"value,omitempty"`   // The value of the tag. Not applicable to custom tags.  If a tag does have a separate key and value (in the textual representation they are split by the colon ‘:’), this field is set with the actual value. Key-value pairs can occur for automatically imported tags and tags set by rules if extractors are used.
	Context *UniversalTagContext `json:"context,omitempty"` // The origin of the tag, such as AWS or Cloud Foundry. For custom tags use the `CONTEXTLESS` value.  The context is set for tags that are automatically imported by OneAgent (for example, from the AWS console or environment variables). It’s useful for determining the origin of tags when not manually defined, and it also helps to prevent clashes with other existing tags. If the tag is not automatically imported, `CONTEXTLESS` set.
}

func (me *UniversalTag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The key of the tag. For custom tags, put the tag value here. The key allows categorization of multiple tags. It is possible that there are multiple values for a single key which will all be represented as standalone tags. Therefore, the key does not have the semantic of a map key but is more like a key of a key-value tuple. In some cases, for example custom tags, the key represents the actual tag value and the value field is not set – those are called valueless tags",
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value of the tag. Not applicable to custom tags. If a tag does have a separate key and value (in the textual representation they are split by the colon ‘:’), this field is set with the actual value. Key-value pairs can occur for automatically imported tags and tags set by rules if extractors are used",
		},
		"tag_key": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "has no documentation",
			Elem:        &hcl.Resource{Schema: new(UniversalTagKey).Schema()},
		},
		"context": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The origin of the tag, such as AWS or Cloud Foundry. For custom tags use the `CONTEXTLESS` value. The context is set for tags that are automatically imported by OneAgent (for example, from the AWS console or environment variables). It’s useful for determining the origin of tags when not manually defined, and it also helps to prevent clashes with other existing tags. If the tag is not automatically imported, `CONTEXTLESS` set. Possible values are `AWS`, `AWS_GENERIC`, `AZURE`, `CLOUD_FOUNDRY`, `CONTEXTLESS`, `ENVIRONMENT`, `GOOGLE_COMPUTE_ENGINE` and `KUBERNETES`",
		},
	}
}

func (me *UniversalTag) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"key":     me.Key,
		"value":   me.Value,
		"tag_key": me.TagKey,
		"context": me.Context,
	})
}

func (me *UniversalTag) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"key":     &me.Key,
		"value":   &me.Value,
		"tag_key": &me.TagKey,
		"context": &me.Context,
	})
}

func (me *UniversalTag) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"key":     me.Key,
		"value":   me.Value,
		"tagKey":  me.TagKey,
		"context": me.Context,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *UniversalTag) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"key":     me.Key,
		"value":   me.Value,
		"tagKey":  me.TagKey,
		"context": me.Context,
	})
}

// UniversalTagContext The origin of the tag, such as AWS or Cloud Foundry. For custom tags use the `CONTEXTLESS` value.
// The context is set for tags that are automatically imported by OneAgent (for example, from the AWS console or environment variables). It’s useful for determining the origin of tags when not manually defined, and it also helps to prevent clashes with other existing tags. If the tag is not automatically imported, `CONTEXTLESS` set.
type UniversalTagContext string

// UniversalTagContexts offers the known enum values
var UniversalTagContexts = struct {
	AWS                 UniversalTagContext
	AWSGeneric          UniversalTagContext
	Azure               UniversalTagContext
	CloudFoundry        UniversalTagContext
	Contextless         UniversalTagContext
	Environment         UniversalTagContext
	GoogleComputeEngine UniversalTagContext
	Kubernetes          UniversalTagContext
}{
	"AWS",
	"AWS_GENERIC",
	"AZURE",
	"CLOUD_FOUNDRY",
	"CONTEXTLESS",
	"ENVIRONMENT",
	"GOOGLE_COMPUTE_ENGINE",
	"KUBERNETES",
}
