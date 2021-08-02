package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// AnsibleTowerConfig Configuration of the Ansible Tower notification.
type AnsibleTowerConfig struct {
	BaseNotificationConfig
	AcceptAnyCertificate bool    `json:"acceptAnyCertificate"` // Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	CustomMessage        string  `json:"customMessage"`        // The custom message of the notification.   This message will be displayed in the extra variables **Message** field of your job template.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	JobTemplateID        int32   `json:"jobTemplateID"`        // The ID of the target Ansible Tower job template.
	JobTemplateURL       string  `json:"jobTemplateURL"`       // The URL of the target Ansible Tower job template.
	Password             *string `json:"password,omitempty"`   // The password for the Ansible Tower account.
	Username             string  `json:"username"`             // The username of the Ansible Tower account.
}

func (me *AnsibleTowerConfig) GetType() Type {
	return Types.Ansibletower
}

func (me *AnsibleTowerConfig) Schema() map[string]*hcl.Schema {
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
		"accept_any_certificate": {
			Type:        hcl.TypeBool,
			Description: "Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates",
			Required:    true,
		},
		"custom_message": {
			Type:        hcl.TypeString,
			Description: "The custom message of the notification.   This message will be displayed in the extra variables **Message** field of your job template.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"job_template_id": {
			Type:        hcl.TypeInt,
			Description: "The ID of the target Ansible Tower job template",
			Required:    true,
		},
		"job_template_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the target Ansible Tower job template",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "The password for the Ansible Tower account",
			Optional:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the Ansible Tower account",
			Required:    true,
		},
	}
}

func (me *AnsibleTowerConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	result["accept_any_certificate"] = me.AcceptAnyCertificate
	result["custom_message"] = me.CustomMessage
	result["job_template_id"] = int(me.JobTemplateID)
	result["job_template_url"] = me.JobTemplateURL
	if me.Password != nil {
		result["password"] = *me.Password
	} else if v, ok := decoder.GetOk("password"); ok {
		result["password"] = v.(string)
	}
	result["username"] = me.Username
	return result, nil
}

func (me *AnsibleTowerConfig) UnmarshalHCL(decoder hcl.Decoder) error {
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
		delete(me.Unknowns, "accept_any_certificate")
		delete(me.Unknowns, "custom_message")
		delete(me.Unknowns, "job_template_id")
		delete(me.Unknowns, "job_template_url")
		delete(me.Unknowns, "password")
		delete(me.Unknowns, "username")
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
	if value, ok := decoder.GetOk("accept_any_certificate"); ok {
		me.AcceptAnyCertificate = value.(bool)
	}
	if value, ok := decoder.GetOk("custom_message"); ok {
		me.CustomMessage = value.(string)
	}
	if value, ok := decoder.GetOk("job_template_id"); ok {
		me.JobTemplateID = int32(value.(int))
	}
	if value, ok := decoder.GetOk("job_template_url"); ok {
		me.JobTemplateURL = value.(string)
	}
	adapter := hcl.Adapt(decoder)
	me.Password = adapter.GetString("password")
	if value, ok := decoder.GetOk("username"); ok {
		me.Username = value.(string)
	}
	return nil
}

func (me *AnsibleTowerConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":                   me.ID,
		"name":                 me.Name,
		"type":                 me.GetType(),
		"active":               me.Active,
		"alertingProfile":      me.AlertingProfile,
		"acceptAnyCertificate": me.AcceptAnyCertificate,
		"customMessage":        me.CustomMessage,
		"jobTemplateID":        me.JobTemplateID,
		"jobTemplateURL":       me.JobTemplateURL,
		"password":             me.Password,
		"username":             me.Username,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *AnsibleTowerConfig) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"id":                   &me.ID,
		"name":                 &me.Name,
		"type":                 &me.Type,
		"active":               &me.Active,
		"alertingProfile":      &me.AlertingProfile,
		"acceptAnyCertificate": &me.AcceptAnyCertificate,
		"customMessage":        &me.CustomMessage,
		"jobTemplateID":        &me.JobTemplateID,
		"jobTemplateURL":       &me.JobTemplateURL,
		"password":             &me.Password,
		"username":             &me.Username,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
