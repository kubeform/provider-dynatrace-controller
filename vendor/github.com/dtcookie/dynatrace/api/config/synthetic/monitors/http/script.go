package http

import "github.com/dtcookie/hcl"

type Script struct {
	Version  string   `json:"version"`  // Script version—use the `1.0` value here
	Requests Requests `json:"requests"` // A list of HTTP requests to be performed by the monitor.\n\nThe requests are executed in the order in which they appear in the script
}

func (me *Script) GetVersion() string {
	return "1.0"
}

func (me *Script) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"request": {
			Type:        hcl.TypeList,
			Description: "A HTTP request to be performed by the monitor.",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Request).Schema()},
		},
	}
}

func (me *Script) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.Requests) > 0 {
		entries := []interface{}{}
		for _, request := range me.Requests {
			if marshalled, err := request.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["request"] = entries
	}
	return result, nil
}

func (me *Script) UnmarshalHCL(decoder hcl.Decoder) error {
	me.Version = me.GetVersion()
	if result, ok := decoder.GetOk("request.#"); ok {
		me.Requests = Requests{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Request)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "request", idx)); err != nil {
				return err
			}
			me.Requests = append(me.Requests, entry)
		}
	}
	return nil
}
