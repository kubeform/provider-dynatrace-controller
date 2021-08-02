package customservices

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// CustomService has no documentation
type CustomService struct {
	ID                  *string                    `json:"id,omitempty"`                  // The ID of the custom service (UUID)
	Name                string                     `json:"name"`                          // The name of the custom service, displayed in the UI
	Order               *string                    `json:"order,omitempty"`               // The order string. Sorting custom services alphabetically by their order string determines their relative ordering. Typically this is managed by Dynatrace internally and will not be present in GET responses
	Enabled             bool                       `json:"enabled"`                       // Custom service enabled/disabled
	Rules               []*DetectionRule           `json:"rules,omitempty"`               // The list of rules defining the custom service
	QueueEntryPoint     *bool                      `json:"queueEntryPoint"`               // The queue entry point flag. Set to `true` for custom messaging services
	QueueEntryPointType *QueueEntryPointType       `json:"queueEntryPointType,omitempty"` // The queue entry point type
	ProcessGroups       []string                   `json:"processGroups,omitempty"`       // The list of process groups the custom service should belong to
	Metadata            *api.ConfigMetadata        `json:"metadata,omitempty"`
	Unknowns            map[string]json.RawMessage `json:"-"`
}

func (me *CustomService) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the custom service, displayed in the UI",
			Required:    true,
		},
		// "order": {
		// 	Type:        hcl.TypeString,
		// 	Description: "The order string. Sorting custom services alphabetically by their order string determines their relative ordering. Typically this is managed by Dynatrace internally and will not be present in GET responses",
		// 	Optional:    true,
		// },
		"technology": {
			Type:        hcl.TypeString,
			Description: "Matcher applying to the file name (ENDS_WITH, EQUALS or STARTS_WITH). Default value is ENDS_WITH (if applicable)",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Custom service enabled/disabled",
			Required:    true,
		},
		"queue_entry_point": {
			Type:        hcl.TypeBool,
			Description: "The queue entry point flag. Set to `true` for custom messaging services",
			Optional:    true,
		},
		"queue_entry_point_type": {
			Type:        hcl.TypeString,
			Description: "The queue entry point type (IBM_MQ, JMS, KAFKA, MSMQ or RABBIT_MQ)",
			Optional:    true,
		},
		"process_groups": {
			Type:        hcl.TypeSet,
			Description: "The list of process groups the custom service should belong to",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"rule": {
			Type:        hcl.TypeList,
			Description: "The list of rules defining the custom service",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(DetectionRule).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *CustomService) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	// if me.Order != nil {
	// 	result["order"] = *me.Order
	// }
	result["enabled"] = me.Enabled
	if len(me.Rules) > 0 {
		entries := []interface{}{}
		for _, entry := range me.Rules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rule"] = entries
	}
	if me.QueueEntryPoint != nil {
		result["queue_entry_point"] = opt.Bool(me.QueueEntryPoint)
	}
	if me.QueueEntryPointType != nil {
		result["queue_entry_point_type"] = string(*me.QueueEntryPointType)
	}
	if me.ProcessGroups != nil {
		result["process_groups"] = me.ProcessGroups
	}
	return result, nil
}

func (me *CustomService) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		// delete(me.Unknowns, "order")
		delete(me.Unknowns, "technology")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "queue_entry_point")
		delete(me.Unknowns, "queue_entry_point_type")
		delete(me.Unknowns, "process_groups")
		delete(me.Unknowns, "rule")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	adapter := hcl.Adapt(decoder)
	// me.Order = adapter.GetString("order")
	me.Enabled = opt.Bool(adapter.GetBool("enabled"))
	if result, ok := decoder.GetOk("rule.#"); ok {
		me.Rules = []*DetectionRule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(DetectionRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "rule", idx)); err != nil {
				return err
			}
			me.Rules = append(me.Rules, entry)
		}
	}
	me.QueueEntryPoint = adapter.GetBool("queue_entry_point")
	if value := adapter.GetString("queue_entry_point_type"); value != nil && len(*value) > 0 {
		me.QueueEntryPointType = QueueEntryPointType(*value).Ref()
	}
	me.ProcessGroups = decoder.GetStringSet("process_groups")
	return nil
}

func (me *CustomService) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("id", me.ID); err != nil {
		return nil, err
	}
	if err := m.Marshal("name", me.Name); err != nil {
		return nil, err
	}
	if err := m.Marshal("order", me.Order); err != nil {
		return nil, err
	}
	if err := m.Marshal("enabled", me.Enabled); err != nil {
		return nil, err
	}
	if err := m.Marshal("rules", me.Rules); err != nil {
		return nil, err
	}
	if err := m.Marshal("queueEntryPoint", me.QueueEntryPoint); err != nil {
		return nil, err
	}
	if err := m.Marshal("queueEntryPointType", me.QueueEntryPointType); err != nil {
		return nil, err
	}
	if err := m.Marshal("processGroups", me.ProcessGroups); err != nil {
		return nil, err
	}
	if err := m.Marshal("metadata", me.Metadata); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *CustomService) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("id", &me.ID); err != nil {
		return err
	}
	if err := m.Unmarshal("name", &me.Name); err != nil {
		return err
	}
	if err := m.Unmarshal("order", &me.Order); err != nil {
		return err
	}
	if err := m.Unmarshal("enabled", &me.Enabled); err != nil {
		return err
	}
	if err := m.Unmarshal("rules", &me.Rules); err != nil {
		return err
	}
	if err := m.Unmarshal("queueEntryPoint", &me.QueueEntryPoint); err != nil {
		return err
	}
	if err := m.Unmarshal("queueEntryPointType", &me.QueueEntryPointType); err != nil {
		return err
	}
	if err := m.Unmarshal("processGroups", &me.ProcessGroups); err != nil {
		return err
	}
	if err := m.Unmarshal("metadata", &me.Metadata); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// QueueEntryPointType has no documentation
type QueueEntryPointType string

func (me QueueEntryPointType) Ref() *QueueEntryPointType {
	return &me
}

// QueueEntryPointTypes offers the known enum values
var QueueEntryPointTypes = struct {
	IBMMQ    QueueEntryPointType
	Jms      QueueEntryPointType
	Kafka    QueueEntryPointType
	MSMQ     QueueEntryPointType
	RabbitMQ QueueEntryPointType
	Values   func() []string
}{
	"IBM_MQ",
	"JMS",
	"KAFKA",
	"MSMQ",
	"RABBIT_MQ",
	func() []string {
		return []string{"IBM_MQ", "JMS", "KAFKA", "MSMQ", "RABBIT_MQ"}
	},
}
