package alerting

import (
	"encoding/json"
	"sort"
	"strings"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// Profile Configuration of an alerting profile.
type Profile struct {
	ID               *string                    `json:"id,omitempty"`               // The ID of the alerting profile.
	DisplayName      string                     `json:"displayName"`                // The name of the alerting profile, displayed in the UI.
	MzID             *string                    `json:"mzId,omitempty"`             // The ID of the management zone to which the alerting profile applies.
	Rules            []*ProfileSeverityRule     `json:"rules,omitempty"`            // A list of severity rules.   The rules are evaluated from top to bottom. The first matching rule applies and further evaluation stops.  If you specify both severity rule and event filter, the AND logic applies.
	EventTypeFilters []*EventTypeFilter         `json:"eventTypeFilters,omitempty"` // The list of event filters.  For all filters that are *negated* inside of these event filters, that is all "Predefined" as well as "Custom" (Title and/or Description) ones the AND logic applies. For all *non-negated* ones the OR logic applies. Between these two groups, negated and non-negated, the AND logic applies.  If you specify both severity rule and event filter, the AND logic applies.
	Metadata         *api.ConfigMetadata        `json:"metadata,omitempty"`         // Metadata useful for debugging
	Unknowns         map[string]json.RawMessage `json:"-"`
}

func (me *Profile) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"display_name": {
			Type:        hcl.TypeString,
			Description: "The name of the alerting profile, displayed in the UI",
			Required:    true,
		},
		"mz_id": {
			Type:        hcl.TypeString,
			Description: "The ID of the management zone to which the alerting profile applies",
			Optional:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "A list of rules for management zone usage.  Each rule is evaluated independently of all other rules",
			Optional:    true,
			MinItems:    1,
			Elem: &hcl.Resource{
				Schema: new(ProfileSeverityRule).Schema(),
			},
		},
		"event_type_filters": {
			Type:        hcl.TypeList,
			Description: "The list of event filters.  For all filters that are *negated* inside of these event filters, that is all `Predefined` as well as `Custom` (Title and/or Description) ones the AND logic applies. For all *non-negated* ones the OR logic applies. Between these two groups, negated and non-negated, the AND logic applies.  If you specify both severity rule and event filter, the AND logic applies",
			Optional:    true,
			MinItems:    1,
			Elem: &hcl.Resource{
				Schema: new(EventTypeFilter).Schema(),
			},
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
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Profile) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["display_name"] = me.DisplayName
	if me.MzID != nil {
		result["mz_id"] = me.MzID
	}
	if me.Rules != nil {
		rules := append([]*ProfileSeverityRule{}, me.Rules...)
		sort.Slice(rules, func(i, j int) bool {
			d1, _ := json.Marshal(rules[i])
			d2, _ := json.Marshal(rules[j])
			cmp := strings.Compare(string(d1), string(d2))
			return (cmp == -1)
		})
		entries := []interface{}{}
		for _, entry := range rules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rules"] = entries
	}
	if me.EventTypeFilters != nil {
		entries := []interface{}{}
		for _, entry := range me.EventTypeFilters {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["event_type_filters"] = entries
	}
	return result, nil
}

func (me *Profile) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "display_name")
		delete(me.Unknowns, "mz_id")
		delete(me.Unknowns, "rules")
		delete(me.Unknowns, "event_type_filters")
		delete(me.Unknowns, "metadata")
		delete(me.Unknowns, "managementZoneId")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("display_name"); ok {
		me.DisplayName = value.(string)
	}
	if value, ok := decoder.GetOk("mz_id"); ok {
		me.MzID = opt.NewString(value.(string))
	}
	if result, ok := decoder.GetOk("rules.#"); ok {
		me.Rules = []*ProfileSeverityRule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(ProfileSeverityRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "rules", idx)); err != nil {
				return err
			}
			me.Rules = append(me.Rules, entry)
		}
	}
	if result, ok := decoder.GetOk("event_type_filters.#"); ok {
		me.EventTypeFilters = []*EventTypeFilter{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(EventTypeFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "event_type_filters", idx)); err != nil {
				return err
			}
			me.EventTypeFilters = append(me.EventTypeFilters, entry)
		}
	}

	return nil
}

func (me *Profile) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	if me.ID != nil {
		rawMessage, err := json.Marshal(me.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.DisplayName)
		if err != nil {
			return nil, err
		}
		m["displayName"] = rawMessage
	}
	if me.MzID != nil {
		rawMessage, err := json.Marshal(me.MzID)
		if err != nil {
			return nil, err
		}
		m["mzId"] = rawMessage
	}
	if len(me.Rules) > 0 {
		rawMessage, err := json.Marshal(me.Rules)
		if err != nil {
			return nil, err
		}
		m["rules"] = rawMessage
	}
	if len(me.EventTypeFilters) > 0 {
		rawMessage, err := json.Marshal(me.EventTypeFilters)
		if err != nil {
			return nil, err
		}
		m["eventTypeFilters"] = rawMessage
	}
	if me.Metadata != nil {
		rawMessage, err := json.Marshal(me.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *Profile) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &me.ID); err != nil {
			return err
		}
	}
	if v, found := m["displayName"]; found {
		if err := json.Unmarshal(v, &me.DisplayName); err != nil {
			return err
		}
	}
	if v, found := m["mzId"]; found {
		if err := json.Unmarshal(v, &me.MzID); err != nil {
			return err
		}
	}
	if v, found := m["rules"]; found {
		if err := json.Unmarshal(v, &me.Rules); err != nil {
			return err
		}
	}
	if v, found := m["eventTypeFilters"]; found {
		if err := json.Unmarshal(v, &me.EventTypeFilters); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &me.Metadata); err != nil {
			return err
		}
	}
	delete(m, "id")
	delete(m, "displayName")
	delete(m, "mzId")
	delete(m, "managementZoneId")
	delete(m, "rules")
	delete(m, "eventTypeFilters")
	delete(m, "metadata")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
