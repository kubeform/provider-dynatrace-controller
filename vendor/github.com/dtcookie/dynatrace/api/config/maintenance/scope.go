package maintenance

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Scope The scope of the maintenance window.
//  The scope restricts the alert/problem detection suppression to certain Dynatrace entities. It can contain a list of entities and/or matching rules for dynamic formation of the scope.
//  If no scope is specified, the alert/problem detection suppression applies to the entire environment.
type Scope struct {
	Entities []string                   `json:"entities"` // A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs.
	Matches  []*Filter                  `json:"matches"`  // A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies.
	Unknowns map[string]json.RawMessage `json:"-"`
}

func (me *Scope) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entities": {
			Type:        hcl.TypeSet,
			Description: "A list of Dynatrace entities (for example, hosts or services) to be included in the scope.  Allowed values are Dynatrace entity IDs",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"matches": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "A list of matching rules for dynamic scope formation.  If several rules are set, the OR logic applies",
			Elem: &hcl.Resource{
				Schema: new(Filter).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Scope) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if len(me.Entities) > 0 {
		result["entities"] = me.Entities
	}
	if len(me.Matches) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Matches {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["matches"] = entries
	}
	return result, nil
}

func (me *Scope) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "entities")
		delete(me.Unknowns, "matches")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	me.Entities = decoder.GetStringSet("entities")
	if result, ok := decoder.GetOk("matches.#"); ok {
		me.Matches = []*Filter{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Filter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "matches", idx)); err != nil {
				return err
			}
			me.Matches = append(me.Matches, entry)
		}
	}
	return nil
}

func (me *Scope) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if len(me.Entities) == 0 {
		data, err := json.Marshal([]string{})
		if err != nil {
			return nil, err
		}
		m["entities"] = data
	} else {
		if err := m.Marshal("entities", me.Entities); err != nil {
			return nil, err
		}
	}
	if len(me.Matches) == 0 {
		data, err := json.Marshal([]string{})
		if err != nil {
			return nil, err
		}
		m["matches"] = data
	} else {
		if err := m.Marshal("matches", me.Matches); err != nil {
			return nil, err
		}
	}
	return json.Marshal(m)
}

func (me *Scope) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("entities", &me.Entities); err != nil {
		return err
	}
	if err := m.Unmarshal("matches", &me.Matches); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
