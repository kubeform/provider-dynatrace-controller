package request

import (
	"github.com/dtcookie/hcl"
)

type HeadersSection struct {
	Headers      Headers  `json:"addHeaders"`
	Restrictions []string `json:"toRequests,omitempty"`
}

func (me *HeadersSection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"header": {
			Type:        hcl.TypeList,
			Description: "contains an HTTP header of the request",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Header).Schema()},
		},
		"restrictions": {
			Type:        hcl.TypeSet,
			Description: "Restrict applying headers to a set of URLs",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *HeadersSection) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me.Headers) > 0 {
		entries := []interface{}{}
		for _, header := range me.Headers {
			if marshalled, err := header.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["header"] = entries
	}
	if len(me.Restrictions) > 0 {
		result["restrictions"] = me.Restrictions
	}
	return result, nil
}

func (me *HeadersSection) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("header.#"); ok {
		me.Headers = Headers{}
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Header)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "header", idx)); err != nil {
				return err
			}
			me.Headers = append(me.Headers, entry)
		}
	}
	if err := decoder.Decode("restrictions", &me.Restrictions); err != nil {
		return err
	}
	return nil
}

// Headers is a list of request headers
type Headers []*Header

func (me *Headers) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"header": {
			Type:        hcl.TypeList,
			Description: "contains an HTTP header of the request",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Header).Schema()},
		},
	}
}

func (me Headers) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, header := range me {
			if marshalled, err := header.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["header"] = entries
	}
	return result, nil
}

func (me *Headers) UnmarshalHCL(decoder hcl.Decoder) error {
	if result, ok := decoder.GetOk("header.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Header)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "header", idx)); err != nil {
				return err
			}
			*me = append(*me, entry)
		}
	}
	return nil
}

// Header contains an HTTP header of the request
type Header struct {
	Name  string `json:"name"`  // The key of the header
	Value string `json:"value"` // The value of the header
}

func (me *Header) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The key of the header",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the header",
			Required:    true,
		},
	}
}

func (me *Header) MarshalHCL() (map[string]interface{}, error) {
	return map[string]interface{}{"name": me.Name, "value": me.Value}, nil
}

func (me *Header) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("name", &me.Name); err != nil {
		return err
	}
	if err := decoder.Decode("value", &me.Value); err != nil {
		return err
	}
	return nil
}
