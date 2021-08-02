package customservices

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// DetectionRule the defining rules for a CustomService
type DetectionRule struct {
	ID               *string                    `json:"id,omitempty"`              // The ID of the detection rule
	Enabled          bool                       `json:"enabled"`                   // Rule enabled/disabled
	FileName         *string                    `json:"fileName,omitempty"`        // The PHP file containing the class or methods to instrument. Required for PHP custom service. Not applicable to Java and .NET
	FileNameMatcher  *FileNameMatcher           `json:"fileNameMatcher,omitempty"` // Matcher applying to the file name. Default value is `ENDS_WITH` (if applicable)
	ClassName        *string                    `json:"className,omitempty"`       // The fully qualified class or interface to instrument. Required for Java and .NET custom services. Not applicable to PHP
	ClassNameMatcher *ClassNameMatcher          `json:"matcher,omitempty"`         // Matcher applying to the class name. `STARTS_WITH` can only be used if there is at least one annotation defined. Default value is `EQUALS`
	MethodRules      []*MethodRule              `json:"methodRules"`               // List of methods to instrument
	Annotations      []string                   `json:"annotations"`               // Additional annotations filter of the rule. Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented. nNot applicable to PHP
	Unknowns         map[string]json.RawMessage `json:"-"`
}

func (me *DetectionRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Description: "The ID of the detection rule",
			Computed:    true,
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "Rule enabled/disabled",
			Required:    true,
		},
		"file": {
			Type:        hcl.TypeList,
			Description: "The PHP file containing the class or methods to instrument. Required for PHP custom service. Not applicable to Java and .NET",
			Optional:    true,
			MaxItems:    1,
			Elem: &hcl.Resource{
				Schema: new(FileSection).Schema(),
			},
		},
		"class": {
			Type:        hcl.TypeList,
			Description: "The fully qualified class or interface to instrument (or a substring if matching to a string). Required for Java and .NET custom services. Not applicable to PHP",
			Optional:    true,
			MaxItems:    1,
			Elem: &hcl.Resource{
				Schema: new(ClassSection).Schema(),
			},
		},
		"annotations": {
			Type:        hcl.TypeSet,
			Description: "Additional annotations filter of the rule. Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented. Not applicable to PHP",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"method": {
			Type:        hcl.TypeList,
			Description: "methods to instrument",
			Required:    true,
			Elem: &hcl.Resource{
				Schema: new(MethodRule).Schema(),
			},
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *DetectionRule) MarshalHCL() (map[string]interface{}, error) {
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
	result["enabled"] = me.Enabled
	if me.FileName != nil || me.FileNameMatcher != nil {
		fileSection := &FileSection{
			Name:  me.FileName,
			Match: me.FileNameMatcher,
		}
		if marshalled, err := fileSection.MarshalHCL(); err == nil {
			result["file"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.ClassName != nil || me.ClassNameMatcher != nil {
		classSection := &ClassSection{
			Name:  me.ClassName,
			Match: me.ClassNameMatcher,
		}
		if marshalled, err := classSection.MarshalHCL(); err == nil {
			result["class"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.MethodRules) > 0 {
		entries := []interface{}{}
		for _, entry := range me.MethodRules {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["method"] = entries
	}
	if len(me.Annotations) > 0 {
		result["annotations"] = me.Annotations
	}
	return result, nil
}

func (me *DetectionRule) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "id")
		delete(me.Unknowns, "enabled")
		delete(me.Unknowns, "file")
		delete(me.Unknowns, "class")
		delete(me.Unknowns, "annotations")
		delete(me.Unknowns, "method")
		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	adapter := hcl.Adapt(decoder)
	me.ID = adapter.GetString("id")
	me.Enabled = opt.Bool(adapter.GetBool("enabled"))
	if _, ok := decoder.GetOk("file.#"); ok {
		fileSection := new(FileSection)
		if err := fileSection.UnmarshalHCL(hcl.NewDecoder(decoder, "file", 0)); err != nil {
			return err
		}
		me.FileName = fileSection.Name
		me.FileNameMatcher = fileSection.Match
	}
	if _, ok := decoder.GetOk("class.#"); ok {
		classSection := new(ClassSection)
		if err := classSection.UnmarshalHCL(hcl.NewDecoder(decoder, "class", 0)); err != nil {
			return err
		}
		me.ClassName = classSection.Name
		me.ClassNameMatcher = classSection.Match
	}
	if result, ok := decoder.GetOk("method.#"); ok {
		me.MethodRules = []*MethodRule{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(MethodRule)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "method", idx)); err != nil {
				return err
			}
			me.MethodRules = append(me.MethodRules, entry)
		}
	}
	me.Annotations = decoder.GetStringSet("annotations")
	return nil
}

func (me *DetectionRule) MarshalJSON() ([]byte, error) {
	m := xjson.NewProperties(me.Unknowns)
	if err := m.Marshal("id", me.ID); err != nil {
		return nil, err
	}
	if err := m.Marshal("enabled", me.Enabled); err != nil {
		return nil, err
	}
	if err := m.Marshal("fileName", me.FileName); err != nil {
		return nil, err
	}
	if err := m.Marshal("fileNameMatcher", me.FileNameMatcher); err != nil {
		return nil, err
	}
	if err := m.Marshal("className", me.ClassName); err != nil {
		return nil, err
	}
	if err := m.Marshal("matcher", me.ClassNameMatcher); err != nil {
		return nil, err
	}
	if err := m.Marshal("methodRules", me.MethodRules); err != nil {
		return nil, err
	}
	if err := m.Marshal("annotations", me.Annotations); err != nil {
		return nil, err
	}
	return json.Marshal(m)
}

func (me *DetectionRule) UnmarshalJSON(data []byte) error {
	m := xjson.Properties{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if err := m.Unmarshal("id", &me.ID); err != nil {
		return err
	}
	if err := m.Unmarshal("enabled", &me.Enabled); err != nil {
		return err
	}
	if err := m.Unmarshal("fileName", &me.FileName); err != nil {
		return err
	}
	if err := m.Unmarshal("fileNameMatcher", &me.FileNameMatcher); err != nil {
		return err
	}
	if err := m.Unmarshal("className", &me.ClassName); err != nil {
		return err
	}
	if err := m.Unmarshal("matcher", &me.ClassNameMatcher); err != nil {
		return err
	}
	if err := m.Unmarshal("methodRules", &me.MethodRules); err != nil {
		return err
	}
	if err := m.Unmarshal("annotations", &me.Annotations); err != nil {
		return err
	}

	if len(m) > 0 {
		me.Unknowns = m
	}
	return nil
}

// FileNameMatcher has no documentation
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

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *FileNameMatcher) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = FileNameMatcher(name)
	return nil
}

// ClassNameMatcher has no documentation
type ClassNameMatcher string

func (me ClassNameMatcher) Ref() *ClassNameMatcher {
	return &me
}

// ClassNameMatchers offers the known enum values
var ClassNameMatchers = struct {
	EndsWith   ClassNameMatcher
	Equals     ClassNameMatcher
	StartsWith ClassNameMatcher
}{
	"ENDS_WITH",
	"EQUALS",
	"STARTS_WITH",
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *ClassNameMatcher) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = ClassNameMatcher(name)
	return nil
}
