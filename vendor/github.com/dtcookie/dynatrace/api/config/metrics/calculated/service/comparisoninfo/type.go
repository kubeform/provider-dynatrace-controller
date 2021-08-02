package comparisoninfo

// Type Defines the actual set of fields depending on the value. See one of the following objects:
// * `STRING` -> StringComparisonInfo
// * `NUMBER` -> NumberComparisonInfo
// * `BOOLEAN` -> BooleanComparisonInfo
// * `HTTP_METHOD` -> HttpMethodComparisonInfo
// * `STRING_REQUEST_ATTRIBUTE` -> StringRequestAttributeComparisonInfo
// * `NUMBER_REQUEST_ATTRIBUTE` -> NumberRequestAttributeComparisonInfo
// * `ZOS_CALL_TYPE` -> ZosComparisonInfo
// * `IIB_INPUT_NODE_TYPE` -> IIBInputNodeTypeComparisonInfo
// * `ESB_INPUT_NODE_TYPE` -> ESBInputNodeTypeComparisonInfo
// * `FAILED_STATE` -> FailedStateComparisonInfo
// * `FLAW_STATE` -> FlawStateComparisonInfo
// * `FAILURE_REASON` -> FailureReasonComparisonInfo
// * `HTTP_STATUS_CLASS` -> HttpStatusClassComparisonInfo
// * `TAG` -> TagComparisonInfo
// * `FAST_STRING` -> FastStringComparisonInfo
// * `SERVICE_TYPE` -> ServiceTypeComparisonInfo
type Type string

// Types offers the known enum values
var Types = struct {
	Boolean                Type
	ESBInputNodeType       Type
	FailedState            Type
	FailureReason          Type
	FastString             Type
	FlawState              Type
	HTTPMethod             Type
	HTTPStatusClass        Type
	IIBInputNodeType       Type
	Number                 Type
	NumberRequestAttribute Type
	ServiceType            Type
	String                 Type
	StringRequestAttribute Type
	Tag                    Type
	ZosCallType            Type
}{
	"BOOLEAN",
	"ESB_INPUT_NODE_TYPE",
	"FAILED_STATE",
	"FAILURE_REASON",
	"FAST_STRING",
	"FLAW_STATE",
	"HTTP_METHOD",
	"HTTP_STATUS_CLASS",
	"IIB_INPUT_NODE_TYPE",
	"NUMBER",
	"NUMBER_REQUEST_ATTRIBUTE",
	"SERVICE_TYPE",
	"STRING",
	"STRING_REQUEST_ATTRIBUTE",
	"TAG",
	"ZOS_CALL_TYPE",
}
