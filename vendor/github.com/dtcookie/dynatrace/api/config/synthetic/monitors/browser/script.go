package browser

import (
	"github.com/dtcookie/hcl"
)

type Script struct {
	Version       string        `json:"version"`                 // Script version—use the `1.0` value here
	Type          ScriptType    `json:"type"`                    // The type of monitor. Possible values are `clickpath` for clickpath monitors and `availability` for single-URL browser monitors. These monitors are only allowed to have one event of the `navigate` type
	Configuration *ScriptConfig `json:"configuration,omitempty"` // The setup of the monitor
	Events        Events        `json:"events,omitempty"`        // Steps of the clickpath—the first step must always be of the `navigate` type
}

func (me *Script) GetVersion() string {
	return "1.0"
}

func (me *Script) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of monitor. Possible values are `clickpath` for clickpath monitors and `availability` for single-URL browser monitors. These monitors are only allowed to have one event of the `navigate` type",
			Required:    true,
		},
		"configuration": {
			Type:        hcl.TypeList,
			Description: "The setup of the monitor",
			Optional:    true,
			MaxItems:    1,
			Elem: &hcl.Resource{
				Schema: new(ScriptConfig).Schema(),
			},
		},
		"events": {
			Type:        hcl.TypeList,
			Description: "Steps of the clickpath—the first step must always be of the `navigate` type",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Events).Schema()},
		},
	}
}

func (me *Script) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.Type)
	if me.Configuration != nil {
		if marshalled, err := me.Configuration.MarshalHCL(); err == nil {
			result["configuration"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Events) > 0 {
		if marshalled, err := me.Events.MarshalHCL(); err == nil {
			result["events"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Script) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Version = me.GetVersion()
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("configuration", &me.Configuration); err != nil {
		return err
	}
	if err := decoder.Decode("events", &me.Events); err != nil {
		return err
	}
	return nil
}
