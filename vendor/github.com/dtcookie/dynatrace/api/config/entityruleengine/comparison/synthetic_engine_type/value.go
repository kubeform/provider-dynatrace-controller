package synthetic_engine_type

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
	Classic Value
	Custom  Value
}{
	"CLASSIC",
	"CUSTOM",
}
