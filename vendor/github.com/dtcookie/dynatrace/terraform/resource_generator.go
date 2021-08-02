package terraform

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceFor has no documentation
func ResourceFor(v interface{}) *schema.Resource {
	if v == nil {
		panic("cannot generate a resource for `nil`")
	}
	resource := structFor("", v, reflect.TypeOf(v)).Schema.Elem.(*schema.Resource)
	delete(resource.Schema, "id")
	return resource
}

func propName(field reflect.StructField) string {
	if !startsWithUpper(field.Name) {
		return ""
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

func revisedPropName(field reflect.StructField, value reflect.Value) string {
	if !startsWithUpper(field.Name) {
		return ""
	}
	kind := field.Type.Kind()
	if kind == reflect.Interface {
		return unCamel(unref(reflect.TypeOf(value.Interface())).Name())
		// log.Println(fmt.Sprintf("[%s] kind: %v, type: %v, typeName: %v", field.Name, kind, typ.String(), typ.Name()))
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

func optional(field reflect.StructField) bool {
	if jsonValue, ok := field.Tag.Lookup("json"); ok {
		if len(jsonValue) > 0 {
			return strings.Contains(jsonValue, "omitempty")
		}
	}
	return false
}

func base(iface reflect.Type) reflect.Type {
	if method, ok := iface.MethodByName("Initialize"); ok {
		methodType := method.Type
		if methodType.NumOut() != 0 {
			return nil
		}
		if methodType.NumIn() != 1 {
			return nil
		}
		paramType := methodType.In(0)
		if !paramType.Implements(iface) {
			return nil
		}
		return paramType
	}
	return nil
}

func implementors(baseType reflect.Type, iface reflect.Type) []reflect.Type {
	if method, ok := baseType.MethodByName("Implementors"); ok {
		methodType := method.Type
		if methodType.NumIn() != 1 {
			// fmt.Printf("   numIn != , actual: %v (%v)\n", methodType.NumIn(), methodType.In(0))
			return nil
		}
		if methodType.NumOut() != 1 {
			return nil
		}
		returnType := methodType.Out(0)
		if returnType != reflect.SliceOf(iface) {
			return nil
		}
		baseInstance := reflect.New(unref(baseType))
		vMethod := baseInstance.MethodByName("Implementors")
		vImplementors := vMethod.Call([]reflect.Value{})[0]
		implementors := []reflect.Type{}
		for i := 0; i < vImplementors.Len(); i++ {
			implementors = append(implementors, vImplementors.Index(i).Elem().Type())
		}
		return implementors
	}
	return nil
}

func collectProperties(hint string, schemaMap map[string]*schema.Schema, field reflect.StructField) {
	// fmt.Println("collectProperties", hint)
	if field.Anonymous {
		if unref(field.Type).Kind() == reflect.Struct {
			anonResource := structFor(field.Name, nil, field.Type)
			for v, k := range anonResource.Schema.Elem.(*schema.Resource).Schema {
				schemaMap[v] = k
			}
		}
		return
	}
	if !startsWithUpper(field.Name) {
		return
	}
	propertyName := propName(field)
	if (len(propertyName) > 0) && (propertyName != "-") {
		sch := structFor(fmt.Sprintf("%v.%v", hint, propertyName), nil, field.Type)
		// sch.Optional = optional(field)
		// sch.Required = !sch.Optional
		if sch.Schemata != nil {
			for k, v := range sch.Schemata {
				schemaMap[k] = v
				v.Optional = true
				v.Required = false

			}
		} else {
			sch.Schema.Optional = true
			sch.Schema.Required = false
			schemaMap[propertyName] = sch.Schema
		}
	}

}

// AnnotatedSchema has no documentation
type AnnotatedSchema struct {
	Schema   *schema.Schema
	Schemata map[string]*schema.Schema
}

func structFor(hint string, v interface{}, t reflect.Type) *AnnotatedSchema {
	// log.Printf("structFor(hint: %v, v: %v, t: %v)\n", hint, v, t)
	switch t.Kind() {
	case reflect.Ptr:
		if v == nil {
			v = reflect.New(t.Elem()).Interface()
		}
		rv := reflect.ValueOf(v)
		return structFor(hint, rv.Elem().Interface(), rv.Elem().Type())
	case reflect.Struct:
		resource := new(schema.Resource)
		resource.Schema = map[string]*schema.Schema{}
		for i := 0; i < t.NumField(); i++ {
			collectProperties(hint, resource.Schema, t.Field(i))
		}
		return &AnnotatedSchema{Schema: &schema.Schema{Type: schema.TypeList, MaxItems: 1, Elem: resource}}
	case reflect.Slice:
		elemSchema := structFor(hint, nil, t.Elem())
		if elemSchema.Schemata != nil {
			schemata := map[string]*schema.Schema{}
			for k, v := range elemSchema.Schemata {
				if (v.Type == schema.TypeList) && (v.MaxItems == 1) {
					switch v.Elem.(type) {
					case *schema.Resource:
						v.MaxItems = 0
						schemata[k] = v
					default:
					}
				}
				schemata[k] = &schema.Schema{Type: schema.TypeList, Elem: v.Elem}
			}
			return &AnnotatedSchema{Schemata: schemata}
		}

		if (elemSchema.Schema.Type == schema.TypeList) && (elemSchema.Schema.MaxItems == 1) {
			switch elemSchema.Schema.Elem.(type) {
			case *schema.Resource:
				elemSchema.Schema.MaxItems = 0
				return elemSchema
			default:
			}
		}
		return &AnnotatedSchema{Schema: &schema.Schema{Type: schema.TypeList, Elem: elemSchema.Schema}}
	case reflect.String:
		return &AnnotatedSchema{Schema: &schema.Schema{Type: schema.TypeString}}
	case reflect.Int, reflect.Uint, reflect.Int8, reflect.Int16, reflect.Int64, reflect.Int32, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &AnnotatedSchema{Schema: &schema.Schema{Type: schema.TypeInt}}
	case reflect.Float32, reflect.Float64:
		return &AnnotatedSchema{Schema: &schema.Schema{Type: schema.TypeFloat}}
	case reflect.Bool:
		return &AnnotatedSchema{Schema: &schema.Schema{Type: schema.TypeBool}}
	case reflect.Interface:
		baseType := base(t)
		if baseType == nil {
			panic(fmt.Errorf("no base type found (type: %v, kind: %v)", t, t.Kind()))
		}
		implementors := implementors(baseType, t)
		if implementors == nil {
			panic(fmt.Errorf("no implementors found (type: %v, kind: %v, base: %v)", t, t.Kind(), baseType))
		}
		schemata := map[string]*schema.Schema{}
		for _, implementor := range implementors {
			resource := &schema.Resource{Schema: map[string]*schema.Schema{}}
			implementorStruct := structFor(fmt.Sprintf("%v[%v]", hint, implementor), nil, implementor)

			for k, v := range implementorStruct.Schema.Elem.(*schema.Resource).Schema {
				resource.Schema[k] = v
				v.Required = false
				v.Optional = true
				schemata[unCamel(implementor.Elem().Name())] = &schema.Schema{
					Type: schema.TypeList,
					// MaxItems: 1,
					Elem: resource,
				}
			}
			baseStruct := structFor(fmt.Sprintf("%v[%v]", hint, baseType), nil, baseType)
			for k, v := range baseStruct.Schema.Elem.(*schema.Resource).Schema {
				resource.Schema[k] = v
				v.Required = false
				v.Optional = true
				schemata[unCamel(baseType.Elem().Name())] = &schema.Schema{
					Type: schema.TypeList,
					// MaxItems: 1,
					Elem: resource,
				}
			}
		}
		return &AnnotatedSchema{
			Schemata: schemata,
		}
	case reflect.Map:
		sch := &schema.Schema{
			Type: schema.TypeList,
			Elem: structFor(fmt.Sprintf("%v[%v]", hint, t.Elem()), v, t.Elem()),
		}
		switch typedElem := sch.Elem.(type) {
		case *schema.Resource:
			typedElem.Schema["key"] = &schema.Schema{Type: schema.TypeString, Required: true}
		case *schema.Schema:
			if (typedElem.Type == schema.TypeList) && (typedElem.MaxItems == 1) {
				if typedTypedElem, ok := typedElem.Elem.(*schema.Resource); ok {
					if _, exists := typedTypedElem.Schema["key"]; !exists {
						typedTypedElem.Schema["key"] = &schema.Schema{Type: schema.TypeString, Required: true}
						typedElem.MaxItems = 0
						return &AnnotatedSchema{Schema: typedElem}
					}
				}
			}
			typedElem.Required = true
			sch.Elem = &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key":    {Type: schema.TypeString, Required: true},
					"values": typedElem,
				},
			}
		case *AnnotatedSchema:
			if (typedElem.Schema.Type == schema.TypeList) && (typedElem.Schema.MaxItems == 1) {
				if typedTypedElem, ok := typedElem.Schema.Elem.(*schema.Resource); ok {
					if _, exists := typedTypedElem.Schema["key"]; !exists {
						typedTypedElem.Schema["key"] = &schema.Schema{Type: schema.TypeString, Required: true}
						typedElem.Schema.MaxItems = 0
						return &AnnotatedSchema{Schema: typedElem.Schema}
					}
				}
			}
			typedElem.Schema.Required = true
			sch.Elem = &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key":    {Type: schema.TypeString, Required: true},
					"values": typedElem.Schema,
				},
			}
		default:
			panic(fmt.Errorf("unexpected type found within schema.Schema.Elem: %T", sch.Elem))
		}
		return &AnnotatedSchema{Schema: sch}
	default:
		panic(fmt.Errorf("unsupported type %v (kind: %v)", t, t.Kind()))
	}
}

var defIndent = "\t"

// DumpResource has no documentation
func DumpResource(w io.Writer, resource *schema.Resource, indent string) {
	fmt.Fprintln(w, fmt.Sprintf("%s&schema.Resource{", indent))
	dumpResource(w, resource, indent+defIndent)
	fmt.Fprintln(w, fmt.Sprintf("%s}", indent))
}

func dumpResource(w io.Writer, resource *schema.Resource, indent string) {
	fmt.Fprintln(w, fmt.Sprintf("%sSchema: map[string]*schema.Schema{", indent))
	for k, v := range resource.Schema {
		if isSingleLineSchema(v.Type) {
			fmt.Fprint(w, fmt.Sprintf(`%s"%s": {`, indent+defIndent, k))
			dumpSchemaSingleLine(w, v)
			fmt.Fprintln(w, "},")
		} else {
			fmt.Fprintln(w, fmt.Sprintf(`%s"%s": {`, indent+defIndent, k))
			dumpSchema(w, v, indent+defIndent+defIndent)
			fmt.Fprintln(w, fmt.Sprintf(`%s},`, indent+defIndent))
		}
	}
	fmt.Fprintln(w, fmt.Sprintf("%s},", indent+defIndent))
}

func typestr(t schema.ValueType) string {
	switch t {
	case schema.TypeString:
		return "schema.TypeString"
	case schema.TypeSet:
		return "schema.TypeSet"
	case schema.TypeInt:
		return "schema.TypeInt"
	case schema.TypeList:
		return "schema.TypeList"
	case schema.TypeFloat:
		return "schema.TypeFloat"
	case schema.TypeBool:
		return "schema.TypeBool"
	default:
		panic(fmt.Errorf("unknown ValueType %v", t))
	}
}

func isSingleLineSchema(t schema.ValueType) bool {
	switch t {
	case schema.TypeList, schema.TypeSet, schema.TypeMap:
		return false
	default:
		return true
	}
}

func dumpSchemaSingleLine(w io.Writer, sch *schema.Schema) {
	fmt.Fprint(w, fmt.Sprintf("Type: %v", typestr(sch.Type)))
	if sch.MaxItems != 0 {
		fmt.Fprint(w, fmt.Sprintf(", MaxItems: %v", sch.MaxItems))
	}
	if sch.Optional {
		fmt.Fprint(w, fmt.Sprintf(", Optional: %v", sch.Optional))
	}
	if sch.Required {
		fmt.Fprint(w, fmt.Sprintf(", Required: %v", sch.Required))
	}
}

func dumpSchema(w io.Writer, sch *schema.Schema, indent string) {
	switch typedElem := sch.Elem.(type) {
	case *AnnotatedSchema:
		if typedElem.Schemata != nil {
			for _, v := range typedElem.Schemata {
				dumpSchema(w, v, indent)
			}
			return
		}
	default:
	}

	fmt.Fprintln(w, fmt.Sprintf("%sType: %v,", indent, typestr(sch.Type)))
	if sch.MaxItems != 0 {
		fmt.Fprintln(w, fmt.Sprintf("%vMaxItems: %v,", indent, sch.MaxItems))
	}
	if sch.Optional {
		fmt.Fprintln(w, fmt.Sprintf("%sOptional: %v,", indent, sch.Optional))
	}
	if sch.Required {
		fmt.Fprintln(w, fmt.Sprintf("%sRequired: %v,", indent, sch.Required))
	}
	if sch.Elem != nil {
		switch typedElem := sch.Elem.(type) {
		case *schema.Resource:
			fmt.Fprintln(w, fmt.Sprintf("%sElem: &schema.Resource{", indent+defIndent))
			dumpResource(w, typedElem, indent+defIndent+defIndent)
			fmt.Fprintln(w, fmt.Sprintf("%s},", indent+defIndent))
		case *schema.Schema:
			if isSingleLineSchema(typedElem.Type) {
				fmt.Fprint(w, fmt.Sprintf("%sElem: &schema.Schema{", indent+defIndent))
				dumpSchemaSingleLine(w, typedElem)
				fmt.Fprintln(w, "},")
			} else {
				fmt.Fprintln(w, fmt.Sprintf("%sElem: &schema.Schema{", indent+defIndent))
				dumpSchema(w, typedElem, indent+defIndent+defIndent)
				fmt.Fprintln(w, fmt.Sprintf("%s},", indent+defIndent))
			}
		case *AnnotatedSchema:
			if isSingleLineSchema(typedElem.Schema.Type) {
				fmt.Fprint(w, fmt.Sprintf("%sElem: &schema.Schema{", indent+defIndent))
				dumpSchemaSingleLine(w, typedElem.Schema)
				fmt.Fprintln(w, "},")
			} else {
				fmt.Fprintln(w, fmt.Sprintf("%sElem: &schema.Schema{", indent+defIndent))
				dumpSchema(w, typedElem.Schema, indent+defIndent+defIndent)
				fmt.Fprintln(w, fmt.Sprintf("%s},", indent+defIndent))
			}
		default:
			panic(fmt.Errorf("unsupported elem type %T", sch.Elem))
		}
	}
}
