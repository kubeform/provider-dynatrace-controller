package api

// StubList is an ordered list of short representations of Dynatrace entities
type EntityRefs struct {
	Values []*EntityRef `json:"values"` // An ordered list of short representations of Dynatrace entities
}
