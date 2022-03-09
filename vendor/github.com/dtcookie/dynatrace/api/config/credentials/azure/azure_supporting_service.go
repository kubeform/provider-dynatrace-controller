package azure

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AzureSupportingService A supporting service to be monitored.
type AzureSupportingService struct {
	Name             *string                    `json:"name,omitempty"`             // The name of the supporting service.
	MonitoredMetrics []*AzureMonitoredMetric    `json:"monitoredMetrics,omitempty"` // A list of metrics to be monitored for this service. It must include all the recommended metrics.
	Unknowns         map[string]json.RawMessage `json:"-"`
}

func (ass *AzureSupportingService) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the supporting service.",
			Optional:    true,
		},
		"monitored_metrics": {
			Type:        hcl.TypeList,
			Description: "A list of Azure tags to be monitored.  You can specify up to 10 tags. A resource tagged with *any* of the specified tags is monitored.  Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(AzureMonitoredMetric).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (ass *AzureSupportingService) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(ass.Unknowns) > 0 {
		for k, v := range ass.Unknowns {
			m[k] = v
		}
	}
	if ass.Name != nil {
		rawMessage, err := json.Marshal(ass.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	if ass.MonitoredMetrics != nil {
		rawMessage, err := json.Marshal(ass.MonitoredMetrics)
		if err != nil {
			return nil, err
		}
		m["monitoredMetrics"] = rawMessage
	}
	return json.Marshal(m)
}

func (ass *AzureSupportingService) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &ass.Name); err != nil {
			return err
		}
	}
	if v, found := m["monitoredMetrics"]; found {
		if err := json.Unmarshal(v, &ass.MonitoredMetrics); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "monitoredMetrics")

	if len(m) > 0 {
		ass.Unknowns = m
	}
	return nil
}

func (ass *AzureSupportingService) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(ass.Unknowns) > 0 {
		data, err := json.Marshal(ass.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}

	if ass.Name != nil {
		result["name"] = ass.Name
	}
	if ass.MonitoredMetrics != nil {
		entries := []interface{}{}
		for _, entry := range ass.MonitoredMetrics {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["monitored_metrics"] = entries
	}
	return result, nil
}

func (ass *AzureSupportingService) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), ass); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &ass.Unknowns); err != nil {
			return err
		}
		delete(ass.Unknowns, "name")
		delete(ass.Unknowns, "monitored_metrics")

		if len(ass.Unknowns) == 0 {
			ass.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		ass.Name = opt.NewString(value.(string))
	}
	if result, ok := decoder.GetOk("monitored_metrics.#"); ok {
		ass.MonitoredMetrics = []*AzureMonitoredMetric{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(AzureMonitoredMetric)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "monitored_metrics", idx)); err != nil {
				return err
			}
			ass.MonitoredMetrics = append(ass.MonitoredMetrics, entry)
		}
	}
	return nil
}
