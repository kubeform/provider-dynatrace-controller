package browser

import "github.com/dtcookie/hcl"

type JavascriptEvent struct {
	EventBase
	Javascript string         `json:"javaScript"`       // The JavaScript code to be executed in this event
	Wait       *WaitCondition `json:"wait,omitempty"`   // The wait condition for the event—defines how long Dynatrace should wait before the next action is executed
	Target     *Target        `json:"target,omitempty"` // The tab on which the page should open
}

func (me *JavascriptEvent) GetType() EventType {
	return EventTypes.Javascript
}

func (me *JavascriptEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"code": {
			Type:        hcl.TypeString,
			Description: "The JavaScript code to be executed in this event",
			Required:    true,
		},
		"wait": {
			Type:        hcl.TypeList,
			Description: "The wait condition for the event—defines how long Dynatrace should wait before the next action is executed",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(WaitCondition).Schema()},
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

func (me *JavascriptEvent) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["code"] = me.Javascript
	if me.Wait != nil {
		if marshalled, err := me.Wait.MarshalHCL(); err == nil {
			result["wait"] = []interface{}{marshalled}
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

func (me *JavascriptEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = EventTypes.Tap
	if err := decoder.Decode("code", &me.Javascript); err != nil {
		return err
	}
	if err := decoder.Decode("wait", &me.Wait); err != nil {
		return err
	}
	if err := decoder.Decode("target", &me.Target); err != nil {
		return err
	}
	return nil
}
