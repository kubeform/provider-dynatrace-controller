package web

import "github.com/dtcookie/hcl"

// ApplicationDataPrivacy represents data privacy settings of the application
type ApplicationDataPrivacy struct {
	WebApplicationID                *string                           `json:"identifier,omitempty"`               // Dynatrace entity ID of the web application
	DataCaptureOptInEnabled         bool                              `json:"dataCaptureOptInEnabled"`            // Set to `true` to disable data capture and cookies until JavaScriptAPI `dtrum.enable()` is called
	PersistentCookieForUserTracking bool                              `json:"persistentCookieForUserTracking"`    // Set to `true` to set persistent cookie in order to recognize returning devices
	DoNotTrackBehaviour             DoNotTrackBehaviour               `json:"doNotTrackBehaviour"`                // How to handle the \"Do Not Track\" header: \n\n* `IGNORE_DO_NOT_TRACK`: ignore the header and capture the data. \n* `CAPTURE_ANONYMIZED`: capture the data but do not tie it to the user. \n* `DO_NOT_CAPTURE`: respect the header and do not capture.
	SessionReplayDataPrivacy        *SessionReplayDataPrivacySettings `json:"sessionReplayDataPrivacy,omitempty"` // Data privacy settings for Session Replay
}

func (me *ApplicationDataPrivacy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"web_application_id": {
			Type:        hcl.TypeString,
			Description: "Dynatrace entity ID of the web application",
			Required:    true,
		},
		"data_capture_opt_in": {
			Type:        hcl.TypeBool,
			Description: "Set to `true` to disable data capture and cookies until JavaScriptAPI `dtrum.enable()` is called",
			Optional:    true,
		},
		"persistent_cookie_for_user_tracking": {
			Type:        hcl.TypeBool,
			Description: "Set to `true` to set persistent cookie in order to recognize returning devices",
			Optional:    true,
		},
		"do_not_track_behaviour": {
			Type:        hcl.TypeString,
			Description: "How to handle the \"Do Not Track\" header: \n\n* `IGNORE_DO_NOT_TRACK`: ignore the header and capture the data. \n* `CAPTURE_ANONYMIZED`: capture the data but do not tie it to the user. \n* `DO_NOT_CAPTURE`: respect the header and do not capture.",
			Required:    true,
		},
		"session_replay_data_privacy": {
			Type:        hcl.TypeList,
			Description: "Data privacy settings for Session Replay",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(SessionReplayDataPrivacySettings).Schema()},
		},
	}
}

func (me *ApplicationDataPrivacy) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"web_application_id":                  me.WebApplicationID,
		"data_capture_opt_in":                 me.DataCaptureOptInEnabled,
		"persistent_cookie_for_user_tracking": me.PersistentCookieForUserTracking,
		"do_not_track_behaviour":              me.DoNotTrackBehaviour,
		"session_replay_data_privacy":         me.SessionReplayDataPrivacy,
	})
}

func (me *ApplicationDataPrivacy) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"web_application_id":                  &me.WebApplicationID,
		"data_capture_opt_in":                 &me.DataCaptureOptInEnabled,
		"persistent_cookie_for_user_tracking": &me.PersistentCookieForUserTracking,
		"do_not_track_behaviour":              &me.DoNotTrackBehaviour,
		"session_replay_data_privacy":         &me.SessionReplayDataPrivacy,
	})
}
