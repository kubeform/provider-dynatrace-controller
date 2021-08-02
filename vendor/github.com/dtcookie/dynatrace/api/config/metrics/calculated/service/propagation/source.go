package propagation

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// Source Defines valid sources of request attributes for conditions or placeholders.
type Source struct {
	ManagementZone *string                    `json:"managementZone,omitempty"` // Use only request attributes from services that belong to this management zone.. Use either this or `serviceTag`
	ServiceTag     *UniversalTag              `json:"serviceTag,omitempty"`     // Use only request attributes from services that have this tag. Use either this or `managementZone`
	Unknowns       map[string]json.RawMessage `json:"-"`
}

func (me *Source) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"management_zone": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "Use only request attributes from services that belong to this management zone.. Use either this or `serviceTag`",
		},
		"service_tag": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Use only request attributes from services that have this tag. Use either this or `managementZone`",
			Elem:        &hcl.Resource{Schema: new(UniversalTag).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Source) MarshalHCL() (map[string]interface{}, error) {
	properties, err := hcl.NewProperties(me, me.Unknowns)
	if err != nil {
		return nil, err
	}
	return properties.EncodeAll(map[string]interface{}{
		"management_zone": me.ManagementZone,
		"service_tag":     me.ServiceTag,
		"unknowns":        me.Unknowns,
	})
}

func (me *Source) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"management_zone": &me.ManagementZone,
		"service_tag":     &me.ServiceTag,
		"unknowns":        &me.Unknowns,
	})
}

func (me *Source) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"managementZone": me.ManagementZone,
		"serviceTag":     me.ServiceTag,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Source) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		"managementZone": &me.ManagementZone,
		"serviceTag":     &me.ServiceTag,
	})
}
