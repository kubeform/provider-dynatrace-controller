package browser

import (
	"encoding/json"
	"fmt"

	"github.com/dtcookie/hcl"
)

type Events []Event

type EventWrapper struct {
	Event Event
}

func (me *EventWrapper) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the event to appear in the UI",
			Required:    true,
		},
		"select": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a key strokes event. ",
			MaxItems:    1,
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(SelectOptionEvent).Schema()},
		},
		"navigate": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a navigation event",
			Optional:    true,
			MaxItems:    1,
			// ConflictsWith: []string{"select"},
			Elem: &hcl.Resource{Schema: new(NavigateEvent).Schema()},
		},
		"keystrokes": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a key strokes event",
			Optional:    true,
			MaxItems:    1,
			// ConflictsWith: []string{"navigate", "select"},
			Elem: &hcl.Resource{Schema: new(KeyStrokesEvent).Schema()},
		},
		"javascript": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a javascript event",
			Optional:    true,
			MaxItems:    1,
			// ConflictsWith: []string{"keystrokes", "navigate", "select"},
			Elem: &hcl.Resource{Schema: new(JavascriptEvent).Schema()},
		},
		"cookie": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a cookie event",
			Optional:    true,
			MaxItems:    1,
			// ConflictsWith: []string{"javascript", "keystrokes", "navigate", "select"},
			Elem: &hcl.Resource{Schema: new(CookieEvent).Schema()},
		},
		"tap": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a tap event",
			Optional:    true,
			MaxItems:    1,
			// ConflictsWith: []string{"cookie", "javascript", "keystrokes", "navigate", "select"},
			Elem: &hcl.Resource{Schema: new(TapEvent).Schema()},
		},
		"click": {
			Type:        hcl.TypeList,
			Description: "Properties specified for a click event",
			Optional:    true,
			MaxItems:    1,
			// ConflictsWith: []string{"tap", "cookie", "javascript", "keystrokes", "navigate", "select"},
			Elem: &hcl.Resource{Schema: new(ClickEvent).Schema()},
		},
	}
}

func (me *EventWrapper) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["description"] = me.Event.GetDescription()
	if marshalled, err := me.Event.MarshalHCL(); err == nil {
		if me.Event.GetType() == EventTypes.Click {
			result["click"] = []interface{}{marshalled}
		} else if me.Event.GetType() == EventTypes.Tap {
			result["tap"] = []interface{}{marshalled}
		} else if me.Event.GetType() == EventTypes.Cookie {
			result["cookie"] = []interface{}{marshalled}
		} else if me.Event.GetType() == EventTypes.Javascript {
			result["javascript"] = []interface{}{marshalled}
		} else if me.Event.GetType() == EventTypes.KeyStrokes {
			result["keystrokes"] = []interface{}{marshalled}
		} else if me.Event.GetType() == EventTypes.Navigate {
			result["navigate"] = []interface{}{marshalled}
		} else if me.Event.GetType() == EventTypes.SelectOption {
			result["select"] = []interface{}{marshalled}
		} else {
			return nil, fmt.Errorf("events of type %s are not supported", me.Event.GetType())
		}
	}
	return result, nil
}

func (me *EventWrapper) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("click.#"); ok && result.(int) != 0 {
		evt := new(ClickEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "click", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.Click
		me.Event = evt
	}
	if result, ok := decoder.GetOk("tap.#"); ok && result.(int) != 0 {
		evt := new(TapEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "tap", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.Tap
		me.Event = evt
	}
	if result, ok := decoder.GetOk("cookie.#"); ok && result.(int) != 0 {
		evt := new(CookieEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "cookie", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.Cookie
		me.Event = evt
	}
	if result, ok := decoder.GetOk("javascript.#"); ok && result.(int) != 0 {
		evt := new(JavascriptEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "javascript", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.Javascript
		me.Event = evt
	}
	if result, ok := decoder.GetOk("keystrokes.#"); ok && result.(int) != 0 {
		evt := new(KeyStrokesEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "keystrokes", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.KeyStrokes
		me.Event = evt
	}
	if result, ok := decoder.GetOk("navigate.#"); ok && result.(int) != 0 {
		evt := new(NavigateEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "navigate", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.Navigate
		me.Event = evt
	}
	if result, ok := decoder.GetOk("select.#"); ok && result.(int) != 0 {
		evt := new(SelectOptionEvent)
		if err := evt.UnmarshalHCL(hcl.NewDecoder(decoder, "select", 0)); err != nil {
			return err
		}
		evt.Type = EventTypes.SelectOption
		me.Event = evt
	}
	if me.Event != nil {
		if v, ok := decoder.GetOk("description"); ok {
			me.Event.SetDescription(v.(string))
		}
	}
	return nil
}

func (me *Events) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"event": {
			Type:        hcl.TypeList,
			Description: "An event",
			Optional:    true,
			Elem:        &hcl.Resource{Schema: new(EventWrapper).Schema()},
		},
	}
}

func (me *Events) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("event.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			evtw := new(EventWrapper)
			if err := evtw.UnmarshalHCL(hcl.NewDecoder(decoder, "event", idx)); err != nil {
				return err
			}
			*me = append(*me, evtw.Event)
		}
	}
	return nil
}

func (me Events) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	for _, event := range me {
		evtw := &EventWrapper{Event: event}
		if marshalled, err := evtw.MarshalHCL(); err == nil {
			entries = append(entries, marshalled)
		} else {
			return nil, err
		}
	}
	if len(entries) > 0 {
		result["event"] = entries
	}
	return result, nil
}

type evt struct {
	Type EventType `json:"type"`
}

func (me *Events) UnmarshalJSON(data []byte) error {
	records := []json.RawMessage{}
	if err := json.Unmarshal(data, &records); err != nil {
		return err
	}
	for _, record := range records {
		var e evt
		if err := json.Unmarshal(record, &e); err != nil {
			return err
		}
		if e.Type == EventTypes.Click {
			var re ClickEvent
			if err := json.Unmarshal(record, &re); err != nil {
				return err
			}
			*me = append(*me, &re)
		} else if e.Type == EventTypes.Cookie {
			var re CookieEvent
			if err := json.Unmarshal(record, &re); err != nil {
				return err
			}
			*me = append(*me, &re)
		} else if e.Type == EventTypes.Javascript {
			var re JavascriptEvent
			if err := json.Unmarshal(record, &re); err != nil {
				return err
			}
			*me = append(*me, &re)
		} else if e.Type == EventTypes.KeyStrokes {
			var re KeyStrokesEvent
			if err := json.Unmarshal(record, &re); err != nil {
				return err
			}
			*me = append(*me, &re)
		} else if e.Type == EventTypes.Navigate {
			var re NavigateEvent
			if err := json.Unmarshal(record, &re); err != nil {
				return err
			}
			*me = append(*me, &re)
		} else if e.Type == EventTypes.Tap {
			var re TapEvent
			if err := json.Unmarshal(record, &re); err != nil {
				return err
			}
			*me = append(*me, &re)
		}
	}

	return nil
}

// func (me Events) MarshalJSON() ([]byte, error) {
// }

type Event interface {
	GetType() EventType
	GetDescription() string
	SetDescription(string)
	MarshalHCL() (map[string]interface{}, error)
}

type EventBase struct {
	Type        EventType `json:"type"`        // The type of synthetic event
	Description string    `json:"description"` // A short description of the event to appear in the UI
}

func (me *EventBase) GetType() EventType {
	return me.Type
}

func (me *EventBase) GetDescription() string {
	return me.Description
}

func (me *EventBase) SetDescription(description string) {
	me.Description = description
}

func (me *EventBase) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"description": {
			Type:        hcl.TypeString,
			Description: "A short description of the event to appear in the UI",
			Required:    true,
		},
	}
}
