package gojson

import (
	"reflect"
	"strings"
)

// evalTag inspects the given StructField regarding JSON specific tagging
// If no tag is specified it returns the field name as the property name
// The second return value tells whether that property is optional or not
func evalTag(field reflect.StructField) (string, bool) {
	name := field.Name
	var tagValue string
	var found bool
	if tagValue, found = field.Tag.Lookup("json"); !found {
		return name, false
	}
	tagValue = strings.TrimSpace(tagValue)
	if empty(tagValue) {
		return name, false
	}
	if !strings.Contains(tagValue, ",") {
		return tagValue, false
	}
	parts := strings.Split(tagValue, ",")
	tagName := strings.TrimSpace(parts[0])
	if !empty(tagName) {
		name = tagName
	}
	omitEmpty := false
	for i := 1; i < len(parts); i++ {
		part := strings.TrimSpace(parts[i])
		if part == "omitempty" {
			omitEmpty = true
		}
	}

	return name, omitEmpty
}
