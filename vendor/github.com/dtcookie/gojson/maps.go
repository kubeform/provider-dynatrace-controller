package gojson

import (
	"encoding/json"
	"reflect"
)

func (enc *encoder) marshalMap(v reflect.Value) ([]byte, error) {
	if v.IsNil() {
		return []byte("null"), nil
	}
	var result string
	rawProps := rawProperties{}
	var err error
	if rawProps, err = enc.marshalMapEntries(v, nil); err != nil {
		return nil, err
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

func (enc *encoder) marshalMapEntries(v reflect.Value, except []string) (rawProperties, error) {
	rawProps := rawProperties{}
	if v.IsNil() || (v.Len() == 0) {
		return nil, nil
	}
	for _, vKey := range v.MapKeys() {
		key := vKey.Interface().(string)
		if contains(except, key) {
			continue
		}
		var err error
		var data []byte
		entry := v.MapIndex(vKey)
		if entry.IsZero() {
			data = []byte("null")
		} else if data, err = enc.marshal(entry); err != nil {
			return nil, err
		}
		rawProps.Add(rawProperty{Name: key, Bytes: data})
	}
	return rawProps, nil
}
