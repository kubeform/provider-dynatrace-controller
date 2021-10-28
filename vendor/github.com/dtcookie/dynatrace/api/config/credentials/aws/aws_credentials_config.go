package aws

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AWSCredentialsConfig Configuration of an AWS credentials.
type AWSCredentialsConfig struct {
	ID                          *string                       `json:"id,omitempty"`                          // The unique ID of the credentials.
	Label                       string                        `json:"label"`                                 // The name of the credentials.
	Metadata                    *api.ConfigurationMetadata    `json:"metadata,omitempty"`                    // Metadata useful for debugging
	SupportingServicesToMonitor []*AWSSupportingServiceConfig `json:"supportingServicesToMonitor,omitempty"` // A list of supporting services to be monitored.
	TaggedOnly                  *bool                         `json:"taggedOnly"`                            // Monitor only resources which have specified AWS tags (`true`) or all resources (`false`).
	AuthenticationData          *AWSAuthenticationData        `json:"authenticationData"`                    // A credentials for the AWS authentication.
	PartitionType               PartitionType                 `json:"partitionType"`                         // The type of the AWS partition.
	TagsToMonitor               []*AWSConfigTag               `json:"tagsToMonitor"`                         // A list of AWS tags to be monitored.  You can specify up to 10 tags.  Only applicable when the **taggedOnly** parameter is set to `true`.
	ConnectionStatus            *ConnectionStatus             `json:"connectionStatus,omitempty"`            // The status of the connection to the AWS environment.   * `CONNECTED`: There was a connection within last 10 minutes.  * `DISCONNECTED`: A problem occurred with establishing connection using these credentials. Check whether the data is correct.  * `UNINITIALIZED`: The successful connection has never been established for these credentials.
	Unknowns                    map[string]json.RawMessage    `json:"-"`
}

func (awscc *AWSCredentialsConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"label": {
			Type:        hcl.TypeString,
			Description: "The name of the credentials",
			Optional:    true,
		},
		"tagged_only": {
			Type:        hcl.TypeBool,
			Description: "Monitor only resources which have specified AWS tags (`true`) or all resources (`false`)",
			Required:    true,
		},
		"partition_type": {
			Type:        hcl.TypeString,
			Description: "The type of the AWS partition",
			Required:    true,
		},
		"authentication_data": {
			Type:        hcl.TypeList,
			Description: "credentials for the AWS authentication",
			Required:    true,
			MaxItems:    1,
			Elem: &hcl.Resource{
				Schema: new(AWSAuthenticationData).Schema(),
			},
		},
		"tags_to_monitor": {
			Type:        hcl.TypeList,
			Description: "AWS tags to be monitored. You can specify up to 10 tags. Only applicable when the **tagged_only** parameter is set to `true`",
			Optional:    true,
			MaxItems:    10,
			Elem: &hcl.Resource{
				Schema: new(AWSConfigTag).Schema(),
			},
		},
		"supporting_services_to_monitor": {
			Type:        hcl.TypeList,
			Description: "supporting services to be monitored",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(AWSSupportingServiceConfig).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (awscc *AWSCredentialsConfig) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(awscc.Unknowns) > 0 {
		for k, v := range awscc.Unknowns {
			m[k] = v
		}
	}
	if awscc.ID != nil {
		rawMessage, err := json.Marshal(awscc.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(awscc.Label)
		if err != nil {
			return nil, err
		}
		m["label"] = rawMessage
	}
	if awscc.Metadata != nil {
		rawMessage, err := json.Marshal(awscc.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	if awscc.SupportingServicesToMonitor != nil {
		rawMessage, err := json.Marshal(awscc.SupportingServicesToMonitor)
		if err != nil {
			return nil, err
		}
		m["supportingServicesToMonitor"] = rawMessage
	}

	if rawMessage, err := json.Marshal(opt.Bool(awscc.TaggedOnly)); err == nil {
		m["taggedOnly"] = rawMessage
	} else {
		return nil, err
	}

	if awscc.AuthenticationData != nil {
		rawMessage, err := json.Marshal(awscc.AuthenticationData)
		if err != nil {
			return nil, err
		}
		m["authenticationData"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(awscc.PartitionType)
		if err != nil {
			return nil, err
		}
		m["partitionType"] = rawMessage
	}
	if awscc.TagsToMonitor != nil {
		rawMessage, err := json.Marshal(awscc.TagsToMonitor)
		if err != nil {
			return nil, err
		}
		m["tagsToMonitor"] = rawMessage
	}
	if awscc.ConnectionStatus != nil {
		rawMessage, err := json.Marshal(awscc.ConnectionStatus)
		if err != nil {
			return nil, err
		}
		m["connectionStatus"] = rawMessage
	}
	return json.Marshal(m)
}

// UnmarshalJSON provides custom JSON deserialization
func (awscc *AWSCredentialsConfig) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &awscc.ID); err != nil {
			return err
		}
	}
	if v, found := m["label"]; found {
		if err := json.Unmarshal(v, &awscc.Label); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &awscc.Metadata); err != nil {
			return err
		}
	}
	if v, found := m["supportingServicesToMonitor"]; found {
		if err := json.Unmarshal(v, &awscc.SupportingServicesToMonitor); err != nil {
			return err
		}
	}
	if v, found := m["taggedOnly"]; found {
		if err := json.Unmarshal(v, &awscc.TaggedOnly); err != nil {
			return err
		}
	} else {
		awscc.TaggedOnly = opt.NewBool(false)
	}
	if v, found := m["authenticationData"]; found {
		if err := json.Unmarshal(v, &awscc.AuthenticationData); err != nil {
			return err
		}
	}
	if v, found := m["partitionType"]; found {
		if err := json.Unmarshal(v, &awscc.PartitionType); err != nil {
			return err
		}
	}
	if v, found := m["tagsToMonitor"]; found {
		if err := json.Unmarshal(v, &awscc.TagsToMonitor); err != nil {
			return err
		}
	}
	if v, found := m["connectionStatus"]; found {
		if err := json.Unmarshal(v, &awscc.ConnectionStatus); err != nil {
			return err
		}
	}
	delete(m, "id")
	delete(m, "label")
	delete(m, "metadata")
	delete(m, "supportingServicesToMonitor")
	delete(m, "taggedOnly")
	delete(m, "authenticationData")
	delete(m, "partitionType")
	delete(m, "tagsToMonitor")
	delete(m, "connectionStatus")
	if len(m) > 0 {
		awscc.Unknowns = m
	}
	return nil
}

func (awscc *AWSCredentialsConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if awscc.Unknowns != nil {
		data, err := json.Marshal(awscc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if awscc.SupportingServicesToMonitor != nil {
		entries := []interface{}{}
		for _, entry := range awscc.SupportingServicesToMonitor {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["supporting_services_to_monitor"] = entries
	}
	result["label"] = awscc.Label
	result["tagged_only"] = opt.Bool(awscc.TaggedOnly)

	if awscc.AuthenticationData != nil {
		if marshalled, err := awscc.AuthenticationData.MarshalHCL(); err == nil {
			result["authentication_data"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	result["partition_type"] = awscc.PartitionType
	if awscc.TagsToMonitor != nil {
		entries := []interface{}{}
		for _, entry := range awscc.TagsToMonitor {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["tags_to_monitor"] = entries
	}
	return result, nil
}

func (awscc *AWSCredentialsConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), awscc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &awscc.Unknowns); err != nil {
			return err
		}
		delete(awscc.Unknowns, "supporting_services_to_monitor")
		delete(awscc.Unknowns, "label")
		delete(awscc.Unknowns, "tagged_only")
		delete(awscc.Unknowns, "authentication_data")
		delete(awscc.Unknowns, "partition_type")
		delete(awscc.Unknowns, "tags_to_monitor")
		if len(awscc.Unknowns) == 0 {
			awscc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("label"); ok {
		awscc.Label = value.(string)
	}
	if result, ok := decoder.GetOk("supporting_services_to_monitor.#"); ok {
		awscc.SupportingServicesToMonitor = []*AWSSupportingServiceConfig{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(AWSSupportingServiceConfig)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "supporting_services_to_monitor", idx)); err != nil {
				return err
			}
			awscc.SupportingServicesToMonitor = append(awscc.SupportingServicesToMonitor, entry)
		}
	}
	if _, value := decoder.GetChange("tagged_only"); value != nil {
		awscc.TaggedOnly = opt.NewBool(value.(bool))
	}
	if _, ok := decoder.GetOk("authentication_data.#"); ok {
		awscc.AuthenticationData = new(AWSAuthenticationData)
		if err := awscc.AuthenticationData.UnmarshalHCL(hcl.NewDecoder(decoder, "authentication_data", 0)); err != nil {
			return err
		}
	}
	if value, ok := decoder.GetOk("partition_type"); ok {
		awscc.PartitionType = PartitionType(value.(string))
	}
	if result, ok := decoder.GetOk("tags_to_monitor.#"); ok {
		awscc.TagsToMonitor = []*AWSConfigTag{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(AWSConfigTag)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "tags_to_monitor", idx)); err != nil {
				return err
			}
			awscc.TagsToMonitor = append(awscc.TagsToMonitor, entry)
		}
	}
	return nil
}
