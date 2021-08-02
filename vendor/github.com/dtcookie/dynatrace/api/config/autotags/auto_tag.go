package autotags

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
)

// AutoTag Configuration of an auto-tag. It defines the conditions of tag usage and the tag value.
type AutoTag struct {
	ID                       *string                    `json:"id,omitempty"`                       // The ID of the auto-tag.
	Name                     string                     `json:"name"`                               // The name of the auto-tag, which is applied to entities.  Additionally you can specify a **valueFormat** in the tag rule. In that case the tag is used in the `name:valueFormat` format.  For example you can extend the `Infrastructure` tag to `Infrastructure:Windows` and `Infrastructure:Linux`.
	Rules                    []*Rule                    `json:"rules,omitempty"`                    // The list of rules for tag usage. When there are multiple rules, the OR logic applies.
	EntitySelectorBasedRules []*EntitySelectorBasedRule `json:"entitySelectorBasedRules,omitempty"` // A list of entity-selector based rules for auto tagging usage.\n\nIf several rules are specified, the **OR** logic applies
	Metadata                 *api.ConfigMetadata        `json:"metadata,omitempty"`                 // Metadata useful for debugging
	Unknowns                 map[string]json.RawMessage `json:"-"`
}

func (me *AutoTag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the auto-tag, which is applied to entities.  Additionally you can specify a **valueFormat** in the tag rule. In that case the tag is used in the `name:valueFormat` format.  For example you can extend the `Infrastructure` tag to `Infrastructure:Windows` and `Infrastructure:Linux`.",
			Required:    true,
		},
		"metadata": {
			Type:        hcl.TypeList,
			MaxItems:    1,
			Description: "`metadata` exists for backwards compatibility but shouldn't get specified anymore",
			Deprecated:  "`metadata` exists for backwards compatibility but shouldn't get specified anymore",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(api.ConfigMetadata).Schema(),
			},
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "A list of rules for management zone usage.  Each rule is evaluated independently of all other rules",
			Optional:    true,
			MinItems:    1,
			Elem: &hcl.Resource{
				Schema: new(Rule).Schema(),
			},
		},
		"entity_selector_based_rule": {
			Type:        hcl.TypeList,
			Description: "A list of entity-selector based rules for management zone usage. If several rules are specified, the `or` logic applies",
			Optional:    true,
			MinItems:    1,
			Elem: &hcl.Resource{
				Schema: new(EntitySelectorBasedRule).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *AutoTag) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	if me.Rules != nil {
		entries := []interface{}{}
		for _, entry := range me.Rules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rules"] = entries
	}
	if me.EntitySelectorBasedRules != nil {
		entries := []interface{}{}
		for _, entry := range me.EntitySelectorBasedRules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["entity_selector_based_rule"] = entries
	}
	return result, nil
}

func (me *AutoTag) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "rules")
		delete(me.Unknowns, "entity_selector_based_rule")
		delete(me.Unknowns, "metadata")
		delete(me.Unknowns, "name")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if result, ok := decoder.GetOk("rules.#"); ok {
		me.Rules = []*Rule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Rule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "rules", idx)); err != nil {
				return err
			}
			me.Rules = append(me.Rules, entry)
		}
	}
	if result, ok := decoder.GetOk("entity_selector_based_rule.#"); ok {
		me.EntitySelectorBasedRules = []*EntitySelectorBasedRule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(EntitySelectorBasedRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "entity_selector_based_rule", idx)); err != nil {
				return err
			}
			me.EntitySelectorBasedRules = append(me.EntitySelectorBasedRules, entry)
		}
	}
	return nil
}

func (me *AutoTag) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	if me.ID != nil {
		rawMessage, err := json.Marshal(me.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	if me.Metadata != nil {
		rawMessage, err := json.Marshal(me.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	if len(me.Rules) > 0 {
		rawMessage, err := json.Marshal(me.Rules)
		if err != nil {
			return nil, err
		}
		m["rules"] = rawMessage
	}
	if len(me.EntitySelectorBasedRules) > 0 {
		rawMessage, err := json.Marshal(me.EntitySelectorBasedRules)
		if err != nil {
			return nil, err
		}
		m["entitySelectorBasedRules"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *AutoTag) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &me.ID); err != nil {
			return err
		}
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &me.Name); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &me.Metadata); err != nil {
			return err
		}
	}
	if v, found := m["rules"]; found {
		if err := json.Unmarshal(v, &me.Rules); err != nil {
			return err
		}
	}
	if v, found := m["entitySelectorBasedRules"]; found {
		if err := json.Unmarshal(v, &me.EntitySelectorBasedRules); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "metadata")
	delete(m, "rules")
	delete(m, "id")
	delete(m, "entitySelectorBasedRules")
	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
