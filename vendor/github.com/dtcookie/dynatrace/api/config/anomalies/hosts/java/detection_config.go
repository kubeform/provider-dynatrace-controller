package java

import (
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/java/oom"
	"github.com/dtcookie/dynatrace/api/config/anomalies/hosts/java/oot"
	"github.com/dtcookie/hcl"
)

type DetectionConfig struct {
	OutOfMemoryDetection  *oom.DetectionConfig `json:"outOfMemoryDetection"`  // Configuration of Java out of memory problems detection.
	OutOfThreadsDetection *oot.DetectionConfig `json:"outOfThreadsDetection"` // Configuration of Java out of threads problems detection.
}

func (me *DetectionConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"out_of_threads": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of Java out of threads problems detection",
			Elem:        &hcl.Resource{Schema: new(oot.DetectionConfig).Schema()},
		},
		"out_of_memory": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of Java out of memory problems detection",
			Elem:        &hcl.Resource{Schema: new(oom.DetectionConfig).Schema()},
		},
	}
}

func (me *DetectionConfig) IsConfigured() bool {
	if me.OutOfMemoryDetection != nil && me.OutOfMemoryDetection.Enabled {
		return true
	}
	if me.OutOfThreadsDetection != nil && me.OutOfThreadsDetection.Enabled {
		return true
	}
	return false
}

func (me *DetectionConfig) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.OutOfMemoryDetection != nil && me.OutOfMemoryDetection.Enabled {
		if marshalled, err := me.OutOfMemoryDetection.MarshalHCL(hcl.NewDecoder(decoder, "out_of_memory", 0)); err == nil {
			result["out_of_memory"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.OutOfThreadsDetection != nil && me.OutOfThreadsDetection.Enabled {
		if marshalled, err := me.OutOfThreadsDetection.MarshalHCL(hcl.NewDecoder(decoder, "out_of_threads", 0)); err == nil {
			result["out_of_threads"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *DetectionConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	me.OutOfMemoryDetection = &oom.DetectionConfig{Enabled: false}
	me.OutOfThreadsDetection = &oot.DetectionConfig{Enabled: false}

	if _, ok := decoder.GetOk("out_of_threads.#"); ok {
		me.OutOfThreadsDetection = new(oot.DetectionConfig)
		if err := me.OutOfThreadsDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "out_of_threads", 0)); err != nil {
			return err
		}
	}
	if _, ok := decoder.GetOk("out_of_memory.#"); ok {
		me.OutOfMemoryDetection = new(oom.DetectionConfig)
		if err := me.OutOfMemoryDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "out_of_memory", 0)); err != nil {
			return err
		}
	}
	return nil
}
