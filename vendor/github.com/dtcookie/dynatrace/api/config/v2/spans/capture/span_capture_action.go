package capture

type SpanCaptureAction string

var SpanEntryPointActions = struct {
	Capture SpanCaptureAction
	Ignore  SpanCaptureAction
}{
	Capture: SpanCaptureAction("CAPTURE"),
	Ignore:  SpanCaptureAction("IGNORE"),
}
