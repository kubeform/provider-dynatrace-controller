package api

import "encoding/json"

// MarshalJSON provides a string representation in JSON format for debugging purposes
func MarshalJSON(v interface{}) string {
	var result string
	if bytes, err := json.MarshalIndent(v, "", "  "); err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}
	return result
}
