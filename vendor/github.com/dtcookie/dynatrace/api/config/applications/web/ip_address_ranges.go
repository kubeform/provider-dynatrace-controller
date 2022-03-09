package web

import "github.com/dtcookie/hcl"

type IPAddressRanges []*IPAddressRange

func (me *IPAddressRanges) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"range": {
			Type:        hcl.TypeList,
			Description: "The IP address or the IP address range to be mapped to the location",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(IPAddressRange).Schema()},
		},
	}
}

func (me IPAddressRanges) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["range"] = entries
	}
	return result, nil
}

func (me *IPAddressRanges) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("range", me); err != nil {
		return err
	}
	return nil
}

// IPAddressRange The IP address or the IP address range to be mapped to the location
type IPAddressRange struct {
	SubNetMask *int32  `json:"subnetMask,omitempty"` // The subnet mask of the IP address range. Valid values range from 0 to 128.
	Address    string  `json:"address"`              // The IP address to be mapped. \n\nFor an IP address range, this is the **from** address.
	ToAddress  *string `json:"addressTo,omitempty"`  // The **to** address of the IP address range.
}

func (me *IPAddressRange) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"subnet_mask": {
			Type:        hcl.TypeInt,
			Description: "The subnet mask of the IP address range. Valid values range from 0 to 128.",
			Optional:    true,
		},
		"address": {
			Type:        hcl.TypeString,
			Description: "The IP address to be mapped. \n\nFor an IP address range, this is the **from** address.",
			Required:    true,
		},
		"address_to": {
			Type:        hcl.TypeString,
			Description: "The **to** address of the IP address range.",
			Optional:    true,
		},
	}
}

func (me *IPAddressRange) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"subnet_mask": me.SubNetMask,
		"address":     me.Address,
		"address_to":  me.ToAddress,
	})
}

func (me *IPAddressRange) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"subnet_mask": &me.SubNetMask,
		"address":     &me.Address,
		"address_to":  &me.ToAddress,
	})
}
