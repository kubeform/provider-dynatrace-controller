package bitness

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

// Values offers the known enum values
var Values = struct {
	V32 Value
	V64 Value
}{
	"32",
	"64",
}

func (bcv *Value) String() string {
	return string(*bcv)
}
