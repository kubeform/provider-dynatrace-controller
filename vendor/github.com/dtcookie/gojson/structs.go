package gojson

import (
	"encoding/json"
	"reflect"
)

func (enc *encoder) marshalStruct(v reflect.Value) ([]byte, error) {
	var result string
	var err error
	rawProps := rawProperties{}
	var vUnknownProperties *reflect.Value
	if rawProps, vUnknownProperties, err = enc.marshalStructFields(v); err != nil {
		return nil, err
	}
	if vUnknownProperties != nil {
		innerRawProps := rawProperties{}
		if innerRawProps, err = enc.marshalMapEntries(*vUnknownProperties, rawProps.Names()); err != nil {
			return nil, err
		}
		rawProps.Merge(innerRawProps)
	}
	if (rawProps == nil) || len(rawProps) == 0 {
		return []byte("{}"), nil
	}
	var data []byte
	for _, p := range rawProps {
		if !empty(result) {
			result = result + ", "
		}
		if data, err = json.Marshal(p.Name); err != nil {
			return nil, err
		}
		result = result + string(data) + ": " + string(p.Bytes)
	}
	return []byte("{ " + result + " }"), nil
}

// marshalStructFields produces three results
// * The marshalled fields of the Struct represented by the given Value v
// * If found within the Struct (or within an embedded Struct), a Value representing the UnknownProperties Struct Field
// * an error in case something went wrong
func (enc *encoder) marshalStructFields(v reflect.Value) (rawProperties, *reflect.Value, error) {
	var err error
	rawProps := rawProperties{}
	var vUnknownProperties *reflect.Value
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			innerRawProps := []rawProperty{}
			if innerRawProps, vUnknownProperties, err = enc.marshalStructFields(v.Field(i)); err != nil {
				return nil, nil, err
			}
			rawProps.Merge(innerRawProps)
		}
	}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			continue
		}
		if !startsWithUpper(field.Name) {
			continue
		}
		propName := field.Name
		omitEmpty := false
		propName, omitEmpty = evalTag(field)
		if propName == "-" {
			// TODO: pointer to UnknownProperties should also be allowed
			if field.Type == tUnknownProperties {
				vField := v.Field(i)
				vUnknownProperties = &vField
			}
			continue
		}
		vField := v.Field(i)
		if !omitEmpty || !vField.IsZero() {
			var data []byte
			if data, err = json.Marshal(propName); err != nil {
				return nil, nil, err
			}
			if data, err = enc.marshal(vField); err != nil {
				return nil, nil, err
			}
			rawProps.Add(rawProperty{Name: propName, Bytes: data})
		}
	}
	return rawProps, vUnknownProperties, nil
}
