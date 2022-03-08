package web

import "github.com/dtcookie/hcl"

// SessionReplayMaskingSetting represents configuration of the Session Replay masking
type SessionReplayMaskingSetting struct {
	Preset MaskingPreset `json:"maskingPreset"`          // The type of the masking: \n\n* `MASK_ALL`: Mask all texts, user input, and images. \n* `MASK_USER_INPUT`: Mask all data that is provided through user input \n* `ALLOW_LIST`: Only elements, specified in **maskingRules** are shown, everything else is masked. \n* `BLOCK_LIST`: Elements, specified in **maskingRules** are masked, everything else is shown.
	Rules  MaskingRules  `json:"maskingRules,omitempty"` // A list of masking rules
}

func (me *SessionReplayMaskingSetting) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"preset": {
			Type:        hcl.TypeString,
			Description: "The type of the masking: \n\n* `MASK_ALL`: Mask all texts, user input, and images. \n* `MASK_USER_INPUT`: Mask all data that is provided through user input \n* `ALLOW_LIST`: Only elements, specified in **maskingRules** are shown, everything else is masked. \n* `BLOCK_LIST`: Elements, specified in **maskingRules** are masked, everything else is shown",
			Required:    true,
		},
		"rules": {
			Type:        hcl.TypeList,
			Description: "A list of masking rules",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(MaskingRules).Schema()},
		},
	}
}

func (me *SessionReplayMaskingSetting) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"preset": me.Preset,
		"rules":  me.Rules,
	})
}

func (me *SessionReplayMaskingSetting) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"preset": &me.Preset,
		"rules":  &me.Rules,
	})
}

type MaskingPreset string

var MaskingPresets = struct {
	AllowList     MaskingPreset
	BlockList     MaskingPreset
	MaskAll       MaskingPreset
	MaskUserInput MaskingPreset
}{
	"ALLOW_LIST",
	"BLOCK_LIST",
	"MASK_ALL",
	"MASK_USER_INPUT",
}
