package auth

// Type specifies the type of authentication. `basic` or `webform`
type Type string

// Types offers the known enum values
var Types = struct {
	Basic   Type
	WebForm Type
}{
	"basic",
	"webform",
}
