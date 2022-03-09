package web

import "github.com/dtcookie/hcl"

// IPAddressRestrictionSettings Settings for restricting certain ip addresses and for introducing subnet mask. It also restricts the mode
type IPAddressRestrictionSettings struct {
	Mode         RestrictionMode `json:"mode"`                            // The mode of the list of ip address restrictions. Possible values area `EXCLUDE` and `INCLUDE`.
	Restrictions IPAddressRanges `json:"ipAddressRestrictions,omitempty"` // The IP addresses or the IP address ranges to be mapped to the location
}

func (me *IPAddressRestrictionSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"mode": {
			Type:        hcl.TypeString,
			Description: "The mode of the list of ip address restrictions. Possible values area `EXCLUDE` and `INCLUDE`.",
			Required:    true,
		},
		"restrictions": {
			Type:        hcl.TypeList,
			Description: "The IP addresses or the IP address ranges to be mapped to the location",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(IPAddressRanges).Schema()},
		},
	}
}

func (me *IPAddressRestrictionSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"mode":         me.Mode,
		"restrictions": me.Restrictions,
	})
}

func (me *IPAddressRestrictionSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"mode":         &me.Mode,
		"restrictions": &me.Restrictions,
	})
}
