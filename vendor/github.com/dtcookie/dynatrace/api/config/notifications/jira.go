package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// JiraConfig Configuration of the Jira notification.
type JiraConfig struct {
	BaseNotificationConfig
	IssueType   string  `json:"issueType"`          // The type of the Jira issue to be created by this notification.
	Password    *string `json:"password,omitempty"` // The password for the Jira profile.
	ProjectKey  string  `json:"projectKey"`         // The project key of the Jira issue to be created by this notification.
	Summary     string  `json:"summary"`            // The summary of the Jira issue to be created by this notification.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	URL         string  `json:"url"`                // The URL of the Jira API endpoint.
	Username    string  `json:"username"`           // The username of the Jira profile.
	Description string  `json:"description"`        // The description of the Jira issue to be created by this notification.   You can use same placeholders as in issue summary.
}

func (me *JiraConfig) GetType() Type {
	return Types.Jira
}

func (me *JiraConfig) Schema() map[string]*hcl.Schema {
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
		"issue_type": {
			Type:        hcl.TypeString,
			Description: "The type of the Jira issue to be created by this notification",
			Required:    true,
		},
		"password": {
			Type:        hcl.TypeString,
			Description: "The password for the Jira profile",
			Optional:    true,
		},
		"project_key": {
			Type:        hcl.TypeString,
			Description: "The project key of the Jira issue to be created by this notification",
			Required:    true,
		},
		"summary": {
			Type:        hcl.TypeString,
			Description: "The summary of the Jira issue to be created by this notification.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the Jira API endpoint",
			Required:    true,
		},
		"username": {
			Type:        hcl.TypeString,
			Description: "The username of the Jira profile",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "The description of the Jira issue to be created by this notification.   You can use same placeholders as in issue summary",
			Required:    true,
		},
	}
}

func (me *JiraConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	result["issue_type"] = me.IssueType
	if me.Password != nil {
		result["password"] = *me.Password
	} else if v, ok := decoder.GetOk("password"); ok {
		result["password"] = v.(string)
	}
	result["project_key"] = me.ProjectKey
	result["summary"] = me.Summary
	result["url"] = me.URL
	result["username"] = me.Username
	result["description"] = me.Description
	return result, nil
}

func (me *JiraConfig) UnmarshalHCL(decoder hcl.Decoder) error {
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
		delete(me.Unknowns, "issue_type")
		delete(me.Unknowns, "password")
		delete(me.Unknowns, "project_key")
		delete(me.Unknowns, "summary")
		delete(me.Unknowns, "url")
		delete(me.Unknowns, "username")
		delete(me.Unknowns, "description")
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
	if value, ok := decoder.GetOk("issue_type"); ok {
		me.IssueType = value.(string)
	}
	adapter := hcl.Adapt(decoder)
	me.Password = adapter.GetString("password")
	if value, ok := decoder.GetOk("project_key"); ok {
		me.ProjectKey = value.(string)
	}
	if value, ok := decoder.GetOk("summary"); ok {
		me.Summary = value.(string)
	}
	if value, ok := decoder.GetOk("url"); ok {
		me.URL = value.(string)
	}
	if value, ok := decoder.GetOk("username"); ok {
		me.Username = value.(string)
	}
	if value, ok := decoder.GetOk("description"); ok {
		me.Description = value.(string)
	}
	return nil
}

func (me *JiraConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":              me.ID,
		"name":            me.Name,
		"type":            me.GetType(),
		"active":          me.Active,
		"alertingProfile": me.AlertingProfile,
		"issueType":       me.IssueType,
		"password":        me.Password,
		"projectKey":      me.ProjectKey,
		"summary":         me.Summary,
		"url":             me.URL,
		"username":        me.Username,
		"description":     me.Description,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *JiraConfig) UnmarshalJSON(data []byte) error {
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
		"issueType":       &me.IssueType,
		"password":        &me.Password,
		"projectKey":      &me.ProjectKey,
		"summary":         &me.Summary,
		"url":             &me.URL,
		"username":        &me.Username,
		"description":     &me.Description,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
