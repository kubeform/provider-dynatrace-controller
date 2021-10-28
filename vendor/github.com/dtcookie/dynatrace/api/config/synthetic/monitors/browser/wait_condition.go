package browser

import "github.com/dtcookie/hcl"

type WaitCondition struct {
	WaitFor               string      `json:"waitFor"`                         // The time to wait before the next event is triggered. Possible values are `page_complete` (wait for the page to load completely), `network` (wait for background network activity to complete), `next_action` (wait for the next action), `time` (wait for a specified periodof time) and `validation` (wait for a specific element to appear)
	Milliseconds          *int        `json:"milliseconds,omitempty"`          // The time to wait, in millisencods. The maximum allowed value is `60000`. Required for the type `time`, not applicable otherwise.
	TimeoutInMilliseconds *int        `json:"timeoutInMilliseconds,omitempty"` // The maximum amount of time to wait for a certain element to appear, in milliseconds—if exceeded, the action is marked as failed.\nThe maximum allowed value is 60000. Required for the type `validation`, not applicable otherwise.
	Validation            *Validation `json:"validation,omitempty"`            // The element to wait for. Required for the `validation` type, not applicable otherwise.
}

func (me *WaitCondition) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"wait_for": {
			Type:        hcl.TypeString,
			Description: "The time to wait before the next event is triggered. Possible values are `page_complete` (wait for the page to load completely), `network` (wait for background network activity to complete), `next_action` (wait for the next action), `time` (wait for a specified periodof time) and `validation` (wait for a specific element to appear)",
			Required:    true,
		},
		"milliseconds": {
			Type:        hcl.TypeInt,
			Description: "The time to wait, in millisencods. The maximum allowed value is `60000`. Required for the type `time`, not applicable otherwise.",
			Optional:    true,
		},
		"timeout": {
			Type:        hcl.TypeInt,
			Description: "he maximum amount of time to wait for a certain element to appear, in milliseconds—if exceeded, the action is marked as failed.\nThe maximum allowed value is 60000. Required for the type `validation`, not applicable otherwise..",
			Optional:    true,
		},
		"validation": {
			Type:        hcl.TypeList,
			Description: "The elements to wait for. Required for the `validation` type, not applicable otherwise.",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Validation).Schema()},
		},
	}
}

func (me *WaitCondition) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["wait_for"] = me.WaitFor
	if me.Milliseconds != nil {
		result["milliseconds"] = *me.Milliseconds
	}
	if me.TimeoutInMilliseconds != nil {
		result["timeout"] = *me.TimeoutInMilliseconds
	}
	if me.Validation != nil {
		if marshalled, err := me.Validation.MarshalHCL(); err == nil {
			result["validation"] = []interface{}{marshalled}
		}
	}
	return result, nil
}

func (me *WaitCondition) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("wait_for", &me.WaitFor); err != nil {
		return err
	}
	if err := decoder.Decode("milliseconds", &me.Milliseconds); err != nil {
		return err
	}
	if err := decoder.Decode("timeout", &me.TimeoutInMilliseconds); err != nil {
		return err
	}
	if err := decoder.Decode("validation", &me.Validation); err != nil {
		return err
	}
	return nil
}
