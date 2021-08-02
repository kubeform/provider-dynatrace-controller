package match

import "github.com/dtcookie/hcl"

type SpanMatcher struct {
	Source        Source     `json:"source"`
	SpanKindValue *SpanKind  `json:"spanKindValue"`
	Type          Comparison `json:"type"`
	SourceKey     *string    `json:"sourceKey"`
	Value         *string    `json:"value"`
	CaseSensitive bool       `json:"caseSensitive"`
}

func (me *SpanMatcher) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"comparison": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Possible values are `EQUALS`, `CONTAINS`, `STARTS_WITH`, `ENDS_WITH`, `DOES_NOT_EQUAL`, `DOES_NOT_CONTAIN`, `DOES_NOT_START_WITH` and `DOES_NOT_END_WITH`.",
		},
		"source": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "What to match against. Possible values are `SPAN_NAME`, `SPAN_KIND`, `ATTRIBUTE`, `INSTRUMENTATION_LIBRARY_NAME` and `INSTRUMENTATION_LIBRARY_VERSION`",
		},
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "Whether to match strings case sensitively or not",
		},
		"value": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The value to compare against. When `source` is `SPAN_KIND` the only allowed values are `INTERNAL`, `SERVER`, `CLIENT`, `PRODUCER` and `CONSUMER`",
		},
		"key": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The name of the attribute if `source` is `ATTRIBUTE`",
		},
	}
}

func (me *SpanMatcher) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	m := map[string]interface{}{
		"source":         me.Source,
		"comparison":     me.Type,
		"key":            me.SourceKey,
		"case_sensitive": me.CaseSensitive,
	}

	if me.Source == Sources.SpanKind && me.SpanKindValue != nil {
		m["value"] = string(*me.SpanKindValue)
	} else {
		m["value"] = *me.Value
	}
	if !me.CaseSensitive {
		delete(m, "case_sensitive")
	}
	return properties.EncodeAll(m)
}

func (me *SpanMatcher) UnmarshalHCL(decoder hcl.Decoder) error {
	m := map[string]interface{}{
		"source":         &me.Source,
		"comparison":     &me.Type,
		"key":            &me.SourceKey,
		"case_sensitive": &me.CaseSensitive,
		"value":          &me.Value,
	}
	if err := decoder.DecodeAll(m); err != nil {
		return err
	}

	if me.Source == Sources.SpanKind {
		if me.Value != nil {
			me.SpanKindValue = SpanKind(*me.Value).Ref()
		}
		me.Value = nil
	}
	return nil
}

type SpanMatchers []*SpanMatcher

func (me *SpanMatchers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"match": {
			Type:        hcl.TypeList,
			Required:    true,
			Description: "Matching strategies for the Span",
			Elem:        &hcl.Resource{Schema: new(SpanMatcher).Schema()},
		},
	}
}

func (me SpanMatchers) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("match", me)
}

func (me *SpanMatchers) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("match", me); err != nil {
		return err
	}
	return nil
}
