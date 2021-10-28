package browser

import (
	"github.com/dtcookie/hcl"
)

type ClickEvent struct {
	EventBase
	Button   int            `json:"button"`             // the mouse button to be used for the click
	Wait     *WaitCondition `json:"wait,omitempty"`     // The wait condition for the event—defines how long Dynatrace should wait before the next action is executed
	Validate Validations    `json:"validate,omitempty"` // The validation rule for the event—helps you verify that your browser monitor loads the expected page content or page element
	Target   *Target        `json:"target,omitempty"`   // The tab on which the page should open
}

func (me *ClickEvent) GetType() EventType {
	return EventTypes.Click
}

func (me *ClickEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"button": {
			Type:        hcl.TypeInt,
			Description: "the mouse button to be used for the click",
			Required:    true,
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

func (me *ClickEvent) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["button"] = me.Button
	if me.Wait != nil {
		if marshalled, err := me.Wait.MarshalHCL(); err == nil {
			result["wait"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.Validate) > 0 {
		if marshalled, err := me.Validate.MarshalHCL(); err == nil {
			result["validate"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Target != nil {
		if marshalled, err := me.Target.MarshalHCL(); err == nil {
			result["target"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *ClickEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = EventTypes.Click
	if err := decoder.Decode("button", &me.Button); err != nil {
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
