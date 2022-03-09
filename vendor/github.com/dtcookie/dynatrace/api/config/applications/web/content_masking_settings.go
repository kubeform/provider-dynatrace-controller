package web

import "github.com/dtcookie/hcl"

// ContentMaskingSettings represents content masking settings for Session Replay. \n\nFor more details, see [Configure Session Replay](https://dt-url.net/0m03slq) in Dynatrace Documentation
type ContentMaskingSettings struct {
	RecordingMaskingSettingsVersion int32                        `json:"recordingMaskingSettingsVersion"` // The version of the content masking. \n\nYou can use this API only with the version 2. \n\nIf you're using version 1, set this field to `2` in the PUT request to switch to version 2
	RecordingMaskingSettings        *SessionReplayMaskingSetting `json:"recordingMaskingSettings"`        // Configuration of the Session Replay masking during Recording
	PlaybackMaskingSettings         *SessionReplayMaskingSetting `json:"playbackMaskingSettings"`         // Configuration of the Session Replay masking during Playback
}

func (me *ContentMaskingSettings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"recording": {
			Type:        hcl.TypeList,
			Description: "Configuration of the Session Replay masking during Recording",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(SessionReplayMaskingSetting).Schema()},
		},
		"playback": {
			Type:        hcl.TypeList,
			Description: "Configuration of the Session Replay masking during Playback",
			Required:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(SessionReplayMaskingSetting).Schema()},
		},
	}
}

func (me *ContentMaskingSettings) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"recording": me.RecordingMaskingSettings,
		"playback":  me.PlaybackMaskingSettings,
	})
}

func (me *ContentMaskingSettings) UnmarshalHCL(decoder hcl.Decoder) error {
	me.RecordingMaskingSettingsVersion = 2
	return decoder.DecodeAll(map[string]interface{}{
		"recording": &me.RecordingMaskingSettings,
		"playback":  &me.PlaybackMaskingSettings,
	})
}
