package capture

import (
	"github.com/dtcookie/hcl"
)

// SpanCaptureSetting OpenTelemetry/OpenTracing spans can start new PurePaths. Define rules that define which spans should not be considered as entry points.\n\nNote: This config does not apply to Trace ingest
type SpanCaptureSetting struct {
	SpanCaptureRule *SpanCaptureRule `json:"spanCaptureRule"`
}

func (me *SpanCaptureSetting) Schema() map[string]*hcl.Schema {
	return new(SpanCaptureRule).Schema()
}

func (me *SpanCaptureSetting) MarshalHCL() (map[string]interface{}, error) {
	return me.SpanCaptureRule.MarshalHCL()
}

func (me *SpanCaptureSetting) UnmarshalHCL(decoder hcl.Decoder) error {
	me.SpanCaptureRule = new(SpanCaptureRule)
	return me.SpanCaptureRule.UnmarshalHCL(decoder)
}
