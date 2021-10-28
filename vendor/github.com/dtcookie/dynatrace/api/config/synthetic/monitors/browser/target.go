package browser

import "github.com/dtcookie/hcl"

type Target struct {
	Window   *string  `json:"window,omitempty"`   // The tab of the target
	Locators Locators `json:"locators,omitempty"` // The list of locators identifying the desired element
}

func (me *Target) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"window": {
			Type:        hcl.TypeString,
			Description: "The tab of the target",
			Optional:    true,
		},
		"locators": {
			Type:        hcl.TypeList,
			Description: "The list of locators identifying the desired element",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Locators).Schema()},
		},
	}
}

func (me *Target) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.Window != nil {
		result["window"] = string(*me.Window)
	}
	if len(me.Locators) > 0 {
		if marshalled, err := me.Locators.MarshalHCL(); err == nil {
			result["locators"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Target) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("window", &me.Window); err != nil {
		return err
	}
	if err := decoder.Decode("locators", &me.Locators); err != nil {
		return err
	}
	return nil
}
