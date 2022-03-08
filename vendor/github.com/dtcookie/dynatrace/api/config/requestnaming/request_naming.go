package requestnaming

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/metrics/calculated/service"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

var o Order

type RequestNaming struct {
	// Order           *string                    `json:"order,omitempty"`           // The order string. Sorting request namings alphabetically by their order string determines their relative ordering.\n\nTypically this is managed by Dynatrace internally and will not be present in GET responses nor used if present in PUT/POST requests, except where noted otherwise
	Enabled         bool                       `json:"enabled"`                   // The rule is enabled (`true`) or disabled (`false`)
	NamingPattern   string                     `json:"namingPattern"`             // The name to be assigned to matching requests
	ManagementZones []string                   `json:"managementZones,omitempty"` // Specifies the management zones for which this rule should be applied
	Conditions      service.Conditions         `json:"conditions"`                // The set of conditions for the request naming rule usage. \n\n You can specify several conditions. The request has to match **all** the specified conditions for the rule to trigger
	Placeholders    service.Placeholders       `json:"placeholders"`              // The list of custom placeholders to be used in the naming pattern. \n\n It enables you to extract a request attribute value or other request attribute and use it in the request naming pattern.
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *RequestNaming) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		// "order": {
		// 	Type:        hcl.TypeString,
		// 	Optional:    true,
		// 	Description: "The order string. Sorting request namings alphabetically by their order string determines their relative ordering.\n\nTypically this is managed by Dynatrace internally and will not be present in GET responses nor used if present in PUT/POST requests, except where noted otherwise",
		// },
		"enabled": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The rule is enabled (`true`) or disabled (`false`)",
		},
		"naming_pattern": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The name to be assigned to matching requests",
		},
		"management_zones": {
			Type:        hcl.TypeSet,
			Optional:    true,
			Description: "Specifies the management zones for which this rule should be applied",
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"conditions": {
			Type:        hcl.TypeList,
			Required:    true,
			Description: "The set of conditions for the request naming rule usage. \n\n You can specify several conditions. The request has to match **all** the specified conditions for the rule to trigger",
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(service.Conditions).Schema()},
		},
		"placeholders": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "The list of custom placeholders to be used in the naming pattern. \n\n It enables you to extract a request attribute value or other request attribute and use it in the request naming pattern.",
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(service.Placeholders).Schema()},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *RequestNaming) MarshalHCL() (map[string]interface{}, error) {
	if properties, err := hcl.NewProperties(me, me.Unknowns); err != nil {
		return nil, err
	} else {
		return properties.EncodeAll(map[string]interface{}{
			// "order":            me.Order,
			"enabled":          me.Enabled,
			"naming_pattern":   me.NamingPattern,
			"management_zones": me.ManagementZones,
			"conditions":       me.Conditions,
			"placeholders":     me.Placeholders,
			"unknowns":         me.Unknowns,
		})
	}
}

func (me *RequestNaming) UnmarshalHCL(decoder hcl.Decoder) error {
	err := decoder.DecodeAll(map[string]interface{}{
		// "order":            &me.Order,
		"enabled":          &me.Enabled,
		"naming_pattern":   &me.NamingPattern,
		"management_zones": &me.ManagementZones,
		"conditions":       &me.Conditions,
		"placeholders":     &me.Placeholders,
		"unknowns":         &me.Unknowns,
	})
	return err
}

func (me *RequestNaming) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		// "order":           me.Order,
		"enabled":         me.Enabled,
		"namingPattern":   me.NamingPattern,
		"managementZones": me.ManagementZones,
		"conditions":      me.Conditions,
		"placeholders":    me.Placeholders,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *RequestNaming) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	return properties.UnmarshalAll(map[string]interface{}{
		// "order":           &me.Order,
		"enabled":         &me.Enabled,
		"namingPattern":   &me.NamingPattern,
		"managementZones": &me.ManagementZones,
		"conditions":      &me.Conditions,
		"placeholders":    &me.Placeholders,
	})
}
