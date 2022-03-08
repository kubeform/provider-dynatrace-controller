package web

import "github.com/dtcookie/hcl"

// BrowserRestrictionSettings Settings for restricting certain browser type, version, platform and, comparator. It also restricts the mode
type BrowserRestrictionSettings struct {
	Mode                RestrictionMode     `json:"mode"`                          // The mode of the list of browser restrictions. Possible values area `EXCLUDE` and `INCLUDE`.
	BrowserRestrictions BrowserRestrictions `json:"browserRestrictions,omitempty"` // A list of browser restrictions
}

func (me *BrowserRestrictionSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"mode": {
			Type:        hcl.TypeString,
			Description: "The mode of the list of browser restrictions. Possible values area `EXCLUDE` and `INCLUDE`.",
			Required:    true,
		},
		"restrictions": {
			Type:        hcl.TypeList,
			Description: "A list of browser restrictions",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(BrowserRestrictions).Schema()},
		},
	}
}

func (me *BrowserRestrictionSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"mode":                me.Mode,
		"browserRestrictions": me.BrowserRestrictions,
	})
}

func (me *BrowserRestrictionSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"mode":                &me.Mode,
		"browserRestrictions": &me.BrowserRestrictions,
	})
}
