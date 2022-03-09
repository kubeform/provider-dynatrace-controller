package web

import "github.com/dtcookie/hcl"

// DestinationDetails Configuration of a destination-based conversion goal
type DestinationDetails struct {
	URLOrPath     string     `json:"urlOrPath"`               // The path to be reached to hit the conversion goal
	MatchType     *MatchType `json:"matchType,omitempty"`     // The operator of the match. Possible values are `Begins`, `Contains` and `Ends`.
	CaseSensitive bool       `json:"caseSensitive,omitempty"` // The match is case-sensitive (`true`) or (`false`)
}

func (me *DestinationDetails) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"url_or_path": {
			Type:        hcl.TypeString,
			Description: "The path to be reached to hit the conversion goal",
			Required:    true,
		},
		"match_type": {
			Type:        hcl.TypeString,
			Description: "The operator of the match. Possible values are `Begins`, `Contains` and `Ends`.",
			Optional:    true,
		},
		"case_sensitive": {
			Type:        hcl.TypeBool,
			Description: "The match is case-sensitive (`true`) or (`false`)",
			Optional:    true,
		},
	}
}

func (me *DestinationDetails) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"url_or_path":    me.URLOrPath,
		"match_type":     me.MatchType,
		"case_sensitive": me.CaseSensitive,
	})
}

func (me *DestinationDetails) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"url_or_path":    &me.URLOrPath,
		"match_type":     &me.MatchType,
		"case_sensitive": &me.CaseSensitive,
	})
}
