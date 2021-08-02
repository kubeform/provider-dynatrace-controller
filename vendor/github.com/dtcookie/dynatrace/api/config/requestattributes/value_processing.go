package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// ValueProcessing Process values as specified.
type ValueProcessing struct {
	ExtractSubstring    *ExtractSubstring          `json:"extractSubstring,omitempty"`    // Preprocess by extracting a substring from the original value.
	SplitAt             *string                    `json:"splitAt,omitempty"`             // Split (preprocessed) string values at this separator.
	Trim                *bool                      `json:"trim"`                          // Prune Whitespaces. Defaults to false.
	ValueCondition      *ValueCondition            `json:"valueCondition,omitempty"`      // IBM integration bus label node name condition for which the value is captured.
	ValueExtractorRegex *string                    `json:"valueExtractorRegex,omitempty"` // Extract value from captured data per regex.
	Unknowns            map[string]json.RawMessage `json:"-"`
}

func (me *ValueProcessing) IsZero() bool {
	if me.ExtractSubstring != nil && !me.ExtractSubstring.IsZero() {
		return false
	}
	if me.SplitAt != nil && len(*me.SplitAt) > 0 {
		return false
	}
	if opt.Bool(me.Trim) {
		return false
	}
	if me.ValueCondition != nil && !me.ValueCondition.IsZero() {
		return false
	}
	if me.ValueExtractorRegex != nil && len(*me.ValueExtractorRegex) > 0 {
		return false
	}
	if len(me.Unknowns) > 0 {
		return false
	}
	return true
}

func (me *ValueProcessing) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"extract_substring": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Preprocess by extracting a substring from the original value",
			Elem: &hcl.Resource{
				Schema: new(ExtractSubstring).Schema(),
			},
		},
		"split_at": {
			Type:        hcl.TypeString,
			Description: "Split (preprocessed) string values at this separator",
			Optional:    true,
		},
		"trim": {
			Type:        hcl.TypeBool,
			Description: "Prune Whitespaces. Defaults to false",
			Optional:    true,
		},
		"value_condition": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "IBM integration bus label node name condition for which the value is captured",
			Elem: &hcl.Resource{
				Schema: new(ValueCondition).Schema(),
			},
		},
		"value_extractor_regex": {
			Type:        hcl.TypeString,
			Description: "Extract value from captured data per regex",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ValueProcessing) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.ExtractSubstring != nil {
		if marshalled, err := me.ExtractSubstring.MarshalHCL(); err == nil {
			result["extract_substring"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.SplitAt != nil {
		result["split_at"] = *me.SplitAt
	}
	if me.Trim != nil {
		result["trim"] = opt.Bool(me.Trim)
	}
	if me.ValueCondition != nil {
		if marshalled, err := me.ValueCondition.MarshalHCL(); err == nil {
			result["value_condition"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.ValueExtractorRegex != nil {
		result["value_extractor_regex"] = *me.ValueExtractorRegex
	}
	return result, nil
}

func (me *ValueProcessing) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "extract_substring")
		delete(me.Unknowns, "split_at")
		delete(me.Unknowns, "trim")
		delete(me.Unknowns, "value_condition")
		delete(me.Unknowns, "value_extractor_regex")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, ok := decoder.GetOk("extract_substring.#"); ok {
		me.ExtractSubstring = new(ExtractSubstring)
		if err := me.ExtractSubstring.UnmarshalHCL(hcl.NewDecoder(decoder, "extract_substring", 0)); err != nil {
			return err
		}
	}
	adapter := hcl.Adapt(decoder)
	me.SplitAt = adapter.GetString("split_at")
	me.Trim = adapter.GetBool("trim")
	if _, ok := decoder.GetOk("value_condition.#"); ok {
		me.ValueCondition = new(ValueCondition)
		if err := me.ValueCondition.UnmarshalHCL(hcl.NewDecoder(decoder, "value_condition", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("value_extractor_regex"); ok {
		me.ValueExtractorRegex = opt.NewString(value.(string))
	}
	return nil
}

func (me *ValueProcessing) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("extractSubstring", me.ExtractSubstring); err != nil {
		return nil, err
	}
	if err := m.Marshal("splitAt", me.SplitAt); err != nil {
		return nil, err
	}
	if err := m.Marshal("trim", opt.Bool(me.Trim)); err != nil {
		return nil, err
	}
	if err := m.Marshal("valueCondition", me.ValueCondition); err != nil {
		return nil, err
	}
	if err := m.Marshal("valueExtractorRegex", me.ValueExtractorRegex); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *ValueProcessing) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("extractSubstring", &me.ExtractSubstring); err != nil {
		return err
	}
	if err := m.Unmarshal("splitAt", &me.SplitAt); err != nil {
		return err
	}
	if err := m.Unmarshal("trim", &me.Trim); err != nil {
		return err
	}
	if err := m.Unmarshal("valueCondition", &me.ValueCondition); err != nil {
		return err
	}
	if err := m.Unmarshal("valueExtractorRegex", &me.ValueExtractorRegex); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
