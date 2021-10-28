package monitors

import (
	"sort"

	"github.com/dtcookie/hcl"
)

type TagsWithSourceInfo []*TagWithSourceInfo

func (me *TagsWithSourceInfo) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"tag": {
			Type:        hcl.TypeList,
			Description: "Tag with source of a Dynatrace entity.",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(TagWithSourceInfo).Schema(),
			},
		},
	}
}

func (me TagsWithSourceInfo) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	sorted := TagsWithSourceInfo{}
	sorted = append(sorted, me...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})
	for _, tag := range sorted {
		if marshalled, err := tag.MarshalHCL(); err == nil {
			entries = append(entries, marshalled)
		} else {
			return nil, err
		}
	}
	result["tag"] = entries
	return result, nil
}

func (me *TagsWithSourceInfo) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("tag.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(TagWithSourceInfo)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "tag", idx)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Tag with source of a Dynatrace entity
type TagWithSourceInfo struct {
	Source  *TagSource `json:"source,omitempty"` // The source of the tag, such as USER, RULE_BASED or AUTO
	Context TagContext `json:"context"`          // The origin of the tag, such as AWS or Cloud Foundry. \n\n Custom tags use the `CONTEXTLESS` value
	Key     string     `json:"key"`              // The key of the tag. \n\n Custom tags have the tag value here
	Value   *string    `json:"value"`            // The value of the tag. \n\n Not applicable to custom tags
}

func (me *TagWithSourceInfo) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"source": {
			Type:        hcl.TypeString,
			Description: "The source of the tag. Supported values are `USER`, `RULE_BASED` and `AUTO`.",
			Optional:    true,
		},
		"context": {
			Type:        hcl.TypeString,
			Description: "The origin of the tag. Supported values are `AWS`, `AWS_GENERIC`, `AZURE`, `CLOUD_FOUNDRY`, `CONTEXTLESS`, `ENVIRONMENT`, `GOOGLE_CLOUD` and `KUBERNETES`.\n\nCustom tags use the `CONTEXTLESS` value.",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "The key of the tag.\n\nCustom tags have the tag value here.",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: " The value of the tag.\n\nNot applicable to custom tags.",
			Optional:    true,
		},
	}
}

func (me TagWithSourceInfo) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.Source != nil {
		result["source"] = string(*me.Source)
	}
	result["context"] = string(me.Context)
	result["key"] = me.Key
	if me.Value != nil {
		result["value"] = *me.Value
	}
	return result, nil
}

func (me *TagWithSourceInfo) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("source", &me.Source); err != nil {
		return err
	}
	if err := decoder.Decode("context", &me.Context); err != nil {
		return err
	}
	if err := decoder.Decode("key", &me.Key); err != nil {
		return err
	}
	if err := decoder.Decode("value", &me.Value); err != nil {
		return err
	}
	return nil
}

// TagSource The source of the tag, such as USER, RULE_BASED or AUTO
type TagSource string

// TagSources offers the known enum values
var TagSources = struct {
	Auto      TagSource
	RuleBased TagSource
	User      TagSource
}{
	"AUTO",
	"RULE_BASED",
	"USER",
}

// TagContext The origin of the tag, such as AWS or Cloud Foundry. \n\n Custom tags use the `CONTEXTLESS` value
type TagContext string

// TagContexts offers the known enum values
var TagContexts = struct {
	AWS          TagContext
	AWSGeneric   TagContext
	Azure        TagContext
	CloudFoundry TagContext
	ContextLess  TagContext
	Environment  TagContext
	GoogleCloud  TagContext
	Kubernetes   TagContext
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
