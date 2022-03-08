package keyrequests

import (
	"github.com/dtcookie/hcl"
)

// KeyRequest has no documentation
type KeyRequest struct {
	Names     []string `json:"names"`
	ServiceID string   `json:"serviceID"`
}

func (me *KeyRequest) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"service": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Whether to create an entry point or not",
		},
		"names": {
			Type:        hcl.TypeSet,
			Optional:    true,
			MinItems:    1,
			Description: "The names of the key requests",
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *KeyRequest) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}

	return properties.EncodeAll(map[string]interface{}{
		"names":   me.Names,
		"service": me.ServiceID,
	})
}

func (me *KeyRequest) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"names":   &me.Names,
		"service": &me.ServiceID,
	})
}
