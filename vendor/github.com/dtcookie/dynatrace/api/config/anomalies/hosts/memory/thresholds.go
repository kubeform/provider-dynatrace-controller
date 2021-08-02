package memory

import (
	"github.com/dtcookie/hcl"
)

// Thresholds Custom thresholds for high memory usage. If not set then the automatic mode is used.
//  **Both** conditions must be met to trigger an alert.
type Thresholds struct {
	UsedMemoryPercentageWindows    int32 `json:"usedMemoryPercentageWindows"`    // Memory usage is higher than *X*% on Windows.
	UsedMemoryPercentageNonWindows int32 `json:"usedMemoryPercentageNonWindows"` // Memory usage is higher than *X*% on Linux.
	PageFaultsPerSecondNonWindows  int32 `json:"pageFaultsPerSecondNonWindows"`  // Memory page fault rate is higher than *X* faults per second on Linux.
	PageFaultsPerSecondWindows     int32 `json:"pageFaultsPerSecondWindows"`     // Memory page fault rate is higher than *X* faults per second on Windows.
}

type osThresholds struct {
	UsedMemoryPercentage int32 // Memory usage is higher than *X*%.
	PageFaultsPerSecond  int32 // Memory page fault rate is higher than *X* faults per second.
}

func (me *osThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"usage": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Memory usage is higher than *X*%",
		},
		"page_faults": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Memory page fault rate is higher than *X* faults per second",
		},
	}
}

func (me *osThresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["usage"] = int(me.UsedMemoryPercentage)
	result["page_faults"] = int(me.PageFaultsPerSecond)
	return result, nil
}

func (me *osThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("usage"); ok {
		me.UsedMemoryPercentage = int32(value.(int))
	}
	if value, ok := decoder.GetOk("page_faults"); ok {
		me.PageFaultsPerSecond = int32(value.(int))
	}
	return nil
}

func (me *Thresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"windows": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Custom thresholds for Windows",
			Elem:        &hcl.Resource{Schema: new(osThresholds).Schema()},
		},
		"linux": {
			Type:        hcl.TypeList,
			Required:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Custom thresholds for Linux",
			Elem:        &hcl.Resource{Schema: new(osThresholds).Schema()},
		},
	}
}

func (me *Thresholds) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	thrs := &osThresholds{UsedMemoryPercentage: me.UsedMemoryPercentageNonWindows, PageFaultsPerSecond: me.PageFaultsPerSecondNonWindows}
	if marshalled, err := thrs.MarshalHCL(hcl.NewDecoder(decoder, "linux", 0)); err == nil {
		result["linux"] = []interface{}{marshalled}
	} else {
		return nil, err
	}
	thrs = &osThresholds{
		UsedMemoryPercentage: me.UsedMemoryPercentageWindows,
		PageFaultsPerSecond:  me.PageFaultsPerSecondWindows,
	}
	if marshalled, err := thrs.MarshalHCL(hcl.NewDecoder(decoder, "windows", 0)); err == nil {
		result["windows"] = []interface{}{marshalled}
	} else {
		return nil, err
	}
	return result, nil
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("windows.#"); ok {
		cfg := new(osThresholds)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "windows", 0)); err != nil {
			return err
		}
		me.PageFaultsPerSecondWindows = cfg.PageFaultsPerSecond
		me.UsedMemoryPercentageWindows = cfg.UsedMemoryPercentage
	}
	if _, ok := decoder.GetOk("linux.#"); ok {
		cfg := new(osThresholds)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "linux", 0)); err != nil {
			return err
		}
		me.PageFaultsPerSecondNonWindows = cfg.PageFaultsPerSecond
		me.UsedMemoryPercentageNonWindows = cfg.UsedMemoryPercentage
	}
	return nil
}
