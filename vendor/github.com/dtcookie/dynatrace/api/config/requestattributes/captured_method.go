package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// CapturedMethod has no documentation
type CapturedMethod struct {
	ArgumentIndex    *int32                     `json:"argumentIndex,omitempty"`    // The index of the argument to capture. Set `0` to capture the return value, `1` or higher to capture a mehtod argument.   Required if the **capture** is set to `ARGUMENT`.  Not applicable in other cases.
	Capture          Capture                    `json:"capture"`                    // What to capture from the method.
	DeepObjectAccess *string                    `json:"deepObjectAccess,omitempty"` // The getter chain to apply to the captured object. It is required in one of the following cases:  The **capture** is set to `THIS`.    The **capture** is set to `ARGUMENT`, and the argument is not a primitive, a primitive wrapper class, a string, or an array.   Not applicable in other cases.
	Method           *MethodReference           `json:"method"`                     // Configuration of a method to be captured.
	Unknowns         map[string]json.RawMessage `json:"-"`
}

func (me *CapturedMethod) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"argument_index": {
			Type:        hcl.TypeInt,
			Description: "The index of the argument to capture. Set `0` to capture the return value, `1` or higher to capture a mehtod argument.   Required if the **capture** is set to `ARGUMENT`.  Not applicable in other cases",
			Optional:    true,
		},
		"capture": {
			Type:        hcl.TypeString,
			Description: "What to capture from the method",
			Required:    true,
		},
		"deep_object_access": {
			Type:        hcl.TypeString,
			Description: "The getter chain to apply to the captured object. It is required in one of the following cases:  The **capture** is set to `THIS`.    The **capture** is set to `ARGUMENT`, and the argument is not a primitive, a primitive wrapper class, a string, or an array.   Not applicable in other cases",
			Optional:    true,
		},
		"method": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of a method to be captured",
			Elem: &hcl.Resource{
				Schema: new(MethodReference).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *CapturedMethod) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if me.ArgumentIndex != nil {
		result["argument_index"] = int(opt.Int32(me.ArgumentIndex))
	}
	result["capture"] = string(me.Capture)
	if me.DeepObjectAccess != nil {
		result["deep_object_access"] = *me.DeepObjectAccess
	}
	if me.Method != nil {
		if marshalled, err := me.Method.MarshalHCL(); err == nil {
			result["method"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *CapturedMethod) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "argument_index")
		delete(me.Unknowns, "capture")
		delete(me.Unknowns, "deep_object_access")
		delete(me.Unknowns, "method")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("argument_index"); ok {
		me.ArgumentIndex = opt.NewInt32(int32(value.(int)))
	}
	if value, ok := decoder.GetOk("capture"); ok {
		me.Capture = Capture(value.(string))
	}
	if value, ok := decoder.GetOk("deep_object_access"); ok {
		me.DeepObjectAccess = opt.NewString(value.(string))
	}
	if _, ok := decoder.GetOk("method.#"); ok {
		me.Method = new(MethodReference)
		if err := me.Method.UnmarshalHCL(hcl.NewDecoder(decoder, "method", 0)); err != nil {
			return err
		}
	}
	return nil
}

func (me *CapturedMethod) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("argumentIndex", me.ArgumentIndex); err != nil {
		return nil, err
	}
	if err := m.Marshal("capture", me.Capture); err != nil {
		return nil, err
	}
	if err := m.Marshal("deepObjectAccess", me.DeepObjectAccess); err != nil {
		return nil, err
	}
	if err := m.Marshal("method", me.Method); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *CapturedMethod) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("argumentIndex", &me.ArgumentIndex); err != nil {
		return err
	}
	if err := m.Unmarshal("capture", &me.Capture); err != nil {
		return err
	}
	if err := m.Unmarshal("deepObjectAccess", &me.DeepObjectAccess); err != nil {
		return err
	}
	if err := m.Unmarshal("method", &me.Method); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// Capture What to capture from the method.
type Capture string

// Captures offers the known enum values
var Captures = struct {
	Argument        Capture
	ClassName       Capture
	MethodName      Capture
	Occurrences     Capture
	SimpleClassName Capture
	This            Capture
}{
	"ARGUMENT",
	"CLASS_NAME",
	"METHOD_NAME",
	"OCCURRENCES",
	"SIMPLE_CLASS_NAME",
	"THIS",
}
