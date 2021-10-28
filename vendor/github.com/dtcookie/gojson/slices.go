package gojson

import "reflect"

func (enc *encoder) marshalSlice(v reflect.Value) ([]byte, error) {
	var result string
	var data []byte
	var err error
	if data, err = enc.marshalSliceElems(v); err != nil {
		return nil, err
	}
	if data != nil {
		result = result + string(data)
	}
	return []byte("[ " + result + " ]"), nil
}

func (enc *encoder) marshalSliceElems(v reflect.Value) ([]byte, error) {
	var result string
	if ((v.Type().Kind() == reflect.Slice) && v.IsNil()) || (v.Len() == 0) {
		return nil, nil
	}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		var err error
		var data []byte
		elemKind := elem.Type().Kind()
		if (elemKind == reflect.Ptr || elemKind == reflect.Array || elemKind == reflect.Interface || elemKind == reflect.Slice || elemKind == reflect.Map) && elem.IsZero() {
			data = []byte("null")
		} else if data, err = enc.marshal(elem); err != nil {
			return nil, err
		}
		if !empty(result) {
			result = result + ", "
		}
		result = result + string(data)
	}
	return []byte(result), nil
}
