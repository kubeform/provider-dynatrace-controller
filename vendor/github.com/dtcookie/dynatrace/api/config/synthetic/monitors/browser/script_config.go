package browser

import (
	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors/request"
	"github.com/dtcookie/hcl"
)

// ScriptConfig contains the setup of the monitor
type ScriptConfig struct {
	UserAgent          *string                 `json:"userAgent,omitempty"`          // The user agent of the request
	Device             *Device                 `json:"device,omitempty"`             // The emulated device of the monitor—holds either the parameters of the custom device or the name and orientation of the preconfigured device.\n\nIf not set, then the Desktop preconfigured device is used
	Bandwidth          *Bandwidth              `json:"bandwidth,omitempty"`          // The emulated network conditions of the monitor.\n\nIf not set, then the full available bandwidth is used
	RequestHeaders     *request.HeadersSection `json:"requestHeaders,omitempty"`     // The list of HTTP headers to be sent with requests of the monitor
	Cookies            request.Cookies         `json:"cookies,omitempty"`            // These cookies are added before execution of the first step
	BlockRequests      []string                `json:"blockRequests,omitempty"`      // Block these URLs
	BypassCSP          bool                    `json:"bypassCSP"`                    // Bypass Content Security Policy of monitored pages
	MonitorFrames      *MonitorFrames          `json:"monitorFrames,omitempty"`      // Capture performance metrics for pages loaded in frames
	JavascriptSettings *JavascriptSettings     `json:"javaScriptSettings,omitempty"` // Custom JavaScript Agent settings
	DisableWebSecurity bool                    `json:"disableWebSecurity"`           // No documentation available
	IgnoredErrorCodes  *IgnoredErrorCodes      `json:"ignoredErrorCodes"`            // Ignore specific status codes
}

func (me *ScriptConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"disable_web_security": {
			Type:        hcl.TypeBool,
			Description: "No documentation available",
			Optional:    true,
		},
		"bypass_csp": {
			Type:        hcl.TypeBool,
			Description: "Bypass Content Security Policy of monitored pages",
			Optional:    true,
		},
		"monitor_frames": {
			Type:        hcl.TypeBool,
			Description: "Capture performance metrics for pages loaded in frames",
			Optional:    true,
		},
		"user_agent": {
			Type:        hcl.TypeString,
			Description: "The user agent of the request",
			Optional:    true,
		},
		"ignored_error_codes": {
			Type:        hcl.TypeList,
			Description: "Ignore specific status codes",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(IgnoredErrorCodes).Schema()},
		},
		"javascript_setttings": {
			Type:        hcl.TypeList,
			Description: "Custom JavaScript Agent settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(JavascriptSettings).Schema()},
		},
		"device": {
			Type:        hcl.TypeList,
			Description: "The emulated device of the monitor—holds either the parameters of the custom device or the name and orientation of the preconfigured device.\n\nIf not set, then the Desktop preconfigured device is used",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Device).Schema()},
		},
		"bandwidth": {
			Type:        hcl.TypeList,
			Description: "The emulated device of the monitor—holds either the parameters of the custom device or the name and orientation of the preconfigured device.\n\nIf not set, then the Desktop preconfigured device is used",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(Bandwidth).Schema()},
		},
		"headers": {
			Type:        hcl.TypeList,
			Description: "The list of HTTP headers to be sent with requests of the monitor",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(request.HeadersSection).Schema()},
		},
		"block": {
			Type:        hcl.TypeSet,
			Description: "Block these URLs",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"cookies": {
			Type:        hcl.TypeList,
			Description: "These cookies are added before execution of the first step",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(request.Cookies).Schema()},
		},
	}
}

func (me *ScriptConfig) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["disable_web_security"] = me.DisableWebSecurity
	result["bypass_csp"] = me.BypassCSP
	if me.MonitorFrames != nil && me.MonitorFrames.Enabled {
		result["monitor_frames"] = me.MonitorFrames.Enabled
	} else {
		result["monitor_frames"] = false
	}
	if me.UserAgent != nil {
		result["user_agent"] = me.UserAgent
	}
	if me.IgnoredErrorCodes != nil {
		if marshalled, err := me.IgnoredErrorCodes.MarshalHCL(); err == nil {
			result["ignored_error_codes"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.JavascriptSettings != nil {
		if marshalled, err := me.JavascriptSettings.MarshalHCL(); err == nil {
			result["javascript_setttings"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Device != nil {
		if marshalled, err := me.Device.MarshalHCL(); err == nil {
			result["device"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Bandwidth != nil {
		if marshalled, err := me.Bandwidth.MarshalHCL(); err == nil {
			result["bandwidth"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.RequestHeaders != nil {
		if marshalled, err := me.RequestHeaders.MarshalHCL(); err == nil {
			result["headers"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if len(me.BlockRequests) > 0 {
		result["block"] = me.BlockRequests
	}
	if len(me.Cookies) > 0 {
		if marshalled, err := me.Cookies.MarshalHCL(); err == nil {
			result["cookies"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *ScriptConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("disable_web_security", &me.DisableWebSecurity); err != nil {
		return err
	}
	me.BlockRequests = decoder.GetStringSet("block")
	if err := decoder.Decode("bypass_csp", &me.BypassCSP); err != nil {
		return err
	}
	monitorFrames := false
	if err := decoder.Decode("monitor_frames", &monitorFrames); err != nil {
		return err
	}
	if monitorFrames {
		me.MonitorFrames = &MonitorFrames{Enabled: true}
	}
	if err := decoder.Decode("user_agent", &me.UserAgent); err != nil {
		return err
	}
	if err := decoder.Decode("ignored_error_codes", &me.IgnoredErrorCodes); err != nil {
		return err
	}
	if err := decoder.Decode("javascript_setttings", &me.JavascriptSettings); err != nil {
		return err
	}
	if err := decoder.Decode("device", &me.Device); err != nil {
		return err
	}
	if err := decoder.Decode("bandwidth", &me.Bandwidth); err != nil {
		return err
	}
	if _, ok := decoder.GetOk("headers.#"); ok {
		me.RequestHeaders = new(request.HeadersSection)
		if err := me.RequestHeaders.UnmarshalHCL(hcl.NewDecoder(decoder, "headers", 0)); err != nil {
			return err
		}
	}
	if err := decoder.Decode("headers", &me.RequestHeaders); err != nil {
		return err
	}
	if err := decoder.Decode("cookies", &me.Cookies); err != nil {
		return err
	}
	return nil
}

type IgnoredErrorCodes struct {
	StatusCodes              string  `json:"statusCodes"`                        // You can use exact number, range or status class mask. Multiple values can be separated by comma, i.e. 404, 405-410, 5xx
	MatchingDocumentRequests *string `json:"matchingDocumentRequests,omitempty"` // Only apply to document request matching this regex
}

func (me *IgnoredErrorCodes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"status_codes": {
			Type:        hcl.TypeString,
			Description: "You can use exact number, range or status class mask. Multiple values can be separated by comma, i.e. 404, 405-410, 5xx",
			Required:    true,
		},
		"matching_document_requests": {
			Type:        hcl.TypeString,
			Description: "Only apply to document request matching this regex",
			Optional:    true,
		},
	}
}

func (me *IgnoredErrorCodes) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["status_codes"] = me.StatusCodes
	if me.MatchingDocumentRequests != nil {
		result["matching_document_requests"] = *me.MatchingDocumentRequests
	}
	return result, nil
}

func (me *IgnoredErrorCodes) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("status_codes", &me.StatusCodes); err != nil {
		return err
	}
	if err := decoder.Decode("matching_document_requests", &me.MatchingDocumentRequests); err != nil {
		return err
	}
	return nil
}

type JavascriptSettings struct {
	TimeoutSettings         *TimeoutSettings         `json:"timeoutSettings"`
	CustomProperties        *string                  `json:"customProperties"`
	VisuallyCompleteOptions *VisuallyCompleteOptions `json:"visuallyCompleteOptions"`
}

func (me *JavascriptSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"timeout_settings": {
			Type:        hcl.TypeList,
			Description: "Custom JavaScript Agent settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(TimeoutSettings).Schema()},
		},
		"custom_properties": {
			Type:        hcl.TypeString,
			Description: "Additional Javascript Agent Properties",
			Optional:    true,
		},
		"visually_complete_options": {
			Type:        hcl.TypeList,
			Description: "Parameters for Visually complete and Speed index calculation",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(VisuallyCompleteOptions).Schema()},
		},
	}
}

func (me *JavascriptSettings) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.TimeoutSettings != nil {
		if marshalled, err := me.TimeoutSettings.MarshalHCL(); err == nil {
			result["timeout_settings"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.VisuallyCompleteOptions != nil {
		if marshalled, err := me.VisuallyCompleteOptions.MarshalHCL(); err == nil {
			result["visually_complete_options"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.CustomProperties != nil && len(*me.CustomProperties) > 0 {
		result["custom_properties"] = *me.CustomProperties
	}
	return result, nil
}

func (me *JavascriptSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("timeout_settings", &me.TimeoutSettings); err != nil {
		return err
	}
	if err := decoder.Decode("visually_complete_options", &me.VisuallyCompleteOptions); err != nil {
		return err
	}
	if err := decoder.Decode("custom_properties", &me.CustomProperties); err != nil {
		return err
	}
	return nil
}

type VisuallyCompleteOptions struct {
	ImageSizeThreshold int      `json:"imageSizeThreshold"` // Use this setting to define the minimum visible area per element (in pixels) for an element to be counted towards Visually complete and Speed index
	InactivityTimeout  int      `json:"inactivityTimeout"`  // The time the Visually complete module waits for inactivity and no further mutations on the page after the load action
	MutationTimeout    int      `json:"mutationTimeout"`    // The time the Visually complete module waits after an XHR or custom action closes to start the calculation
	ExcludedURLs       []string `json:"excludedUrls"`       // Use regular expressions to define URLs for images and iFrames to exclude from detection by the Visually complete module
	ExcludedElements   []string `json:"excludedElements"`   // Query CSS selectors to specify mutation nodes (elements that change) to ignore in Visually complete and Speed index calculation
}

func (me *VisuallyCompleteOptions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"image_size_threshold": {
			Type:        hcl.TypeInt,
			Description: "Use this setting to define the minimum visible area per element (in pixels) for an element to be counted towards Visually complete and Speed index",
			Required:    true,
		},
		"inactivity_timeout": {
			Type:        hcl.TypeInt,
			Description: "The time the Visually complete module waits for inactivity and no further mutations on the page after the load action",
			Required:    true,
		},
		"mutation_timeout": {
			Type:        hcl.TypeInt,
			Description: "The time the Visually complete module waits after an XHR or custom action closes to start the calculation",
			Required:    true,
		},
		"excluded_urls": {
			Type:        hcl.TypeList,
			Description: "Parameters for Visually complete and Speed index calculation",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"excluded_elements": {
			Type:        hcl.TypeList,
			Description: "Query CSS selectors to specify mutation nodes (elements that change) to ignore in Visually complete and Speed index calculation",
			Optional:    true,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
	}
}

func (me *VisuallyCompleteOptions) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["image_size_threshold"] = me.ImageSizeThreshold
	result["inactivity_timeout"] = me.InactivityTimeout
	result["mutation_timeout"] = me.MutationTimeout
	if len(me.ExcludedURLs) > 0 {
		result["excluded_urls"] = me.ExcludedURLs
	}
	if len(me.ExcludedElements) > 0 {
		result["excluded_elements"] = me.ExcludedElements
	}
	return result, nil
}

func (me *VisuallyCompleteOptions) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("image_size_threshold", &me.ImageSizeThreshold); err != nil {
		return err
	}
	if err := decoder.Decode("inactivity_timeout", &me.InactivityTimeout); err != nil {
		return err
	}
	if err := decoder.Decode("mutation_timeout", &me.MutationTimeout); err != nil {
		return err
	}
	if value, ok := decoder.GetOk("excluded_urls"); ok {
		me.ExcludedURLs = []string{}
		for _, e := range value.([]interface{}) {
			me.ExcludedURLs = append(me.ExcludedURLs, e.(string))
		}
	}
	if value, ok := decoder.GetOk("excluded_elements"); ok {
		me.ExcludedElements = []string{}
		for _, e := range value.([]interface{}) {
			me.ExcludedElements = append(me.ExcludedElements, e.(string))
		}
	}
	return nil
}

type TimeoutSettings struct {
	TemporaryActionLimit        int `json:"temporaryActionLimit"`        // Track up to n cascading setTimeout calls
	TemporaryActionTotalTimeout int `json:"temporaryActionTotalTimeout"` // Limit cascading timeouts cumulatively to n ms
}

func (me *TimeoutSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"action_limit": {
			Type:        hcl.TypeInt,
			Description: "Track up to n cascading setTimeout calls",
			Required:    true,
		},
		"total_timeout": {
			Type:        hcl.TypeInt,
			Description: "Limit cascading timeouts cumulatively to n ms",
			Required:    true,
		},
	}
}

func (me *TimeoutSettings) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["action_limit"] = int(me.TemporaryActionLimit)
	result["total_timeout"] = int(me.TemporaryActionTotalTimeout)
	return result, nil
}

func (me *TimeoutSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("action_limit", &me.TemporaryActionLimit); err != nil {
		return err
	}
	if err := decoder.Decode("total_timeout", &me.TemporaryActionTotalTimeout); err != nil {
		return err
	}
	return nil
}

type MonitorFrames struct {
	Enabled bool `json:"enabled"`
}
