package scope

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// ManagementZone A scope filter for a management zone identifier.
type ManagementZone struct {
	BaseAlertingScope
	ID *string `json:"mzId,omitempty"` // The management zone id to match on.
}

func (me *ManagementZone) GetType() FilterType {
	return FilterTypes.ManagementZone
}

func (me *ManagementZone) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeString,
			Optional:    true,
			Description: "The management zone id to match on",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *ManagementZone) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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

func (me *ManagementZone) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "filterType")
		delete(me.Unknowns, "mzId")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("id"); ok {
		me.ID = opt.NewString(value.(string))
	}
	return nil
}

func (me *ManagementZone) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.GetType(),
		"mzId":       me.ID,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *ManagementZone) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
		"mzId":       &me.ID,
	}); err != nil {
		return err
	}
	return nil
}
