package web

import "github.com/dtcookie/hcl"

// UserActionDetails Configuration of a user action-based conversion goal
type UserActionDetails struct {
	Value         *string      `json:"value,omitempty"`         // The value to be matched to hit the conversion goal
	CaseSensitive bool         `json:"caseSensitive,omitempty"` // The match is case-sensitive (`true`) or (`false`)
	MatchType     *MatchType   `json:"matchType,omitempty"`     // The operator of the match. Possible values are `Begins`, `Contains` and `Ends`.
	MatchEntity   *MatchEntity `json:"matchEntity,omitempty"`   // The type of the entity to which the rule applies. Possible values are `ActionName`, `CssSelector`, `JavaScriptVariable`, `MetaTag`, `PagePath`, `PageTitle`, `PageUrl`, `UrlAnchor` and `XhrUrl`.
	ActionType    *ActionType  `json:"actionType,omitempty"`    // Type of the action to which the rule applies. Possible values are `Custom`, `Load` and `Xhr`.
}

func (me *UserActionDetails) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"value": {
			Type:        hcl.TypeString,
			Description: "The value to be matched to hit the conversion goal",
			Optional:    true,
		},
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "The match is case-sensitive (`true`) or (`false`)",
			Optional:    true,
		},
		"match_type": {
			Type:        hcl.TypeString,
			Description: "The operator of the match. Possible values are `Begins`, `Contains` and `Ends`.",
			Optional:    true,
		},
		"match_entity": {
			Type:        hcl.TypeString,
			Description: "The type of the entity to which the rule applies. Possible values are `ActionName`, `CssSelector`, `JavaScriptVariable`, `MetaTag`, `PagePath`, `PageTitle`, `PageUrl`, `UrlAnchor` and `XhrUrl`.",
			Optional:    true,
		},
		"action_type": {
			Type:        hcl.TypeString,
			Description: "Type of the action to which the rule applies. Possible values are `Custom`, `Load` and `Xhr`.",
			Optional:    true,
		},
	}
}

func (me *UserActionDetails) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"value":          me.Value,
		"case_sensitive": me.CaseSensitive,
		"match_type":     me.MatchType,
		"match_entity":   me.MatchEntity,
		"action_type":    me.ActionType,
	})
}

func (me *UserActionDetails) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"value":          &me.Value,
		"case_sensitive": &me.CaseSensitive,
		"match_type":     &me.MatchType,
		"match_entity":   &me.MatchEntity,
		"action_type":    &me.ActionType,
	})
}
