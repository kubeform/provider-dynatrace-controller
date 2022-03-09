package web

import (
	"github.com/dtcookie/hcl"
)

// SessionReplayDataPrivacySettings Data privacy settings for Session Replay
type SessionReplayDataPrivacySettings struct {
	OptIn                  bool                    `json:"optInModeEnabled,omitempty"`       // If `true`, session recording is disabled until JavaScriptAPI `dtrum.enableSessionReplay()` is called
	URLExclusionRules      []string                `json:"urlExclusionRules,omitempty"`      // A list of URLs to be excluded from recording
	ContentMaskingSettings *ContentMaskingSettings `json:"contentMaskingSettings,omitempty"` // Content masking settings for Session Replay. \n\nFor more details, see [Configure Session Replay](https://dt-url.net/0m03slq) in Dynatrace Documentation
}

func (me *SessionReplayDataPrivacySettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"opt_in": {
			Type:        hcl.TypeBool,
			Description: "If `true`, session recording is disabled until JavaScriptAPI `dtrum.enableSessionReplay()` is called",
			Optional:    true,
		},
		"url_exclusion_rules": {
			Type:        hcl.TypeList,
			Description: "A list of URLs to be excluded from recording",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"content_masking_settings": {
			Type:        hcl.TypeList,
			Description: "Content masking settings for Session Replay. \n\nFor more details, see [Configure Session Replay](https://dt-url.net/0m03slq) in Dynatrace Documentation",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(ContentMaskingSettings).Schema()},
		},
	}
}

func (me *SessionReplayDataPrivacySettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"opt_in":                   me.OptIn,
		"url_exclusion_rules":      me.URLExclusionRules,
		"content_masking_settings": me.ContentMaskingSettings,
	})
}

func (me *SessionReplayDataPrivacySettings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"opt_in":                   &me.OptIn,
		"url_exclusion_rules":      &me.URLExclusionRules,
		"content_masking_settings": &me.ContentMaskingSettings,
	})
}
