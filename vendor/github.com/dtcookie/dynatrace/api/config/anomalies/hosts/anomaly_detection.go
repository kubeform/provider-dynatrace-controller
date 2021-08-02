package hosts

import (
	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/connection"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/cpu"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks/inodes"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks/slow"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks/space"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/gc"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/java"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/java/oom"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/java/oot"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/memory"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/droppedpackets"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/errors"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/retransmission"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/tcp"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/network/utilization"
	"github.com/dtcookie/hcl"
)

// AnomalyDetection Configuration of anomaly detection for hosts.
type AnomalyDetection struct {
	NetworkDroppedPacketsDetection     *droppedpackets.DetectionConfig `json:"networkDroppedPacketsDetection"`     // Configuration of high number of dropped packets detection.
	HighNetworkDetection               *utilization.DetectionConfig    `json:"highNetworkDetection"`               // Configuration of high network utilization detection.
	NetworkHighRetransmissionDetection *retransmission.DetectionConfig `json:"networkHighRetransmissionDetection"` // Configuration of high retransmission rate detection.
	NetworkTcpProblemsDetection        *tcp.DetectionConfig            `json:"networkTcpProblemsDetection"`        // Configuration of TCP connectivity problems detection.
	NetworkErrorsDetection             *errors.DetectionConfig         `json:"networkErrorsDetection"`             // Configuration of high number of network errors detection.
	HighMemoryDetection                *memory.DetectionConfig         `json:"highMemoryDetection"`                // Configuration of high memory usage detection.
	HighCPUSaturationDetection         *cpu.DetectionConfig            `json:"highCpuSaturationDetection"`         // Configuration of high CPU saturation detection
	OutOfMemoryDetection               *oom.DetectionConfig            `json:"outOfMemoryDetection"`               // Configuration of Java out of memory problems detection.
	OutOfThreadsDetection              *oot.DetectionConfig            `json:"outOfThreadsDetection"`              // Configuration of Java out of threads problems detection.
	HighGcActivityDetection            *gc.DetectionConfig             `json:"highGcActivityDetection"`            // Configuration of high Garbage Collector activity detection.
	ConnectionLostDetection            *connection.LostDetectionConfig `json:"connectionLostDetection"`            // Configuration of lost connection detection.
	DiskSlowWritesAndReadsDetection    *slow.DetectionConfig           `json:"diskSlowWritesAndReadsDetection"`    // Configuration of slow running disks detection.
	DiskLowSpaceDetection              *space.DetectionConfig          `json:"diskLowSpaceDetection"`              // Configuration of low disk space detection.
	DiskLowInodesDetection             *inodes.DetectionConfig         `json:"diskLowInodesDetection"`             // Configuration of low disk inodes number detection.
	Metadata                           *api.ConfigMetadata             `json:"metadata,omitempty"`                 // Metadata useful for debugging
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"memory": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high memory usage detection",
			Elem:        &hcl.Resource{Schema: new(memory.DetectionConfig).Schema()},
		},
		"cpu": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high CPU saturation detection",
			Elem:        &hcl.Resource{Schema: new(cpu.DetectionConfig).Schema()},
		},
		"gc": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of high Garbage Collector activity detection",
			Elem:        &hcl.Resource{Schema: new(gc.DetectionConfig).Schema()},
		},
		"connections": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of lost connection detection",
			Elem:        &hcl.Resource{Schema: new(connection.LostDetectionConfig).Schema()},
		},
		"network": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of network related anomalies",
			Elem:        &hcl.Resource{Schema: new(network.DetectionConfig).Schema()},
		},
		"disks": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of disk related anomalies",
			Elem:        &hcl.Resource{Schema: new(disks.DetectionConfig).Schema()},
		},
		"java": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of Java related anomalies",
			Elem:        &hcl.Resource{Schema: new(java.DetectionConfig).Schema()},
		},
	}
}

func (me *AnomalyDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.HighMemoryDetection != nil && me.HighMemoryDetection.Enabled {
		if marshalled, err := me.HighMemoryDetection.MarshalHCL(hcl.NewDecoder(decoder, "memory", 0)); err == nil {
			result["memory"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.HighCPUSaturationDetection != nil && me.HighCPUSaturationDetection.Enabled {
		if marshalled, err := me.HighCPUSaturationDetection.MarshalHCL(hcl.NewDecoder(decoder, "cpu_saturation", 0)); err == nil {
			result["cpu"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}

	if me.HighGcActivityDetection != nil && me.HighGcActivityDetection.Enabled {
		if marshalled, err := me.HighGcActivityDetection.MarshalHCL(hcl.NewDecoder(decoder, "gc", 0)); err == nil {
			result["gc"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.ConnectionLostDetection != nil && me.ConnectionLostDetection.Enabled {
		if marshalled, err := me.ConnectionLostDetection.MarshalHCL(hcl.NewDecoder(decoder, "connections", 0)); err == nil {
			result["connections"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	jdc := &java.DetectionConfig{
		OutOfMemoryDetection:  me.OutOfMemoryDetection,
		OutOfThreadsDetection: me.OutOfThreadsDetection,
	}
	if jdc.IsConfigured() {
		if marshalled, err := jdc.MarshalHCL(hcl.NewDecoder(decoder, "java", 0)); err == nil {
			result["java"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	ndc := &network.DetectionConfig{
		NetworkDroppedPacketsDetection:     me.NetworkDroppedPacketsDetection,
		HighNetworkDetection:               me.HighNetworkDetection,
		NetworkHighRetransmissionDetection: me.NetworkHighRetransmissionDetection,
		NetworkTcpProblemsDetection:        me.NetworkTcpProblemsDetection,
		NetworkErrorsDetection:             me.NetworkErrorsDetection,
	}
	if ndc.IsConfigured() {
		if marshalled, err := ndc.MarshalHCL(hcl.NewDecoder(decoder, "network", 0)); err == nil {
			result["network"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	ddc := &disks.DetectionConfig{
		Speed:  me.DiskSlowWritesAndReadsDetection,
		Space:  me.DiskLowSpaceDetection,
		Inodes: me.DiskLowInodesDetection,
	}
	if ddc.IsConfigured() {
		if marshalled, err := ddc.MarshalHCL(hcl.NewDecoder(decoder, "disks", 0)); err == nil {
			result["disks"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *AnomalyDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	me.NetworkDroppedPacketsDetection = &droppedpackets.DetectionConfig{Enabled: false}
	me.HighNetworkDetection = &utilization.DetectionConfig{Enabled: false}
	me.NetworkHighRetransmissionDetection = &retransmission.DetectionConfig{Enabled: false}
	me.NetworkTcpProblemsDetection = &tcp.DetectionConfig{Enabled: false}
	me.NetworkErrorsDetection = &errors.DetectionConfig{Enabled: false}
	me.DiskSlowWritesAndReadsDetection = &slow.DetectionConfig{Enabled: false}
	me.DiskLowSpaceDetection = &space.DetectionConfig{Enabled: false}
	me.DiskLowInodesDetection = &inodes.DetectionConfig{Enabled: false}
	me.HighMemoryDetection = &memory.DetectionConfig{Enabled: false}
	me.HighCPUSaturationDetection = &cpu.DetectionConfig{Enabled: false}
	me.OutOfMemoryDetection = &oom.DetectionConfig{Enabled: false}
	me.OutOfThreadsDetection = &oot.DetectionConfig{Enabled: false}
	me.HighGcActivityDetection = &gc.DetectionConfig{Enabled: false}
	me.ConnectionLostDetection = &connection.LostDetectionConfig{Enabled: false}

	if _, ok := decoder.GetOk("network.#"); ok {
		cfg := new(network.DetectionConfig)

		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "network", 0)); err != nil {
			return err
		}
		me.NetworkDroppedPacketsDetection = cfg.NetworkDroppedPacketsDetection
		me.HighNetworkDetection = cfg.HighNetworkDetection
		me.NetworkHighRetransmissionDetection = cfg.NetworkHighRetransmissionDetection
		me.NetworkTcpProblemsDetection = cfg.NetworkTcpProblemsDetection
		me.NetworkErrorsDetection = cfg.NetworkErrorsDetection
	}
	if _, ok := decoder.GetOk("disks.#"); ok {
		cfg := new(disks.DetectionConfig)

		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "disks", 0)); err != nil {
			return err
		}
		me.DiskSlowWritesAndReadsDetection = cfg.Speed
		me.DiskLowSpaceDetection = cfg.Space
		me.DiskLowInodesDetection = cfg.Inodes
	}
	if _, ok := decoder.GetOk("memory.#"); ok {
		me.HighMemoryDetection = new(memory.DetectionConfig)
		if err := me.HighMemoryDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "memory", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("cpu.#"); ok {
		me.HighCPUSaturationDetection = new(cpu.DetectionConfig)
		if err := me.HighCPUSaturationDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "cpu", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("gc.#"); ok {
		me.HighGcActivityDetection = new(gc.DetectionConfig)
		if err := me.HighGcActivityDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "gc", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("connections.#"); ok {
		me.ConnectionLostDetection = new(connection.LostDetectionConfig)
		if err := me.ConnectionLostDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "connections", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("java.#"); ok {
		cfg := new(java.DetectionConfig)

		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "java", 0)); err != nil {
			return err
		}
		me.OutOfMemoryDetection = cfg.OutOfMemoryDetection
		me.OutOfThreadsDetection = cfg.OutOfThreadsDetection
	}
	return nil
}
