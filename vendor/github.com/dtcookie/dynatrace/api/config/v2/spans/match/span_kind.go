package match

type SpanKind string

func (me SpanKind) Ref() *SpanKind {
	return &me
}

var SpanKinds = struct {
	Internal SpanKind
	Server   SpanKind
	Client   SpanKind
	Producer SpanKind
	Consumer SpanKind
}{
	Internal: SpanKind("INTERNAL"),
	Server:   SpanKind("SERVER"),
	Client:   SpanKind("CLIENT"),
	Producer: SpanKind("PRODUCER"),
	Consumer: SpanKind("CONSUMER"),
}
