package azure

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AzureMonitoredMetric A metric of supporting service to be monitored.
type AzureMonitoredMetric struct {
	Name       *string                    `json:"name,omitempty"`       // The name of the metric of the supporting service.
	Dimensions []string                   `json:"dimensions,omitempty"` // A list of metric's dimensions names. It must include all the recommended dimensions.
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (amm *AzureMonitoredMetric) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "the name of the metric of the supporting service",
			Optional:    true,
		},
		"dimensions": {
			Type:        hcl.TypeList,
			Description: "a list of metric's dimensions names",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

// UnmarshalJSON provides custom JSON deserialization
func (amm *AzureMonitoredMetric) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &amm.Name); err != nil {
			return err
		}
	}
	if v, found := m["dimensions"]; found {
		if err := json.Unmarshal(v, &amm.Dimensions); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "dimensions")
	if len(m) > 0 {
		amm.Unknowns = m
	}
	return nil
}

// MarshalJSON provides custom JSON serialization
func (amm *AzureMonitoredMetric) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(amm.Unknowns) > 0 {
		for k, v := range amm.Unknowns {
			m[k] = v
		}
	}
	if amm.Name != nil {
		rawMessage, err := json.Marshal(amm.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	if amm.Dimensions != nil {
		rawMessage, err := json.Marshal(amm.Dimensions)
		if err != nil {
			return nil, err
		}
		m["dimensions"] = rawMessage
	}
	return json.Marshal(m)
}

func (amm *AzureMonitoredMetric) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(amm.Unknowns) > 0 {
		data, err := json.Marshal(amm.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = *amm.Name
	if amm.Dimensions != nil {
		entries := []interface{}{}
		for _, dimension := range amm.Dimensions {
			entries = append(entries, dimension)
		}
		result["dimensions"] = entries
	}
	return result, nil
}

func (amm *AzureMonitoredMetric) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), amm); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &amm.Unknowns); err != nil {
			return err
		}
		delete(amm.Unknowns, "name")
		delete(amm.Unknowns, "dimensions")
		if len(amm.Unknowns) == 0 {
			amm.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		amm.Name = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("dimensions.#"); ok {
		amm.Dimensions = []string{}
		if dims, ok := decoder.GetOk("dimensions"); ok {
			for _, dim := range dims.([]interface{}) {
				amm.Dimensions = append(amm.Dimensions, dim.(string))
			}
		}
	}
	return nil
}
