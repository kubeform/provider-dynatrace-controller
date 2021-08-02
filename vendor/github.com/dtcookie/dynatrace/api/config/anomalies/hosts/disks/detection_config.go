package disks

import (
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks/inodes"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks/slow"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/disks/space"
	"github.com/dtcookie/hcl"
)

type DetectionConfig struct {
	Speed  *slow.DetectionConfig   `json:"diskSlowWritesAndReadsDetection"` // Configuration of slow running disks detection.
	Space  *space.DetectionConfig  `json:"diskLowSpaceDetection"`           // Configuration of low disk space detection.
	Inodes *inodes.DetectionConfig `json:"diskLowInodesDetection"`          // Configuration of low disk inodes number detection.
}

func (me *DetectionConfig) IsConfigured() bool {
	if me.Speed != nil && me.Speed.Enabled {
		return true
	}
	if me.Space != nil && me.Space.Enabled {
		return true
	}
	if me.Inodes != nil && me.Inodes.Enabled {
		return true
	}
	return false
}

func (me *DetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"space": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of low disk space detection",
			Elem:        &hcl.Resource{Schema: new(space.DetectionConfig).Schema()},
		},
		"speed": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of slow running disks detection",
			Elem:        &hcl.Resource{Schema: new(slow.DetectionConfig).Schema()},
		},
		"inodes": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of low disk inodes number detection",
			Elem:        &hcl.Resource{Schema: new(inodes.DetectionConfig).Schema()},
		},
	}
}

func (me *DetectionConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Space != nil {
		if marshalled, err := me.Space.MarshalHCL(hcl.NewDecoder(decoder, "space", 0)); err == nil {
			result["space"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Speed != nil {
		if marshalled, err := me.Speed.MarshalHCL(hcl.NewDecoder(decoder, "speed", 0)); err == nil {
			result["speed"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Inodes != nil {
		if marshalled, err := me.Inodes.MarshalHCL(hcl.NewDecoder(decoder, "inodes", 0)); err == nil {
			result["inodes"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *DetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Speed = &slow.DetectionConfig{Enabled: false}
	me.Space = &space.DetectionConfig{Enabled: false}
	me.Inodes = &inodes.DetectionConfig{Enabled: false}

	if _, ok := decoder.GetOk("space.#"); ok {
		me.Space = new(space.DetectionConfig)
		if err := me.Space.UnmarshalHCL(hcl.NewDecoder(decoder, "space", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("speed.#"); ok {
		me.Speed = new(slow.DetectionConfig)
		if err := me.Speed.UnmarshalHCL(hcl.NewDecoder(decoder, "speed", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("inodes.#"); ok {
		me.Inodes = new(inodes.DetectionConfig)
		if err := me.Inodes.UnmarshalHCL(hcl.NewDecoder(decoder, "inodes", 0)); err != nil {
			return err
		}
	}
	return nil
}
