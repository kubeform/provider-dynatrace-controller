package web

import "github.com/dtcookie/hcl"

// VisitNumActionDetails Configuration of a number of user actions-based conversion goal
type VisitNumActionDetails struct {
	NumUserActions *int32 `json:"numUserActions,omitempty"` // The number of user actions to hit the conversion goal
}

func (me *VisitNumActionDetails) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"num_user_actions": {
			Type:        hcl.TypeInt,
			Description: "The number of user actions to hit the conversion goal",
			Optional:    true,
		},
	}
}

func (me *VisitNumActionDetails) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"num_user_actions": me.NumUserActions,
	})
}

func (me *VisitNumActionDetails) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"num_user_actions": &me.NumUserActions,
	})
}
