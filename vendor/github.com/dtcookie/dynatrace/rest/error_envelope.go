package rest

import "encoding/json"

// ErrorEnvelope represents error messages returned from the REST API
type ErrorEnvelope struct {
	Error Error `json:"error"`
}

// Error represents the error description and details of an error returned by the REST API
type Error struct {
	Code                 int32                 `json:"code"`
	Message              string                `json:"message"`
	ConstraintViolations []ConstraintViolation `json:"constraintViolations"`
}

func (e *Error) Error() string {
	if data, err := json.MarshalIndent(e, "", "  "); err == nil {
		return string(data)
	}
	return "no error message available"
}

// ConstraintViolation holds the details of a constraint violation
type ConstraintViolation struct {
	ParameterLocation string `json:"parameterLocation"`
	Location          string `json:"location"`
	Message           string `json:"message"`
	Path              string `json:"path"`
}
