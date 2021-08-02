package entrypoints

import (
	"github.com/dtcookie/hcl"
)

// SpanEntryPoint OpenTelemetry/OpenTracing spans can start new PurePaths. Define rules that define which spans should not be considered as entry points.\n\nNote: This config does not apply to Trace ingest
type SpanEntryPoint struct {
	EntryPointRule *SpanEntrypointRule `json:"entryPointRule"`
}

func (me *SpanEntryPoint) Schema() map[string]*hcl.Schema {
	return new(SpanEntrypointRule).Schema()
}

func (me *SpanEntryPoint) MarshalHCL() (map[string]interface{}, error) {
	return me.EntryPointRule.MarshalHCL()
}

func (me *SpanEntryPoint) UnmarshalHCL(decoder hcl.Decoder) error {
	me.EntryPointRule = new(SpanEntrypointRule)
	return me.EntryPointRule.UnmarshalHCL(decoder)
}
