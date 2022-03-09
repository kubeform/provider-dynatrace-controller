package users

import (
	"github.com/dtcookie/hcl"
)

// The configuration of the user
type UserConfig struct {
	UserName  string   `json:"id"`               // User ID
	Email     string   `json:"email"`            // User's email address
	FirstName string   `json:"firstName"`        // User's first name
	LastName  string   `json:"lastName"`         // User's last name
	Groups    []string `json:"groups,omitempty"` // List of user's user group IDs
}

func (me *UserConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"user_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The User Name",
		},
		"email": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "User's email address",
		},
		"first_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "User's first name",
		},
		"last_name": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "User's last name",
		},
		"groups": {
			Type:        hcl.TypeSet,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Optional:    true,
			Description: "List of user's user group IDs",
		},
	}
}

func (me *UserConfig) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"user_name":  me.UserName,
		"email":      me.Email,
		"first_name": me.FirstName,
		"last_name":  me.LastName,
		"groups":     me.Groups,
	})
}

func (me *UserConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"user_name":  &me.UserName,
		"email":      &me.Email,
		"first_name": &me.FirstName,
		"last_name":  &me.LastName,
		"groups":     &me.Groups,
	})
}
