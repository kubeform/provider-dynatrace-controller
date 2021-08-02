package azure_compute_mode

// Value The value to compare to.
type Value string

func (v *Value) String() string {
	return string(*v)
}

func (v Value) Ref() *Value {
	return &v
}

// Values offers the known enum values
var Values = struct {
	Dedicated Value
	Shared    Value
}{
	"DEDICATED",
	"SHARED",
}
