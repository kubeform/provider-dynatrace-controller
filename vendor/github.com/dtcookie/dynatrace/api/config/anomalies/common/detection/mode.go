package detection

// Mode How to detect response time degradation: automatically, or based on fixed thresholds, or do not detect.
type Mode string

// Modes offers the known enum values
var Modes = struct {
	DetectAutomatically        Mode
	DetectUsingFixedThresholds Mode
	DontDetect                 Mode
}{
	"DETECT_AUTOMATICALLY",
	"DETECT_USING_FIXED_THRESHOLDS",
	"DONT_DETECT",
}
