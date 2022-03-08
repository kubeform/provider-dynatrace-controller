package web

import "github.com/dtcookie/hcl"

type ConversionGoals []*ConversionGoal

func (me *ConversionGoals) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"goal": {
			Type:        hcl.TypeList,
			Description: "A conversion goal of the application",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(ConversionGoal).Schema()},
		},
	}
}

func (me ConversionGoals) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if len(me) > 0 {
		entries := []interface{}{}
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["goal"] = entries
	}
	return result, nil
}

func (me *ConversionGoals) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("goal", me); err != nil {
		return err
	}
	return nil
}

// ConversionGoal A conversion goal of the application
type ConversionGoal struct {
	Name                  string                 `json:"name"`                            // The name of the conversion goal. Valid length within 1 and 50 characters.
	ID                    *string                `json:"id,omitempty"`                    // The ID of conversion goal. \n\n Omit it while creating a new conversion goal
	Type                  *ConversionGoalType    `json:"type,omitempty"`                  // The type of the conversion goal. Possible values are `Destination`, `UserAction`, `VisitDuration` and `VisitNumActions`
	DestinationDetails    *DestinationDetails    `json:"destinationDetails,omitempty"`    // Configuration of a destination-based conversion goal
	UserActionDetails     *UserActionDetails     `json:"userActionDetails,omitempty"`     // Configuration of a user action-based conversion goal
	VisitDurationDetails  *VisitDurationDetails  `json:"visitDurationDetails,omitempty"`  // Configuration of a visit duration-based conversion goal
	VisitNumActionDetails *VisitNumActionDetails `json:"visitNumActionDetails,omitempty"` // Configuration of a number of user actions-based conversion goal
}

func (me *ConversionGoal) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the conversion goal. Valid length within 1 and 50 characters.",
			Required:    true,
		},
		"id": {
			Type:        hcl.TypeString,
			Description: "The ID of conversion goal. \n\n Omit it while creating a new conversion goal",
			Optional:    true,
			Computed:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the conversion goal. Possible values are `Destination`, `UserAction`, `VisitDuration` and `VisitNumActions`",
			Optional:    true,
		},
		"destination": {
			Type:        hcl.TypeList,
			Description: "Configuration of a destination-based conversion goal",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(DestinationDetails).Schema()},
		},
		"user_action": {
			Type:        hcl.TypeList,
			Description: "Configuration of a destination-based conversion goal",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionDetails).Schema()},
		},
		"visit_duration": {
			Type:        hcl.TypeList,
			Description: "Configuration of a destination-based conversion goal",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(VisitDurationDetails).Schema()},
		},
		"visit_num_action": {
			Type:        hcl.TypeList,
			Description: "Configuration of a destination-based conversion goal",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(VisitNumActionDetails).Schema()},
		},
	}
}

func (me *ConversionGoal) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":             me.Name,
		"id":               me.ID,
		"type":             me.Type,
		"destination":      me.DestinationDetails,
		"user_action":      me.UserActionDetails,
		"visit_duration":   me.VisitDurationDetails,
		"visit_num_action": me.VisitNumActionDetails,
	})
}

func (me *ConversionGoal) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":             &me.Name,
		"id":               &me.ID,
		"type":             &me.Type,
		"destination":      &me.DestinationDetails,
		"user_action":      &me.UserActionDetails,
		"visit_duration":   &me.VisitDurationDetails,
		"visit_num_action": &me.VisitNumActionDetails,
	})
}
