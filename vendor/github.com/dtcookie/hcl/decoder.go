package hcl

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/dtcookie/opt"
)

// Decoder has no documentation
type MinDecoder interface {
	GetOk(key string) (interface{}, bool)
	Get(key string) interface{}
	GetChange(key string) (interface{}, interface{})
	GetOkExists(key string) (interface{}, bool)
	HasChange(key string) bool
}

// Decoder has no documentation
type Decoder interface {
	GetOk(key string) (interface{}, bool)
	Get(key string) interface{}
	GetChange(key string) (interface{}, interface{})
	GetStringSet(key string) []string
	GetOkExists(key string) (interface{}, bool)
	Reader(unkowns ...map[string]json.RawMessage) Reader
	HasChange(key string) bool
	MarshalAll(items map[string]interface{}) (Properties, error)

	Decode(key string, v interface{}) error
	DecodeAll(map[string]interface{}) error
	DecodeAny(map[string]interface{}) (interface{}, error)

	DecodeSlice(key string, v interface{}) error
}

type mindecoder struct {
	parent MinDecoder
}

func DecoderFrom(m MinDecoder) Decoder {
	return &mindecoder{parent: m}
}

func (d *mindecoder) Decode(key string, v interface{}) error {
	return NewDecoder(d).Decode(key, v)
}

func (d *mindecoder) DecodeAll(m map[string]interface{}) error {
	return NewDecoder(d).DecodeAll(m)
}

func (d *mindecoder) DecodeSlice(key string, v interface{}) error {
	return NewDecoder(d).DecodeSlice(key, v)
}

func (d *mindecoder) DecodeAny(m map[string]interface{}) (interface{}, error) {
	return NewDecoder(d).DecodeAny(m)
}

func (d *decoder) DecodeAny(m map[string]interface{}) (interface{}, error) {
	if len(m) == 0 {
		return nil, nil
	}
	for k, v := range m {
		found, err := d.decode(k, v)
		if err != nil {
			return nil, err
		}
		if found {
			return v, nil
		}
	}
	return nil, nil
}

func (d *decoder) DecodeAll(m map[string]interface{}) error {
	if len(m) == 0 {
		return nil
	}
	for k, v := range m {
		if err := d.Decode(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (d *decoder) DecodeSlice(key string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Type().Kind() != reflect.Ptr || rv.Type().Elem().Kind() != reflect.Slice {
		return fmt.Errorf("decoding slices requires a pointer to a slice to be specified. %T doesn't qualify", v)
	}
	elemType := rv.Type().Elem().Elem()
	if !elemType.Implements(reflect.TypeOf((*Unmarshaler)(nil)).Elem()) {
		return fmt.Errorf("decoding slices requires a pointer to a slice of elements that implement hcl.Unmarshaler to be specified. %T doesn't qualify (%v is not implementing %v)", v, elemType, reflect.TypeOf((*Unmarshaler)(nil)).Elem())
	}
	vSlice := rv.Elem()
	if result, ok := d.GetOk(fmt.Sprintf("%v.#", key)); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := reflect.New(elemType.Elem()).Interface()
			if err := entry.(Unmarshaler).UnmarshalHCL(NewDecoder(d, key, idx)); err != nil {
				return err
			}
			vSlice.Set(reflect.Append(vSlice, reflect.ValueOf(entry)))
		}
	}

	return nil
}

func (d *decoder) Decode(key string, v interface{}) error {
	_, err := d.decode(key, v)
	return err
}

var stringType = reflect.TypeOf("")
var stringSliceType = reflect.TypeOf([]string{})
var marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()

func (d *decoder) decode(key string, v interface{}) (bool, error) {
	vTarget := reflect.ValueOf(v)
	if !vTarget.IsValid() || vTarget.IsNil() {
		return false, errors.New("passed an invalid target value to Decode()")
	}
	if vTarget.Type().Kind() == reflect.Ptr {
		valueType := vTarget.Type()
		valueType = valueType.Elem()
		if valueType.Kind() == reflect.Ptr {
			structValueType := valueType.Elem()
			if structValueType.Kind() == reflect.Struct {
				if valueType.Implements(marshalerType) {
					// log.Printf("%v implements %v ==> %v", valueType, marshalerType, true)
					newValue := reflect.New(structValueType)
					// log.Printf("new value has type %v (valueType: %v)", newValue.Type(), valueType)
					newValueIface := newValue.Interface()
					if unmarshaler, ok := newValueIface.(Unmarshaler); ok {
						if _, ok := d.GetOk(fmt.Sprintf("%v.#", key)); ok {
							vTarget = vTarget.Elem()
							vTarget.Set(newValue)
							if err := unmarshaler.UnmarshalHCL(NewDecoder(d, key, 0)); err != nil {
								return true, err
							}
							return true, nil
						}
					}
				}
			}
		}
	}
	if unmarshaler, ok := v.(Unmarshaler); ok {
		if _, ok := d.GetOk(fmt.Sprintf("%v.#", key)); ok {
			if err := unmarshaler.UnmarshalHCL(NewDecoder(d, key, 0)); err != nil {
				return true, err
			}
			return true, nil
		}
	}
	if vTarget.Type().Kind() != reflect.Ptr {
		return false, fmt.Errorf("Decode (%v) requires a pointer to store results into", key)
	}
	if result, ok := d.GetOk(key); ok {
		switch vActual := v.(type) {
		case *[]string:
			set := result.(Set)
			strs := []string{}
			for _, elem := range set.List() {
				strs = append(strs, elem.(string))
			}
			*vActual = strs
			return true, nil
		case *[]float64:
			set := result.(Set)
			strs := []float64{}
			for _, elem := range set.List() {
				strs = append(strs, elem.(float64))
			}
			*vActual = strs
			return true, nil
		case *string:
			*vActual = result.(string)
			return true, nil
		case **string:
			*vActual = opt.NewString(result.(string))
			return true, nil
		case *bool:
			*vActual = result.(bool)
			return true, nil
		case **bool:
			*vActual = opt.NewBool(result.(bool))
			return true, nil
		case *int32:
			*vActual = int32(result.(int))
			return true, nil
		case **int32:
			*vActual = opt.NewInt32(int32(result.(int)))
			return true, nil
		case *int64:
			*vActual = int64(result.(int))
			return true, nil
		case **int64:
			*vActual = opt.NewInt64(int64(result.(int)))
			return true, nil
		case *int8:
			*vActual = int8(result.(int))
			return true, nil
		case **int8:
			*vActual = opt.NewInt8(int8(result.(int)))
			return true, nil
		case *int16:
			*vActual = int16(result.(int))
			return true, nil
		case **int16:
			*vActual = opt.NewInt16(int16(result.(int)))
			return true, nil
		case *int:
			*vActual = int(result.(int))
			return true, nil
		case **int:
			*vActual = opt.NewInt(int(result.(int)))
			return true, nil
		case *uint32:
			*vActual = uint32(result.(int))
			return true, nil
		case **uint32:
			*vActual = opt.NewUInt32(uint32(result.(int)))
			return true, nil
		case *uint64:
			*vActual = uint64(result.(int))
			return true, nil
		case **uint64:
			*vActual = opt.NewUInt64(uint64(result.(int)))
			return true, nil
		case *uint8:
			*vActual = uint8(result.(int))
			return true, nil
		case **uint8:
			*vActual = opt.NewUInt8(uint8(result.(int)))
			return true, nil
		case *uint16:
			*vActual = uint16(result.(int))
			return true, nil
		case **uint16:
			*vActual = opt.NewUInt16(uint16(result.(int)))
			return true, nil
		case *uint:
			*vActual = uint(result.(int))
			return true, nil
		case **uint:
			*vActual = opt.NewUint(uint(result.(int)))
			return true, nil
		case *float64:
			*vActual = float64(result.(float64))
			return true, nil
		case **float64:
			*vActual = opt.NewFloat64(float64(result.(float64)))
			return true, nil
		default:
			vTarget := reflect.ValueOf(v)
			tTarget := vTarget.Type()
			if tTarget.Kind() == reflect.Ptr {
				tElem := tTarget.Elem()
				if tElem.Kind() == reflect.String {
					vTarget := vTarget.Elem()
					vResult := reflect.ValueOf(result)
					tTarget := vTarget.Type()
					vTarget.Set(vResult.Convert(tTarget))
					return true, nil
				} else if tElem.Kind() == reflect.Ptr {
					tElem := tElem.Elem()
					if tElem.Kind() == reflect.String {
						vTarget := vTarget.Elem()
						vNewEnumPtr := reflect.New(vTarget.Type().Elem())
						vNewEnum := vNewEnumPtr.Elem()
						valueToSet := reflect.ValueOf(result).Convert(vTarget.Type().Elem())
						vNewEnum.Set(valueToSet)
						vTarget.Set(vNewEnumPtr)
						return true, nil
					}
				} else if tElem.Kind() == reflect.Slice {
					tSliceElem := tElem.Elem()
					if tSliceElem.Kind() == reflect.String {
						enumType := tElem.Elem()
						enumSliceType := reflect.SliceOf(enumType)
						vEnumSlicePtr := reflect.New(enumSliceType)
						vEnumSlice := vEnumSlicePtr.Elem()
						for _, iString := range result.(Set).List() {
							vEnumSlice = reflect.Append(vEnumSlice, reflect.ValueOf(iString).Convert(enumType))
						}
						vTarget.Elem().Set(reflect.ValueOf(vEnumSlice.Interface()))
						return true, nil
					}
				}
			}
		}
		vTarget := vTarget.Elem()
		vResult := reflect.ValueOf(result)
		tResult := vResult.Type()
		tTarget := vTarget.Type()
		// tOrigTarget := reflect.ValueOf(v).Type()
		if tResult == stringType {
			if tTarget.Kind() == reflect.String {
				if tTarget != stringType {
					vTarget.Set(vResult.Convert(tTarget))
					// log.Printf("%v %v covered", tOrigTarget, key)
					return true, nil
				}
			}
			if tTarget.Kind() == reflect.Ptr {
				tTarget = tTarget.Elem()
				if tTarget.Kind() == reflect.String {
					if tTarget != stringType {
						tEnum := reflect.ValueOf(v).Type().Elem().Elem()
						vEnumPtr := reflect.New(tEnum)
						vEnum := vEnumPtr.Elem()
						vEnum.Set(vResult.Convert(tEnum))
						vTarget.Set(vEnumPtr)
						// log.Printf("%v %v covered", tOrigTarget, key)
						return true, nil
					} else {
						vTarget.Set(reflect.ValueOf(opt.NewString(result.(string))))
						// log.Printf("%v %v covered", tOrigTarget, key)
						return true, nil
					}
				}
			}
		}
		setResult, ok := result.(Set)
		if ok && vTarget.Type() == stringSliceType {
			entries := []string{}
			for _, entry := range setResult.List() {
				entries = append(entries, entry.(string))
			}
			vTarget.Set(reflect.ValueOf(entries))
			return true, nil
		}
		if vResult.Type().AssignableTo(vTarget.Type()) {
			vTarget.Set(vResult)
			return true, nil
		} else {
			log.Printf("[WARN] %v %v NOT covered", reflect.ValueOf(v).Type(), key)
		}
	}
	return false, nil
}

func (d *mindecoder) GetStringSet(key string) []string {
	result := []string{}
	if value, ok := d.GetOk(key); ok {
		if arr, ok := value.([]interface{}); ok {
			for _, elem := range arr {
				result = append(result, elem.(string))
			}
		} else if set, ok := value.(Set); ok {
			if set.Len() > 0 {
				for _, elem := range set.List() {
					result = append(result, elem.(string))
				}
			}
		}
	}
	return result
}

func (d *mindecoder) Reader(unkowns ...map[string]json.RawMessage) Reader {
	if len(unkowns) > 0 {
		return NewReader(d, unkowns[0])
	}
	return NewReader(d, nil)
}

func (d *mindecoder) MarshalAll(items map[string]interface{}) (Properties, error) {
	properties := Properties{}
	if err := properties.MarshalAll(d, items); err != nil {
		return nil, err
	}
	return properties, nil
}

func (d *mindecoder) GetOk(key string) (interface{}, bool) {
	return d.parent.GetOk(key)
}

func (d *mindecoder) HasChange(key string) bool {
	return d.parent.HasChange(key)
}

func (d *mindecoder) GetOkExists(key string) (interface{}, bool) {
	return d.parent.GetOkExists(key)
}

func (d *mindecoder) GetChange(key string) (interface{}, interface{}) {
	return d.parent.GetChange(key)
}

func (d *mindecoder) Get(key string) interface{} {
	return d.parent.Get(key)
}

// NewDecoder has no documentation
func NewDecoder(parent MinDecoder, address ...interface{}) Decoder {
	joined := ""
	sep := ""
	for _, part := range address {
		joined = joined + sep + fmt.Sprintf("%v", part)
		sep = "."
	}
	return &decoder{parent: parent, address: joined}
}

type decoder struct {
	parent  MinDecoder
	address string
}

func (d *decoder) Reader(unkowns ...map[string]json.RawMessage) Reader {
	if len(unkowns) > 0 {
		return NewReader(d, unkowns[0])
	}
	return NewReader(d, nil)
}

func (d *decoder) MarshalAll(items map[string]interface{}) (Properties, error) {
	properties := Properties{}
	if err := properties.MarshalAll(d, items); err != nil {
		return nil, err
	}
	return properties, nil
}

func (d *decoder) HasChange(key string) bool {
	if d.address == "" {
		return d.parent.HasChange(key)
	}
	return d.parent.HasChange(d.address + "." + key)
}

func (d *decoder) GetStringSet(key string) []string {
	result := []string{}
	if value, ok := d.GetOk(key); ok {
		if arr, ok := value.([]interface{}); ok {
			for _, elem := range arr {
				result = append(result, elem.(string))
			}
		} else if set, ok := value.(Set); ok {
			if set.Len() > 0 {
				for _, elem := range set.List() {
					result = append(result, elem.(string))
				}
			}
		}
	}
	return result
}

// GetOk returns the data for the given key and whether or not the key
// has been set to a non-zero value at some point.
//
// The first result will not necessarilly be nil if the value doesn't exist.
// The second result should be checked to determine this information.
func (d *decoder) GetOk(key string) (interface{}, bool) {
	if d.address == "" {
		return d.parent.GetOk(key)
	}
	return d.parent.GetOk(d.address + "." + key)
}

func (d *decoder) GetOkExists(key string) (interface{}, bool) {
	if d.address == "" {
		return d.parent.GetOkExists(key)
	}
	return d.parent.GetOkExists(d.address + "." + key)
}

func (d *decoder) GetChange(key string) (interface{}, interface{}) {
	if d.address == "" {
		return d.parent.GetChange(key)
	}
	return d.parent.GetChange(d.address + "." + key)
}

// Get returns the data for the given key, or nil if the key doesn't exist
// in the schema.
//
// If the key does exist in the schema but doesn't exist in the configuration,
// then the default value for that type will be returned. For strings, this is
// "", for numbers it is 0, etc.
//
// If you want to test if something is set at all in the configuration,
// use GetOk.
func (d *decoder) Get(key string) interface{} {
	if d.address == "" {
		return d.parent.Get(key)
	}
	return d.parent.Get(d.address + "." + key)
}

func VoidDecoder() Decoder {
	return &voidDecoder{}
}

type voidDecoder struct{}

func (d *voidDecoder) DecodeAny(m map[string]interface{}) (interface{}, error) {
	return nil, nil
}

func (vd *voidDecoder) GetOk(key string) (interface{}, bool) {
	return nil, false
}

func (vd *voidDecoder) Get(key string) interface{} {
	return nil
}

func (vd *voidDecoder) GetChange(key string) (interface{}, interface{}) {
	return nil, false
}

func (vd *voidDecoder) GetStringSet(key string) []string {
	return nil
}

func (vd *voidDecoder) GetOkExists(key string) (interface{}, bool) {
	return nil, false
}

func (vd *voidDecoder) Decode(key string, v interface{}) error {
	return nil
}

func (d *voidDecoder) DecodeAll(m map[string]interface{}) error {
	return nil
}

func (vd *voidDecoder) Reader(unkowns ...map[string]json.RawMessage) Reader {
	if len(unkowns) > 0 {
		return NewReader(vd, unkowns[0])
	}
	return NewReader(vd, nil)
}

func (vd *voidDecoder) HasChange(key string) bool {
	return false
}

func (vd *voidDecoder) MarshalAll(items map[string]interface{}) (Properties, error) {
	properties := Properties{}
	if err := properties.MarshalAll(vd, items); err != nil {
		return nil, err
	}
	return properties, nil
}

func (vd *voidDecoder) DecodeSlice(key string, v interface{}) error {
	return nil
}
