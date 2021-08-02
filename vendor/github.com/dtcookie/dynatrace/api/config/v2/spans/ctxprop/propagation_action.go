package ctxprop

type PropagationAction string

var PropagationActions = struct {
	Propagate     PropagationAction
	DontPropagate PropagationAction
}{
	Propagate:     PropagationAction("PROPAGATE"),
	DontPropagate: PropagationAction("DONT_PROPAGATE"),
}
