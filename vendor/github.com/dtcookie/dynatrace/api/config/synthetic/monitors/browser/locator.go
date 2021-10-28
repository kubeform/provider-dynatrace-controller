package browser

import "github.com/dtcookie/hcl"

// Locators a list of locators identifying the desired element
type Locators []*Locator

func (me *Locators) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"locator": {
			Type:        hcl.TypeList,
			Description: "A locator dentifyies the desired element",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Locator).Schema()},
		},
	}
}

func (me Locators) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	for _, entry := range me {
		if marshalled, err := entry.MarshalHCL(); err == nil {
			entries = append(entries, marshalled)
		} else {
			return nil, err
		}
		result["locator"] = entries
	}
	return result, nil
}

func (me *Locators) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("locator", me); err != nil {
		return err
	}
	return nil
}

type Locator struct {
	Type  LocatorType `json:"type"`  // Defines where to look for an element. `css` (CSS Selector) or `dom` (Javascript code)
	Value string      `json:"value"` // The name of the element to be found
}

func (me *Locator) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "Defines where to look for an element. `css` (CSS Selector) or `dom` (Javascript code)",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The name of the element to be found",
			Required:    true,
		},
	}
}

func (me *Locator) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.Type)
	result["value"] = me.Value
	return result, nil
}

func (me *Locator) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("value", &me.Value); err != nil {
		return err
	}
	return nil
}

// LocatorType defines where to look for an element. `css` (CSS Selector) or `dom` (Javascript code)
type LocatorType string

// LocatorTypes offers the known enum values
var LocatorTypes = struct {
	ContentMatch LocatorType
	ElementMatch LocatorType
}{
	"css",
	"dom",
}
