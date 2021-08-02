package load

// Threshold Minimal service load to detect response time degradation.
//  Response time degradation of services with smaller load won't trigger alerts.
type Threshold string

// Thresholds offers the known enum values
var Thresholds = struct {
	FifteenRequestsPerMinute Threshold
	FiveRequestsPerMinute    Threshold
	OneRequestPerMinute      Threshold
	TenRequestsPerMinute     Threshold
}{
	"FIFTEEN_REQUESTS_PER_MINUTE",
	"FIVE_REQUESTS_PER_MINUTE",
	"ONE_REQUEST_PER_MINUTE",
	"TEN_REQUESTS_PER_MINUTE",
}
