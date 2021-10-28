package auth

import (
	"github.com/dtcookie/hcl"
)

// Credentials The login credentials to bypass the browser login mask during a Navigate event
type Credentials struct {
	Type       string     `json:"type"`       // The type of authentication
	Credential Credential `json:"credential"` // A reference to the entry within the credential vault
}

func (me *Credentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of authentication",
			Required:    true,
		},
		"creds": {
			Type:        hcl.TypeString,
			Description: "A reference to the entry within the credential vault",
			Required:    true,
		},
	}
}

func (me *Credentials) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = me.Type
	result["creds"] = me.Credential.ID
	return result, nil
}

func (me *Credentials) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	cred := new(Credential)
	if err := decoder.Decode("creds", &cred.ID); err != nil {
		return err
	}
	if len(cred.ID) > 0 {
		me.Credential = *cred
	}
	return nil
}

type Credential struct {
	ID string `json:"id"`
}
