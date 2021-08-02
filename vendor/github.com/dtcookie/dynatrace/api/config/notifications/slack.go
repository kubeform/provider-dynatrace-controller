package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// SlackConfig Configuration of the Slack notification.
type SlackConfig struct {
	BaseNotificationConfig
	Title   string  `json:"title"`         // The content of the message.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	URL     *string `json:"url,omitempty"` // The URL of the Slack WebHook.  This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests.
	Channel string  `json:"channel"`       // The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to.
}

func (me *SlackConfig) GetType() Type {
	return Types.Slack
}

func (me *SlackConfig) Schema() map[string]*hcl.Schema {
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
		"title": {
			Type:        hcl.TypeString,
			Description: "The content of the message.  You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the Slack WebHook.  This is confidential information, therefore GET requests return this field with the `null` value, and it is optional for PUT requests",
			Optional:    true,
		},
		"channel": {
			Type:        hcl.TypeString,
			Description: "The channel (for example, `#general`) or the user (for example, `@john.smith`) to send the message to",
			Required:    true,
		},
	}
}

func (me *SlackConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	result["title"] = me.Title
	if me.URL != nil {
		result["url"] = *me.URL
	}
	if me.URL != nil {
		result["url"] = *me.URL
	} else if v, ok := decoder.GetOk("url"); ok {
		result["url"] = v.(string)
	}
	result["channel"] = me.Channel
	return result, nil
}

func (me *SlackConfig) UnmarshalHCL(decoder hcl.Decoder) error {
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
		delete(me.Unknowns, "title")
		delete(me.Unknowns, "url")
		delete(me.Unknowns, "channel")
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
	if value, ok := decoder.GetOk("title"); ok {
		me.Title = value.(string)
	}
	adapter := hcl.Adapt(decoder)
	me.URL = adapter.GetString("url")
	if value, ok := decoder.GetOk("channel"); ok {
		me.Channel = value.(string)
	}
	return nil
}

func (me *SlackConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":              me.ID,
		"name":            me.Name,
		"type":            me.GetType(),
		"active":          me.Active,
		"alertingProfile": me.AlertingProfile,
		"title":           me.Title,
		"url":             me.URL,
		"channel":         me.Channel,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *SlackConfig) UnmarshalJSON(data []byte) error {
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
		"title":           &me.Title,
		"url":             &me.URL,
		"channel":         &me.Channel,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
