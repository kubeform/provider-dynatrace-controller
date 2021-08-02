package ostype

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
	AIX     Value
	Darwin  Value
	Hpux    Value
	Linux   Value
	Solaris Value
	Windows Value
	Zos     Value
}{
	"AIX",
	"DARWIN",
	"HPUX",
	"LINUX",
	"SOLARIS",
	"WINDOWS",
	"ZOS",
}
