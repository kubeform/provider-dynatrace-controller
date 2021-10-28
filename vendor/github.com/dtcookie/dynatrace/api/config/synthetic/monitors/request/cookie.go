package request

import "github.com/dtcookie/hcl"

// Cookies contains the list of cookies to be created for the monitor. Every cookie must be unique within the list. However, you can use the same cookie again in other event
type Cookies []*Cookie

func (me *Cookies) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"cookie": {
			Type:        hcl.TypeList,
			Description: "A request cookie",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(Cookie).Schema()},
		},
	}
}

func (me Cookies) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	for _, cookie := range me {
		if marshalled, err := cookie.MarshalHCL(); err == nil {
			entries = append(entries, marshalled)
		} else {
			return nil, err
		}
	}
	result["cookie"] = entries
	return result, nil
}

func (me *Cookies) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("cookie", me); err != nil {
		return err
	}
	return nil
}

// Cookie a request cookie
type Cookie struct {
	Name   string  `json:"name"`           // The name of the cookie. The following cookie names are now allowed: `dtCookie`, `dtLatC`, `dtPC`, `rxVisitor`, `rxlatency`, `rxpc`, `rxsession` and `rxvt`
	Value  string  `json:"value"`          // The value of the cookie. The following symbols are not allowed: `;`, `,`, `\` and `"`.
	Domain string  `json:"domain"`         // The domain of the cookie
	Path   *string `json:"path,omitempty"` // The path of the cookie
}

func (me *Cookie) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the cookie. The following cookie names are now allowed: `dtCookie`, `dtLatC`, `dtPC`, `rxVisitor`, `rxlatency`, `rxpc`, `rxsession` and `rxvt`",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the cookie. The following symbols are not allowed: `;`, `,`, `\\` and `\"`.",
			Required:    true,
		},
		"domain": {
			Type:        hcl.TypeString,
			Description: "The domain of the cookie.",
			Required:    true,
		},
		"path": {
			Type:        hcl.TypeString,
			Description: "The path of the cookie.",
			Optional:    true,
		},
	}
}

func (me *Cookie) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["name"] = me.Name
	result["value"] = me.Value
	result["domain"] = me.Domain
	if me.Path != nil {
		result["path"] = *me.Path
	}
	return result, nil
}

func (me *Cookie) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("name", &me.Name); err != nil {
		return err
	}
	if err := decoder.Decode("value", &me.Value); err != nil {
		return err
	}
	if err := decoder.Decode("domain", &me.Domain); err != nil {
		return err
	}
	if err := decoder.Decode("path", &me.Path); err != nil {
		return err
	}
	return nil
}
