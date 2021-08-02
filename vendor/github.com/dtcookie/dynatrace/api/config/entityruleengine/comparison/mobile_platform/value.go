package mobile_platform

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

func (mpcv *Value) String() string {
	return string(*mpcv)
}

// Values offers the known enum values
var Values = struct {
	Android Value
	Ios     Value
	Linux   Value
	MacOS   Value
	Other   Value
	Tvos    Value
	Windows Value
}{
	"ANDROID",
	"IOS",
	"LINUX",
	"MAC_OS",
	"OTHER",
	"TVOS",
	"WINDOWS",
}
