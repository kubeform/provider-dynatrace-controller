package paastype

// Value The value to compare to.
type Value string

func (v Value) Ref() *Value {
	return &v
}

func (ptcv *Value) String() string {
	return string(*ptcv)
}

// Values offers the known enum values
var Values = struct {
	AWSECSEC2       Value
	AWSECSFargate   Value
	AWSLambda       Value
	AzureFunctions  Value
	AzureWebsites   Value
	CloudFoundry    Value
	GoogleAppEngine Value
	Heroku          Value
	Kubernetes      Value
	Openshift       Value
}{
	"AWS_ECS_EC2",
	"AWS_ECS_FARGATE",
	"AWS_LAMBDA",
	"AZURE_FUNCTIONS",
	"AZURE_WEBSITES",
	"CLOUD_FOUNDRY",
	"GOOGLE_APP_ENGINE",
	"HEROKU",
	"KUBERNETES",
	"OPENSHIFT",
}
