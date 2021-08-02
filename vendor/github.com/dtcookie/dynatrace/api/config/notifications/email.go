package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// EmailConfig Configuration of the email notification.
type EmailConfig struct {
	BaseNotificationConfig
	BccReceivers []string `json:"bccReceivers,omitempty"` // The list of the email BCC-recipients.
	Body         string   `json:"body"`                   // The template of the email notification.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
	CcReceivers  []string `json:"ccReceivers,omitempty"`  // The list of the email CC-recipients.
	Receivers    []string `json:"receivers"`              // The list of the email recipients.
	Subject      string   `json:"subject"`                // The subject of the email notifications.
}

func (me *EmailConfig) GetType() Type {
	return Types.Email
}

func (me *EmailConfig) Schema() map[string]*hcl.Schema {
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
		"bcc_receivers": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The list of the email BCC-recipients",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"cc_receivers": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The list of the email CC-recipients",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"receivers": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The list of the email recipients",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"body": {
			Type:        hcl.TypeString,
			Description: "The template of the email notification.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
		"subject": {
			Type:        hcl.TypeString,
			Description: "The subject of the email notifications",
			Required:    true,
		},
	}
}

func (me *EmailConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	if len(me.BccReceivers) > 0 {
		result["bcc_receivers"] = me.BccReceivers
	}
	if len(me.CcReceivers) > 0 {
		result["cc_receivers"] = me.CcReceivers
	}
	if len(me.Receivers) > 0 {
		result["receivers"] = me.Receivers
	}
	result["body"] = me.Body
	result["subject"] = me.Subject
	return result, nil
}

func (me *EmailConfig) UnmarshalHCL(decoder hcl.Decoder) error {
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
	me.BccReceivers = decoder.GetStringSet("bcc_receivers")
	me.CcReceivers = decoder.GetStringSet("cc_receivers")
	me.Receivers = decoder.GetStringSet("receivers")
	if value, ok := decoder.GetOk("body"); ok {
		me.Body = value.(string)
	}
	if value, ok := decoder.GetOk("subject"); ok {
		me.Subject = value.(string)
	}
	return nil
}

func (me *EmailConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":              me.ID,
		"name":            me.Name,
		"type":            me.GetType(),
		"active":          me.Active,
		"alertingProfile": me.AlertingProfile,
		"bccReceivers":    me.BccReceivers,
		"ccReceivers":     me.CcReceivers,
		"receivers":       me.Receivers,
		"body":            me.Body,
		"subject":         me.Subject,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *EmailConfig) UnmarshalJSON(data []byte) error {
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
		"bccReceivers":    &me.BccReceivers,
		"ccReceivers":     &me.CcReceivers,
		"receivers":       &me.Receivers,
		"body":            &me.Body,
		"subject":         &me.Subject,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
