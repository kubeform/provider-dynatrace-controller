package vault

import "github.com/dtcookie/gojson"

// CredentialsList A list of credentials sets for Synthetic monitors.
type CredentialsList struct {
	Credentials []CredentialsResponseElement `json:"credentials"` // A list of credentials sets for Synthetic monitors.
}

// UnmarshalJSON provides custom JSON deserialization
func (cl *CredentialsList) UnmarshalJSON(data []byte) error {
	return gojson.Unmarshal(data, cl)
}

// MarshalJSON provides custom JSON serialization
func (cl *CredentialsList) MarshalJSON() ([]byte, error) {
	return gojson.Marshal(cl)
}
