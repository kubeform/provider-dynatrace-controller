package aws

// Type The type of the authentication: role-based or key-based.
type Type string

// Types offers the known enum values
var Types = struct {
	Keys Type
	Role Type
}{
	"KEYS",
	"ROLE",
}
