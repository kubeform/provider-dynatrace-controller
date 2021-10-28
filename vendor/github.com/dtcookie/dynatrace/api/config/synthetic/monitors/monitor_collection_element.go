package monitors

// MonitorCollectionElement is the short representation of a synthetic monitor
type MonitorCollectionElement struct {
	// "required" : [ "enabled", "entityId", "name", "type" ],
	Name     string `json:"name"`     // The name of a synthetic object
	EntityID string `json:"entityId"` // The ID of a synthetic object
	Type     Type   `json:"type"`     // The type of a synthetic monitor
	Enabled  bool   `json:"enabled"`  // The state of a synthetic monitor
}

// Type The type of a synthetic monitor
// BROWSER: A Browser Monitor
// HTTP: A HTTP Monitor
type Type string

// Types offers the known enum values
var Types = struct {
	Browser Type
	HTTP    Type
}{
	"BROWSER",
	"HTTP",
}
