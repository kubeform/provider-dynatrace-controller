package terraform

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Address has no documentation
type Address string

func (a Address) add(s string) Address {
	if len(a) == 0 {
		return Address(s)
	}
	return Address(string(a) + "." + s)
}

func (a Address) idx(i int) Address {
	s := fmt.Sprintf("%d", i)
	if i < 10 {
		s = "0" + s
	}
	if i < 100 {
		s = "0" + s
	}
	return Address(fmt.Sprintf("%v.%s", a, s))
}

// ResData has no documentation
type ResData interface {
	Set(string, interface{}) error
}

func struggle(value interface{}) bool {
	// slogger.Println("struggle", value)
	switch o := value.(type) {
	case []interface{}:
		for _, elem := range o {
			if struggle(elem) {
				sort.SliceStable(o, func(i, j int) bool {
					a := o[i].(map[string]interface{})
					b := o[j].(map[string]interface{})
					return strings.Compare(a["key"].(string), b["key"].(string)) == -1
				})
			}
		}
		return false
	case map[string]interface{}:
		for _, v := range o {
			struggle(v)
		}
		if _, found := o["key"]; found {
			return true
		}
		return false
	default:
		return false
	}
}

// ToTerraform has no documentation
func ToTerraform(v interface{}, resourceData ResData) error {
	res := &ngResource{values: map[string]interface{}{}}
	if err := toTerraform(v, res, ""); err != nil {
		return err
	}
	struggle(res.values)

	for k, v := range res.values {
		// if k != "id" && k != "metadata" {
		if k != "id" {
			resourceData.Set(k, v)
		}
	}
	return nil
}

func pName(value reflect.Value, field reflect.StructField) string {
	if !startsWithUpper(field.Name) {
		return ""
	}
	kind := field.Type.Kind()
	if kind == reflect.Interface {
		return unCamel(unref(reflect.TypeOf(value.Interface())).Name())
	}
	propertyName := unCamel(field.Name)
	if jsonValue, ok := field.Tag.Lookup("json"); ok {
		if len(jsonValue) > 0 {
			if strings.Contains(jsonValue, ",") {
				parts := strings.Split(jsonValue, ",")
				part0 := strings.TrimSpace(parts[0])
				if len(part0) > 0 {
					propertyName = unCamel(part0)
				}
			} else {
				propertyName = unCamel(jsonValue)
			}
		}
	}
	if propertyName == "-" {
		return ""
	}
	return propertyName
}

// FieldKind has no documentation
type FieldKind string

const (
	// Unknown has no documentation
	Unknown FieldKind = FieldKind("Unknown")
	// Primitive has no documentation
	Primitive = FieldKind("Primitive")
	// StructSlice has no documentation
	StructSlice = FieldKind("StructSlice")
	// PrimitiveSlice has no documentation
	PrimitiveSlice = FieldKind("PrimitiveSlice")
	// InterfaceSlice has no documentation
	InterfaceSlice = FieldKind("InterfaceSlice")
	// Map has no documentation
	Map = FieldKind("Map")
	// Interface has no documentation
	Interface = FieldKind("Interface")
	// Struct has no documentation
	Struct = FieldKind("Struct")
)

func unrefValue(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Ptr:
		return unrefValue(v.Elem())
	default:
		return v
	}
}

func getFieldKind(value reflect.Value, field reflect.StructField) (FieldKind, error) {
	fieldType := unref(field.Type)
	switch fieldType.Kind() {
	case reflect.String, reflect.Bool, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Float32, reflect.Float64:
		return Primitive, nil
	case reflect.Map:
		return Map, nil
	case reflect.Struct:
		return Struct, nil
	case reflect.Interface:
		return Interface, nil
	case reflect.Slice:
		elemType := unref(fieldType.Elem())
		switch elemType.Kind() {
		case reflect.String, reflect.Bool, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Float32, reflect.Float64:
			return PrimitiveSlice, nil
		case reflect.Interface:
			return InterfaceSlice, nil
		case reflect.Struct:
			return StructSlice, nil
		default:
			return Unknown, fmt.Errorf("getFieldKind: recognized slice, but elem kind %v is not supported", elemType.Kind())
		}
	}
	return Unknown, fmt.Errorf("getFieldKind: not supported kind: %v (%v)", fieldType.Kind(), field.Name)
}

func toTerraform(v interface{}, rd NGResourceCollector, address Address) error {
	if v == nil {
		return nil
	}
	rv := reflect.ValueOf(v)
	typ := rv.Type()
	kind := typ.Kind()
	switch kind {
	case reflect.Ptr:
		if !rv.IsZero() && !rv.IsNil() {
			return toTerraform(rv.Elem().Interface(), rd, address)
		}
		return nil
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			field := typ.Field(i)
			vField := rv.Field(i)
			if len(pName(vField, field)) == 0 {
				continue
			}
			var fieldKind FieldKind
			var err error
			if fieldKind, err = getFieldKind(vField, field); err != nil {
				return err
			}
			switch fieldKind {
			case Primitive:
				propertyName := pName(vField, field)
				if err := toTerraform(vField.Interface(), rd, address.add(unCamel(propertyName))); err != nil {
					return err
				}
			case Struct:
				if field.Anonymous {
					if err := toTerraform(vField.Interface(), rd, address); err != nil {
						return err
					}
				} else {
					propertyName := pName(vField, field)
					if err := toTerraform(vField.Interface(), rd, address.add(unCamel(propertyName)).idx(0)); err != nil {
						return err
					}
				}
			case PrimitiveSlice:
				propertyName := pName(vField, field)
				for idx := 0; idx < vField.Len(); idx++ {
					if err := toTerraform(vField.Index(idx).Interface(), rd, address.add(unCamel(propertyName)).idx(idx)); err != nil {
						return err
					}
				}
			case InterfaceSlice:
				cnt := map[string]int{}
				for idx := 0; idx < vField.Len(); idx++ {
					vElem := vField.Index(idx)
					elemType := unref(reflect.TypeOf(vElem.Interface()))
					propertyName := unCamel(elemType.Name())
					pIdx := 0
					found := false
					if pIdx, found = cnt[propertyName]; found {
						pIdx++
						cnt[propertyName] = pIdx
					} else {
						pIdx = 0
						cnt[propertyName] = 0
					}
					if err := toTerraform(vElem.Interface(), rd, address.add(unCamel(propertyName)).idx(pIdx)); err != nil {
						return err
					}
				}
			case StructSlice:
				propertyName := pName(vField, field)
				for idx := 0; idx < vField.Len(); idx++ {
					if err := toTerraform(vField.Index(idx).Interface(), rd, address.add(unCamel(propertyName)).idx(idx)); err != nil {
						return err
					}
				}
			case Map:
				propertyName := pName(vField, field)
				for idx, vMapKey := range vField.MapKeys() {
					vMapEntry := vField.MapIndex(vMapKey)
					if err := toTerraform(vMapKey.Interface(), rd, address.add(unCamel(propertyName)).idx(idx).add("key")); err != nil {
						return err
					}
					if err := toTerraform(vMapEntry.Interface(), rd, address.add(unCamel(propertyName)).idx(idx)); err != nil {
						return err
					}
				}
			case Interface:
				fieldType := unref(reflect.TypeOf(vField.Interface()))
				propertyName := unCamel(fieldType.Name())
				if err := toTerraform(vField.Interface(), rd, address.add(unCamel(propertyName)).idx(0)); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unable to deal with field kind %v (%v.%v)", fieldKind, address, field.Name)
			}
		}
		return nil
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i)
			if err := toTerraform(elem.Interface(), rd, address.idx(i)); err != nil {
				return err
			}
		}
		return nil
	case reflect.String, reflect.Bool, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Float32, reflect.Float64:
		rd.Set(string(address), v)
		return nil
	case reflect.Map:
		for idx, vMapKey := range rv.MapKeys() {
			vMapEntry := rv.MapIndex(vMapKey)
			if err := toTerraform(vMapKey.Interface(), rd, address.add("values").idx(idx).add("key")); err != nil {
				return err
			}
			entryType := unref(reflect.TypeOf(vMapEntry.Interface()))
			switch entryType.Kind() {
			case reflect.Slice:
				if err := toTerraform(vMapEntry.Interface(), rd, address.add("values").idx(idx).add("values")); err != nil {
					return err
				}
			default:
				if err := toTerraform(vMapEntry.Interface(), rd, address.add("values").idx(idx)); err != nil {
					return err
				}
			}
		}
		return nil
	default:
		return fmt.Errorf("objects of kind %v are not supported (%v)", kind, address)
	}
}

type tlogger string

func (tl tlogger) Println(v ...interface{}) {
	if tl == "enabled" {
		log.Println(v...)
	}
}

var logger = tlogger("disabled")
var slogger = tlogger("enabled")

func hide(v interface{}) {}

// ValueType has no documentation
type ValueType string

var valueTypes = ValueType("seed")

func (vt ValueType) of(t schema.ValueType) ValueType {
	switch t {
	case schema.TypeInvalid:
		return TypeInvalid
	case schema.TypeBool:
		return TypeBool
	case schema.TypeInt:
		return TypeInt
	case schema.TypeFloat:
		return TypeFloat
	case schema.TypeString:
		return TypeString
	case schema.TypeList:
		return TypeList
	case schema.TypeMap:
		return TypeMap
	case schema.TypeSet:
		return TypeSet
	default:
		panic(fmt.Sprintf("unsupported value type %v", t))
	}
}

const (
	// TypeInvalid has no documentation
	TypeInvalid ValueType = ValueType("TypeInvalid")
	// TypeBool has no documentation
	TypeBool ValueType = ValueType("TypeBool")
	// TypeInt has no documentation
	TypeInt ValueType = ValueType("TypeInt")
	// TypeFloat has no documentation
	TypeFloat ValueType = ValueType("TypeFloat")
	// TypeString has no documentation
	TypeString ValueType = ValueType("TypeString")
	// TypeList has no documentation
	TypeList ValueType = ValueType("TypeList")
	// TypeMap has no documentation
	TypeMap ValueType = ValueType("TypeMap")
	// TypeSet has no documentation
	TypeSet ValueType = ValueType("TypeSet")
)

// Resource has no documentation
type Resource struct {
	Schema map[string]*Schema `json:"schema"`
}

func (res *Resource) of(other *schema.Resource) *Resource {
	if other == nil {
		return nil
	}
	this := new(Resource)
	this.Schema = map[string]*Schema{}
	for k, v := range other.Schema {
		this.Schema[k] = new(Schema).of(v)
	}
	return this
}

// Schema has no documentation
type Schema struct {
	Type     ValueType   `json:"type"`
	Optional bool        `json:"optional,omitempty"`
	Required bool        `json:"required,omitempty"`
	Elem     interface{} `json:"elem,omitempty"`
	MaxItems int         `json:"maxItems,omitempty"`
	MinItems int         `json:"minItems,omitempty"`
}

func (sch *Schema) of(other *schema.Schema) *Schema {
	if other == nil {
		return nil
	}
	this := new(Schema)
	this.Type = valueTypes.of(other.Type)
	this.Optional = other.Optional
	this.Required = other.Required
	this.MaxItems = other.MaxItems
	this.MinItems = other.MinItems
	if other.Elem != nil {
		switch typedElem := other.Elem.(type) {
		case *schema.Schema:
			this.Elem = new(Schema).of(typedElem)
		case *schema.Resource:
			this.Elem = new(Resource).of(typedElem)
		default:
			panic(fmt.Sprintf("unsupported elem type %T", typedElem))
		}
	}
	return this
}

// ResourceToJSON has no documentation
func ResourceToJSON(res *schema.Resource) string {
	resource := new(Resource).of(res)
	var bytes []byte
	var err error
	if bytes, err = json.MarshalIndent(resource, "", "  "); err != nil {
		return err.Error()
	}
	return string(bytes)
}
