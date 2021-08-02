package terraform

import (
	"reflect"
	"strings"
)

func isPrimitive(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Map:
		return false
	case reflect.Ptr:
		return isPrimitive(*(v.(*interface{})))
	default:
		return true
	}
}

func toInt64Slice(v interface{}) []int64 {
	result := []int64{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, int64(elem.(int)))
	}
	return result
}

func toInt32Slice(v interface{}) []int32 {
	result := []int32{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, int32(elem.(int)))
	}
	return result
}
func toIntSlice(v interface{}) []int {
	result := []int{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, elem.(int))
	}
	return result
}

func toUIntSlice(v interface{}) []uint {
	result := []uint{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, uint(elem.(int)))
	}
	return result
}
func toInt8Slice(v interface{}) []int8 {
	result := []int8{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, int8(elem.(int)))
	}
	return result
}

func toInt16Slice(v interface{}) []int16 {
	result := []int16{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, int16(elem.(int)))
	}
	return result
}

func toUInt64Slice(v interface{}) []uint64 {
	result := []uint64{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, uint64(elem.(int)))
	}
	return result
}

func toUInt32Slice(v interface{}) []uint32 {
	result := []uint32{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, uint32(elem.(int)))
	}
	return result
}

func toUInt8Slice(v interface{}) []uint8 {
	result := []uint8{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, uint8(elem.(int)))
	}
	return result
}

func toUInt16Slice(v interface{}) []uint16 {
	result := []uint16{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, uint16(elem.(int)))
	}
	return result
}

func toFloat32Slice(v interface{}) []float32 {
	result := []float32{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, elem.(float32))
	}
	return result
}

func toFloat64Slice(v interface{}) []float64 {
	result := []float64{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, elem.(float64))
	}
	return result
}

func toBoolSlice(v interface{}) []bool {
	result := []bool{}
	slice := v.([]interface{})
	for _, elem := range slice {
		result = append(result, elem.(bool))
	}
	return result
}

func toStringSlice(v interface{}, sliceType reflect.Type) interface{} {
	entries := v.([]interface{})
	vSlice := reflect.MakeSlice(sliceType, len(entries), len(entries))
	for idx, entry := range entries {
		vEntry := vSlice.Index(idx)
		vEntry.Set(reflect.ValueOf(entry).Convert(sliceType.Elem()))
	}
	return vSlice.Interface()
}

func startsWithUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	c := s[0:1]
	return c == strings.ToUpper(c)
}

func isPrimitiveType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Ptr:
		return isPrimitiveType(t.Elem())
	case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func unref(t reflect.Type) reflect.Type {
	if t == nil {
		return nil
	}
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
