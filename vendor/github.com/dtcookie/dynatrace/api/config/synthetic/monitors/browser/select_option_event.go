package browser

import (
	"github.com/dtcookie/hcl"
)

type SelectOptionEvent struct {
	EventBase
	Selections ListOptions    `json:"selections"`         // The options to be selected
	Wait       *WaitCondition `json:"wait,omitempty"`     // The wait condition for the event—defines how long Dynatrace should wait before the next action is executed
	Validate   Validations    `json:"validate,omitempty"` // The validation rule for the event—helps you verify that your browser monitor loads the expected page content or page element
	Target     *Target        `json:"target,omitempty"`   // The tab on which the page should open
}

func (me *SelectOptionEvent) GetType() EventType {
	return EventTypes.SelectOption
}

func (me *SelectOptionEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"selections": {
			Type:        hcl.TypeList,
			Description: "The options to be selected",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ListOptions).Schema()},
		},
		"wait": {
			Type:        hcl.TypeList,
			Description: "The wait condition for the event—defines how long Dynatrace should wait before the next action is executed",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(WaitCondition).Schema()},
		},
		"validate": {
			Type:        hcl.TypeList,
			Description: "The validation rules for the event—helps you verify that your browser monitor loads the expected page content or page element",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Validations).Schema()},
		},
		"target": {
			Type:        hcl.TypeList,
			Description: "The tab on which the page should open",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Target).Schema()},
		},
	}
}

func (me *SelectOptionEvent) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.Target != nil {
		if marshalled, err := me.Target.MarshalHCL(); err == nil {
			result["target"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Wait != nil {
		if marshalled, err := me.Wait.MarshalHCL(); err == nil {
			result["wait"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Validate) > 0 {
		if marshalled, err := me.Wait.MarshalHCL(); err == nil {
			result["validate"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Selections) > 0 {
		if marshalled, err := me.Selections.MarshalHCL(); err == nil {
			result["selections"] = []interface{}{marshalled}
		} else {
			return nil, err
		}

	}

	return result, nil
}

func (me *SelectOptionEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = EventTypes.Tap
	if err := decoder.Decode("selections", &me.Selections); err != nil {
		return err
	}
	if err := decoder.Decode("wait", &me.Wait); err != nil {
		return err
	}
	if err := decoder.Decode("validate", &me.Validate); err != nil {
		return err
	}
	if err := decoder.Decode("target", &me.Target); err != nil {
		return err
	}
	return nil
}
