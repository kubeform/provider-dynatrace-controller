package dashboards

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// CustomFilterChartSeriesDimensionConfig Configuration of the charted metric splitting
type CustomFilterChartSeriesDimensionConfig struct {
	ID              string                     `json:"id"`             // The ID of the dimension by which the metric is split
	Name            *string                    `json:"name,omitempty"` // The name of the dimension by which the metric is split
	Values          []string                   `json:"values"`         // The splitting value
	EntityDimension *bool                      `json:"entityDimension,omitempty"`
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *CustomFilterChartSeriesDimensionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Description: "The ID of the dimension by which the metric is split",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the dimension by which the metric is split",
			Optional:    true,
		},
		"values": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The splitting value",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"entity_dimension": {
			Type:        hcl.TypeBool,
			Description: "",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *CustomFilterChartSeriesDimensionConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["id"] = me.ID
	if me.Name != nil {
		result["name"] = opt.String(me.Name)
	}
	if len(me.Values) > 0 {
		result["values"] = me.Values
	}
	if me.EntityDimension != nil {
		result["entity_dimension"] = opt.Bool(me.EntityDimension)
	}
	return result, nil
}

func (me *CustomFilterChartSeriesDimensionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "id")
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "values")
		delete(me.Unknowns, "entity_dimension")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("id"); ok {
		me.ID = value.(string)
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = opt.NewString(value.(string))
	}
	me.Values = decoder.GetStringSet("values")
	if _, value := decoder.GetChange("entity_dimension"); value != nil {
		me.EntityDimension = opt.NewBool(value.(bool))
	}
	return nil
}

func (me *CustomFilterChartSeriesDimensionConfig) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(me.Unknowns) > 0 {
		for k, v := range me.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(me.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	if me.Name != nil {
		rawMessage, err := json.Marshal(me.Name)
		if err != nil {
			return nil, err
		}
		m["name"] = rawMessage
	}
	values := me.Values
	if values == nil {
		values = []string{}
	}
	{
		rawMessage, err := json.Marshal(values)
		if err != nil {
			return nil, err
		}
		m["values"] = rawMessage
	}
	if me.EntityDimension != nil {
		rawMessage, err := json.Marshal(me.EntityDimension)
		if err != nil {
			return nil, err
		}
		m["entityDimension"] = rawMessage
	}
	return json.Marshal(m)
}

func (me *CustomFilterChartSeriesDimensionConfig) UnmarshalJSON(data []byte) error {
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
	if err := m.Unmarshal("values", &me.Values); err != nil {
		return err
	}
	if err := m.Unmarshal("entityDimension", &me.EntityDimension); err != nil {
		return err
	}
	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
