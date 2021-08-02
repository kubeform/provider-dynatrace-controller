package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// PagerDutyConfig Configuration of the PagerDuty notification.
type PagerDutyConfig struct {
	BaseNotificationConfig
	Account       string  `json:"account"`                 // The name of the PagerDuty account.
	ServiceAPIKey *string `json:"serviceApiKey,omitempty"` // The API key to access PagerDuty.
	ServiceName   string  `json:"serviceName"`             // The name of the service.
}

func (me *PagerDutyConfig) GetType() Type {
	return Types.PagerDuty
}

func (me *PagerDutyConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the notification configuration",
			Required:    true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "The configuration is enabled (`true`) or disabled (`false`)",
			Required:    true,
		},
		"alerting_profile": {
			Type:        hcl.TypeString,
			Description: "The ID of the associated alerting profile",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
		"account": {
			Type:        hcl.TypeString,
			Description: "The name of the PagerDuty account",
			Required:    true,
		},
		"service_api_key": {
			Type:        hcl.TypeString,
			Description: "The API key to access PagerDuty",
			Optional:    true,
		},
		"service_name": {
			Type:        hcl.TypeString,
			Description: "The name of the service",
			Required:    true,
		},
	}
}

func (me *PagerDutyConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	result["active"] = me.Active
	result["alerting_profile"] = me.AlertingProfile
	result["account"] = me.Account
	if me.ServiceAPIKey != nil {
		result["service_api_key"] = *me.ServiceAPIKey
	} else if v, ok := decoder.GetOk("service_api_key"); ok {
		result["service_api_key"] = v.(string)
	}

	result["service_name"] = me.ServiceName

	return result, nil
}

func (me *PagerDutyConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "active")
		delete(me.Unknowns, "alerting_profile")
		delete(me.Unknowns, "account")
		delete(me.Unknowns, "service_api_key")
		delete(me.Unknowns, "service_name")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("active"); ok {
		me.Active = value.(bool)
	}
	if value, ok := decoder.GetOk("alerting_profile"); ok {
		me.AlertingProfile = value.(string)
	}
	if value, ok := decoder.GetOk("account"); ok {
		me.Account = value.(string)
	}
	adapter := hcl.Adapt(decoder)
	me.ServiceAPIKey = adapter.GetString("service_api_key")
	if value, ok := decoder.GetOk("service_name"); ok {
		me.ServiceName = value.(string)
	}
	return nil
}

func (me *PagerDutyConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":              me.ID,
		"name":            me.Name,
		"type":            me.GetType(),
		"active":          me.Active,
		"alertingProfile": me.AlertingProfile,
		"account":         me.Account,
		"serviceApiKey":   me.ServiceAPIKey,
		"serviceName":     me.ServiceName,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *PagerDutyConfig) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"id":              &me.ID,
		"name":            &me.Name,
		"type":            &me.Type,
		"active":          &me.Active,
		"alertingProfile": &me.AlertingProfile,
		"account":         &me.Account,
		"serviceApiKey":   &me.ServiceAPIKey,
		"serviceName":     &me.ServiceName,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
