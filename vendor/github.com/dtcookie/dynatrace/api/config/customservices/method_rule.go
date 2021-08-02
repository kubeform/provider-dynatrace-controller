package customservices

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// MethodRule TODO: documentation
type MethodRule struct {
	ID            *string                    `json:"id,omitempty"`  // The ID of the method rule
	MethodName    string                     `json:"methodName"`    // The method to instrument
	ArgumentTypes []string                   `json:"argumentTypes"` // Fully qualified types of argument the method expects
	ReturnType    *string                    `json:"returnType"`    // Fully qualified type the method returns
	Unknowns      map[string]json.RawMessage `json:"-"`
}

func (me *MethodRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Description: "The ID of the method rule",
			Computed:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "The method to instrument",
			Required:    true,
		},
		"returns": {
			Type:        hcl.TypeString,
			Description: "Fully qualified type the method returns",
			Optional:    true,
		},
		"arguments": {
			Type:        hcl.TypeList,
			Description: "Fully qualified types of argument the method expects",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *MethodRule) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.ID != nil {
		result["id"] = opt.String(me.ID)
	}
	result["name"] = me.MethodName
	if me.ReturnType != nil {
		result["returns"] = *me.ReturnType
	}
	if len(me.ArgumentTypes) > 0 {
		result["arguments"] = me.ArgumentTypes
	}
	return result, nil
}

func (me *MethodRule) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "id")
		delete(me.Unknowns, "name")
		delete(me.Unknowns, "returns")
		delete(me.Unknowns, "arguments")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	adapter := hcl.Adapt(decoder)
	me.ID = adapter.GetString("id")
	me.MethodName = opt.String(adapter.GetString("name"))
	if value, ok := decoder.GetOk("arguments"); ok {
		me.ArgumentTypes = []string{}
		for _, v := range value.([]interface{}) {
			me.ArgumentTypes = append(me.ArgumentTypes, v.(string))
		}
	}
	me.ReturnType = adapter.GetString("returns")
	return nil
}

func (me *MethodRule) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("id", me.ID); err != nil {
		return nil, err
	}
	if err := m.Marshal("methodName", me.MethodName); err != nil {
		return nil, err
	}
	if err := m.Marshal("argumentTypes", me.ArgumentTypes); err != nil {
		return nil, err
	}
	if err := m.Marshal("returnType", me.ReturnType); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *MethodRule) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("id", &me.ID); err != nil {
		return err
	}
	if err := m.Unmarshal("methodName", &me.MethodName); err != nil {
		return err
	}
	if err := m.Unmarshal("argumentTypes", &me.ArgumentTypes); err != nil {
		return err
	}
	if err := m.Unmarshal("returnType", &me.ReturnType); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}
