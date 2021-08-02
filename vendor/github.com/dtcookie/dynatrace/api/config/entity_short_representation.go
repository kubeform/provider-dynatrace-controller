package api

// EntityShortRepresentation is the short representation of a Dynatrace entity
type EntityShortRepresentation struct {
	ID          string `json:"id"`                    // the ID of the Dynatrace entity
	Name        string `json:"name,omitempty"`        // the name of the Dynatrace entity
	Description string `json:"description,omitempty"` // a short description of the Dynatrace entity
}

// String provides a string representation in JSON format for debugging purposes
func (esr EntityShortRepresentation) String() string {
	return MarshalJSON(&esr)
}
