package match

type Source string

var Sources = struct {
	SpanName                      Source
	SpanKind                      Source
	Attribute                     Source
	InstrumentationLibraryName    Source
	InstrumentationLibraryVersion Source
}{
	SpanName:                      Source("SPAN_NAME"),
	SpanKind:                      Source("SPAN_KIND"),
	Attribute:                     Source("ATTRIBUTE"),
	InstrumentationLibraryName:    Source("INSTRUMENTATION_LIBRARY_NAME"),
	InstrumentationLibraryVersion: Source("INSTRUMENTATION_LIBRARY_VERSION"),
}
