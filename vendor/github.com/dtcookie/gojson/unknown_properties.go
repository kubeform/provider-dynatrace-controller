package gojson

import "reflect"

// UnknownProperties represents the type of a Struct Field that will be used
// for any properties within JSON code that don't match any of the Fields found
// within the Struct.
// When marshalling a Struct all of the contents of this map will get added to
// the JSON representation.
type UnknownProperties map[string]interface{}

var tUnknownProperties = reflect.TypeOf(UnknownProperties{})
