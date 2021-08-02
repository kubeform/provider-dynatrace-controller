package aws

import (
	"encoding/json"
)

// KeyBasedAuthentication The credentials for the key-based authentication.
type KeyBasedAuthentication struct {
	AccessKey string                     `json:"accessKey"` // The ID of the access key
	SecretKey *string                    `json:"secretKey"` // The secret access key
	Unknowns  map[string]json.RawMessage `json:"-"`
}

// UnmarshalJSON provides custom JSON deserialization
func (kba *KeyBasedAuthentication) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["accessKey"]; found {
		if err := json.Unmarshal(v, &kba.AccessKey); err != nil {
			return err
		}
	}
	if v, found := m["secretKey"]; found {
		if err := json.Unmarshal(v, &kba.SecretKey); err != nil {
			return err
		}
	}
	delete(m, "accessKey")
	delete(m, "secretKey")
	if len(m) > 0 {
		kba.Unknowns = m
	}
	return nil
}

func (kba *KeyBasedAuthentication) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(kba.Unknowns) > 0 {
		for k, v := range kba.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(kba.AccessKey)
		if err != nil {
			return nil, err
		}
		m["accessKey"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(kba.SecretKey)
		if err != nil {
			return nil, err
		}
		m["secretKey"] = rawMessage
	}
	return json.Marshal(m)
}
