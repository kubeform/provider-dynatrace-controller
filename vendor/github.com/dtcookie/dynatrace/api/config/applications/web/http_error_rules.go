package web

import "github.com/dtcookie/hcl"

type HTTPErrorRules []*HTTPErrorRule

func (me *HTTPErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"rule": {
			Type:        hcl.TypeList,
			Description: "Configuration of the HTTP error in the web application",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(HTTPErrorRule).Schema()},
		},
	}
}

func (me HTTPErrorRules) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["rule"] = entries
	}
	return result, nil
}

// HTTPErrorRule represents configuration of the HTTP error in the web application
type HTTPErrorRule struct {
	ConsiderUnknownErrorCode bool                 `json:"considerUnknownErrorCode"` // If `true`, match by errors that have unknown HTTP status code
	ConsiderBlockedRequests  bool                 `json:"considerBlockedRequests"`  // If `true`, match by errors that have CSP Rule violations
	ErrorCodes               *string              `json:"errorCodes,omitempty"`     // The HTTP status code or status code range to match by. \n\nThis field is required if **considerUnknownErrorCode** AND **considerBlockedRequests** are both set to `false`
	FilterByURL              bool                 `json:"filterByUrl"`              // If `true`, filter errors by URL
	Filter                   *HTTPErrorRuleFilter `json:"filter,omitempty"`         // The matching rule for the URL. Popssible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH` and `EQUALS`.
	URL                      *string              `json:"url,omitempty"`            // The URL to look for
	Capture                  bool                 `json:"capture"`                  // Capture (`true`) or ignore (`false`) the error
	ImpactApdex              bool                 `json:"impactApdex"`              // Include (`true`) or exclude (`false`) the error in Apdex calculation
	ConsiderForAI            bool                 `json:"considerForAi"`            // Include (`true`) or exclude (`false`) the error in Davis AI [problem detection and analysis](https://dt-url.net/a963kd2)
}

func (me *HTTPErrorRule) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"consider_unknown_error_code": {
			Type:        hcl.TypeBool,
			Description: "If `true`, match by errors that have unknown HTTP status code",
			Optional:    true,
		},
		"consider_blocked_requests": {
			Type:        hcl.TypeBool,
			Description: "If `true`, match by errors that have CSP Rule violations",
			Optional:    true,
		},
		"error_codes": {
			Type:        hcl.TypeString,
			Description: "The HTTP status code or status code range to match by. \n\nThis field is required if **considerUnknownErrorCode** AND **considerBlockedRequests** are both set to `false`",
			Optional:    true,
		},
		"filter_by_url": {
			Type:        hcl.TypeBool,
			Description: "If `true`, filter errors by URL",
			Optional:    true,
		},
		"filter": {
			Type:        hcl.TypeString,
			Description: "The matching rule for the URL. Popssible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH` and `EQUALS`.",
			Optional:    true,
		},
		"url": {
			Type:        hcl.TypeString,
			Description: "The URL to look for",
			Optional:    true,
		},
		"capture": {
			Type:        hcl.TypeBool,
			Description: "Capture (`true`) or ignore (`false`) the error",
			Optional:    true,
		},
		"impact_apdex": {
			Type:        hcl.TypeBool,
			Description: "Include (`true`) or exclude (`false`) the error in Apdex calculation",
			Optional:    true,
		},
		"consider_for_ai": {
			Type:        hcl.TypeBool,
			Description: "Include (`true`) or exclude (`false`) the error in Davis AI [problem detection and analysis](https://dt-url.net/a963kd2)",
			Optional:    true,
		},
	}
}

func (me *HTTPErrorRule) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"consider_unknown_error_code": me.ConsiderUnknownErrorCode,
		"consider_blocked_requests":   me.ConsiderBlockedRequests,
		"error_codes":                 me.ErrorCodes,
		"filter_by_url":               me.FilterByURL,
		"filter":                      me.Filter,
		"url":                         me.URL,
		"capture":                     me.Capture,
		"impact_apdex":                me.ImpactApdex,
		"consider_for_ai":             me.ConsiderForAI,
	})
}

func (me *HTTPErrorRule) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"consider_unknown_error_code": &me.ConsiderUnknownErrorCode,
		"consider_blocked_requests":   &me.ConsiderBlockedRequests,
		"error_codes":                 &me.ErrorCodes,
		"filter_by_url":               &me.FilterByURL,
		"filter":                      &me.Filter,
		"url":                         &me.URL,
		"capture":                     &me.Capture,
		"impact_apdex":                &me.ImpactApdex,
		"consider_for_ai":             &me.ConsiderForAI,
	})
}
