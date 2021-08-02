package terraform

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type address string

func (a address) append(v string) address {
	if len(a) == 0 {
		return address(v)
	}
	return address(fmt.Sprintf("%v.%v", a, v))
}

func (a address) replace(v string) address {
	if len(a) == 0 {
		return address(v)
	}
	s := string(a)
	idx := strings.LastIndex(s, ".")
	if idx == -1 {
		return address(v)
	}
	return address(fmt.Sprintf("%v.%v", s[:idx], v))
}

func (a address) backtrack() (address, string) {
	if len(a) == 0 {
		return a, ""
	}
	s := string(a)
	idx := strings.LastIndex(s, ".")
	if idx == -1 {
		return a, ""
	}
	return address(s[:idx]), s[idx+1:]
}

func (a address) index(v int) address {
	if len(a) == 0 {
		// panic("indexing root key not allowed")
		return a
	}
	return address(fmt.Sprintf("%v.%v", a, v))
}

func toJSON(v interface{}) string {
	var result string
	if bytes, err := json.Marshal(v); err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}
	return result
}

func evalDiscrKey(t reflect.Type) string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if jsonXValue, ok := field.Tag.Lookup("json-x"); ok {
			if strings.Contains(jsonXValue, "discriminator") {
				propertyName := propName(field)
				if len(propertyName) > 0 {
					return propertyName
				}
			}
		}
	}
	return ""
}

// Resolver has no documentation
type Resolver interface {
	Resolve(t reflect.Type) (interface{}, error)
	resolve(hint address, t reflect.Type, sliced bool) (interface{}, error)
	GetOk(key address) (interface{}, bool)
	Count(key address) (int, bool)
	HasCount(key address) bool
}

type resolver struct {
	ResourceData ResourceData
}

// NewResolver has no documentation
func NewResolver(source interface{}) (Resolver, error) {
	var resourceData ResourceData
	var err error
	if resourceData, err = NewResourceData(source); err != nil {
		return nil, err
	}
	return &resolver{ResourceData: resourceData}, nil
}

// GetOk has no documentation
func (res *resolver) GetOk(key address) (interface{}, bool) {
	return res.ResourceData.GetOk(fmt.Sprintf("%v", key))
}

// GetString has no documentation
func (res *resolver) GetString(key address) (string, bool) {
	if result, ok := res.ResourceData.GetOk(fmt.Sprintf("%v", key)); ok {
		if strResult, ok := result.(string); ok {
			return strResult, ok
		}
		panic(fmt.Errorf("key '%v' expected to be string, but was '%T'", key, result))
	}
	return "", false
}

// Count has no documentation
func (res *resolver) Count(key address) (int, bool) {
	if cnt, ok := res.ResourceData.GetOk(fmt.Sprintf("%v.#", key)); ok {
		return cnt.(int), ok
	}
	return 0, false
}

// HasCount has no documentation
func (res *resolver) HasCount(key address) bool {
	if _, ok := res.ResourceData.GetOk(fmt.Sprintf("%v.#", key)); ok {
		return ok
	}
	return false
}

// Resolve has no documentation
func (res *resolver) Resolve(t reflect.Type) (interface{}, error) {
	return res.resolve("", t, false)
}

func (res *resolver) resolve(hint address, t reflect.Type, sliced bool) (interface{}, error) {
	if t == nil {
		return nil, errors.New("cannot resolve objects of nil type")
	}
	switch t.Kind() {
	case reflect.Ptr:
		var result interface{}
		var err error
		if result, err = res.resolve(hint, t.Elem(), sliced); err != nil {
			return nil, err
		}
		if result == nil {
			return nil, nil
		}

		vPtr := reflect.New(t.Elem())
		vPtr.Elem().Set(reflect.ValueOf(result).Convert(vPtr.Type().Elem()))
		return vPtr.Interface(), nil
	case reflect.String:
		if result, ok := res.GetOk(hint); ok {
			return result.(string), nil
		}
		return nil, nil
	case reflect.Bool:
		if result, ok := res.GetOk(hint); ok {
			return result.(bool), nil
		}
		return nil, nil
	case reflect.Uint:
		if result, ok := res.GetOk(hint); ok {
			return uint(result.(int)), nil
		}
		return nil, nil
	case reflect.Uint8:
		if result, ok := res.GetOk(hint); ok {
			return uint8(result.(int)), nil
		}
		return nil, nil
	case reflect.Uint16:
		if result, ok := res.GetOk(hint); ok {
			return uint16(result.(int)), nil
		}
		return nil, nil
	case reflect.Uint32:
		if result, ok := res.GetOk(hint); ok {
			return uint32(result.(int)), nil
		}
		return nil, nil
	case reflect.Uint64:
		if result, ok := res.GetOk(hint); ok {
			return uint64(result.(int)), nil
		}
		return nil, nil
	case reflect.Int:
		if result, ok := res.GetOk(hint); ok {
			return result.(int), nil
		}
		return nil, nil
	case reflect.Int8:
		if result, ok := res.GetOk(hint); ok {
			return int8(result.(int)), nil
		}
		return nil, nil
	case reflect.Int16:
		if result, ok := res.GetOk(hint); ok {
			return int16(result.(int)), nil
		}
		return nil, nil
	case reflect.Int32:
		if result, ok := res.GetOk(hint); ok {
			return int32(result.(int)), nil
		}
		return nil, nil
	case reflect.Int64:
		if result, ok := res.GetOk(hint); ok {
			return int64(result.(int)), nil
		}
		return nil, nil
	case reflect.Float64:
		if result, ok := res.GetOk(hint); ok {
			return result.(float64), nil
		}
		return nil, nil
	case reflect.Float32:
		if result, ok := res.GetOk(hint); ok {
			return float32(result.(float64)), nil
		}
		return nil, nil
	case reflect.Slice:
		if isPrimitiveType(t.Elem()) {
			if untypedResult, ok := res.GetOk(hint); ok || untypedResult != nil {
				switch t.Elem().Kind() {
				case reflect.String:
					return toStringSlice(untypedResult, t), nil
				case reflect.Bool:
					return toBoolSlice(untypedResult), nil
				case reflect.Int:
					return toIntSlice(untypedResult), nil
				case reflect.Int8:
					return toInt8Slice(untypedResult), nil
				case reflect.Int16:
					return toInt16Slice(untypedResult), nil
				case reflect.Int32:
					return toInt32Slice(untypedResult), nil
				case reflect.Int64:
					return toInt64Slice(untypedResult), nil
				case reflect.Uint:
					return toUIntSlice(untypedResult), nil
				case reflect.Uint8:
					return toUInt8Slice(untypedResult), nil
				case reflect.Uint16:
					return toUInt16Slice(untypedResult), nil
				case reflect.Uint32:
					return toUInt32Slice(untypedResult), nil
				case reflect.Uint64:
					return toUInt64Slice(untypedResult), nil
				case reflect.Float32:
					return toFloat32Slice(untypedResult), nil
				case reflect.Float64:
					return toFloat64Slice(untypedResult), nil
				}
			}
			return nil, nil
		}
		result := reflect.MakeSlice(t, 0, 0)
		if t.Elem().Kind() == reflect.Interface {
			baseType := base(t.Elem())
			for _, implementorType := range implementors(baseType, t.Elem()) {
				matchHint := hint.replace(unCamel(unref(implementorType).Name()))
				var implementorResults interface{}
				var err error
				if implementorResults, err = res.resolve(matchHint, reflect.SliceOf(implementorType), false); err != nil {
					return nil, err
				}
				vImplementorResults := reflect.ValueOf(implementorResults)
				for i := 0; i < vImplementorResults.Len(); i++ {
					result = reflect.Append(result, vImplementorResults.Index(i))
				}
			}
		}

		var err error
		if cnt, ok := res.Count(hint); ok {
			for i := 0; i < cnt; i++ {
				var entry interface{}
				if entry, err = res.resolve(hint.index(i), t.Elem(), true); err != nil {
					return nil, err
				}
				if entry != nil {
					result = reflect.Append(result, reflect.ValueOf(entry))
				}
			}
		}
		return result.Interface(), nil
	case reflect.Map:
		result := reflect.MakeMap(t)
		var err error
		if cnt, ok := res.Count(hint); ok {
			for i := 0; i < cnt; i++ {
				var key string
				var ok bool
				if key, ok = res.GetString(hint.index(i).append("key")); !ok {
					return nil, errors.New("expected an attribute 'key' but didn't find it")
				}
				var entry interface{}
				if unref(t.Elem()).Kind() == reflect.Struct {
					if entry, err = res.resolve(hint.index(i), t.Elem(), true); err != nil {
						return nil, err
					}
				} else {
					if entry, err = res.resolve(hint.index(i).append("values"), t.Elem(), true); err != nil {
						return nil, err
					}
				}
				if entry != nil {
					result.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(entry))
				}
			}
		}
		return result.Interface(), nil
	case reflect.Struct:
		origHint := hint
		if !sliced && (len(hint) > 0) {
			var cnt int
			var ok bool
			if cnt, ok = res.Count(hint); !ok {
				return nil, nil
			}
			hint = hint.index(0)
			if cnt == 1 {
				if _, ok := res.GetOk(hint); !ok {
					return reflect.New(t).Elem().Interface(), nil
				}
			}
		}
		if _, ok := res.GetOk(hint); !ok {
			return nil, nil
		}
		vResult := reflect.New(t).Elem()
		for i := 0; i < vResult.NumField(); i++ {
			vField := vResult.Field(i)
			field := t.Field(i)
			if field.Anonymous {
				var err error
				var member interface{}
				if member, err = res.resolve(origHint, field.Type, sliced); err != nil {
					return nil, err
				}
				if member != nil {
					vField.Set(reflect.ValueOf(member))
				}
			} else {
				if propertyName := propName(field); len(propertyName) > 0 {
					var propertyValue interface{}
					var err error
					if propertyValue, err = res.resolve(hint.append(propertyName), field.Type, false); err != nil {
						return nil, err
					}
					if propertyValue != nil {
						vField.Set(reflect.ValueOf(propertyValue).Convert(vField.Type()))
					} else if !optional(field) {
						vField.Set(reflect.ValueOf(reflect.New(vField.Type()).Elem().Interface()))
					}
				}
			}
		}
		return vResult.Interface(), nil
	case reflect.Interface:
		baseType := base(t)
		for _, implementorType := range implementors(baseType, t) {
			matchHint := hint
			matches := res.matches(hint.index(0), implementorType)
			if !matches {
				mHint := hint.replace(unCamel(unref(implementorType).Name()))
				// log.Println("matching against address", mHint.index(0), "for type", implementorType)
				if matches = res.matches(mHint.index(0), implementorType); matches {
					// log.Println("  ", "matched")
					matchHint = mHint
				}
			}
			if matches {
				var result interface{}
				var err error
				if result, err = res.resolve(matchHint, implementorType.Elem(), sliced); err != nil {
					return nil, err
				}
				if result == nil {
					return nil, nil
				}
				vPtr := reflect.New(implementorType.Elem())
				vPtr.Elem().Set(reflect.ValueOf(result))
				return vPtr.Interface(), nil
			}
		}
		var result interface{}
		var err error
		if result, err = res.resolve(hint, baseType.Elem(), sliced); err != nil {
			return nil, err
		}
		if result == nil {
			if result, err = res.resolve(hint.replace(unCamel(unref(baseType.Elem()).Name())), baseType.Elem(), sliced); err != nil {
				return nil, err
			}
			if result == nil {
				return nil, nil
			}
		}
		vPtr := reflect.New(baseType.Elem())
		vPtr.Elem().Set(reflect.ValueOf(result))
		return vPtr.Interface(), nil
	default:
		return nil, fmt.Errorf("[Resolve] unsupported type %v (kind: %v)", t, t.Kind())
	}
}

func (res *resolver) matches(addr address, t reflect.Type) bool {
	debug := string(addr) == "rules.0.conditions.0.service_type_comparison"
	debug = false
	t = unref(t)
	discrKey := ""
	discrValues := []string{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.Anonymous {
			continue
		}
		if jsonXValue, ok := field.Tag.Lookup("json"); ok {
			jsonXValue = strings.TrimSpace(jsonXValue)
			if len(jsonXValue) == 0 {
				continue
			}
			if !strings.HasPrefix(jsonXValue, ",") {
				continue
			}
			jsonXValue = jsonXValue[1:]
			if strings.Contains(jsonXValue, "=") {
				parts := strings.Split(jsonXValue, "=")
				discrKey = strings.TrimSpace(parts[0])
				discrValues = strings.Split(strings.TrimSpace(parts[1]), "|")
			} else {
				discrValues = strings.Split(jsonXValue, "|")
			}
			if len(discrKey) == 0 {
				discrKey = evalDiscrKey(field.Type)
			}
			if debug {
				log.Println("discrValues", discrValues)
				log.Println("discrKey", discrKey)
			}
			if (len(discrKey) > 0) && (len(discrValues) > 0) {
				var discrValueFound interface{}
				var sDiscrValue string
				var ok bool
				appAddr := addr.append(discrKey)
				if debug {
					appAddr = addr.index(0).append(discrKey)
					log.Println("res.GetOk", appAddr)
				}
				if discrValueFound, ok = res.GetOk(appAddr); ok {
					if debug {
						log.Println("  discrValueFound", discrValueFound)
					}
					if sDiscrValue, ok = discrValueFound.(string); ok {
						for _, discrValue := range discrValues {
							if sDiscrValue == strings.TrimSpace(discrValue) {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}
