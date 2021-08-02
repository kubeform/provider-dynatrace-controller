package managementzones

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// ManagementZone The configuration of the management zone. It defines how the management zone applies.
type ManagementZone struct {
	ID                       *string                    `json:"id,omitempty"`                       // The ID of the management zone
	Description              *string                    `json:"description,omitempty"`              // The description of the management zone
	Metadata                 *api.ConfigMetadata        `json:"metadata,omitempty"`                 // Metadata useful for debugging
	Name                     string                     `json:"name"`                               // The name of the management zone.
	Rules                    []*Rule                    `json:"rules,omitempty"`                    // A list of rules for management zone usage. Each rule is evaluated independently of all other rules.
	DimensionalRules         []*DimensionalRule         `json:"dimensionalRules,omitempty"`         // A list of dimensional data rules for management zone usage. If several rules are specified, the **OR** logic applies
	EntitySelectorBasedRules []*EntitySelectorBasedRule `json:"entitySelectorBasedRules,omitempty"` // A list of entity-selector based rules for management zone usage. If several rules are specified, the **OR** logic applies
	Unknowns                 map[string]json.RawMessage `json:"-"`
}

func (mz *ManagementZone) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the management zone",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "The description of the management zone",
			Optional:    true,
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
		"dimensional_rule": {
			Type:        hcl.TypeList,
			Description: "A list of dimensional data rules for management zone usage. If several rules are specified, the `or` logic applies",
			Optional:    true,
			MinItems:    1,
			Elem: &hcl.Resource{
				Schema: new(DimensionalRule).Schema(),
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

func (mz *ManagementZone) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(mz.Unknowns) > 0 {
		data, err := json.Marshal(mz.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = mz.Name
	if mz.Description != nil && len(*mz.Description) > 0 {
		result["description"] = *mz.Description
	}
	if mz.Rules != nil {
		entries := []interface{}{}
		for _, entry := range mz.Rules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rules"] = entries
	}
	if len(mz.DimensionalRules) > 0 {
		entries := []interface{}{}
		for _, entry := range mz.DimensionalRules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["dimensional_rule"] = entries
	}
	if len(mz.EntitySelectorBasedRules) > 0 {
		entries := []interface{}{}
		for _, entry := range mz.EntitySelectorBasedRules {
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

func (mz *ManagementZone) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), mz); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &mz.Unknowns); err != nil {
			return err
		}
		delete(mz.Unknowns, "rules")
		delete(mz.Unknowns, "dimensional_rule")
		delete(mz.Unknowns, "entity_selector_based_rule")
		delete(mz.Unknowns, "metadata")
		delete(mz.Unknowns, "name")
		delete(mz.Unknowns, "description")
		if len(mz.Unknowns) == 0 {
			mz.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		mz.Name = value.(string)
	}
	if value, ok := decoder.GetOk("description"); ok {
		mz.Description = opt.NewString(value.(string))
	}
	if result, ok := decoder.GetOk("rules.#"); ok {
		mz.Rules = []*Rule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Rule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "rules", idx)); err != nil {
				return err
			}
			mz.Rules = append(mz.Rules, entry)
		}
	}
	if result, ok := decoder.GetOk("dimensional_rule.#"); ok {
		mz.DimensionalRules = []*DimensionalRule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(DimensionalRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "dimensional_rule", idx)); err != nil {
				return err
			}
			mz.DimensionalRules = append(mz.DimensionalRules, entry)
		}
	}
	if result, ok := decoder.GetOk("entity_selector_based_rule.#"); ok {
		mz.EntitySelectorBasedRules = []*EntitySelectorBasedRule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(EntitySelectorBasedRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "entity_selector_based_rule", idx)); err != nil {
				return err
			}
			mz.EntitySelectorBasedRules = append(mz.EntitySelectorBasedRules, entry)
		}
	}
	return nil
}

func (mz *ManagementZone) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(mz.Unknowns) > 0 {
		for k, v := range mz.Unknowns {
			m[k] = v
		}
	}
	if mz.ID != nil {
		rawMessage, err := json.Marshal(mz.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(mz.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	if mz.Description != nil {
		rawMessage, err := json.Marshal(mz.Description)
		if err != nil {
			return nil, err
		}
		m["description"] = rawMessage
	}
	if mz.Metadata != nil {
		rawMessage, err := json.Marshal(mz.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	if len(mz.Rules) > 0 {
		rawMessage, err := json.Marshal(mz.Rules)
		if err != nil {
			return nil, err
		}
		m["rules"] = rawMessage
	}
	if len(mz.DimensionalRules) > 0 {
		rawMessage, err := json.Marshal(mz.DimensionalRules)
		if err != nil {
			return nil, err
		}
		m["dimensionalRules"] = rawMessage
	}
	if len(mz.EntitySelectorBasedRules) > 0 {
		rawMessage, err := json.Marshal(mz.EntitySelectorBasedRules)
		if err != nil {
			return nil, err
		}
		m["entitySelectorBasedRules"] = rawMessage
	}
	return json.Marshal(m)
}

func (mz *ManagementZone) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &mz.ID); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &mz.Metadata); err != nil {
			return err
		}
	}
	if v, found := m["description"]; found {
		if err := json.Unmarshal(v, &mz.Description); err != nil {
			return err
		}
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &mz.Name); err != nil {
			return err
		}
	}
	if v, found := m["rules"]; found {
		if err := json.Unmarshal(v, &mz.Rules); err != nil {
			return err
		}
	}
	if v, found := m["dimensionalRules"]; found {
		if err := json.Unmarshal(v, &mz.DimensionalRules); err != nil {
			return err
		}
	}
	if v, found := m["entitySelectorBasedRules"]; found {
		if err := json.Unmarshal(v, &mz.EntitySelectorBasedRules); err != nil {
			return err
		}
	}
	delete(m, "id")
	delete(m, "name")
	delete(m, "description")
	delete(m, "metadata")
	delete(m, "rules")
	delete(m, "dimensionalRules")
	delete(m, "entitySelectorBasedRules")

	if len(m) > 0 {
		mz.Unknowns = m
	}
	return nil
}
