package hcl

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
)

func sk(s string) string {
	if s == "name" {
		return "00" + s
	}
	if s == "description" {
		return "10" + s
	}
	if s == "type" {
		return "20" + s
	}
	if s == "enabled" {
		return "30" + s
	}
	return "90" + s
}

type exportEntry interface {
	Write(w io.Writer, indent string) error
	IsOptional() bool
	IsDefault() bool
	IsLessThan(other exportEntry) bool
}

type exportEntries []exportEntry

func (e exportEntries) Less(i, j int) bool {
	return e[i].IsLessThan(e[j])
}

func resOpt0(key string, bc string, sch *Schema) bool {
	if sch == nil {
		return false
	}
	switch sch.Type {
	case TypeBool:
		return sch.Optional
	case TypeInt:
		return sch.Optional
	case TypeFloat:
		return sch.Optional
	case TypeString:
		return sch.Optional
	case TypeList:
		switch v := sch.Elem.(type) {
		case *Resource:
			return resOpt(bc, v.Schema)
		// case *Schema:
		// 	return resOpt(bc, v)
		default:
			return sch.Optional
		}

	case TypeMap:
		return false
	case TypeSet:
		return false
	default:
		return false

	}
}

func resOpt(bc string, sch map[string]*Schema) bool {
	bc = strings.TrimPrefix(bc, ".")
	if strings.Contains(bc, ".") {
		idx := strings.Index(bc, ".")
		return resOpt0(bc[:idx], bc[idx+1:], sch[bc[:idx]])
	}
	return resOpt0(bc, "", sch[bc])
}

func (e *exportEntries) eval(key string, value interface{}, breadCrumbs string, schema map[string]*Schema) {
	if value == nil {
		return
	}
	switch v := value.(type) {
	case string, bool, int, int32, int64, int8, int16, uint, uint32, uint64, uint8, uint16, float32, float64:
		entry := &primitiveEntry{Key: key, Value: value, BreadCrumbs: breadCrumbs, Optional: resOpt(breadCrumbs, schema)}
		*e = append(*e, entry)
	case *string, *bool, *int, *int32, *int64, *int8, *int16, *uint, *uint32, *uint64, *uint8, *uint16, *float32, *float64:
		if v == nil {
			return
		}
		entry := &primitiveEntry{Key: key, Value: v, BreadCrumbs: breadCrumbs, Optional: resOpt(breadCrumbs, schema)}
		*e = append(*e, entry)
	case []interface{}:
		if len(v) == 0 {
			return
		}
		switch typedElem := v[0].(type) {
		case map[string]interface{}:
			for _, elem := range v {
				entry := &resourceEntry{Key: key, Entries: exportEntries{}}
				entry.Entries.handle(elem.(map[string]interface{}), breadCrumbs, schema)
				*e = append(*e, entry)
			}
		case string, bool, int, int32, int64, int8, int16, uint, uint32, uint64, uint8, uint16, float32, float64:
			entry := &primitiveEntry{Key: key, Value: value}
			*e = append(*e, entry)
		default:
			panic(fmt.Sprintf("unsupported elem type %T", typedElem))
		}
	case []string:
		if len(v) == 0 {
			return
		}
		entry := &primitiveEntry{Key: key, Value: value, Optional: resOpt(breadCrumbs, schema)}
		*e = append(*e, entry)
	case StringSet:
		if len(v) == 0 {
			return
		}
		entry := &primitiveEntry{Key: key, Value: value, Optional: resOpt(breadCrumbs, schema)}
		*e = append(*e, entry)
	case []float64:
		if len(v) == 0 {
			return
		}
		entry := &primitiveEntry{Key: key, Value: value, Optional: resOpt(breadCrumbs, schema)}
		*e = append(*e, entry)
	case map[string]interface{}:
		if len(v) == 0 {
			return
		}
		entry := &resourceEntry{Key: key, Entries: exportEntries{}}
		for xk, xv := range v {
			entry.Entries = append(entry.Entries, &primitiveEntry{Key: xk, Value: xv, Optional: resOpt(breadCrumbs, schema)})
		}
		*e = append(*e, entry)
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.String:
			e.eval(key, fmt.Sprintf("%v", v), breadCrumbs, schema)
		default:
			panic(fmt.Sprintf(">>>>> [%v] type %T not supported yet\n", key, v))
		}

	}
}

func (e *exportEntries) handle(m map[string]interface{}, breadCrumbs string, schema map[string]*Schema) {
	for k, v := range m {
		e.eval(k, v, breadCrumbs+"."+k, schema)
	}
}

type Schemer interface {
	Schema() map[string]*Schema
}

func ExportOpt(marshaler Marshaler, w io.Writer) error {
	var m map[string]interface{}
	var err error
	if m, err = marshaler.MarshalHCL(); err != nil {
		return err
	}
	var schema map[string]*Schema
	if schemer, ok := marshaler.(Schemer); ok {
		schema = schemer.Schema()
	}
	// if schema != nil {
	// 	data, _ := json.MarshalIndent(schema, "", "  ")
	// 	fmt.Println(string(data))
	// }
	return export(m, w, schema)
}

func Export(marshaler Marshaler, w io.Writer) error {
	var m map[string]interface{}
	var err error
	if m, err = marshaler.MarshalHCL(); err != nil {
		return err
	}
	var schema map[string]*Schema
	return export(m, w, schema)
}

func ExtExport(marshaler ExtMarshaler, w io.Writer) error {
	var m map[string]interface{}
	var err error
	if m, err = marshaler.MarshalHCL(&voidDecoder{}); err != nil {
		return err
	}
	return export(m, w, map[string]*Schema{})
}

func export(m map[string]interface{}, w io.Writer, schema map[string]*Schema) error {
	var err error
	ents := exportEntries{}
	ents.handle(m, "", schema)
	sort.SliceStable(ents, ents.Less)
	for _, entry := range ents {
		if !(entry.IsOptional() && entry.IsDefault()) {
			if err := entry.Write(w, "  "); err != nil {
				return err
			}
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		} else {
			if err := entry.Write(w, "  # "); err != nil {
				return err
			}
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}
	}
	return err
}

type primitiveEntry struct {
	Indent      string
	Key         string
	Optional    bool
	BreadCrumbs string
	Value       interface{}
}

func jsonenc(v interface{}, indent string) string {
	switch rv := v.(type) {
	case string:
		if strings.Contains(rv, "\n") {
			return "<<-EOT\n" + indent + "  " + strings.ReplaceAll(rv, "\n", "\n"+indent+"  ") + "\n" + indent + "EOT"
		}
	case *string:
		erv := *rv
		if strings.Contains(erv, "\n") {
			return "<<-EOT\n" + indent + "  " + strings.ReplaceAll(erv, "\n", "\n"+indent+"  ") + "\n" + indent + "EOT"
		}
	default:
	}
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func (pe *primitiveEntry) Write(w io.Writer, indent string) error {
	_, err := w.Write([]byte(fmt.Sprintf("%s%v = %v ", indent, pe.Key, jsonenc(pe.Value, indent))))
	return err
}

func (pe *primitiveEntry) IsOptional() bool {
	return pe.Optional
}

func (pe *primitiveEntry) IsDefault() bool {
	switch rv := pe.Value.(type) {
	case bool:
		return !rv
	case string:
		return len(rv) == 0
	default:
		return false
	}
}

func (pe *primitiveEntry) IsLessThan(other exportEntry) bool {
	switch ro := other.(type) {
	case *primitiveEntry:
		return strings.Compare(sk(pe.Key), sk(ro.Key)) < 0
	case *resourceEntry:
		return true
	}
	return false
}

type resourceEntry struct {
	Indent      string
	Key         string
	BreadCrumbs string
	Optional    bool
	Entries     exportEntries
}

func (pe *resourceEntry) IsOptional() bool {
	return false
}

func (pe *resourceEntry) IsDefault() bool {
	return false
}

func (re *resourceEntry) IsLessThan(other exportEntry) bool {
	switch ro := other.(type) {
	case *primitiveEntry:
		return false
	case *resourceEntry:
		return strings.Compare(re.Key, ro.Key) < 0
	}
	return false
}
func (re *resourceEntry) Write(w io.Writer, indent string) error {
	s := fmt.Sprintf("%s%v {\n", indent, re.Key)
	if _, err := w.Write([]byte(s)); err != nil {
		return err
	}
	sort.SliceStable(re.Entries, re.Entries.Less)
	for _, entry := range re.Entries {
		if !(entry.IsOptional() && entry.IsDefault()) {
			if err := entry.Write(w, indent+"  "); err != nil {
				return err
			}
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		} else {
			if err := entry.Write(w, indent+"  # "); err != nil {
				return err
			}
			if _, err := w.Write([]byte("\n")); err != nil {
				return err
			}
		}
	}
	if _, err := w.Write([]byte(indent + "}")); err != nil {
		return err
	}
	return nil
}
