package entrypoints

type SpanEntrypointAction string

var SpanEntryPointActions = struct {
	CreateEntrypoint     SpanEntrypointAction
	DontCreateEntrypoint SpanEntrypointAction
}{
	CreateEntrypoint:     SpanEntrypointAction("CREATE_ENTRYPOINT"),
	DontCreateEntrypoint: SpanEntrypointAction("DONT_CREATE_ENTRYPOINT"),
}
