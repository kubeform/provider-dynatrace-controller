package gojson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

// Marshal has no documentation
func Marshal(v interface{}) ([]byte, error) {
	enc := &encoder{buf: new(bytes.Buffer)}
	var data json.RawMessage
	var err error
	if data, err = enc.marshal(reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

// MarshalIndent has no documentation
func MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	enc := &encoder{buf: new(bytes.Buffer)}
	var data json.RawMessage
	var err error
	if data, err = enc.marshal(reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return json.MarshalIndent(data, prefix, indent)
}

type encoder struct {
	buf *bytes.Buffer
}

func (enc *encoder) marshal(v reflect.Value) ([]byte, error) {
	switch v.Type().Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return []byte("null"), nil
		}
		return enc.marshal(v.Elem())
	case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return json.Marshal(v.Interface())
	case reflect.Map:
		return enc.marshalMap(v)
	case reflect.Slice:
		if v.IsNil() {
			return []byte("null"), nil
		}
		fallthrough
	case reflect.Array:
		return enc.marshalSlice(v)
	case reflect.Struct:
		return enc.marshalStruct(v)
	case reflect.Interface:
		if v.IsNil() {
			return []byte("null"), nil
		}

		return enc.marshal(v.Elem())
	default:
		return nil, fmt.Errorf("unsupported type %v (kind: %v)", v.Type(), v.Type().Kind())
	}
}
