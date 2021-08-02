package osarch

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
	Arm    Value
	Ia64   Value
	Parisc Value
	Ppc    Value
	Ppcle  Value
	S390   Value
	Sparc  Value
	X86    Value
	Zos    Value
}{
	"ARM",
	"IA64",
	"PARISC",
	"PPC",
	"PPCLE",
	"S390",
	"SPARC",
	"X86",
	"ZOS",
}
