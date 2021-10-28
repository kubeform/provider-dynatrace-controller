package validation

import (
	"github.com/dtcookie/hcl"
)

// Settings helps you verify that your HTTP monitor loads the expected content
type Settings struct {
	Rules Rules `json:"rules,omitempty"` // A list of validation rules
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Description: "A list of validation rules",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Rule).Schema()},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.Rules) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Rules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rule"] = entries
	}
	return result, nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("rule.#"); ok {
		me.Rules = Rules{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Rule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "rule", idx)); err != nil {
				return err
			}
			me.Rules = append(me.Rules, entry)
		}
	}
	return nil
}
