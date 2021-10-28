package browser

import (
	"github.com/dtcookie/hcl"
)

type Bandwidth struct {
	NetworkType *string `json:"networkType,omitempty"` // The type of the preconfigured network—when editing in the browser, press `Crtl+Spacebar` to see the list of available networks
	Latency     *int    `json:"latency,omitempty"`     // The latency of the network, in milliseconds
	Download    *int    `json:"download,omitempty"`    // The download speed of the network, in bytes per second
	Upload      *int    `json:"upload,omitempty"`      // The upload speed of the network, in bytes per second
}

func (me *Bandwidth) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"network_type": {
			Type:        hcl.TypeString,
			Description: "The type of the preconfigured network—when editing in the browser, press `Crtl+Spacebar` to see the list of available networks",
			Optional:    true,
		},
		"latency": {
			Type:        hcl.TypeInt,
			Description: "The latency of the network, in milliseconds",
			Optional:    true,
		},
		"download": {
			Type:        hcl.TypeInt,
			Description: "The download speed of the network, in bytes per second",
			Optional:    true,
		},
		"upload": {
			Type:        hcl.TypeInt,
			Description: "The upload speed of the network, in bytes per second",
			Optional:    true,
		},
	}
}

func (me *Bandwidth) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.NetworkType != nil {
		result["network_type"] = string(*me.NetworkType)
	}
	if me.Latency != nil {
		result["latency"] = int(*me.Latency)
	}
	if me.Download != nil {
		result["download"] = int(*me.Download)
	}
	if me.Upload != nil {
		result["upload"] = int(*me.Upload)
	}
	return result, nil
}

func (me *Bandwidth) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("network_type", &me.NetworkType); err != nil {
		return err
	}
	if err := decoder.Decode("latency", &me.Latency); err != nil {
		return err
	}
	if err := decoder.Decode("download", &me.Download); err != nil {
		return err
	}
	if err := decoder.Decode("upload", &me.Upload); err != nil {
		return err
	}
	return nil
}
