package browser

// ScriptType The type of monitor. Possible values are `clickpath` for clickpath monitors and `availability` for single-URL browser monitors. These monitors are only allowed to have one event of the `navigate` type
type ScriptType string

// ScriptTypes offers the known enum values
var ScriptTypes = struct {
	ClickPath    ScriptType
	Availability ScriptType
}{
	"clickpath",
	"availability",
}
