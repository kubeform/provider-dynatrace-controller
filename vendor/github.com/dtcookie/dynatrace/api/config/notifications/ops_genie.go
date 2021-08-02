package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// OpsGenieConfig Configuration of the OpsGenie notification.
type OpsGenieConfig struct {
	BaseNotificationConfig
	APIKey  *string `json:"apiKey,omitempty"` // The API key to access OpsGenie.
	Domain  string  `json:"domain"`           // The region domain of the OpsGenie.
	Message string  `json:"message"`          // The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.
}

func (me *OpsGenieConfig) GetType() Type {
	return Types.OpsGenie
}

func (me *OpsGenieConfig) Schema() map[string]*hcl.Schema {
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
		"api_key": {
			Type:        hcl.TypeString,
			Description: "The API key to access OpsGenie",
			Optional:    true,
		},
		"domain": {
			Type:        hcl.TypeString,
			Description: "The region domain of the OpsGenie",
			Required:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the message.  You can use the following placeholders:  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem",
			Required:    true,
		},
	}
}

func (me *OpsGenieConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	if me.APIKey != nil {
		result["api_key"] = *me.APIKey
	} else if v, ok := decoder.GetOk("api_key"); ok {
		result["api_key"] = v.(string)
	}

	result["domain"] = me.Domain
	result["message"] = me.Message
	return result, nil
}

func (me *OpsGenieConfig) UnmarshalHCL(decoder hcl.Decoder) error {
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
		delete(me.Unknowns, "api_key")
		delete(me.Unknowns, "domain")
		delete(me.Unknowns, "message")
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
	adapter := hcl.Adapt(decoder)
	me.APIKey = adapter.GetString("api_key")
	if value, ok := decoder.GetOk("domain"); ok {
		me.Domain = value.(string)
	}
	if value, ok := decoder.GetOk("message"); ok {
		me.Message = value.(string)
	}

	return nil
}

func (me *OpsGenieConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":              me.ID,
		"name":            me.Name,
		"type":            me.GetType(),
		"active":          me.Active,
		"alertingProfile": me.AlertingProfile,
		"apiKey":          me.APIKey,
		"domain":          me.Domain,
		"message":         me.Message,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *OpsGenieConfig) UnmarshalJSON(data []byte) error {
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
		"apiKey":          &me.APIKey,
		"domain":          &me.Domain,
		"message":         &me.Message,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
