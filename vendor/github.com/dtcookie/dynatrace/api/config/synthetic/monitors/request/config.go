package request

import (
	"github.com/dtcookie/hcl"
)

// Config contains the setup of the monitor
type Config struct {
	UserAgent            *string `json:"userAgent,omitempty"`            // The User agent of the request
	AcceptAnyCertificate *bool   `json:"acceptAnyCertificate,omitempty"` // If set to `false`, then the monitor fails with invalid SSL certificates.\n\nIf not set, the `false` option is used
	FollowRedirects      *bool   `json:"followRedirects,omitempty"`      // If set to `false`, redirects are reported as successful requests with response code 3xx.\n\nIf not set, the `false` option is used.
	RequestHeaders       Headers `json:"requestHeaders,omitempty"`       // By default, only the `User-Agent` header is set.\n\nYou can't set or modify this header here. Use the `userAgent` field for that.
}

func (me *Config) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"user_agent": {
			Type:        hcl.TypeString,
			Description: "The User agent of the request",
			Optional:    true,
		},
		"accept_any_certificate": {
			Type:        hcl.TypeBool,
			Description: "If set to `false`, then the monitor fails with invalid SSL certificates.\n\nIf not set, the `false` option is used",
			Optional:    true,
		},
		"follow_redirects": {
			Type:        hcl.TypeBool,
			Description: "If set to `false`, redirects are reported as successful requests with response code 3xx.\n\nIf not set, the `false` option is used.",
			Optional:    true,
		},
		"headers": {
			Type:        hcl.TypeList,
			Description: "The setup of the monitor",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Headers).Schema()},
		},
	}
}

func (me *Config) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.UserAgent != nil && len(*me.UserAgent) > 0 {
		result["user_agent"] = *me.UserAgent
	}
	if me.AcceptAnyCertificate != nil && *me.AcceptAnyCertificate {
		result["accept_any_certificate"] = *me.AcceptAnyCertificate
	}
	if me.FollowRedirects != nil && *me.FollowRedirects {
		result["follow_redirects"] = *me.FollowRedirects
	} else {
		result["follow_redirects"] = false
	}
	if len(me.RequestHeaders) > 0 {
		if marshalled, err := me.RequestHeaders.MarshalHCL(); err == nil {
			result["headers"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *Config) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("user_agent", &me.UserAgent); err != nil {
		return err
	}
	if err := decoder.Decode("accept_any_certificate", &me.AcceptAnyCertificate); err != nil {
		return err
	}
	if err := decoder.Decode("follow_redirects", &me.FollowRedirects); err != nil {
		return err
	}
	if _, ok := decoder.GetOk("headers.#"); ok {
		me.RequestHeaders = Headers{}
		if err := me.RequestHeaders.UnmarshalHCL(hcl.NewDecoder(decoder, "headers", 0)); err != nil {
			return err
		}
	}
	return nil
}
