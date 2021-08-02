package requestattributes

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// RequestAttribute has no documentation
type RequestAttribute struct {
	ID                      *string                    `json:"id,omitempty"`            // The ID of the request attribute.
	Name                    string                     `json:"name"`                    // The name of the request attribute.
	SkipPersonalDataMasking *bool                      `json:"skipPersonalDataMasking"` // Personal data masking flag. Set `true` to skip masking.   Warning: This will potentially access personalized data.
	Confidential            *bool                      `json:"confidential"`            // Confidential data flag. Set `true` to treat the captured data as confidential.
	DataSources             []*DataSource              `json:"dataSources"`             // The list of data sources.
	DataType                DataType                   `json:"dataType"`                // The data type of the request attribute.
	Normalization           Normalization              `json:"normalization"`           // String values transformation.   If the **dataType** is not `string`, set the `Original` here.
	Enabled                 *bool                      `json:"enabled"`                 // The request attribute is enabled (`true`) or disabled (`false`).
	Aggregation             Aggregation                `json:"aggregation"`             // Aggregation type for the request values.
	Metadata                *api.ConfigMetadata        `json:"metadata,omitempty"`      // Metadata useful for debugging
	Unknowns                map[string]json.RawMessage `json:"-"`
}

func (me *RequestAttribute) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the request attribute",
			Required:    true,
		},
		"skip_personal_data_masking": {
			Type:        hcl.TypeBool,
			Description: "Personal data masking flag. Set `true` to skip masking.   Warning: This will potentially access personalized data",
			Optional:    true,
		},
		"confidential": {
			Type:        hcl.TypeBool,
			Description: "Confidential data flag. Set `true` to treat the captured data as confidential",
			Optional:    true,
		},
		"data_sources": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "The list of data sources",
			Elem: &hcl.Resource{
				Schema: new(DataSource).Schema(),
			},
		},
		"data_type": {
			Type:        hcl.TypeString,
			Description: "The data type of the request attribute",
			Required:    true,
		},
		"normalization": {
			Type:        hcl.TypeString,
			Description: "String values transformation.   If the **dataType** is not `string`, set the `Original` here",
			Required:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The request attribute is enabled (`true`) or disabled (`false`)",
			Optional:    true,
		},
		"aggregation": {
			Type:        hcl.TypeString,
			Description: "Aggregation type for the request values",
			Required:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *RequestAttribute) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["name"] = me.Name
	if me.SkipPersonalDataMasking != nil {
		result["skip_personal_data_masking"] = opt.Bool(me.SkipPersonalDataMasking)
	}
	if me.Confidential != nil {
		result["confidential"] = opt.Bool(me.Confidential)
	}
	if len(me.DataSources) > 0 {
		entries := []interface{}{}
		for _, entry := range me.DataSources {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["data_sources"] = entries
	}
	result["data_type"] = string(me.DataType)
	result["normalization"] = string(me.Normalization)
	if me.Enabled != nil {
		result["enabled"] = opt.Bool(me.Enabled)
	}
	result["aggregation"] = string(me.Aggregation)
	return result, nil
}

func (me *RequestAttribute) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "skip_personal_data_masking")
		delete(me.Unknowns, "confidential")
		delete(me.Unknowns, "data_sources")
		delete(me.Unknowns, "data_type")
		delete(me.Unknowns, "normalization")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "aggregation")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if _, value := decoder.GetChange("skip_personal_data_masking"); value != nil {
		me.SkipPersonalDataMasking = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("confidential"); value != nil {
		me.Confidential = opt.NewBool(value.(bool))
	}
	if result, ok := decoder.GetOk("data_sources.#"); ok {
		me.DataSources = []*DataSource{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(DataSource)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "data_sources", idx)); err != nil {
				return err
			}
			me.DataSources = append(me.DataSources, entry)
		}
	}
	if value, ok := decoder.GetOk("data_type"); ok {
		me.DataType = DataType(value.(string))
	}
	if value, ok := decoder.GetOk("normalization"); ok {
		me.Normalization = Normalization(value.(string))
	}
	if _, value := decoder.GetChange("enabled"); value != nil {
		me.Enabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("aggregation"); ok {
		me.Aggregation = Aggregation(value.(string))
	}
	return nil
}

func (me *RequestAttribute) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("id", me.ID); err != nil {
		return nil, err
	}
	if err := m.Marshal("name", me.Name); err != nil {
		return nil, err
	}
	if err := m.Marshal("skipPersonalDataMasking", opt.Bool(me.SkipPersonalDataMasking)); err != nil {
		return nil, err
	}
	if err := m.Marshal("confidential", opt.Bool(me.Confidential)); err != nil {
		return nil, err
	}
	if err := m.Marshal("dataSources", me.DataSources); err != nil {
		return nil, err
	}
	if err := m.Marshal("dataType", me.DataType); err != nil {
		return nil, err
	}
	if err := m.Marshal("normalization", me.Normalization); err != nil {
		return nil, err
	}
	if err := m.Marshal("enabled", me.Enabled); err != nil {
		return nil, err
	}
	if err := m.Marshal("aggregation", me.Aggregation); err != nil {
		return nil, err
	}
	if err := m.Marshal("metadata", me.Metadata); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *RequestAttribute) UnmarshalJSON(data []byte) error {
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
	if err := m.Unmarshal("skipPersonalDataMasking", &me.SkipPersonalDataMasking); err != nil {
		return err
	}
	if err := m.Unmarshal("confidential", &me.Confidential); err != nil {
		return err
	}
	if err := m.Unmarshal("dataSources", &me.DataSources); err != nil {
		return err
	}
	if err := m.Unmarshal("dataType", &me.DataType); err != nil {
		return err
	}
	if err := m.Unmarshal("normalization", &me.Normalization); err != nil {
		return err
	}
	if err := m.Unmarshal("enabled", &me.Enabled); err != nil {
		return err
	}
	if err := m.Unmarshal("aggregation", &me.Aggregation); err != nil {
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

// DataType The data type of the request attribute.
type DataType string

// DataTypes offers the known enum values
var DataTypes = struct {
	Double  DataType
	Integer DataType
	String  DataType
}{
	"DOUBLE",
	"INTEGER",
	"STRING",
}

// Normalization String values transformation.
//  If the **dataType** is not `string`, set the `Original` here.
type Normalization string

// Normalizations offers the known enum values
var Normalizations = struct {
	Original    Normalization
	ToLowerCase Normalization
	ToUpperCase Normalization
}{
	"ORIGINAL",
	"TO_LOWER_CASE",
	"TO_UPPER_CASE",
}

// Aggregation Aggregation type for the request values.
type Aggregation string

// Aggregations offers the known enum values
var Aggregations = struct {
	AllDistinctValues   Aggregation
	Average             Aggregation
	CountDistinctValues Aggregation
	CountValues         Aggregation
	First               Aggregation
	Last                Aggregation
	Maximum             Aggregation
	Minimum             Aggregation
	Sum                 Aggregation
}{
	"ALL_DISTINCT_VALUES",
	"AVERAGE",
	"COUNT_DISTINCT_VALUES",
	"COUNT_VALUES",
	"FIRST",
	"LAST",
	"MAXIMUM",
	"MINIMUM",
	"SUM",
}
