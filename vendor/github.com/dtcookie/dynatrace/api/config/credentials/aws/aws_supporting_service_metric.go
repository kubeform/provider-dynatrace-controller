package aws

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

// AWSSupportingServiceMetric A metric of supporting service to be monitored.
type AWSSupportingServiceMetric struct {
	Name       string                     `json:"name"`       // The name of the metric of the supporting service.
	Statistic  Statistic                  `json:"statistic"`  // The statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM
	Dimensions []string                   `json:"dimensions"` // A list of metric's dimensions names.
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (assm *AWSSupportingServiceMetric) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "the name of the metric of the supporting service",
			Optional:    true,
		},
		"statistic": {
			Type:        hcl.TypeString,
			Description: "the statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM",
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
func (assm *AWSSupportingServiceMetric) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["name"]; found {
		if err := json.Unmarshal(v, &assm.Name); err != nil {
			return err
		}
	}
	if v, found := m["statistic"]; found {
		if err := json.Unmarshal(v, &assm.Statistic); err != nil {
			return err
		}
	}
	if v, found := m["dimensions"]; found {
		if err := json.Unmarshal(v, &assm.Dimensions); err != nil {
			return err
		}
	}
	delete(m, "name")
	delete(m, "statistic")
	delete(m, "dimensions")
	if len(m) > 0 {
		assm.Unknowns = m
	}
	return nil
}

// MarshalJSON provides custom JSON serialization
func (assm *AWSSupportingServiceMetric) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(assm.Unknowns) > 0 {
		for k, v := range assm.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(assm.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(assm.Statistic)
		if err != nil {
			return nil, err
		}
		m["statistic"] = rawMessage
	}
	if assm.Dimensions != nil {
		rawMessage, err := json.Marshal(assm.Dimensions)
		if err != nil {
			return nil, err
		}
		m["dimensions"] = rawMessage
	}
	return json.Marshal(m)
}

func (assm *AWSSupportingServiceMetric) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(assm.Unknowns) > 0 {
		data, err := json.Marshal(assm.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = assm.Name
	result["statistic"] = assm.Statistic
	if assm.Dimensions != nil {
		entries := []interface{}{}
		for _, dimension := range assm.Dimensions {
			entries = append(entries, dimension)
		}
		result["dimensions"] = entries
	}
	return result, nil
}

func (assm *AWSSupportingServiceMetric) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), assm); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &assm.Unknowns); err != nil {
			return err
		}
		delete(assm.Unknowns, "name")
		delete(assm.Unknowns, "statistic")
		delete(assm.Unknowns, "dimensions")
		if len(assm.Unknowns) == 0 {
			assm.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		assm.Name = value.(string)
	}
	if value, ok := decoder.GetOk("statistic"); ok {
		assm.Statistic = Statistic(value.(string))
	}
	if _, ok := decoder.GetOk("dimensions.#"); ok {
		assm.Dimensions = []string{}
		if dims, ok := decoder.GetOk("dimensions"); ok {
			for _, dim := range dims.([]interface{}) {
				assm.Dimensions = append(assm.Dimensions, dim.(string))
			}
		}
	}
	return nil
}
