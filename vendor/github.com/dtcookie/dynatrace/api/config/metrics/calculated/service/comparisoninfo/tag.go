package comparisoninfo

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Tag Comparison for `TAG` attributes.
type Tag struct {
	BaseComparisonInfo
	Comparison TagComparison `json:"comparison"`       // Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Value      *TagInfo      `json:"value,omitempty"`  // Tag of a Dynatrace entity.
	Values     TagInfos      `json:"values,omitempty"` // The values to compare to.
}

func (me *Tag) GetType() Type {
	return Types.Tag
}

func (me *Tag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"values": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "The values to compare to",
			Elem:        &hcl.Resource{Schema: new(TagInfos).Schema()},
		},
		"value": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Description: "The values to compare to",
			Elem:        &hcl.Resource{Schema: new(TagInfo).Schema()},
		},
		"operator": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Operator of the comparison. You can reverse it by setting `negate` to `true`. Possible values are `EQUALS`, `EQUALS_ANY_OF`, `TAG_KEY_EQUALS` and `TAG_KEY_EQUALS_ANY_OF`",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Tag) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"values":   me.Values,
		"value":    me.Value,
		"operator": me.Comparison,
		"unknowns": me.Unknowns,
	})
}

func (me *Tag) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"values":   &me.Values,
		"value":    &me.Value,
		"operator": &me.Comparison,
		"unknowns": &me.Unknowns,
	})
}

func (me *Tag) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":       me.GetType(),
		"negate":     me.Negate,
		"values":     me.Values,
		"value":      me.Value,
		"comparison": me.Comparison,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Tag) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"negate":     &me.Negate,
		"values":     &me.Values,
		"value":      &me.Value,
		"comparison": &me.Comparison,
	})
}

// TagComparison Operator of the comparision. You can reverse it by setting **negate** to `true`.
type TagComparison string

// TagComparisons offers the known enum values
var TagComparisons = struct {
	Equals            TagComparison
	EqualsAnyOf       TagComparison
	TagKeyEquals      TagComparison
	TagKeyEqualsAnyOf TagComparison
}{
	"EQUALS",
	"EQUALS_ANY_OF",
	"TAG_KEY_EQUALS",
	"TAG_KEY_EQUALS_ANY_OF",
}
