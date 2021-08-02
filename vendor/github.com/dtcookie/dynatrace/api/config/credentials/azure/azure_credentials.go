package azure

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AzureCredentials Configuration of Azure app-level credentials.
type AzureCredentials struct {
	ID                        *string                    `json:"id,omitempty"`                 // The Dynatrace entity ID of the Azure credentials configuration.
	Label                     string                     `json:"label"`                        // The unique name of the Azure credentials configuration.  Allowed characters are letters, numbers, and spaces. Also the special characters `.+-_` are allowed.
	DirectoryID               string                     `json:"directoryId"`                  // The Directory ID (also referred to as Tenant ID)  The combination of Application ID and Directory ID must be unique.
	AutoTagging               *bool                      `json:"autoTagging"`                  // The automatic capture of Azure tags is on (`true`) or off (`false`).
	Metadata                  *api.ConfigurationMetadata `json:"metadata,omitempty"`           // Metadata useful for debugging
	MonitorOnlyTagPairs       []*CloudTag                `json:"monitorOnlyTagPairs"`          // A list of Azure tags to be monitored.  You can specify up to 10 tags. A resource tagged with *any* of the specified tags is monitored.  Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`.
	Active                    *bool                      `json:"active,omitempty"`             // The monitoring is enabled (`true`) or disabled (`false`).  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	AppID                     string                     `json:"appId"`                        // The Application ID (also referred to as Client ID)  The combination of Application ID and Directory ID must be unique.
	Key                       *string                    `json:"key,omitempty"`                // The secret key associated with the Application ID.  For security reasons, GET requests return this field as `null`.   Submit your key on creation or update of the configuration. If the field is omitted during an update, the old value remains unaffected.
	MonitorOnlyTaggedEntities *bool                      `json:"monitorOnlyTaggedEntities"`    // Monitor only resources that have specified Azure tags (`true`) or all resources (`false`).
	SupportingServices        []*AzureSupportingService  `json:"supportingServices,omitempty"` // A list of Azure supporting services to be monitored. For each service there's a sublist of its metrics and the metrics' dimensions that should be monitored. All of these elements (services, metrics, dimensions) must have corresponding static definitions on the server.
	Unknowns                  map[string]json.RawMessage `json:"-"`
}

func (ac *AzureCredentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"label": {
			Type:        hcl.TypeString,
			Description: "The unique name of the Azure credentials configuration.  Allowed characters are letters, numbers, and spaces. Also the special characters `.+-_` are allowed",
			Optional:    true,
		},
		"directory_id": {
			Type:        hcl.TypeString,
			Description: "The Directory ID (also referred to as Tenant ID)  The combination of Application ID and Directory ID must be unique",
			Optional:    true,
		},
		"app_id": {
			Type:        hcl.TypeString,
			Description: "The Application ID (also referred to as Client ID)  The combination of Application ID and Directory ID must be unique",
			Optional:    true,
		},
		"auto_tagging": {
			Type:        hcl.TypeBool,
			Description: "The automatic capture of Azure tags is on (`true`) or off (`false`)",
			Optional:    true,
		},
		"monitor_only_tag_pairs": {
			Type:        hcl.TypeList,
			Description: "A list of Azure tags to be monitored.  You can specify up to 10 tags. A resource tagged with *any* of the specified tags is monitored.  Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`",
			Optional:    true,
			MaxItems:    10,
			Elem: &hcl.Resource{
				Schema: new(CloudTag).Schema(),
			},
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "The monitoring is enabled (`true`) or disabled (`false`).  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: " The secret key associated with the Application ID.  For security reasons, GET requests return this field as `null`.   Submit your key on creation or update of the configuration. If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"monitor_only_tagged_entities": {
			Type:        hcl.TypeBool,
			Description: "Monitor only resources that have specified Azure tags (`true`) or all resources (`false`).",
			Required:    true,
		},
		"supporting_services": {
			Type:        hcl.TypeList,
			Description: "A list of Azure supporting services to be monitored. For each service there's a sublist of its metrics and the metrics' dimensions that should be monitored. All of these elements (services, metrics, dimensions) must have corresponding static definitions on the server.",
			Optional:    true,
			MaxItems:    10,
			Elem: &hcl.Resource{
				Schema: new(AzureSupportingService).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (ac *AzureCredentials) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(ac.Unknowns) > 0 {
		for k, v := range ac.Unknowns {
			m[k] = v
		}
	}
	if ac.ID != nil {
		rawMessage, err := json.Marshal(ac.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ac.Label)
		if err != nil {
			return nil, err
		}
		m["label"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(ac.DirectoryID)
		if err != nil {
			return nil, err
		}
		m["directoryId"] = rawMessage
	}
	if rawMessage, err := json.Marshal(opt.Bool(ac.AutoTagging)); err == nil {
		m["autoTagging"] = rawMessage
	} else {
		return nil, err
	}
	if ac.Metadata != nil {
		rawMessage, err := json.Marshal(ac.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	if ac.MonitorOnlyTagPairs != nil {
		rawMessage, err := json.Marshal(ac.MonitorOnlyTagPairs)
		if err != nil {
			return nil, err
		}
		m["monitorOnlyTagPairs"] = rawMessage
	}

	if rawMessage, err := json.Marshal(opt.Bool(ac.Active)); err == nil {
		m["active"] = rawMessage
	} else {
		return nil, err
	}

	if rawMessage, err := json.Marshal(ac.AppID); err == nil {
		m["appId"] = rawMessage
	} else {
		return nil, err
	}

	if ac.Key != nil {
		if rawMessage, err := json.Marshal(ac.Key); err == nil {
			m["key"] = rawMessage
		} else {
			return nil, err
		}
	}
	if rawMessage, err := json.Marshal(opt.Bool(ac.MonitorOnlyTaggedEntities)); err == nil {
		m["monitorOnlyTaggedEntities"] = rawMessage
	} else {
		return nil, err
	}
	if ac.SupportingServices != nil {
		rawMessage, err := json.Marshal(ac.SupportingServices)
		if err != nil {
			return nil, err
		}
		m["supportingServices"] = rawMessage
	}
	return json.Marshal(m)
}

func (ac *AzureCredentials) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &ac.ID); err != nil {
			return err
		}
	}
	if v, found := m["label"]; found {
		if err := json.Unmarshal(v, &ac.Label); err != nil {
			return err
		}
	}
	if v, found := m["directoryId"]; found {
		if err := json.Unmarshal(v, &ac.DirectoryID); err != nil {
			return err
		}
	}
	if v, found := m["autoTagging"]; found {
		if err := json.Unmarshal(v, &ac.AutoTagging); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &ac.Metadata); err != nil {
			return err
		}
	}
	if v, found := m["monitorOnlyTagPairs"]; found {
		if err := json.Unmarshal(v, &ac.MonitorOnlyTagPairs); err != nil {
			return err
		}
	}
	if v, found := m["active"]; found {
		if err := json.Unmarshal(v, &ac.Active); err != nil {
			return err
		}
	}
	if v, found := m["appId"]; found {
		if err := json.Unmarshal(v, &ac.AppID); err != nil {
			return err
		}
	}
	if v, found := m["key"]; found {
		if err := json.Unmarshal(v, &ac.Key); err != nil {
			return err
		}
	}
	if v, found := m["monitorOnlyTaggedEntities"]; found {
		if err := json.Unmarshal(v, &ac.MonitorOnlyTaggedEntities); err != nil {
			return err
		}
	}
	if v, found := m["supportingServices"]; found {
		if err := json.Unmarshal(v, &ac.SupportingServices); err != nil {
			return err
		}
	}
	delete(m, "id")
	delete(m, "label")
	delete(m, "directoryId")
	delete(m, "autoTagging")
	delete(m, "metadata")
	delete(m, "monitorOnlyTagPairs")
	delete(m, "active")
	delete(m, "appId")
	delete(m, "key")
	delete(m, "monitorOnlyTaggedEntities")
	delete(m, "supportingServices")

	if len(m) > 0 {
		ac.Unknowns = m
	}
	return nil
}

func (ac *AzureCredentials) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ac.Unknowns) > 0 {
		data, err := json.Marshal(ac.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}

	result["label"] = ac.Label
	result["directory_id"] = ac.DirectoryID
	result["auto_tagging"] = opt.Bool(ac.AutoTagging)
	if ac.MonitorOnlyTagPairs != nil {
		entries := []interface{}{}
		for _, entry := range ac.MonitorOnlyTagPairs {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["monitor_only_tag_pairs"] = entries
	}
	result["active"] = opt.Bool(ac.Active)
	result["app_id"] = ac.AppID
	if ac.Key != nil {
		result["key"] = *ac.Key
	}
	result["monitor_only_tagged_entities"] = opt.Bool(ac.MonitorOnlyTaggedEntities)
	if ac.SupportingServices != nil {
		entries := []interface{}{}
		for _, entry := range ac.SupportingServices {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}

		result["supporting_services"] = entries
	}
	return result, nil
}

func (ac *AzureCredentials) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ac); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ac.Unknowns); err != nil {
			return err
		}
		delete(ac.Unknowns, "label")
		delete(ac.Unknowns, "directory_id")
		delete(ac.Unknowns, "auto_tagging")
		delete(ac.Unknowns, "monitor_only_tag_pairs")
		delete(ac.Unknowns, "active")
		delete(ac.Unknowns, "app_id")
		delete(ac.Unknowns, "key")
		delete(ac.Unknowns, "monitor_only_tagged_entities")
		delete(ac.Unknowns, "supporting_services")
		if len(ac.Unknowns) == 0 {
			ac.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("label"); ok {
		ac.Label = value.(string)
	}
	if value, ok := decoder.GetOk("directory_id"); ok {
		ac.DirectoryID = value.(string)
	}
	if value, ok := decoder.GetOk("auto_tagging"); ok {
		ac.AutoTagging = opt.NewBool(value.(bool))
	}
	if result, ok := decoder.GetOk("monitor_only_tag_pairs.#"); ok {
		ac.MonitorOnlyTagPairs = []*CloudTag{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(CloudTag)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "monitor_only_tag_pairs", idx)); err != nil {
				return err
			}
			ac.MonitorOnlyTagPairs = append(ac.MonitorOnlyTagPairs, entry)
		}
	}
	if value, ok := decoder.GetOk("active"); ok {
		ac.Active = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("app_id"); ok {
		ac.AppID = value.(string)
	}
	if value, ok := decoder.GetOk("key"); ok {
		ac.Key = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("monitor_only_tagged_entities"); ok {
		ac.MonitorOnlyTaggedEntities = opt.NewBool(value.(bool))
	}
	if result, ok := decoder.GetOk("supporting_services.#"); ok {
		ac.SupportingServices = []*AzureSupportingService{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(AzureSupportingService)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "supporting_services", idx)); err != nil {
				return err
			}
			ac.SupportingServices = append(ac.SupportingServices, entry)
		}
	}
	return nil
}
