package web

import "github.com/dtcookie/hcl"

// ApplicationErrorRules represents configuration of error rules in the web application
type ApplicationErrorRules struct {
	WebApplicationID                         string           `json:"-"`                                        // The EntityID of the the WebApplication
	IgnoreJavaScriptErrorsInApdexCalculation bool             `json:"ignoreJavaScriptErrorsInApdexCalculation"` // Exclude (`true`) or include (`false`) JavaScript errors in Apdex calculation
	IgnoreHttpErrorsInApdexCalculation       bool             `json:"ignoreHttpErrorsInApdexCalculation"`       // Exclude (`true`) or include (`false`) HTTP errors listed in **httpErrorRules** in Apdex calculation
	IgnoreCustomErrorsInApdexCalculation     bool             `json:"ignoreCustomErrorsInApdexCalculation"`     // Exclude (`true`) or include (`false`) custom errors listed in **customErrorRules** in Apdex calculation
	HTTPErrors                               HTTPErrorRules   `json:"httpErrorRules"`                           // An ordered list of HTTP errors.\n\n Rules are evaluated from top to bottom; the first matching rule applies
	CustomErrors                             CustomErrorRules `json:"customErrorRules"`                         // An ordered list of custom errors.\n\n Rules are evaluated from top to bottom; the first matching rule applies
}

func (me *ApplicationErrorRules) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"web_application_id": {
			Type:        hcl.TypeString,
			Description: "The EntityID of the the WebApplication",
			Optional:    true,
		},
		"ignore_js_errors_apdex": {
			Type:        hcl.TypeBool,
			Description: "Exclude (`true`) or include (`false`) JavaScript errors in Apdex calculation",
			Optional:    true,
		},
		"ignore_http_errors_apdex": {
			Type:        hcl.TypeBool,
			Description: "Exclude (`true`) or include (`false`) HTTP errors listed in **httpErrorRules** in Apdex calculation",
			Optional:    true,
		},
		"ignore_custom_errors_apdex": {
			Type:        hcl.TypeBool,
			Description: "Exclude (`true`) or include (`false`) custom errors listed in **customErrorRules** in Apdex calculation",
			Optional:    true,
		},
		"http_errors": {
			Type:        hcl.TypeList,
			Description: "An ordered list of HTTP errors.\n\n Rules are evaluated from top to bottom; the first matching rule applies",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(HTTPErrorRules).Schema()},
		},
		"custom_errors": {
			Type:        hcl.TypeList,
			Description: "An ordered list of HTTP errors.\n\n Rules are evaluated from top to bottom; the first matching rule applies",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(CustomErrorRules).Schema()},
		},
	}
}

func (me *ApplicationErrorRules) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"web_application_id":         me.WebApplicationID,
		"ignore_js_errors_apdex":     me.IgnoreJavaScriptErrorsInApdexCalculation,
		"ignore_http_errors_apdex":   me.IgnoreHttpErrorsInApdexCalculation,
		"ignore_custom_errors_apdex": me.IgnoreCustomErrorsInApdexCalculation,
		"http_errors":                me.HTTPErrors,
		"custom_errors":              me.CustomErrors,
	})
}

func (me *ApplicationErrorRules) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"web_application_id":         &me.WebApplicationID,
		"ignore_js_errors_apdex":     &me.IgnoreJavaScriptErrorsInApdexCalculation,
		"ignore_http_errors_apdex":   &me.IgnoreHttpErrorsInApdexCalculation,
		"ignore_custom_errors_apdex": &me.IgnoreCustomErrorsInApdexCalculation,
		"http_errors":                &me.HTTPErrors,
		"custom_errors":              &me.CustomErrors,
	})
}
