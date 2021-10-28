package browser

import "github.com/dtcookie/hcl"

type Validations []*Validation

func (me *Validations) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"validation": {
			Type:        hcl.TypeList,
			Description: "The element to wait for. Required for the `validation` type, not applicable otherwise.",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Validation).Schema()},
		},
	}
}

func (me Validations) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	for _, entry := range me {
		if marshalled, err := entry.MarshalHCL(); err == nil {
			entries = append(entries, marshalled)
		} else {
			return nil, err
		}
	}
	result["validation"] = entries
	return result, nil
}

func (me *Validations) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("validation", me); err != nil {
		return err
	}
	return nil
}

type Validation struct {
	Type        ValidationType `json:"type"`              // The goal of the validation. `content_match` (check page for the specific content. Not allowed for validation inside of wait condition), `element_match` (check page for the specific element)
	Match       string         `json:"match"`             // The content to look for on the page.\nRegular expressions are allowed. In that case set `isRegex` as `true`. Required for `content_match`, optional for `element_match`.
	IsRegex     bool           `json:"isRegex,omitempty"` // Defines whether `match` is plain text (`false`) of a regular expression (`true`)
	FailIfFound bool           `json:"failIfFound"`       // The condition of the validation. `false` means the validation succeeds if the specified content/element is found. `true` means the validation fails if the specified content/element is found
	Target      *Target        `json:"target,omitempty"`  // The elemnt to look for on the page
}

func (me *Validation) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The goal of the validation. `content_match` (check page for the specific content. Not allowed for validation inside of wait condition), `element_match` (check page for the specific element).",
			Required:    true,
		},
		"match": {
			Type:        hcl.TypeString,
			Description: "The content to look for on the page.\nRegular expressions are allowed. In that case set `isRegex` as `true`. Required for `content_match`, optional for `element_match`.",
			Optional:    true,
		},
		"regex": {
			Type:        hcl.TypeBool,
			Description: "Defines whether `match` is plain text (`false`) or a regular expression (`true`)",
			Optional:    true,
		},
		"fail_if_found": {
			Type:        hcl.TypeBool,
			Description: "The condition of the validation. `false` means the validation succeeds if the specified content/element is found. `true` means the validation fails if the specified content/element is found",
			Optional:    true,
		},
		"target": {
			Type:        hcl.TypeList,
			Description: "The elemnt to look for on the page",
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Target).Schema()},
		},
	}
}

func (me *Validation) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.Type)
	if len(me.Match) > 0 {
		result["match"] = me.Match
	}
	result["regex"] = me.IsRegex
	result["fail_if_found"] = me.FailIfFound
	if me.Target != nil {
		if marshalled, err := me.Target.MarshalHCL(); err == nil {
			result["target"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Validation) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("match", &me.Match); err != nil {
		return err
	}
	if err := decoder.Decode("regex", &me.IsRegex); err != nil {
		return err
	}
	if err := decoder.Decode("fail_if_found", &me.FailIfFound); err != nil {
		return err
	}
	if err := decoder.Decode("target", &me.Target); err != nil {
		return err
	}
	return nil
}

// ValidationType The goal of the validation. `content_match` (check page for the specific content. Not allowed for validation inside of wait condition), `element_match` (check page for the specific element)
type ValidationType string

// ValidationTypes offers the known enum values
var ValidationTypes = struct {
	ContentMatch ValidationType
	ElementMatch ValidationType
}{
	`content_match`,
	`element_match`,
}
