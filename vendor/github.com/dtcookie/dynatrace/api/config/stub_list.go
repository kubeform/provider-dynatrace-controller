package api

// StubList is an ordered list of short representations of Dynatrace entities
type StubList struct {
	Values []*EntityShortRepresentation `json:"values"` // An ordered list of short representations of Dynatrace entities
}

// String provides a string representation in JSON format for debugging purposes
func (sl StubList) String() string {
	return MarshalJSON(&sl)
}
