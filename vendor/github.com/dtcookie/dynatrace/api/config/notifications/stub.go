package notifications

// Stub The short representation of a notification.
type Stub struct {
	Description *string `json:"description,omitempty"` // A short description of the Dynatrace entity.
	ID          string  `json:"id"`                    // The ID of the Dynatrace entity.
	Name        *string `json:"name,omitempty"`        // The name of the Dynatrace entity.
	Type        *Type   `json:"type,omitempty"`        // The type of the notification.
}

// StubList has no documentation
type StubList struct {
	Values []Stub `json:"values,omitempty"` // has no documentation
}
