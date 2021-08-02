package aws

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// AWSSupportingServiceConfig A supporting service to be monitored.
type AWSSupportingServiceConfig struct {
	Name             string                        `json:"name"`             // The name of the supporting service.
	MonitoredMetrics []*AWSSupportingServiceMetric `json:"monitoredMetrics"` // A list of metrics to be monitored for this service.
	Unknowns         map[string]json.RawMessage    `json:"-"`
}

func (assc *AWSSupportingServiceConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "the name of the supporting service",
			Optional:    true,
		},
		"monitored_metrics": {
			Type:        hcl.TypeList,
			Description: "a list of metrics to be monitored for this service",
			Optional:    true,
			MaxItems:    10,
			Elem: &hcl.Resource{
				Schema: new(AWSSupportingServiceMetric).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (assc *AWSSupportingServiceConfig) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &assc.Name); err != nil {
			return err
		}
	}
	if v, found := m["monitoredMetrics"]; found {
		if err := json.Unmarshal(v, &assc.MonitoredMetrics); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "monitoredMetrics")
	if len(m) > 0 {
		assc.Unknowns = m
	}
	return nil
}

// MarshalJSON provides custom JSON serialization
func (assc *AWSSupportingServiceConfig) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(assc.Unknowns) > 0 {
		for k, v := range assc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(assc.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	if assc.MonitoredMetrics != nil {
		rawMessage, err := json.Marshal(assc.MonitoredMetrics)
		if err != nil {
			return nil, err
		}
		m["monitoredMetrics"] = rawMessage
	}
	return json.Marshal(m)
}

func (assc *AWSSupportingServiceConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(assc.Unknowns) > 0 {
		data, err := json.Marshal(assc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = assc.Name
	if assc.MonitoredMetrics != nil {
		entries := []interface{}{}
		for _, monitoredMetric := range assc.MonitoredMetrics {
			if marshalled, err := monitoredMetric.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["monitored_metrics"] = entries
	}
	return result, nil
}

func (assc *AWSSupportingServiceConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), assc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &assc.Unknowns); err != nil {
			return err
		}
		delete(assc.Unknowns, "name")
		delete(assc.Unknowns, "monitored_metrics")
		if len(assc.Unknowns) == 0 {
			assc.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		assc.Name = value.(string)
	}
	if result, ok := decoder.GetOk("monitored_metrics.#"); ok {
		assc.MonitoredMetrics = []*AWSSupportingServiceMetric{}
		for idx := 0; idx < result.(int); idx++ {
			monitoredMetric := new(AWSSupportingServiceMetric)
			if err := monitoredMetric.UnmarshalHCL(hcl.NewDecoder(decoder, "monitored_metrics", idx)); err != nil {
				return err
			}
			assc.MonitoredMetrics = append(assc.MonitoredMetrics, monitoredMetric)
		}
	}
	return nil
}
