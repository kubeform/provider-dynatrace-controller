package dimensions

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

type Dimensions []Dimension

func (me Dimensions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"entity": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A filter for the metrics entity dimensions",
			Elem:        &hcl.Resource{Schema: new(Entity).Schema()},
		},
		"string": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A filter for the metrics string dimensions",
			Elem:        &hcl.Resource{Schema: new(String).Schema()},
		},
		"dimension": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A generic definition for a filter",
			Elem:        &hcl.Resource{Schema: new(BaseDimension).Schema()},
		},
	}
}

func (me Dimensions) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	Entitys := []interface{}{}
	Strings := []interface{}{}
	baseDimensions := []map[string]interface{}{}
	for _, dimension := range me {
		switch dim := dimension.(type) {
		case *Entity:
			if marshalled, err := dim.MarshalHCL(hcl.NewDecoder(decoder, "entity", len(Entitys))); err == nil {
				Entitys = append(Entitys, marshalled)
			} else {
				return nil, err
			}
		case *String:
			if marshalled, err := dim.MarshalHCL(hcl.NewDecoder(decoder, "string", len(Strings))); err == nil {
				Strings = append(Strings, marshalled)
			} else {
				return nil, err
			}
		case *BaseDimension:
			if marshalled, err := dim.MarshalHCL(hcl.NewDecoder(decoder, "dimension", len(baseDimensions))); err == nil {
				baseDimensions = append(baseDimensions, marshalled)
			} else {
				return nil, err
			}
		default:
		}
	}
	if len(Entitys) > 0 {
		result["entity"] = Entitys
	}
	if len(Strings) > 0 {
		result["string"] = Strings
	}
	if len(baseDimensions) > 0 {
		result["dimension"] = baseDimensions
	}
	return result, nil
}

func (me *Dimensions) UnmarshalHCL(decoder hcl.Decoder) error {
	nme := Dimensions{}
	if result, ok := decoder.GetOk("entity.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Entity)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "entity", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("string.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(String)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "string", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("dimension.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(BaseDimension)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "dimension", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	*me = nme
	return nil
}

func (me *Dimensions) UnmarshalJSON(data []byte) error {
	dims := Dimensions{}
	rawMessages := []json.RawMessage{}
	if err := json.Unmarshal(data, &rawMessages); err != nil {
		return err
	}
	for _, rawMessage := range rawMessages {
		properties := map[string]json.RawMessage{}
		if err := json.Unmarshal(rawMessage, &properties); err != nil {
			return err
		}
		if rawFilterType, found := properties["filterType"]; found {
			var sFilterType string
			if err := json.Unmarshal(rawFilterType, &sFilterType); err != nil {
				return err
			}
			switch sFilterType {
			case string(FilterTypes.Entity):
				cfg := new(Entity)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.String):
				cfg := new(String)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			default:
				cfg := new(BaseDimension)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			}
		}
		*me = dims
	}
	return nil
}
