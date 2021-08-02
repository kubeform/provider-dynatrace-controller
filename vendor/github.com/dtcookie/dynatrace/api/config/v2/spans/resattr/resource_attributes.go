package resattr

import (
	"github.com/dtcookie/hcl"
)

// ResourceAttributes has no documentation
type ResourceAttributes struct {
	Keys []*RuleItem `json:"attributeKeys"`
}

func (me *ResourceAttributes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"enabled": {
			Type:        hcl.TypeSet,
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Description: "attributes that should get captured",
		},
		"disabled": {
			Type:        hcl.TypeSet,
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Description: "configured attributes that currently shouldn't be taken into consideration",
		},
	}
}

func (me *ResourceAttributes) MarshalHCL() (map[string]interface{}, error) {
	enableds := []string{}
	disableds := []string{}
	for _, item := range me.Keys {
		if item.Enabled {
			enableds = append(enableds, item.AttributeKey)
		} else {
			disableds = append(disableds, item.AttributeKey)
		}
	}
	m := map[string]interface{}{}
	if len(enableds) > 0 {
		m["enabled"] = enableds
	}
	if len(disableds) > 0 {
		m["disabled"] = disableds
	}
	return m, nil
}

func (me *ResourceAttributes) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Keys = []*RuleItem{}
	if enableds, ok := decoder.GetOk("enabled"); ok {
		enabledSet := enableds.(hcl.Set)
		for _, item := range enabledSet.List() {
			me.Keys = append(me.Keys, &RuleItem{Enabled: true, AttributeKey: item.(string)})
		}
	}
	if enableds, ok := decoder.GetOk("disabled"); ok {
		enabledSet := enableds.(hcl.Set)
		for _, item := range enabledSet.List() {
			me.Keys = append(me.Keys, &RuleItem{Enabled: false, AttributeKey: item.(string)})
		}
	}
	return nil
}
