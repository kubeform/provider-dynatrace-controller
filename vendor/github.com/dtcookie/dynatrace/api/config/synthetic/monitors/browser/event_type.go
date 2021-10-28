package browser

// EventType specifies the type of authentication. `basic` or `webform`
type EventType string

// EventTypes offers the known enum values
var EventTypes = struct {
	KeyStrokes   EventType
	Cookie       EventType
	SelectOption EventType
	Javascript   EventType
	Click        EventType
	Tap          EventType
	Navigate     EventType
}{
	"keystrokes",
	"cookie",
	"selectOption",
	"javascript",
	"click",
	"tap",
	"navigate",
}
