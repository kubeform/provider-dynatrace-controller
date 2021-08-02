package application_type

// Value The value to compare to.
type Value string

func (atcv *Value) String() string {
	return string(*atcv)
}

func (atcv Value) Ref() *Value {
	return &atcv
}

// Values offers the known enum values
var Values = struct {
	AgentlessMonitoring Value
	Amp                 Value
	AutoInjected        Value
	Default             Value
	SaasVendor          Value
}{
	"AGENTLESS_MONITORING",
	"AMP",
	"AUTO_INJECTED",
	"DEFAULT",
	"SAAS_VENDOR",
}
