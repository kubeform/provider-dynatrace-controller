package aws

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// AWSAuthenticationData A credentials for the AWS authentication.
type AWSAuthenticationData struct {
	KeyBasedAuthentication  *KeyBasedAuthentication    `json:"keyBasedAuthentication,omitempty"`  // The credentials for the key-based authentication.
	RoleBasedAuthentication *RoleBasedAuthentication   `json:"roleBasedAuthentication,omitempty"` // The credentials for the role-based authentication.
	Type                    Type                       `json:"type"`                              // The type of the authentication: role-based or key-based.
	Unknowns                map[string]json.RawMessage `json:"-"`
}

func (aad *AWSAuthenticationData) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"access_key": {
			Type:        hcl.TypeString,
			Description: "the access key",
			Optional:    true,
		},
		"secret_key": {
			Type:        hcl.TypeString,
			Description: "the secret access key",
			Optional:    true,
		},
		"account_id": {
			Type:        hcl.TypeString,
			Description: "the ID of the Amazon account",
			Optional:    true,
		},
		"external_id": {
			Type:        hcl.TypeString,
			Description: "the external ID token for setting an IAM role. You can obtain it with the `GET /aws/iamExternalId` request",
			Optional:    true,
		},
		"iam_role": {
			Type:        hcl.TypeString,
			Description: "the IAM role to be used by Dynatrace to get monitoring data",
			Optional:    true,
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
	}
}

func (aad *AWSAuthenticationData) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), aad); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &aad.Unknowns); err != nil {
			return err
		}
		delete(aad.Unknowns, "access_key")
		delete(aad.Unknowns, "secret_key")
		delete(aad.Unknowns, "account_id")
		delete(aad.Unknowns, "external_id")
		delete(aad.Unknowns, "iam_role")
		if len(aad.Unknowns) == 0 {
			aad.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("access_key"); ok {
		if aad.KeyBasedAuthentication == nil {
			aad.KeyBasedAuthentication = new(KeyBasedAuthentication)
		}
		aad.Type = Types.Keys
		aad.KeyBasedAuthentication.AccessKey = value.(string)
	}
	if value, ok := decoder.GetOk("secret_key"); ok {
		if aad.KeyBasedAuthentication == nil {
			aad.KeyBasedAuthentication = new(KeyBasedAuthentication)
		}
		aad.Type = Types.Keys
		aad.KeyBasedAuthentication.SecretKey = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("account_id"); ok {
		if aad.RoleBasedAuthentication == nil {
			aad.RoleBasedAuthentication = new(RoleBasedAuthentication)
		}
		aad.Type = Types.Role
		aad.RoleBasedAuthentication.AccountID = value.(string)
	}
	if value, ok := decoder.GetOk("external_id"); ok {
		if aad.RoleBasedAuthentication == nil {
			aad.RoleBasedAuthentication = new(RoleBasedAuthentication)
		}
		aad.Type = Types.Role
		aad.RoleBasedAuthentication.ExternalID = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("iam_role"); ok {
		if aad.RoleBasedAuthentication == nil {
			aad.RoleBasedAuthentication = new(RoleBasedAuthentication)
		}
		aad.Type = Types.Role
		aad.RoleBasedAuthentication.IamRole = value.(string)
	}
	return nil
}

func (aad *AWSAuthenticationData) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(aad.Unknowns) > 0 {
		data, err := json.Marshal(aad.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	if aad.KeyBasedAuthentication != nil {
		result["access_key"] = aad.KeyBasedAuthentication.AccessKey
		if aad.KeyBasedAuthentication.SecretKey != nil {
			if len(*aad.KeyBasedAuthentication.SecretKey) > 0 {
				result["secret_key"] = *aad.KeyBasedAuthentication.SecretKey
			}
		}
	}
	if aad.RoleBasedAuthentication != nil {
		result["account_id"] = aad.RoleBasedAuthentication.AccountID
		if aad.RoleBasedAuthentication.ExternalID != nil {
			result["external_id"] = aad.RoleBasedAuthentication.ExternalID
		}
		result["iam_role"] = aad.RoleBasedAuthentication.IamRole

	}
	return result, nil
}

// UnmarshalJSON provides custom JSON deserialization
func (aad *AWSAuthenticationData) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["keyBasedAuthentication"]; found {
		if err := json.Unmarshal(v, &aad.KeyBasedAuthentication); err != nil {
			return err
		}
	}
	if v, found := m["roleBasedAuthentication"]; found {
		if err := json.Unmarshal(v, &aad.RoleBasedAuthentication); err != nil {
			return err
		}
	}
	if v, found := m["type"]; found {
		if err := json.Unmarshal(v, &aad.Type); err != nil {
			return err
		}
	} else {
		if aad.RoleBasedAuthentication != nil {
			aad.Type = Types.Role
		} else if aad.KeyBasedAuthentication != nil {
			aad.Type = Types.Keys
		}
	}
	delete(m, "keyBasedAuthentication")
	delete(m, "roleBasedAuthentication")
	delete(m, "type")
	if len(m) > 0 {
		aad.Unknowns = m
	}
	return nil
}

// MarshalJSON provides custom JSON serialization
func (aad *AWSAuthenticationData) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(aad.Unknowns) > 0 {
		for k, v := range aad.Unknowns {
			m[k] = v
		}
	}
	if aad.KeyBasedAuthentication != nil {
		rawMessage, err := json.Marshal(aad.KeyBasedAuthentication)
		if err != nil {
			return nil, err
		}
		m["keyBasedAuthentication"] = rawMessage
	}
	if aad.RoleBasedAuthentication != nil {
		rawMessage, err := json.Marshal(aad.RoleBasedAuthentication)
		if err != nil {
			return nil, err
		}
		m["roleBasedAuthentication"] = rawMessage
	}
	rawMessage, err := json.Marshal(aad.Type)
	if err != nil {
		return nil, err
	}
	m["type"] = rawMessage
	return json.Marshal(m)
}
