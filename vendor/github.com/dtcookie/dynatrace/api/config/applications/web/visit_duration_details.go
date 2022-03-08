package web

import "github.com/dtcookie/hcl"

// VisitDurationDetails Configuration of a visit duration-based conversion goal
type VisitDurationDetails struct {
	DurationInMillis int64 `json:"durationInMillis"` // The duration of session to hit the conversion goal, in milliseconds
}

func (me *VisitDurationDetails) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"duration": {
			Type:        hcl.TypeInt,
			Description: "The duration of session to hit the conversion goal, in milliseconds",
			Required:    true,
		},
	}
}

func (me *VisitDurationDetails) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"duration": me.DurationInMillis,
	})
}

func (me *VisitDurationDetails) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"duration": &me.DurationInMillis,
	})
}
