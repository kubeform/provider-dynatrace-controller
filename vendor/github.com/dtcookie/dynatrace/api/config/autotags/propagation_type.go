package autotags

// PropagationType has no documentation
type PropagationType string

// PropagationTypes offers the known enum values
var PropagationTypes = struct {
	AzureToPg                  PropagationType
	AzureToService             PropagationType
	HostToProcessGroupInstance PropagationType
	ProcessGroupToHost         PropagationType
	ProcessGroupToService      PropagationType
	ServiceToHostLike          PropagationType
	ServiceToProcessGroupLike  PropagationType
}{
	"AZURE_TO_PG",
	"AZURE_TO_SERVICE",
	"HOST_TO_PROCESS_GROUP_INSTANCE",
	"PROCESS_GROUP_TO_HOST",
	"PROCESS_GROUP_TO_SERVICE",
	"SERVICE_TO_HOST_LIKE",
	"SERVICE_TO_PROCESS_GROUP_LIKE",
}
