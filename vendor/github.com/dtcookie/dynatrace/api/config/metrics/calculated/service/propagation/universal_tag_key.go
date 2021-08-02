package propagation

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// UniversalTagKey has no documentation
type UniversalTagKey struct {
	Key     *string                 `json:"key,omitempty"`     // has no documentation
	Context *UniversalTagKeyContext `json:"context,omitempty"` // has no documentation
}

func (me *UniversalTagKey) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "has no documentation",
		},
		"context": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "has no documentation",
		},
	}
}

func (me *UniversalTagKey) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"key":     me.Key,
		"context": me.Context,
	})
}

func (me *UniversalTagKey) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"key":     &me.Key,
		"context": &me.Context,
	})
}

func (me *UniversalTagKey) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"key":     me.Key,
		"context": me.Context,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *UniversalTagKey) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"key":     me.Key,
		"context": me.Context,
	})
}

// UniversalTagKeyContext has no documentation
type UniversalTagKeyContext string

// UniversalTagKeyContexts offers the known enum values
var UniversalTagKeyContexts = struct {
	AWS                 UniversalTagKeyContext
	AWSGeneric          UniversalTagKeyContext
	Azure               UniversalTagKeyContext
	CloudFoundry        UniversalTagKeyContext
	Contextless         UniversalTagKeyContext
	Environment         UniversalTagKeyContext
	GoogleComputeEngine UniversalTagKeyContext
	Kubernetes          UniversalTagKeyContext
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
