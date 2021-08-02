package cloud_type

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
	Azure               Value
	EC2                 Value
	GoogleCloudPlatform Value
	OpenStack           Value
	Oracle              Value
	Unrecognized        Value
}{
	"AZURE",
	"EC2",
	"GOOGLE_CLOUD_PLATFORM",
	"OPENSTACK",
	"ORACLE",
	"UNRECOGNIZED",
}
