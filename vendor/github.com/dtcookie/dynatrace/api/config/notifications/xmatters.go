package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// XMattersConfig Configuration of the xMatters notification.
type XMattersConfig struct {
	BaseNotificationConfig
	URL                  string        `json:"url"`                  // The URL of the xMatters WebHook.
	AcceptAnyCertificate bool          `json:"acceptAnyCertificate"` // Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	Headers              []*HTTPHeader `json:"headers,omitempty"`    // A list of the additional HTTP headers.
	Payload              string        `json:"payload"`              // The content of the message.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.
}

func (me *XMattersConfig) GetType() Type {
	return Types.Xmatters
}

func (me *XMattersConfig) Schema() map[string]*hcl.Schema {
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
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL of the xMatters WebHook",
			Required:    true,
		},
		"accept_any_certificate": {
			Type:        hcl.TypeBool,
			Description: "Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates",
			Required:    true,
		},
		"header": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A list of the additional HTTP headers",
			Elem:        &hcl.Resource{Schema: new(HTTPHeader).Schema()},
		},
		"payload": {
			Type:        hcl.TypeString,
			Description: "The content of the notification message.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas",
			Required:    true,
		},
	}
}

func (me *XMattersConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	if len(me.Headers) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Headers {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["header"] = entries
	}
	result["payload"] = me.Payload
	result["url"] = me.URL

	return result, nil
}

func (me *XMattersConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "active")
		delete(me.Unknowns, "alertingProfile")
		delete(me.Unknowns, "acceptAnyCertificate")
		delete(me.Unknowns, "header")
		delete(me.Unknowns, "payload")
		delete(me.Unknowns, "url")
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
	if result, ok := decoder.GetOk("header.#"); ok {
		me.Headers = []*HTTPHeader{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(HTTPHeader)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "header", idx)); err != nil {
				return err
			}
			me.Headers = append(me.Headers, entry)
		}
	}
	if value, ok := decoder.GetOk("payload"); ok {
		me.Payload = value.(string)
	}
	if value, ok := decoder.GetOk("url"); ok {
		me.URL = value.(string)
	}
	return nil
}

func (me *XMattersConfig) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"id":                   me.ID,
		"name":                 me.Name,
		"type":                 me.GetType(),
		"active":               me.Active,
		"alertingProfile":      me.AlertingProfile,
		"acceptAnyCertificate": me.AcceptAnyCertificate,
		"headers":              me.Headers,
		"payload":              me.Payload,
		"url":                  me.URL,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *XMattersConfig) UnmarshalJSON(data []byte) error {
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
		"headers":              &me.Headers,
		"payload":              &me.Payload,
		"url":                  &me.URL,
	}); err != nil {
		return err
	}
	if len(properties) > 0 {
		me.Unknowns = properties
	}
	return nil
}
