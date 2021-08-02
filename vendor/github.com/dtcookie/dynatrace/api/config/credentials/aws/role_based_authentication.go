package aws

import "encoding/json"

// RoleBasedAuthentication The credentials for the role-based authentication.
type RoleBasedAuthentication struct {
	AccountID  string                     `json:"accountId"`            // The ID of the Amazon account.
	ExternalID *string                    `json:"externalId,omitempty"` // The external ID token for setting an IAM role.   You can obtain it with the `GET /aws/iamExternalId` request.
	IamRole    string                     `json:"iamRole"`              // The IAM role to be used by Dynatrace to get monitoring data.
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (rba *RoleBasedAuthentication) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if v, found := m["accountId"]; found {
		if err := json.Unmarshal(v, &rba.AccountID); err != nil {
			return err
		}
	}
	if v, found := m["iamRole"]; found {
		if err := json.Unmarshal(v, &rba.IamRole); err != nil {
			return err
		}
	}
	if rba.ExternalID != nil {
		if v, found := m["externalId"]; found {
			if err := json.Unmarshal(v, &rba.ExternalID); err != nil {
				return err
			}
		}
	}
	delete(m, "accountId")
	delete(m, "iamRole")
	delete(m, "externalId")
	if len(m) > 0 {
		rba.Unknowns = m
	}
	return nil
}

func (rba *RoleBasedAuthentication) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(rba.Unknowns) > 0 {
		for k, v := range rba.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(rba.AccountID)
		if err != nil {
			return nil, err
		}
		m["accountId"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(rba.IamRole)
		if err != nil {
			return nil, err
		}
		m["iamRole"] = rawMessage
	}
	if rba.ExternalID != nil {
		rawMessage, err := json.Marshal(rba.ExternalID)
		if err != nil {
			return nil, err
		}
		m["externalId"] = rawMessage
	}
	return json.Marshal(m)
}
