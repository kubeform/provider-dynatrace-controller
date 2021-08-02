package scope

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// ProcessGroupID A scope filter for a process group identifier.
type ProcessGroupID struct {
	BaseAlertingScope
	ID string `json:"processGroupId"` // The process groups id to match on.
}

func (me *ProcessGroupID) GetType() FilterType {
	return FilterTypes.ProcessGroupID
}

func (me *ProcessGroupID) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The process groups id to match on",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ProcessGroupID) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["id"] = me.ID
	return result, nil
}

func (me *ProcessGroupID) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "filterType")
		delete(me.Unknowns, "processGroupId")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("id"); ok {
		me.ID = value.(string)
	}
	return nil
}

func (me *ProcessGroupID) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType":     me.GetType(),
		"processGroupId": me.ID,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *ProcessGroupID) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType":     &me.FilterType,
		"processGroupId": &me.ID,
	}); err != nil {
		return err
	}
	return nil
}
