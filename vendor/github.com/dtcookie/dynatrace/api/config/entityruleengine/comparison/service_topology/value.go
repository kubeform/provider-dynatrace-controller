package service_topology

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

func (v *Value) String() string {
	return string(*v)
}

// Values offers the known enum values
var Values = struct {
	ExternalService Value
	FullyMonitored  Value
	OpaqueService   Value
}{
	"EXTERNAL_SERVICE",
	"FULLY_MONITORED",
	"OPAQUE_SERVICE",
}
