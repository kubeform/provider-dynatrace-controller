package network

import (
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/droppedpackets"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/errors"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/retransmission"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/tcp"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/utilization"
	"github.com/dtcookie/hcl"
)

type DetectionConfig struct {
	NetworkDroppedPacketsDetection     *droppedpackets.DetectionConfig `json:"networkDroppedPacketsDetection"`     // Configuration of high number of dropped packets detection.
	HighNetworkDetection               *utilization.DetectionConfig    `json:"highNetworkDetection"`               // Configuration of high network utilization detection.
	NetworkHighRetransmissionDetection *retransmission.DetectionConfig `json:"networkHighRetransmissionDetection"` // Configuration of high retransmission rate detection.
	NetworkTcpProblemsDetection        *tcp.DetectionConfig            `json:"networkTcpProblemsDetection"`        // Configuration of TCP connectivity problems detection.
	NetworkErrorsDetection             *errors.DetectionConfig         `json:"networkErrorsDetection"`             // Configuration of high number of network errors detection.
}

func (me *DetectionConfig) IsConfigured() bool {
	if me.NetworkDroppedPacketsDetection != nil && me.NetworkDroppedPacketsDetection.Enabled {
		return true
	}
	if me.HighNetworkDetection != nil && me.HighNetworkDetection.Enabled {
		return true
	}
	if me.NetworkHighRetransmissionDetection != nil && me.NetworkHighRetransmissionDetection.Enabled {
		return true
	}
	if me.NetworkTcpProblemsDetection != nil && me.NetworkTcpProblemsDetection.Enabled {
		return true
	}
	if me.NetworkErrorsDetection != nil && me.NetworkErrorsDetection.Enabled {
		return true
	}
	return false
}

func (me *DetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dropped_packets": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high number of dropped packets detection",
			Elem:        &hcl.Resource{Schema: new(droppedpackets.DetectionConfig).Schema()},
		},
		"utilization": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high network utilization detection",
			Elem:        &hcl.Resource{Schema: new(utilization.DetectionConfig).Schema()},
		},
		"retransmission": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high retransmission rate detection",
			Elem:        &hcl.Resource{Schema: new(retransmission.DetectionConfig).Schema()},
		},
		"connectivity": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of TCP connectivity problems detection",
			Elem:        &hcl.Resource{Schema: new(tcp.DetectionConfig).Schema()},
		},
		"errors": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high number of network errors detection",
			Elem:        &hcl.Resource{Schema: new(errors.DetectionConfig).Schema()},
		},
	}
}

func (me *DetectionConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.NetworkDroppedPacketsDetection != nil && me.NetworkDroppedPacketsDetection.Enabled {
		if marshalled, err := me.NetworkDroppedPacketsDetection.MarshalHCL(hcl.NewDecoder(decoder, "dropped_packets", 0)); err == nil {
			result["dropped_packets"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.HighNetworkDetection != nil && me.HighNetworkDetection.Enabled {
		if marshalled, err := me.HighNetworkDetection.MarshalHCL(hcl.NewDecoder(decoder, "utilization", 0)); err == nil {
			result["utilization"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.NetworkHighRetransmissionDetection != nil && me.NetworkHighRetransmissionDetection.Enabled {
		if marshalled, err := me.NetworkHighRetransmissionDetection.MarshalHCL(hcl.NewDecoder(decoder, "retransmission", 0)); err == nil {
			result["retransmission"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.NetworkTcpProblemsDetection != nil && me.NetworkTcpProblemsDetection.Enabled {
		if marshalled, err := me.NetworkTcpProblemsDetection.MarshalHCL(hcl.NewDecoder(decoder, "connectivity", 0)); err == nil {
			result["connectivity"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.NetworkErrorsDetection != nil && me.NetworkErrorsDetection.Enabled {
		if marshalled, err := me.NetworkErrorsDetection.MarshalHCL(hcl.NewDecoder(decoder, "errors", 0)); err == nil {
			result["errors"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *DetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	me.NetworkDroppedPacketsDetection = &droppedpackets.DetectionConfig{Enabled: false}
	me.HighNetworkDetection = &utilization.DetectionConfig{Enabled: false}
	me.NetworkHighRetransmissionDetection = &retransmission.DetectionConfig{Enabled: false}
	me.NetworkTcpProblemsDetection = &tcp.DetectionConfig{Enabled: false}
	me.NetworkErrorsDetection = &errors.DetectionConfig{Enabled: false}
	if _, ok := decoder.GetOk("dropped_packets.#"); ok {
		me.NetworkDroppedPacketsDetection = new(droppedpackets.DetectionConfig)
		if err := me.NetworkDroppedPacketsDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "dropped_packets", 0)); err != nil {
			return err
		}
	}

	if _, ok := decoder.GetOk("utilization.#"); ok {
		me.HighNetworkDetection = new(utilization.DetectionConfig)
		if err := me.HighNetworkDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "utilization", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("retransmission.#"); ok {
		me.NetworkHighRetransmissionDetection = new(retransmission.DetectionConfig)
		if err := me.NetworkHighRetransmissionDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "retransmission", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("connectivity.#"); ok {
		me.NetworkTcpProblemsDetection = new(tcp.DetectionConfig)
		if err := me.NetworkTcpProblemsDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "connectivity", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("errors.#"); ok {
		me.NetworkErrorsDetection = new(errors.DetectionConfig)
		if err := me.NetworkErrorsDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "errors", 0)); err != nil {
			return err
		}
	}
	return nil
}
