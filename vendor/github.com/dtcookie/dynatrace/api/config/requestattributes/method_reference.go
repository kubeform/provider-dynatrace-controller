package requestattributes

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// MethodReference Configuration of a method to be captured.
type MethodReference struct {
	ReturnType      string                     `json:"returnType"`                // The return type.
	Visibility      Visibility                 `json:"visibility"`                // The visibility of the method to capture.
	ArgumentTypes   []string                   `json:"argumentTypes"`             // The list of argument types.
	ClassName       *string                    `json:"className,omitempty"`       // The class name where the method to capture resides.   Either this or the **fileName** must be set.
	FileName        *string                    `json:"fileName,omitempty"`        // The file name where the method to capture resides.   Either this or **className** must be set.
	FileNameMatcher *FileNameMatcher           `json:"fileNameMatcher,omitempty"` // The operator of the comparison.   If not set, `EQUALS` is used.
	MethodName      string                     `json:"methodName"`                // The name of the method to capture.
	Modifiers       []Modifier                 `json:"modifiers"`                 // The modifiers of the method to capture.
	Unknowns        map[string]json.RawMessage `json:"-"`
}

func (me *MethodReference) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"return_type": {
			Type:        hcl.TypeString,
			Description: "The return type",
			Required:    true,
		},
		"visibility": {
			Type:        hcl.TypeString,
			Description: "The visibility of the method to capture",
			Required:    true,
		},
		"argument_types": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			Description: "Configuration of a method to be captured",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"class_name": {
			Type:        hcl.TypeString,
			Description: "The class name where the method to capture resides.   Either this or the **fileName** must be set",
			Optional:    true,
		},
		"file_name": {
			Type:        hcl.TypeString,
			Description: "The file name where the method to capture resides.   Either this or **className** must be set",
			Optional:    true,
		},
		"file_name_matcher": {
			Type:        hcl.TypeString,
			Description: "The operator of the comparison. If not set, `EQUALS` is used",
			Optional:    true,
		},
		"method_name": {
			Type:        hcl.TypeString,
			Description: "The name of the method to capture",
			Required:    true,
		},
		"modifiers": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The modifiers of the method to capture",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *MethodReference) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["return_type"] = me.ReturnType
	result["visibility"] = Visibility(me.Visibility)
	if len(me.ArgumentTypes) > 0 {
		result["argument_types"] = me.ArgumentTypes
	}
	if me.ClassName != nil {
		result["class_name"] = *me.ClassName
	}
	if me.FileName != nil {
		result["file_name"] = *me.FileName
	}
	if me.FileNameMatcher != nil {
		result["file_name_matcher"] = string(*me.FileNameMatcher)
	}
	result["method_name"] = me.MethodName
	if len(me.Modifiers) > 0 {
		mods := []string{}
		for _, mod := range me.Modifiers {
			mods = append(mods, string(mod))
		}
		result["modifiers"] = mods
	}
	return result, nil
}

func (me *MethodReference) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "return_type")
		delete(me.Unknowns, "visibility")
		delete(me.Unknowns, "argument_types")
		delete(me.Unknowns, "class_name")
		delete(me.Unknowns, "file_name")
		delete(me.Unknowns, "file_name_matcher")
		delete(me.Unknowns, "method_name")
		delete(me.Unknowns, "modifiers")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("return_type"); ok {
		me.ReturnType = value.(string)
	}
	if value, ok := decoder.GetOk("visibility"); ok {
		me.Visibility = Visibility(value.(string))
	}
	if value, ok := decoder.GetOk("argument_types"); ok {
		me.ArgumentTypes = []string{}
		for _, e := range value.([]interface{}) {
			me.ArgumentTypes = append(me.ArgumentTypes, e.(string))
		}
	}
	if value, ok := decoder.GetOk("class_name"); ok {
		me.ClassName = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("file_name"); ok {
		me.FileName = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("file_name_matcher"); ok {
		me.FileNameMatcher = FileNameMatcher(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("method_name"); ok {
		me.MethodName = value.(string)
	}
	if value, ok := decoder.GetOk("modifiers"); ok {
		me.Modifiers = []Modifier{}
		for _, e := range value.([]interface{}) {
			me.Modifiers = append(me.Modifiers, Modifier(e.(string)))
		}
	}
	return nil
}

func (me *MethodReference) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("returnType", me.ReturnType); err != nil {
		return nil, err
	}
	if err := m.Marshal("visibility", me.Visibility); err != nil {
		return nil, err
	}
	if err := m.Marshal("argumentTypes", me.ArgumentTypes); err != nil {
		return nil, err
	}
	if err := m.Marshal("className", me.ClassName); err != nil {
		return nil, err
	}
	if err := m.Marshal("fileName", me.FileName); err != nil {
		return nil, err
	}
	if err := m.Marshal("fileNameMatcher", me.FileNameMatcher); err != nil {
		return nil, err
	}
	if err := m.Marshal("methodName", me.MethodName); err != nil {
		return nil, err
	}
	if err := m.Marshal("modifiers", me.Modifiers); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *MethodReference) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("returnType", &me.ReturnType); err != nil {
		return err
	}
	if err := m.Unmarshal("visibility", &me.Visibility); err != nil {
		return err
	}
	if err := m.Unmarshal("argumentTypes", &me.ArgumentTypes); err != nil {
		return err
	}
	if err := m.Unmarshal("className", &me.ClassName); err != nil {
		return err
	}
	if err := m.Unmarshal("fileName", &me.FileName); err != nil {
		return err
	}
	if err := m.Unmarshal("fileNameMatcher", &me.FileNameMatcher); err != nil {
		return err
	}
	if err := m.Unmarshal("methodName", &me.MethodName); err != nil {
		return err
	}
	if err := m.Unmarshal("modifiers", &me.Modifiers); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// Visibility The visibility of the method to capture.
type Visibility string

// Visibilitys offers the known enum values
var Visibilitys = struct {
	Internal         Visibility
	PackageProtected Visibility
	Private          Visibility
	Protected        Visibility
	Public           Visibility
}{
	"INTERNAL",
	"PACKAGE_PROTECTED",
	"PRIVATE",
	"PROTECTED",
	"PUBLIC",
}

// FileNameMatcher The operator of the comparison.
//  If not set, `EQUALS` is used.
type FileNameMatcher string

func (me FileNameMatcher) Ref() *FileNameMatcher {
	return &me
}

// FileNameMatchers offers the known enum values
var FileNameMatchers = struct {
	EndsWith   FileNameMatcher
	Equals     FileNameMatcher
	StartsWith FileNameMatcher
}{
	"ENDS_WITH",
	"EQUALS",
	"STARTS_WITH",
}

// Modifier has no documentation
type Modifier string

// Modifiers offers the known enum values
var Modifiers = struct {
	Abstract Modifier
	Extern   Modifier
	Final    Modifier
	Native   Modifier
	Static   Modifier
}{
	"ABSTRACT",
	"EXTERN",
	"FINAL",
	"NATIVE",
	"STATIC",
}
