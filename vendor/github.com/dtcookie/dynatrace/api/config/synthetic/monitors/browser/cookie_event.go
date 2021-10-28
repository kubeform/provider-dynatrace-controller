package browser

import (
	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors/request"
	"github.com/dtcookie/hcl"
)

type CookieEvent struct {
	EventBase
	Cookies request.Cookies `json:"cookies"` // Every cookie must be unique within the list. However, you can use the same cookie again in other event
}

func (me *CookieEvent) GetType() EventType {
	return EventTypes.Cookie
}

func (me *CookieEvent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cookies": {
			Type:        hcl.TypeList,
			Description: "Every cookie must be unique within the list. However, you can use the same cookie again in other event",
			Required:    true,
			MaxItems:    1,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(request.Cookies).Schema()},
		},
	}
}

func (me *CookieEvent) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.Cookies) > 0 {
		if marshalled, err := me.Cookies.MarshalHCL(); err == nil {
			result["cookies"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *CookieEvent) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Type = EventTypes.Tap
	if err := decoder.Decode("cookies", &me.Cookies); err != nil {
		return err
	}
	return nil
}
