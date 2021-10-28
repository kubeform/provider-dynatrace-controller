package gojson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// Unmarshal has no documentation
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(v)}
	}
	dec := new(decoder)
	return dec.unmarshal(data, rv)
}

type decoder struct {
	data []byte
}

type property struct {
	Field reflect.StructField
	Value reflect.Value
}

func collectProperties(v reflect.Value, properties map[string]property) UnknownProperties {
	var unknownProperties UnknownProperties
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Anonymous {
			unknownProperties = collectProperties(v.Field(i), properties)
		}
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Anonymous {
			continue
		}
		if !startsWithUpper(field.Name) {
			continue
		}
		propName := field.Name
		propName, _ = evalTag(field)
		if propName == "-" {
			// TODO: pointer to UnknownProperties should also be allowed
			if field.Type == tUnknownProperties {
				unknownProperties = v.Field(i).Interface().(UnknownProperties)
			}
			continue
		}
		properties[propName] = property{field, v.Field(i)}
	}
	return unknownProperties
}

func (d *decoder) unmarshal(data []byte, rv reflect.Value) error {
	var err error
	switch rv.Type().Kind() {
	case reflect.Ptr:
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		return d.unmarshal(data, rv.Elem())
	case reflect.String:
		var v string
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v).Convert(rv.Type()))
		return nil
	case reflect.Bool:
		var v bool
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Int:
		var v int
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Int8:
		var v int8
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Int16:
		var v int16
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Int32:
		var v int32
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Int64:
		var v int64
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Uint:
		var v uint
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Uint8:
		var v uint8
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Uint16:
		var v uint16
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Uint32:
		var v uint32
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Uint64:
		var v uint64
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Float32:
		var v float32
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Float64:
		var v float64
		if err = json.Unmarshal(data, &v); err != nil {
			return err
		}
		rv.Set(reflect.ValueOf(v))
		return nil
	case reflect.Slice:
		var err error
		rawEntries := []json.RawMessage{}
		if err = json.Unmarshal(data, &rawEntries); err != nil {
			return err
		}
		rv.Set(reflect.MakeSlice(rv.Type(), len(rawEntries), len(rawEntries)))
		for i := 0; i < rv.Len(); i++ {
			if err = d.unmarshal(rawEntries[i], rv.Index(i)); err != nil {
				return err
			}
		}
		return nil
	case reflect.Struct:
		var err error
		properties := map[string]property{}
		collectProperties(rv, properties)
		rawDatas := map[string]json.RawMessage{}
		if err = json.Unmarshal(data, &rawDatas); err != nil {
			return err
		}
		for k, jsonRawMessage := range rawDatas {
			if property, ok := properties[k]; ok {
				if err = d.unmarshal(jsonRawMessage, property.Value); err != nil {
					return err
				}
			}
		}
		return nil
	case reflect.Interface:
		baseType := findBaseType(rv)
		if baseType == nil {
			panic(fmt.Errorf("Unsupported type %v (kind: %v) - base type not found", rv.Type(), rv.Type().Kind()))
		}
		implementors := findImplementors(baseType)
		if implementors == nil {
			panic(fmt.Errorf("Unsupported type %v (kind: %v) - implementors not found", rv.Type(), rv.Type().Kind()))
		}
		rawProperties := map[string]json.RawMessage{}
		if err = json.Unmarshal(data, &rawProperties); err != nil {
			return err
		}
		for _, implementor := range implementors {
			key, value := discriminator(reflect.TypeOf(implementor))
			if key == "" || value == "" {
				continue
			}
			if rawValue, found := rawProperties[key]; found {
				decoded := string(rawValue)
				decoded = strings.TrimSpace(decoded[1 : len(decoded)-1])
				if value == decoded {
					inst := reflect.New(reflect.TypeOf(implementor).Elem())
					if err = d.unmarshal(data, inst); err != nil {
						return err
					}
					rv.Set(inst)
					return nil
				}
			}
		}

		inst := reflect.New(baseType.Elem())
		if err = d.unmarshal(data, inst); err != nil {
			return err
		}
		rv.Set(inst)
		return nil

		// fmt.Println("----------")
		// for _, implementor := range implementors {
		// 	key, value := discriminator(reflect.TypeOf(implementor))
		// 	fmt.Println("  ", key, value, reflect.TypeOf(implementor))
		// }
		// fmt.Println("----------")
		// panic(fmt.Errorf("found no implementor: %v", string(data)))
	default:
		panic(fmt.Errorf("Unsupported type %v (kind: %v)", rv.Type(), rv.Type().Kind()))
		// panic(&json.InvalidUnmarshalError{Type: rv.Type()})
		// return &json.InvalidUnmarshalError{Type: rv.Type()}
	}

}

func discriminator(t reflect.Type) (string, string) {
	t = t.Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.Anonymous {
			continue
		}
		var jsonTag string
		var found bool
		if jsonTag, found = field.Tag.Lookup("json"); found {
			if !strings.HasPrefix(jsonTag, ",") {
				continue
			}
			jsonTag = jsonTag[1:]
			if !strings.Contains(jsonTag, "=") {
				continue
			}
			parts := strings.Split(jsonTag, "=")
			return parts[0], parts[1]
		}
	}
	return "", ""
}

func findImplementors(baseType reflect.Type) []interface{} {
	for i := 0; i < baseType.NumMethod(); i++ {
		method := baseType.Method(i)
		if method.Name == "Implementors" {
			if method.Type.NumOut() != 1 {
				return nil
			}
			if method.Type.NumIn() != 1 {
				fmt.Println(method.Type.In(0))
				return nil
			}
			retVal := reflect.New(baseType).Elem().Method(i).Call([]reflect.Value{})[0]
			result := []interface{}{}
			for j := 0; j < retVal.Len(); j++ {
				result = append(result, retVal.Index(j).Interface())
			}
			return result
		}
	}
	return nil
}

func findBaseType(rv reflect.Value) reflect.Type {
	t := rv.Type()
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if method.Name == "Initialize" {
			if method.Type.NumOut() != 0 {
				return nil
			}
			if method.Type.NumIn() != 1 {
				return nil
			}
			return method.Type.In(0)
		}
	}
	return nil
}
