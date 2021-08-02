package custom_application_type

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

func (catcv *Value) String() string {
	return string(*catcv)
}

// Values offers the known enum values
var Values = struct {
	AmazonEcho        Value
	Desktop           Value
	Embedded          Value
	Iot               Value
	MicrosoftHololens Value
	Ufo               Value
}{
	"AMAZON_ECHO",
	"DESKTOP",
	"EMBEDDED",
	"IOT",
	"MICROSOFT_HOLOLENS",
	"UFO",
}
