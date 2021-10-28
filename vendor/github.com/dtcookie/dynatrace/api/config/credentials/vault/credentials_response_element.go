package vault

import "github.com/dtcookie/gojson"

// CredentialsResponseElement Metadata of the credentials set.
type CredentialsResponseElement struct {
	OwnerAccessOnly bool                           `json:"ownerAccessOnly"` // Flag indicating that this credential is visible only to the owner.
	Type            CredentialsResponseElementType `json:"type"`            // The type of the credentials set.
	Description     string                         `json:"description"`     // A short description of the credentials set.
	ID              *string                        `json:"id,omitempty"`    // The ID of the credentials set.
	Name            string                         `json:"name"`            // The name of the credentials set.
	Owner           string                         `json:"owner"`           // The owner of the credential.
}

// UnmarshalJSON provides custom JSON deserialization
func (cre *CredentialsResponseElement) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cre)
}

// MarshalJSON provides custom JSON serialization
func (cre *CredentialsResponseElement) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cre)
}

// CredentialsResponseElementType The type of the credentials set.
type CredentialsResponseElementType string

// CredentialsResponseElementTypes offers the known enum values
var CredentialsResponseElementTypes = struct {
	Certificate      CredentialsResponseElementType
	Token            CredentialsResponseElementType
	Unknown          CredentialsResponseElementType
	UsernamePassword CredentialsResponseElementType
}{
	"CERTIFICATE",
	"TOKEN",
	"UNKNOWN",
	"USERNAME_PASSWORD",
}
