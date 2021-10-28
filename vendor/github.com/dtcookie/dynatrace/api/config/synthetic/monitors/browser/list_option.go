package browser

import (
	"github.com/dtcookie/hcl"
)

type ListOptions []*ListOption

func (me *ListOptions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"option": {
			Type:        hcl.TypeList,
			Description: "The option to be selected",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(ListOption).Schema()},
		},
	}
}

func (me ListOptions) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	for _, entry := range me {
		if marshalled, err := entry.MarshalHCL(); err == nil {
			entries = append(entries, marshalled)
		} else {
			return nil, err
		}
	}
	result["option"] = entries

	return result, nil
}

func (me *ListOptions) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("option", &me); err != nil {
		return err
	}
	return nil
}

type ListOption struct {
	Index int    `json:"index"` // The index of the option to be selected
	Value string `json:"value"` // The value of the option to be selected
}

func (me *ListOption) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"index": {
			Type:        hcl.TypeInt,
			Description: "The index of the option to be selected",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the option to be selected",
			Required:    true,
		},
	}
}

func (me *ListOption) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["index"] = me.Index
	result["value"] = me.Value

	return result, nil
}

func (me *ListOption) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("index", &me.Index); err != nil {
		return err
	}
	if err := decoder.Decode("value", &me.Value); err != nil {
		return err
	}
	return nil
}
