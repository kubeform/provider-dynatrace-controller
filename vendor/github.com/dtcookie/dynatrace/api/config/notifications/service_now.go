package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// ServiceNowConfig Configuration of the ServiceNow notification.
type ServiceNowConfig struct {
	BaseNotificationConfig
	SendEvents    bool    `json:"sendEvents"`             // Send events into ServiceNow ITOM (`true`).
	SendIncidents bool    `json:"sendIncidents"`          // Send incidents into ServiceNow ITSM (`true`).
	URL           *string `json:"url,omitempty"`          // The URL of the on-premise ServiceNow installation.   This field is mutually exclusive with the **instanceName** field. You can only use one of them.
	Username      string  `json:"username"`               // The username of the ServiceNow account.   Make sure that your user account has the `rest_service`, `web_request_admin`, and `x_dynat_ruxit.Integration` roles.
	InstanceName  *string `json:"instanceName,omitempty"` // The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL.   This field is mutually exclusive with the **url** field. You can only use one of them.
	Message       string  `json:"message"`                // The content of the ServiceNow description.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	Password      *string `json:"password,omitempty"`     // The username to the ServiceNow account
}

func (me *ServiceNowConfig) GetType() Type {
	return Types.ServiceNow
}

func (me *ServiceNowConfig) Schema() map[string]*hcl.Schema {
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
		"send_events": {
			Type:        hcl.TypeBool,
			Description: "Send events into ServiceNow ITOM (`true`)",
			Required:    true,
		},
		"send_incidents": {
			Type:        hcl.TypeBool,
			Description: "Send incidents into ServiceNow ITSM (`true`)",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the on-premise ServiceNow installation.   This field is mutually exclusive with the **instanceName** field. You can only use one of them",
			Optional:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the ServiceNow account.   Make sure that your user account has the `rest_service`, `web_request_admin`, and `x_dynat_ruxit.Integration` roles",
			Required:    true,
		},
		"instance_name": {
			Type:        hcl.TypeString,
			Description: "The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL.   This field is mutually exclusive with the **url** field. You can only use one of them",
			Optional:    true,
		},
		"message": {
			Type:        hcl.TypeString,
			Description: "The content of the ServiceNow description.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "The username to the ServiceNow account",
			Optional:    true,
		},
	}
}

func (me *ServiceNowConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	result["send_events"] = me.SendEvents
	result["send_incidents"] = me.SendIncidents
	if me.URL != nil {
		result["url"] = *me.URL
	}
	result["username"] = me.Username
	if me.InstanceName != nil {
		result["instance_name"] = *me.InstanceName
	}
	result["message"] = me.Message
	if me.Password != nil {
		result["password"] = *me.Password
	} else if v, ok := decoder.GetOk("password"); ok {
		result["password"] = v.(string)
	}

	return result, nil
}

func (me *ServiceNowConfig) UnmarshalHCL(decoder hcl.Decoder) error {
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
		delete(me.Unknowns, "send_events")
		delete(me.Unknowns, "send_incidents")
		delete(me.Unknowns, "url")
		delete(me.Unknowns, "username")
		delete(me.Unknowns, "instance_name")
		delete(me.Unknowns, "message")
		delete(me.Unknowns, "password")
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
	if value, ok := decoder.GetOk("send_events"); ok {
		me.SendEvents = value.(bool)
	}
	if value, ok := decoder.GetOk("send_incidents"); ok {
		me.SendIncidents = value.(bool)
	}
	adapter := hcl.Adapt(decoder)
	me.URL = adapter.GetString("url")
	if value, ok := decoder.GetOk("username"); ok {
		me.Username = value.(string)
	}
	me.InstanceName = adapter.GetString("instance_name")
	if value, ok := decoder.GetOk("message"); ok {
		me.Message = value.(string)
	}
	me.Password = adapter.GetString("password")
	return nil
}

func (me *ServiceNowConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":              me.ID,
		"name":            me.Name,
		"type":            me.GetType(),
		"active":          me.Active,
		"alertingProfile": me.AlertingProfile,
		"sendEvents":      me.SendEvents,
		"sendIncidents":   me.SendIncidents,
		"url":             xjson.Nil(me.URL),
		"username":        me.Username,
		"message":         me.Message,
		"instanceName":    xjson.Nil(me.InstanceName),
		"password":        xjson.Nil(me.Password),
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *ServiceNowConfig) UnmarshalJSON(data []byte) error {
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
		"sendEvents":      &me.SendEvents,
		"sendIncidents":   &me.SendIncidents,
		"url":             &me.URL,
		"username":        &me.Username,
		"message":         &me.Message,
		"instanceName":    &me.InstanceName,
		"password":        &me.Password,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
