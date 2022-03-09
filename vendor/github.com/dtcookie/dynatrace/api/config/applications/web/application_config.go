package web

import (
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// ApplicationConfig Configuration of a web application
type ApplicationConfig struct {
	ID                               *string                        `json:"identifier,omitempty"`             // Dynatrace entity ID of the web application
	Name                             string                         `json:"name"`                             // The name of the web application, displayed in the UI
	Type                             ApplicationType                `json:"type"`                             // The type of the web application
	RealUserMonitoringEnabled        bool                           `json:"realUserMonitoringEnabled"`        // Real user monitoring enabled/disabled
	CostControlUserSessionPercentage int                            `json:"costControlUserSessionPercentage"` // Analize *X*% of user sessions
	LoadActionKeyPerformanceMetric   LoadActionKeyPerformanceMetric `json:"loadActionKeyPerformanceMetric"`   // The key performance metric of load actions
	SessionReplayConfig              *SessionReplaySettings         `json:"sessionReplayConfig,omitempty"`    // Session replay settings
	XHRActionKeyPerformanceMetric    XHRActionKeyPerformanceMetric  `json:"xhrActionKeyPerformanceMetric"`    // The key performance metric of XHR actions
	LoadActionApdexSettings          *ApdexSettings                 `json:"loadActionApdexSettings"`          // Defines the Load Action Apdex settings of an application
	XHRActionApdexSettings           *ApdexSettings                 `json:"xhrActionApdexSettings"`           // Defines the XHR Action Apdex settings of an application
	CustomActionApdexSettings        *ApdexSettings                 `json:"customActionApdexSettings"`        // Defines the Custom Action Apdex settings of an application
	WaterfallSettings                *WaterfallSettings             `json:"waterfallSettings"`                // These settings influence the monitoring data you receive for 3rd party, CDN, and 1st party resources
	MonitoringSettings               *MonitoringSettings            `json:"monitoringSettings"`               // Real user monitoring settings
	UserTags                         UserTags                       `json:"userTags"`                         // User tags settings
	UserActionAndSessionProperties   UserActionAndSessionProperties `json:"userActionAndSessionProperties"`   // User action and session properties settings. Empty List means no change
	UserActionNamingSettings         *UserActionNamingSettings      `json:"userActionNamingSettings"`         // The settings of user action naming
	MetaDataCaptureSettings          MetaDataCaptureSettings        `json:"metaDataCaptureSettings"`          // Java script agent meta data capture settings
	ConversionGoals                  ConversionGoals                `json:"conversionGoals"`                  // A list of conversion goals of the application
	URLInjectionPattern              *string                        `json:"urlInjectionPattern,omitempty"`    // URL injection pattern for manual web application
	KeyUserActions                   KeyUserActions                 `json:"-"`
}

func (me *ApplicationConfig) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the web application, displayed in the UI",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the web application. Possible values are `AUTO_INJECTED`, `BROWSER_EXTENSION_INJECTED` and `MANUALLY_INJECTED`",
			Required:    true,
		},
		"real_user_monitoring_enabled": {
			Type:        hcl.TypeBool,
			Description: "Real user monitoring enabled/disabled",
			Optional:    true,
		},
		"cost_control_user_session_percentage": {
			Type:        hcl.TypeInt,
			Description: "Analize *X*% of user sessions",
			Required:    true,
		},
		"load_action_key_performance_metric": {
			Type:        hcl.TypeString,
			Description: "The key performance metric of load actions. Possible values are `ACTION_DURATION`, `CUMULATIVE_LAYOUT_SHIFT`, `DOM_INTERACTIVE`, `FIRST_INPUT_DELAY`, `LARGEST_CONTENTFUL_PAINT`, `LOAD_EVENT_END`, `LOAD_EVENT_START`, `RESPONSE_END`, `RESPONSE_START`, `SPEED_INDEX` and `VISUALLY_COMPLETE`",
			Required:    true,
		},
		"session_replay_config": {
			Type:        hcl.TypeList,
			Description: "Settings regarding Session Replay",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(SessionReplaySettings).Schema()},
		},
		"xhr_action_key_performance_metric": {
			Type:        hcl.TypeString,
			Description: "The key performance metric of XHR actions. Possible values are `ACTION_DURATION`, `RESPONSE_END`, `RESPONSE_START` and `VISUALLY_COMPLETE`.",
			Required:    true,
		},
		"load_action_apdex_settings": {
			Type:        hcl.TypeList,
			Description: "Defines the Load Action Apdex settings of an application",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ApdexSettings).Schema()},
		},
		"xhr_action_apdex_settings": {
			Type:        hcl.TypeList,
			Description: "Defines the XHR Action Apdex settings of an application",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ApdexSettings).Schema()},
		},
		"custom_action_apdex_settings": {
			Type:        hcl.TypeList,
			Description: "Defines the Custom Action Apdex settings of an application",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ApdexSettings).Schema()},
		},
		"waterfall_settings": {
			Type:        hcl.TypeList,
			Description: "These settings influence the monitoring data you receive for 3rd party, CDN, and 1st party resources",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(WaterfallSettings).Schema()},
		},
		"monitoring_settings": {
			Type:        hcl.TypeList,
			Description: "Real user monitoring settings",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(MonitoringSettings).Schema()},
		},
		"user_tags": {
			Type:        hcl.TypeList,
			Description: "User tags settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserTags).Schema()},
		},
		"user_action_and_session_properties": {
			Type:        hcl.TypeList,
			Description: "User action and session properties settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionAndSessionProperties).Schema()},
		},
		"user_action_naming_settings": {
			Type:        hcl.TypeList,
			Description: "The settings of user action naming",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionNamingSettings).Schema()},
		},
		"meta_data_capture_settings": {
			Type:        hcl.TypeList,
			Description: "Java script agent meta data capture settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(MetaDataCaptureSettings).Schema()},
		},
		"conversion_goals": {
			Type:        hcl.TypeList,
			Description: "A list of conversion goals of the application",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ConversionGoals).Schema()},
		},
		"url_injection_pattern": {
			Type:        hcl.TypeString,
			Description: "URL injection pattern for manual web application",
			Optional:    true,
		},
		"key_user_actions": {
			Type:        hcl.TypeList,
			Description: "User Action names to be flagged as Key User Actions",
			Elem:        &hcl.Resource{Schema: new(KeyUserActions).Schema()},
			Optional:    true,
		},
	}
}

func (me *ApplicationConfig) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":                                 me.Name,
		"type":                                 me.Type,
		"real_user_monitoring_enabled":         me.RealUserMonitoringEnabled,
		"cost_control_user_session_percentage": me.CostControlUserSessionPercentage,
		"load_action_key_performance_metric":   me.LoadActionKeyPerformanceMetric,
		"session_replay_config":                me.SessionReplayConfig,
		"xhr_action_key_performance_metric":    me.XHRActionKeyPerformanceMetric,
		"load_action_apdex_settings":           me.LoadActionApdexSettings,
		"xhr_action_apdex_settings":            me.XHRActionApdexSettings,
		"custom_action_apdex_settings":         me.CustomActionApdexSettings,
		"waterfall_settings":                   me.WaterfallSettings,
		"monitoring_settings":                  me.MonitoringSettings,
		"user_tags":                            me.UserTags,
		"user_action_and_session_properties":   me.UserActionAndSessionProperties,
		"user_action_naming_settings":          me.UserActionNamingSettings,
		"meta_data_capture_settings":           me.MetaDataCaptureSettings,
		"conversion_goals":                     me.ConversionGoals,
		"url_injection_pattern":                me.URLInjectionPattern,
		"key_user_actions":                     me.KeyUserActions,
	})
}

func (me *ApplicationConfig) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeAll(map[string]interface{}{
		"name":                                 &me.Name,
		"type":                                 &me.Type,
		"real_user_monitoring_enabled":         &me.RealUserMonitoringEnabled,
		"cost_control_user_session_percentage": &me.CostControlUserSessionPercentage,
		"load_action_key_performance_metric":   &me.LoadActionKeyPerformanceMetric,
		"session_replay_config":                &me.SessionReplayConfig,
		"xhr_action_key_performance_metric":    &me.XHRActionKeyPerformanceMetric,
		"load_action_apdex_settings":           &me.LoadActionApdexSettings,
		"xhr_action_apdex_settings":            &me.XHRActionApdexSettings,
		"custom_action_apdex_settings":         &me.CustomActionApdexSettings,
		"waterfall_settings":                   &me.WaterfallSettings,
		"monitoring_settings":                  &me.MonitoringSettings,
		"user_tags":                            &me.UserTags,
		"user_action_and_session_properties":   &me.UserActionAndSessionProperties,
		"user_action_naming_settings":          &me.UserActionNamingSettings,
		"meta_data_capture_settings":           &me.MetaDataCaptureSettings,
		"conversion_goals":                     &me.ConversionGoals,
		"url_injection_pattern":                &me.URLInjectionPattern,
		"key_user_actions":                     &me.KeyUserActions,
	}); err != nil {
		return err
	}
	if me.UserTags == nil {
		me.UserTags = UserTags{}
	}
	if me.UserActionAndSessionProperties == nil {
		me.UserActionAndSessionProperties = UserActionAndSessionProperties{}
	}
	if me.ConversionGoals == nil {
		me.ConversionGoals = ConversionGoals{}
	}
	if me.MetaDataCaptureSettings == nil {
		me.MetaDataCaptureSettings = MetaDataCaptureSettings{}
	}
	if me.MonitoringSettings != nil {
		if me.MonitoringSettings.LibraryFileLocation == nil {
			if me.Type == ApplicationTypes.ManuallyInjected {
				me.MonitoringSettings.LibraryFileLocation = nil
			}
			if me.Type == ApplicationTypes.AutoInjected {
				me.MonitoringSettings.LibraryFileLocation = opt.NewString("")
			}
		}
		if me.MonitoringSettings.LibraryFileLocation != nil {
			if me.Type == ApplicationTypes.ManuallyInjected {
				me.MonitoringSettings.LibraryFileLocation = nil
			}
		}
	}
	return nil
}
