package aws

// PartitionType The type of the AWS partition.
type PartitionType string

// PartitionTypes offers the known enum values
var PartitionTypes = struct {
	AWSCn      PartitionType
	AWSDefault PartitionType
	AWSUsGov   PartitionType
}{
	"AWS_CN",
	"AWS_DEFAULT",
	"AWS_US_GOV",
}
