package terraform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"unicode"
)

const indentStr = "\t"

func jenc(s interface{}) string {
	if reflect.TypeOf(s).Kind() != reflect.String {
		panic("string expected")
	}
	if bytes, err := json.Marshal(s); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}

func terraformat(s string) string {
	result := ""
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' {
			result = result + string(r)
		} else {
			result = result + "_"
		}
	}
	return result
}

func tfPrint(buf *bytes.Buffer, key string, v interface{}, indent string) {
	if v == nil {
		panic(fmt.Errorf("%v is nil", key))
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr:
		tfPrint(buf, key, *(v.(*interface{})), indent)
	case reflect.Slice:
		rv := reflect.ValueOf(v)
		if rv.Len() == 0 {
			// fmt.Printf("Slice Type: %v\n", rv.Type())
			if rv.Type() != reflect.TypeOf(ResourceSlice{}) {
				if len(key) > 0 {
					fmt.Fprintln(buf, fmt.Sprintf(`%s%s = []`, indent, key))
				} else {
					fmt.Fprintln(buf, fmt.Sprintf(`%s[]`, indent))
				}
			}
		} else {
			if isPrimitive(rv.Index(0).Interface()) {
				if len(key) > 0 {
					fmt.Fprint(buf, fmt.Sprintf(`%s%s = [`, indent, key))
				} else {
					fmt.Fprint(buf, fmt.Sprintf(`%s[`, indent))
				}
				sep := " "
				for i := 0; i < rv.Len(); i++ {
					value := rv.Index(i)
					if reflect.TypeOf(value.Interface()).Kind() == reflect.String {
						fmt.Fprint(buf, fmt.Sprintf(`%s"%v"`, sep, value.Interface()))
					} else {
						fmt.Fprint(buf, fmt.Sprintf(`%s%v`, sep, value.Interface()))
					}
					sep = ", "
				}
				fmt.Fprintln(buf, ` ]`)
			} else {
				for i := 0; i < rv.Len(); i++ {
					tfPrint(buf, key, rv.Index(i).Interface(), indent)
				}
			}
		}
	case reflect.Map:
		rv := v.(OrderedMap)
		if len(key) > 0 {
			fmt.Fprintln(buf, fmt.Sprintf(`%s%s {`, indent, key))
		} else {
			fmt.Fprintln(buf, fmt.Sprintf(`%s{`, indent))
		}

		for _, jk := range rv.OrderedKeys() {
			tfPrint(buf, jk.Value, rv[jk], indent+indentStr)
		}
		fmt.Fprintln(buf, fmt.Sprintf(`%s}`, indent))
	case reflect.String:
		if len(key) > 0 {

			fmt.Fprintln(buf, fmt.Sprintf(`%s%s = %v`, indent, key, jenc(v)))
		} else {
			fmt.Fprintln(buf, fmt.Sprintf(`%s%s%v`, indent, key, jenc(v)))
		}
	default:
		if len(key) > 0 {
			fmt.Fprintln(buf, fmt.Sprintf(`%s%s = %v`, indent, key, v))
		} else {
			fmt.Fprintln(buf, fmt.Sprintf(`%s%v`, indent, v))
		}
	}
}
