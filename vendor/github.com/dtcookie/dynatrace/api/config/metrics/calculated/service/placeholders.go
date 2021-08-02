package service

import "github.com/dtcookie/hcl"

type Placeholders []*Placeholder

func (me *Placeholders) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"placeholder": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A custom placeholder to be used in a dimension value pattern",
			Elem:        &hcl.Resource{Schema: new(Placeholder).Schema()},
		},
	}
}

func (me Placeholders) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeSlice("placeholder", me)
}

func (me *Placeholders) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("placeholder", me); err != nil {
		return err
	}
	return nil
}
