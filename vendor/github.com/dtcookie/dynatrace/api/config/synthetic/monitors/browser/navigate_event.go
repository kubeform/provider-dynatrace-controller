package browser

import (
	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors/browser/auth"
	"github.com/dtcookie/hcl"
)

type NavigateEvent struct {
	EventBase
	URL            string            `json:"url"`                      // The URL to navigate to
	Wait           *WaitCondition    `json:"wait,omitempty"`           // The wait condition for the event—defines how long Dynatrace should wait before the next action is executed
	Validate       Validations       `json:"validate,omitempty"`       // The validation rule for the event—helps you verify that your browser monitor loads the expected page content or page element
	Target         *Target           `json:"target,omitempty"`         // The tab on which the page should open
	Authentication *auth.Credentials `json:"authentication,omitempty"` // The login credentials to bypass the browser login mask
}

func (me *NavigateEvent) GetType() EventType {
	return EventTypes.Navigate
}

func (me *NavigateEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL to navigate to",
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
		"authentication": {
			Type:        hcl.TypeList,
			Description: "The login credentials to bypass the browser login mask",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(auth.Credentials).Schema()},
		},
	}
}

func (me *NavigateEvent) MarshalHCL() (map[string]interface{}, error) {
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
		if marshalled, err := me.Validate.MarshalHCL(); err == nil {
			result["validate"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	result["url"] = me.URL
	if me.Authentication != nil {
		if marshalled, err := me.Authentication.MarshalHCL(); err == nil {
			result["authentication"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *NavigateEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = EventTypes.Tap
	if err := decoder.Decode("url", &me.URL); err != nil {
		return err
	}
	if err := decoder.Decode("authentication", &me.Authentication); err != nil {
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
