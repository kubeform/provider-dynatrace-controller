package managementzones

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// DimensionalRule represents the dimensional rule of the management zone usage.
// It defines how the management zone applies.
// Each rule is evaluated independently of all other rules
type DimensionalRule struct {
	Enabled    *bool                       `json:"enabled"`    // The rule is enabled (`true`) or disabled (`false`)
	AppliesTo  Application                 `json:"appliesTo"`  // The target of the rule
	Conditions []*DimensionalRuleCondition `json:"conditions"` // A list of conditions for the management zone. \n\n The management zone applies only if **all** conditions are fulfilled
	Unknowns   map[string]json.RawMessage  `json:"-"`
}

func (me *DimensionalRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The rule is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"applies_to": {
			Type:        hcl.TypeString,
			Description: "The target of the rule. Possible values are\n   - `ANY`\n   - `LOG`\n   - `METRIC`",
			Required:    true,
		},
		"condition": {
			Type:        hcl.TypeList,
			Description: "A list of conditions for the management zone. The management zone applies only if **all** conditions are fulfilled",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(DimensionalRuleCondition).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *DimensionalRule) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["enabled"] = opt.Bool(me.Enabled)
	result["applies_to"] = string(me.AppliesTo)
	if len(me.Conditions) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Conditions {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["condition"] = entries
	}

	return result, nil
}

func (me *DimensionalRule) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "applies_to")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "condition")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if _, value := decoder.GetChange("enabled"); value != nil {
		me.Enabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("applies_to"); ok {
		me.AppliesTo = Application(value.(string))
	}
	if result, ok := decoder.GetOk("condition.#"); ok {
		me.Conditions = []*DimensionalRuleCondition{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(DimensionalRuleCondition)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "condition", idx)); err != nil {
				return err
			}
			me.Conditions = append(me.Conditions, entry)
		}
	}
	return nil
}

func (me *DimensionalRule) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	if opt.Bool(me.Enabled) {
		rawMessage, err := json.Marshal(me.Enabled)
		if err != nil {
			return nil, err
		}
		m["enabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(me.AppliesTo)
		if err != nil {
			return nil, err
		}
		m["appliesTo"] = rawMessage
	}
	if len(me.Conditions) > 0 {
		rawMessage, err := json.Marshal(me.Conditions)
		if err != nil {
			return nil, err
		}
		m["conditions"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *DimensionalRule) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["enabled"]; found {
		if err := json.Unmarshal(v, &me.Enabled); err != nil {
			return err
		}
	}
	if v, found := m["appliesTo"]; found {
		if err := json.Unmarshal(v, &me.AppliesTo); err != nil {
			return err
		}
	}
	if v, found := m["conditions"]; found {
		if err := json.Unmarshal(v, &me.Conditions); err != nil {
			return err
		}
	}
	delete(m, "appliesTo")
	delete(m, "enabled")
	delete(m, "conditions")

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// Application The value to compare to.
type Application string

func (v Application) Ref() *Application {
	return &v
}

// Applications offers the known enum values
var Applications = struct {
	Any    Application
	Log    Application
	Metric Application
}{
	"ANY",
	"LOG",
	"METRIC",
}

func (v *Application) String() string {
	return string(*v)
}
