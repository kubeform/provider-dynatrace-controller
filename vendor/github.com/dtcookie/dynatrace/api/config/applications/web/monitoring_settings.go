package web

import (
	"github.com/dtcookie/hcl"
)

// MonitoringSettings Real user monitoring settings
type MonitoringSettings struct {
	FetchRequests                    bool                           `json:"fetchRequests"`                              // `fetch()` request capture enabled/disabled
	XmlHttpRequest                   bool                           `json:"xmlHttpRequest"`                             // `XmlHttpRequest` support enabled/disabled
	JavaScriptFrameworkSupport       *JavaScriptFrameworkSupport    `json:"javaScriptFrameworkSupport"`                 // Support of various JavaScript frameworks
	ContentCapture                   *ContentCapture                `json:"contentCapture"`                             // Settings for content capture
	ExcludeXHRRegex                  string                         `json:"excludeXhrRegex"`                            // You can exclude some actions from becoming XHR actions.\n\nPut a regular expression, matching all the required URLs, here.\n\nIf noting specified the feature is disabled
	CorrelationHeaderInclusionRegex  string                         `json:"correlationHeaderInclusionRegex"`            // To enable RUM for XHR calls to AWS Lambda, define a regular expression matching these calls, Dynatrace can then automatically add a custom header (x-dtc) to each such request to the respective endpoints in AWS.\n\nImportant: These endpoints must accept the x-dtc header, or the requests will fail
	InjectionMode                    InjectionMode                  `json:"injectionMode"`                              // Possible valures are `CODE_SNIPPET`, `CODE_SNIPPET_ASYNC`, `INLINE_CODE` and `JAVASCRIPT_TAG`
	AddCrossOriginAnonymousAttribute *bool                          `json:"addCrossOriginAnonymousAttribute,omitempty"` // Add the cross origin = anonymous attribute to capture JavaScript error messages and W3C resource timings
	ScriptTagCacheDurationInHours    *int32                         `json:"scriptTagCacheDurationInHours,omitempty"`    // Time duration for the cache settings
	LibraryFileLocation              *string                        `json:"libraryFileLocation,omitempty"`              // The location of your application’s custom JavaScript library file. \n\n If nothing specified the root directory of your web server is used. \n\n **Required** for auto-injected applications, not supported by agentless applications. Maximum 512 characters.
	MonitoringDataPath               string                         `json:"monitoringDataPath"`                         // The location to send monitoring data from the JavaScript tag.\n\n Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS. \n\n **Required** for auto-injected applications, optional for agentless applications. Maximum 512 characters.
	CustomConfigurationProperties    string                         `json:"customConfigurationProperties"`              // Additional JavaScript tag properties that are specific to your application. To do this, type key=value pairs separated using a (|) symbol. Maximum 1000 characters.
	ServerRequestPathID              string                         `json:"serverRequestPathId"`                        // Path to identify the server’s request ID. Maximum 150 characters.
	SecureCookieAttribute            bool                           `json:"secureCookieAttribute"`                      // Secure attribute usage for Dynatrace cookies enabled/disabled
	CookiePlacementDomain            string                         `json:"cookiePlacementDomain"`                      // Domain for cookie placement. Maximum 150 characters.
	CacheControlHeaderOptimizations  bool                           `json:"cacheControlHeaderOptimizations"`            // Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled
	AdvancedJavaScriptTagSettings    *AdvancedJavaScriptTagSettings `json:"advancedJavaScriptTagSettings"`              // Advanced JavaScript tag settings
	BrowserRestrictionSettings       *BrowserRestrictionSettings    `json:"browserRestrictionSettings,omitempty"`       // Settings for restricting certain browser type, version, platform and, comparator. It also restricts the mode
	IPAddressRestrictionSettings     *IPAddressRestrictionSettings  `json:"ipAddressRestrictionSettings,omitempty"`     // Settings for restricting certain ip addresses and for introducing subnet mask. It also restricts the mode
	JavaScriptInjectionRules         JavaScriptInjectionRules       `json:"javaScriptInjectionRules,omitempty"`         // Java script injection rules
	AngularPackageName               *string                        `json:"angularPackageName,omitempty"`               // The name of the angular package
}

func (me *MonitoringSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"fetch_requests": {
			Type:        hcl.TypeBool,
			Description: "`fetch()` request capture enabled/disabled",
			Optional:    true,
		},
		"xml_http_request": {
			Type:        hcl.TypeBool,
			Description: "`XmlHttpRequest` support enabled/disabled",
			Optional:    true,
		},
		"javascript_framework_support": {
			Type:        hcl.TypeList,
			Description: "Support of various JavaScript frameworks",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(JavaScriptFrameworkSupport).Schema()},
		},
		"content_capture": {
			Type:        hcl.TypeList,
			Description: "Settings for content capture",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ContentCapture).Schema()},
		},
		"exclude_xhr_regex": {
			Type:        hcl.TypeString,
			Description: "You can exclude some actions from becoming XHR actions.\n\nPut a regular expression, matching all the required URLs, here.\n\nIf noting specified the feature is disabled",
			Optional:    true,
		},
		"correlation_header_inclusion_regex": {
			Type:        hcl.TypeString,
			Description: "To enable RUM for XHR calls to AWS Lambda, define a regular expression matching these calls, Dynatrace can then automatically add a custom header (`x-dtc`) to each such request to the respective endpoints in AWS.\n\nImportant: These endpoints must accept the `x-dtc` header, or the requests will fail",
			Optional:    true,
		},
		"injection_mode": {
			Type:        hcl.TypeString,
			Description: "Possible valures are `CODE_SNIPPET`, `CODE_SNIPPET_ASYNC`, `INLINE_CODE` and `JAVASCRIPT_TAG`.",
			Required:    true,
		},
		"add_cross_origin_anonymous_attribute": {
			Type:        hcl.TypeBool,
			Description: "Add the cross origin = anonymous attribute to capture JavaScript error messages and W3C resource timings",
			Optional:    true,
		},
		"script_tag_cache_duration_in_hours": {
			Type:        hcl.TypeInt,
			Description: "Time duration for the cache settings",
			Optional:    true,
		},
		"library_file_location": {
			Type:        hcl.TypeString,
			Description: "The location of your application’s custom JavaScript library file. \n\n If nothing specified the root directory of your web server is used. \n\n **Required** for auto-injected applications, not supported by agentless applications. Maximum 512 characters.",
			Optional:    true,
		},
		"monitoring_data_path": {
			Type:        hcl.TypeString,
			Description: "The location to send monitoring data from the JavaScript tag.\n\n Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS. \n\n **Required** for auto-injected applications, optional for agentless applications. Maximum 512 characters.",
			Optional:    true,
		},
		"custom_configuration_properties": {
			Type:        hcl.TypeString,
			Description: "The location to send monitoring data from the JavaScript tag.\n\n Specify either a relative or an absolute URL. If you use an absolute URL, data will be sent using CORS. \n\n **Required** for auto-injected applications, optional for agentless applications. Maximum 512 characters.",
			Optional:    true,
		},
		"server_request_path_id": {
			Type:        hcl.TypeString,
			Description: "Path to identify the server’s request ID. Maximum 150 characters.",
			Optional:    true,
		},
		"secure_cookie_attribute": {
			Type:        hcl.TypeBool,
			Description: "Secure attribute usage for Dynatrace cookies enabled/disabled",
			Optional:    true,
		},
		"cookie_placement_domain": {
			Type:        hcl.TypeString,
			Description: "Domain for cookie placement. Maximum 150 characters.",
			Optional:    true,
		},
		"cache_control_header_optimizations": {
			Type:        hcl.TypeBool,
			Description: "Optimize the value of cache control headers for use with Dynatrace real user monitoring enabled/disabled",
			Optional:    true,
		},
		"advanced_javascript_tag_settings": {
			Type:        hcl.TypeList,
			Description: "Advanced JavaScript tag settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(AdvancedJavaScriptTagSettings).Schema()},
		},
		"browser_restriction_settings": {
			Type:        hcl.TypeList,
			Description: "Settings for restricting certain browser type, version, platform and, comparator. It also restricts the mode",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(BrowserRestrictionSettings).Schema()},
		},
		"ip_address_restriction_settings": {
			Type:        hcl.TypeList,
			Description: "Settings for restricting certain ip addresses and for introducing subnet mask. It also restricts the mode",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(IPAddressRestrictionSettings).Schema()},
		},
		"javascript_injection_rules": {
			Type:        hcl.TypeList,
			Description: "Java script injection rules",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(JavaScriptInjectionRules).Schema()},
		},
		"angular_package_name": {
			Type:        hcl.TypeString,
			Description: "The name of the angular package",
			Optional:    true,
		},
	}
}

func (me *MonitoringSettings) MarshalHCL() (map[string]interface{}, error) {
	res, err := hcl.Properties{}.EncodeAll(map[string]interface{}{
		"fetch_requests":                       me.FetchRequests,
		"xml_http_request":                     me.XmlHttpRequest,
		"javascript_framework_support":         me.JavaScriptFrameworkSupport,
		"content_capture":                      me.ContentCapture,
		"exclude_xhr_regex":                    me.ExcludeXHRRegex,
		"correlation_header_inclusion_regex":   me.CorrelationHeaderInclusionRegex,
		"injection_mode":                       me.InjectionMode,
		"add_cross_origin_anonymous_attribute": me.AddCrossOriginAnonymousAttribute,
		"script_tag_cache_duration_in_hours":   me.ScriptTagCacheDurationInHours,
		"library_file_location":                me.LibraryFileLocation,
		"monitoring_data_path":                 me.MonitoringDataPath,
		"custom_configuration_properties":      me.CustomConfigurationProperties,
		"server_request_path_id":               me.ServerRequestPathID,
		"secure_cookie_attribute":              me.SecureCookieAttribute,
		"cookie_placement_domain":              me.CookiePlacementDomain,
		"cache_control_header_optimizations":   me.CacheControlHeaderOptimizations,
		"advanced_javascript_tag_settings":     me.AdvancedJavaScriptTagSettings,
		"browser_restriction_settings":         me.BrowserRestrictionSettings,
		"ip_address_restriction_settings":      me.IPAddressRestrictionSettings,
		"javascript_injection_rules":           me.JavaScriptInjectionRules,
		"angular_package_name":                 me.AngularPackageName,
	})
	if err != nil {
		return nil, err
	}
	if me.BrowserRestrictionSettings != nil {
		if len(me.BrowserRestrictionSettings.BrowserRestrictions) == 0 {
			delete(res, "browser_restriction_settings")
		}
	}
	if me.IPAddressRestrictionSettings != nil {
		if len(me.IPAddressRestrictionSettings.Restrictions) == 0 {
			delete(res, "ip_address_restriction_settings")
		}
	}
	return res, err
}

func (me *MonitoringSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"fetch_requests":                       &me.FetchRequests,
		"xml_http_request":                     &me.XmlHttpRequest,
		"javascript_framework_support":         &me.JavaScriptFrameworkSupport,
		"content_capture":                      &me.ContentCapture,
		"exclude_xhr_regex":                    &me.ExcludeXHRRegex,
		"correlation_header_inclusion_regex":   &me.CorrelationHeaderInclusionRegex,
		"injection_mode":                       &me.InjectionMode,
		"add_cross_origin_anonymous_attribute": &me.AddCrossOriginAnonymousAttribute,
		"script_tag_cache_duration_in_hours":   &me.ScriptTagCacheDurationInHours,
		"library_file_location":                &me.LibraryFileLocation,
		"monitoring_data_path":                 &me.MonitoringDataPath,
		"custom_configuration_properties":      &me.CustomConfigurationProperties,
		"server_request_path_id":               &me.ServerRequestPathID,
		"secure_cookie_attribute":              &me.SecureCookieAttribute,
		"cookie_placement_domain":              &me.CookiePlacementDomain,
		"cache_control_header_optimizations":   &me.CacheControlHeaderOptimizations,
		"advanced_javascript_tag_settings":     &me.AdvancedJavaScriptTagSettings,
		"browser_restriction_settings":         &me.BrowserRestrictionSettings,
		"ip_address_restriction_settings":      &me.IPAddressRestrictionSettings,
		"javascript_injection_rules":           &me.JavaScriptInjectionRules,
		"angular_package_name":                 &me.AngularPackageName,
	})
}
