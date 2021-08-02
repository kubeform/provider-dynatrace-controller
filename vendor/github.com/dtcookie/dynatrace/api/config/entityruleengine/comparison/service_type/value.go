package service_type

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
	BackgroundActivity          Value
	CICSService                 Value
	CustomService               Value
	DatabaseService             Value
	EnterpriseServiceBusService Value
	External                    Value
	IBMIntegrationBusService    Value
	IMSService                  Value
	MessagingService            Value
	QueueListenerService        Value
	RMIService                  Value
	RPCService                  Value
	WebRequestService           Value
	WebService                  Value
}{
	"BACKGROUND_ACTIVITY",
	"CICS_SERVICE",
	"CUSTOM_SERVICE",
	"DATABASE_SERVICE",
	"ENTERPRISE_SERVICE_BUS_SERVICE",
	"EXTERNAL",
	"IBM_INTEGRATION_BUS_SERVICE",
	"IMS_SERVICE",
	"MESSAGING_SERVICE",
	"QUEUE_LISTENER_SERVICE",
	"RMI_SERVICE",
	"RPC_SERVICE",
	"WEB_REQUEST_SERVICE",
	"WEB_SERVICE",
}
