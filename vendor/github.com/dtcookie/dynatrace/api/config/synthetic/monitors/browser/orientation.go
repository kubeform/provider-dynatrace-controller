package browser

// Orientation The orientation of the device— `portrait` or `landscape`
type Orientation string

// Orientations offers the known enum values
var Orientations = struct {
	Portrait  Orientation
	Landscape Orientation
}{
	`portrait`,
	`landscape`,
}
