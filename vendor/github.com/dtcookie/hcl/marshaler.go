package hcl

// Marshaler has no documentation
type Marshaler interface {
	MarshalHCL() (map[string]interface{}, error)
}

type ExtMarshaler interface {
	MarshalHCL(Decoder) (map[string]interface{}, error)
}

type Unmarshaler interface {
	UnmarshalHCL(decoder Decoder) error
}
