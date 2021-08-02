package aws

// Statistic The statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM
type Statistic string

// Statistics offers the known enum values
var Statistics = struct {
	Average     Statistic
	AvgMinMax   Statistic
	Maximum     Statistic
	Minimum     Statistic
	SampleCount Statistic
	Sum         Statistic
}{
	"AVERAGE",
	"AVG_MIN_MAX",
	"MAXIMUM",
	"MINIMUM",
	"SAMPLE_COUNT",
	"SUM",
}
