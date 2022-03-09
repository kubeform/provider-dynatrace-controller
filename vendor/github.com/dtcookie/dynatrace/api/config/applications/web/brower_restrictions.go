package web

import "github.com/dtcookie/hcl"

type BrowserRestrictions []*BrowserRestriction

func (me *BrowserRestrictions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"restriction": {
			Type:        hcl.TypeList,
			Description: "Browser exclusion rules for the browsers that are to be excluded",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(BrowserRestriction).Schema()},
		},
	}
}

func (me BrowserRestrictions) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["restriction"] = entries
	}
	return result, nil
}

func (me *BrowserRestrictions) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("restriction", me); err != nil {
		return err
	}
	return nil
}

// BrowserRestriction Browser exclusion rules for the browsers that are to be excluded
type BrowserRestriction struct {
	BrowserVersion *string     `json:"browserVersion,omitempty"` // The version of the browser that is used
	BrowserType    BrowserType `json:"browserType"`              // The type of the browser that is used. Possible values are `ANDROID_WEBKIT`, `BOTS_SPIDERS`, `CHROME`, `EDGE`, `FIREFOX`, `INTERNET_EXPLORER, `OPERA` and `SAFARI`
	Platform       *Platform   `json:"platform,omitempty"`       // The platform on which the browser is being used. Possible values are `ALL`, `DESKTOP` and `MOBILE`
	Comparator     *Comparator `json:"comparator,omitempty"`     // No documentation available
}

func (me *BrowserRestriction) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"browser_version": {
			Type:        hcl.TypeString,
			Description: "The version of the browser that is used",
			Optional:    true,
		},
		"browser_type": {
			Type:        hcl.TypeString,
			Description: "The type of the browser that is used. Possible values are `ANDROID_WEBKIT`, `BOTS_SPIDERS`, `CHROME`, `EDGE`, `FIREFOX`, `INTERNET_EXPLORER, `OPERA` and `SAFARI`",
			Required:    true,
		},
		"platform": {
			Type:        hcl.TypeString,
			Description: "The platform on which the browser is being used. Possible values are `ALL`, `DESKTOP` and `MOBILE`",
			Optional:    true,
		},
		"comparator": {
			Type:        hcl.TypeString,
			Description: "No documentation available. Possible values are `EQUALS`, `GREATER_THAN_OR_EQUAL` and `LOWER_THAN_OR_EQUAL`.",
			Optional:    true,
		},
	}
}

func (me *BrowserRestriction) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"browser_version": me.BrowserVersion,
		"browser_type":    me.BrowserType,
		"platform":        me.Platform,
		"comparator":      me.Comparator,
	})
}

func (me *BrowserRestriction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"browser_version": &me.BrowserVersion,
		"browser_type":    &me.BrowserType,
		"platform":        &me.Platform,
		"comparator":      &me.Comparator,
	})
}
