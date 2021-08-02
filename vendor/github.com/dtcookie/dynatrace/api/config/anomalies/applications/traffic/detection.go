package traffic

import (
	"github.com/dtcookie/hcl"
)

type Detection struct {
	Drops  *DropDetection  // The configuration of traffic drops detection.
	Spikes *SpikeDetection // The configuration of traffic spikes detection.
}

func (me *Detection) IsEmpty() bool {
	if me.Drops != nil {
		if me.Drops.Enabled {
			return false
		}
		me.Drops = nil
	}
	if me.Spikes != nil {
		if me.Spikes.Enabled {
			return false
		}
		me.Spikes = nil
	}
	return true
}

func (me *Detection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"drops": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration of traffic drops detection",
			Elem:        &hcl.Resource{Schema: new(DropDetection).Schema()},
		},
		"spikes": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The configuration of traffic spikes detection",
			Elem:        &hcl.Resource{Schema: new(SpikeDetection).Schema()},
		},
	}
}

func (me *Detection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if me.Drops != nil && me.Drops.Enabled {
		if marshalled, err := me.Drops.MarshalHCL(hcl.NewDecoder(decoder, "drops", 0)); err == nil {
			result["drops"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Spikes != nil && me.Spikes.Enabled {
		if marshalled, err := me.Spikes.MarshalHCL(hcl.NewDecoder(decoder, "spikes", 0)); err == nil {
			result["spikes"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Detection) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("drops.#"); ok {
		me.Drops = new(DropDetection)
		if err := me.Drops.UnmarshalHCL(hcl.NewDecoder(decoder, "drops", 0)); err != nil {
			return err
		}
		if !me.Drops.Enabled {
			me.Drops = nil
		}
	}
	if _, ok := decoder.GetOk("spikes.#"); ok {
		me.Spikes = new(SpikeDetection)
		if err := me.Spikes.UnmarshalHCL(hcl.NewDecoder(decoder, "spikes", 0)); err != nil {
			return err
		}
		if !me.Spikes.Enabled {
			me.Spikes = nil
		}
	}

	return nil
}
