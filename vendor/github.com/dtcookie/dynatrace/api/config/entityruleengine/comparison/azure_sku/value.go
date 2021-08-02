package azure_sku

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
	Basic    Value
	Dynamic  Value
	Free     Value
	Premium  Value
	Shared   Value
	Standard Value
}{
	"BASIC",
	"DYNAMIC",
	"FREE",
	"PREMIUM",
	"SHARED",
	"STANDARD",
}
