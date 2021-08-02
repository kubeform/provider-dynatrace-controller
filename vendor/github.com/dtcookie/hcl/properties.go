package hcl

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Properties map[string]interface{}

func NewProperties(v interface{}, unknowns ...map[string]json.RawMessage) (Properties, error) {
	properties := Properties{}
	if len(unknowns) == 0 || len(unknowns[0]) == 0 {
		return properties, nil
	}

	// if data, err := json.Marshal(unknowns[0]); err == nil {
	// 	p := xjson.Properties{}
	// 	if err = json.Unmarshal(data, &p); err != nil {
	// 		return nil, err
	// 	}
	// 	for k := range p {
	// 		delete(unknowns[0], k)
	// 	}
	// }
	data, err := json.Marshal(unknowns[0])
	if err != nil {
		return nil, err
	}
	properties["unknowns"] = string(data)
	return properties, nil
}

func (me Properties) MarshalAll(decoder Decoder, items map[string]interface{}) error {
	if items == nil {
		return nil
	}
	for k, v := range items {
		if err := me.Marshal(decoder, k, v); err != nil {
			return err
		}
	}
	return nil
}

func (me Properties) Marshal(decoder Decoder, key string, v interface{}) error {
	switch t := v.(type) {
	case *string:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *bool:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *int:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *int8:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *int16:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *int32:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *int64:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *uint:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *uint16:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *uint8:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *uint32:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *uint64:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *float32:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case *float64:
		if t == nil {
			return nil
		}
		return me.Marshal(decoder, key, *t)
	case string:
		me[key] = t
	case int:
		me[key] = t
	case bool:
		me[key] = t
	case int8:
		me[key] = int(t)
	case int16:
		me[key] = int(t)
	case int32:
		me[key] = int(t)
	case int64:
		me[key] = int(t)
	case uint:
		me[key] = int(t)
	case uint8:
		me[key] = int(t)
	case uint16:
		me[key] = int(t)
	case uint32:
		me[key] = int(t)
	case uint64:
		me[key] = int(t)
	case float32:
		me[key] = float64(t)
	case float64:
		me[key] = float64(t)
	default:
		if marshaller, ok := v.(ExtMarshaler); ok {
			if marshalled, err := marshaller.MarshalHCL(NewDecoder(decoder, key, 0)); err == nil {
				me[key] = []interface{}{marshalled}
				return nil
			} else {
				return err
			}
		}

		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			me[key] = fmt.Sprintf("%v", v)
		case reflect.Ptr:
			switch reflect.TypeOf(v).Elem().Kind() {
			case reflect.String:
				if reflect.ValueOf(v).IsNil() {
					return nil
				}
				if reflect.ValueOf(v).IsZero() {
					return nil
				}
				if !reflect.ValueOf(v).Elem().IsValid() {
					return nil
				}
				return me.Marshal(decoder, key, reflect.ValueOf(v).Elem().Interface())
			}

		default:
			log.Printf("unsupported type %T", v)
		}
	}
	return nil
}

func (me Properties) EncodeSlice(key string, v interface{}) (Properties, error) {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Slice {
		return nil, fmt.Errorf("type %T is not a slice", v)
	}
	if rv.Len() == 0 {
		return me, nil
	}
	entries := []interface{}{}
	for idx := 0; idx < rv.Len(); idx++ {
		vElem := rv.Index(idx)
		elem := vElem.Interface()
		if marshaler, ok := elem.(Marshaler); ok {
			if marshalled, err := marshaler.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("slice entries of type %T are expected to implement hcl.Marshaler but don't", elem)
		}
	}
	me[key] = entries
	return me, nil
}

func (me Properties) EncodeAll(items map[string]interface{}) (Properties, error) {
	if items == nil {
		return me, nil
	}
	for k, v := range items {
		if err := me.Encode(k, v); err != nil {
			return me, err
		}
	}
	return me, nil
}

func (me Properties) Encode(key string, v interface{}) error {
	if v == nil {
		return nil
	}
	switch t := v.(type) {
	case *string:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *bool:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *int:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *int8:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *int16:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *int32:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *int64:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *uint:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *uint16:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *uint8:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *uint32:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *uint64:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *float32:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case *float64:
		if t == nil {
			return nil
		}
		return me.Encode(key, *t)
	case []string:
		if len(t) > 0 {
			me[key] = t
		}
	case string:
		me[key] = t
	case int:
		me[key] = t
	case bool:
		me[key] = t
	case int8:
		me[key] = int(t)
	case int16:
		me[key] = int(t)
	case int32:
		me[key] = int(t)
	case int64:
		me[key] = int(t)
	case uint:
		me[key] = int(t)
	case uint8:
		me[key] = int(t)
	case uint16:
		me[key] = int(t)
	case uint32:
		me[key] = int(t)
	case uint64:
		me[key] = int(t)
	case float32:
		me[key] = float64(t)
	case float64:
		me[key] = float64(t)
	case map[string]json.RawMessage:
		if len(t) == 0 {
			return nil
		}
		data, err := json.Marshal(t)
		if err != nil {
			return err
		}
		me["unknowns"] = string(data)
	default:
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			if reflect.ValueOf(v).Len() == 0 {
				return nil
			}
			if reflect.TypeOf(v).Elem().Kind() == reflect.String {
				entries := []string{}
				vValue := reflect.ValueOf(v)
				for i := 0; i < vValue.Len(); i++ {
					entries = append(entries, fmt.Sprintf("%v", vValue.Index(i).Interface()))
				}
				me[key] = entries
				return nil
			} else if reflect.TypeOf(v).Elem().Kind() == reflect.Float64 {
				entries := []float64{}
				vValue := reflect.ValueOf(v)
				for i := 0; i < vValue.Len(); i++ {
					entries = append(entries, vValue.Index(i).Interface().(float64))
				}
				me[key] = entries
				return nil
			}

		}
		if reflect.TypeOf(v).Kind() == reflect.String {
			me[key] = fmt.Sprintf("%v", v)
			return nil
		} else if marshaller, ok := v.(ExtMarshaler); ok {
			if reflect.ValueOf(v).IsNil() {
				return nil
			}
			if marshalled, err := marshaller.MarshalHCL(VoidDecoder()); err == nil {
				me[key] = []interface{}{marshalled}
				return nil
			} else {
				return err
			}
		} else if marshaller, ok := v.(Marshaler); ok {
			if reflect.ValueOf(v).IsNil() {
				return nil
			}
			if marshalled, err := marshaller.MarshalHCL(); err == nil {
				me[key] = []interface{}{marshalled}
				return nil
			} else {
				return err
			}

		} else if reflect.TypeOf(v).Kind() == reflect.Ptr {
			switch reflect.TypeOf(v).Elem().Kind() {
			case reflect.String:
				if reflect.ValueOf(v).IsNil() {
					return nil
				}
				if reflect.ValueOf(v).IsZero() {
					return nil
				}
				if !reflect.ValueOf(v).Elem().IsValid() {
					return nil
				}
				return me.Encode(key, reflect.ValueOf(v).Elem().Interface())
			}
		}
		panic(fmt.Sprintf("unsupported type %T", v))
	}
	return nil
}
