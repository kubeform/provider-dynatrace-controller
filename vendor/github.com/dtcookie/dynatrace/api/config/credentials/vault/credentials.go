package vault

import (
	"encoding/json"
	"errors"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

type Credentials struct {
	ID                *string            `json:"id,omitempty"`
	Name              string             `json:"name"`                  // The name of the credentials set.
	Description       *string            `json:"description,omitempty"` // A short description of the credentials set..
	Type              CredentialsType    `json:"type"`
	OwnerAccessOnly   bool               `json:"ownerAccessOnly"`             // The credentials set is available to every user (`false`) or to owner only (`true`).
	Scope             Scope              `json:"scope"`                       // The scope of the credentials set
	Token             *string            `json:"token,omitempty"`             // Token in the string format.
	Password          *string            `json:"password,omitempty"`          // The password of the credential.
	Username          *string            `json:"user,omitempty"`              // The username of the credentials set.
	Certificate       *string            `json:"certificate,omitempty"`       // The certificate in the string format.
	CertificateFormat *CertificateFormat `json:"certificateFormat,omitempty"` // The certificate format.
	isPublic          bool               `json:"-"`
}

func (me *Credentials) GetType() CredentialsType {
	if len(me.Type) > 0 {
		return me.Type
	}
	if me.Username != nil {
		return CredentialsTypes.UsernamePassword
	}
	if me.Token != nil {
		return CredentialsTypes.Token
	}
	if me.Certificate != nil {
		if me.isPublic {
			return CredentialsTypes.PublicCertificate
		}
		return CredentialsTypes.Certificate
	}
	return CredentialsTypes.Unknown
}

func (me *Credentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the credentials set",
			Required:    true,
		},
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the credentials set",
			Optional:    true,
		},
		// "type": {
		// 	Type:        hcl.TypeString,
		// 	Description: "The type of the credentials set",
		// 	Optional:    true,
		// },
		"owner_access_only": {
			Type:        hcl.TypeBool,
			Description: "The credentials set is available to every user (`false`) or to owner only (`true`)",
			Optional:    true,
		},
		"public": {
			Type:          hcl.TypeBool,
			Description:   "For certificate authentication specifies whether it's public certificate auth (`true`) or not (`false`).",
			ConflictsWith: []string{"username", "token"},
			Optional:      true,
		},
		"scope": {
			Type:        hcl.TypeString,
			Description: "The scope of the credentials set. Possible values are `ALL`, `EXTENSION` and `SYNTHETIC`",
			Required:    true,
		},
		"token": {
			Type:          hcl.TypeString,
			Description:   "Token in the string format. Specifying a token implies `Token Authentication`.",
			ConflictsWith: []string{"username", "password", "certificate", "format", "public"},
			Sensitive:     true,
			Optional:      true,
		},
		"username": {
			Type:          hcl.TypeString,
			Description:   "The username of the credentials set.",
			ConflictsWith: []string{"token", "public", "certificate"},
			RequiredWith:  []string{"password"},
			Sensitive:     true,
			Optional:      true,
		},
		"password": {
			Type:          hcl.TypeString,
			Description:   "The password of the credential.",
			ConflictsWith: []string{"token"},
			Sensitive:     true,
			Optional:      true,
		},
		"certificate": {
			Type:          hcl.TypeString,
			Description:   "The certificate in the string format.",
			ConflictsWith: []string{"token", "username"},
			RequiredWith:  []string{"format", "password"},
			Optional:      true,
		},
		"format": {
			Type:          hcl.TypeString,
			Description:   "The certificate format. Possible values are `PEM`, `PKCS12` and `UNKNOWN`.",
			ConflictsWith: []string{"token", "username"},
			RequiredWith:  []string{"certificate"},
			Optional:      true,
		},
	}
}

func (me *Credentials) MarshalJSON() ([]byte, error) {
	if me.Username != nil {
		creds := struct {
			ID              *string         `json:"id,omitempty"`
			Name            string          `json:"name"`
			Description     *string         `json:"description,omitempty"`
			Type            CredentialsType `json:"type"`
			OwnerAccessOnly bool            `json:"ownerAccessOnly"`
			Scope           Scope           `json:"scope"`
			Password        string          `json:"password"`
			Username        string          `json:"user"`
		}{
			me.ID,
			me.Name,
			me.Description,
			CredentialsTypes.UsernamePassword,
			me.OwnerAccessOnly,
			me.Scope,
			*me.Password,
			*me.Username,
		}
		return json.Marshal(&creds)
	}
	if me.Token != nil {
		creds := struct {
			ID              *string         `json:"id,omitempty"`
			Name            string          `json:"name"`
			Description     *string         `json:"description,omitempty"`
			Type            CredentialsType `json:"type"`
			OwnerAccessOnly bool            `json:"ownerAccessOnly"`
			Scope           Scope           `json:"scope"`
			Token           string          `json:"token"`
		}{
			me.ID,
			me.Name,
			me.Description,
			CredentialsTypes.Token,
			me.OwnerAccessOnly,
			me.Scope,
			*me.Token,
		}
		return json.Marshal(&creds)
	}
	if me.Certificate != nil {
		creds := struct {
			ID                *string           `json:"id,omitempty"`
			Name              string            `json:"name"`
			Description       *string           `json:"description,omitempty"`
			Type              CredentialsType   `json:"type"`
			OwnerAccessOnly   bool              `json:"ownerAccessOnly"`
			Scope             Scope             `json:"scope"`
			Certificate       string            `json:"certificate"`
			CertificateFormat CertificateFormat `json:"certificateFormat"`
			Password          string            `json:"password"`
		}{
			me.ID,
			me.Name,
			me.Description,
			me.GetType(),
			me.OwnerAccessOnly,
			me.Scope,
			*me.Certificate,
			*me.CertificateFormat,
			*me.Password,
		}
		return json.Marshal(&creds)
	}
	return nil, errors.New("invalid credentials - neither username, token nor certificate were specified")
}

func (me *Credentials) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["name"] = me.Name
	if me.Description != nil && len(*me.Description) > 0 {
		result["description"] = *me.Description
	}
	if me.OwnerAccessOnly {
		result["owner_access_only"] = me.OwnerAccessOnly
	}
	result["scope"] = string(me.Scope)
	if me.Token != nil {
		result["token"] = *me.Token
	}
	if me.Password != nil {
		result["password"] = *me.Password
	}
	if me.Username != nil {
		result["username"] = *me.Username
	}
	if me.Certificate != nil {
		result["certificate"] = *me.Certificate
	}
	if me.CertificateFormat != nil {
		result["format"] = *me.CertificateFormat
	}
	if me.GetType() == CredentialsTypes.PublicCertificate {
		result["public"] = true
	}
	// result["type"] = string(me.GetType())
	return result, nil
}

func (me *Credentials) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("description"); ok {
		me.Description = opt.NewString(value.(string))
		if len(*me.Description) == 0 {
			me.Description = nil
		}
	}
	if value, ok := decoder.GetOk("owner_access_only"); ok {
		me.OwnerAccessOnly = value.(bool)
	}
	if value, ok := decoder.GetOk("scope"); ok {
		me.Scope = Scope(value.(string))
	}
	if value, ok := decoder.GetOk("token"); ok {
		me.Token = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("password"); ok {
		me.Password = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("username"); ok {
		me.Username = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("certificate"); ok {
		me.Certificate = opt.NewString(value.(string))
	}
	if value, ok := decoder.GetOk("format"); ok {
		me.CertificateFormat = CertificateFormat(value.(string)).Ref()
	}
	if value, ok := decoder.GetOk("public"); ok {
		me.isPublic = value.(bool)
	}
	return nil
}

type CredentialsType string

var CredentialsTypes = struct {
	Certificate       CredentialsType
	PublicCertificate CredentialsType
	Token             CredentialsType
	UsernamePassword  CredentialsType
	Unknown           CredentialsType
}{
	"CERTIFICATE",
	"PUBLIC_CERTIFICATE",
	"TOKEN",
	"USERNAME_PASSWORD",
	"UNKNOWN",
}

// CertificateFormat The certificate format.
type CertificateFormat string

// CertificateFormats offers the known enum values
var CertificateFormats = struct {
	Pem     CertificateFormat
	Pkcs12  CertificateFormat
	Unknown CertificateFormat
}{
	"PEM",
	"PKCS12",
	"UNKNOWN",
}

func (me CertificateFormat) Ref() *CertificateFormat {
	return &me
}

type Scope string

// Scopes offers the known enum values
var Scopes = struct {
	All       Scope
	Extension Scope
	Synthetic Scope
}{
	"ALL",
	"EXTENSION",
	"SYNTHETIC",
}
