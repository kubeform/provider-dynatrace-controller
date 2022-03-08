package web

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

type KeyUserActionList struct {
	KeyUserActions KeyUserActions `json:"keyUserActionList,omitempty"` // The list of key user actions in the web application
}

type KeyUserActions []*KeyUserAction

func (me *KeyUserActions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"action": {
			Type:        hcl.TypeList,
			Description: "Configuration of the key user action",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(KeyUserAction).Schema()},
		},
	}
}

func (me KeyUserActions) MarshalHCL() (map[string]interface{}, error) {
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
		result["action"] = entries
	}
	return result, nil
}

func (me *KeyUserActions) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("action", me); err != nil {
		return err
	}
	return nil
}

// KeyUserAction represents configuration of the key user action
type KeyUserAction struct {
	ID     *string           `json:"meIdentifier"`     // The Dynatrace entity ID of the action
	Name   string            `json:"name"`             // The name of the action
	Type   KeyUserActionType `json:"actionType"`       // The type of the action. Possible values are `Custom`, `Load` and `Xhr`.
	Domain *string           `json:"domain,omitempty"` // The domain where the action is performed
}

func (me *KeyUserAction) String() string {
	tmp := struct {
		Name   string            `json:"name"`
		Type   KeyUserActionType `json:"actionType"`
		Domain *string           `json:"domain,omitempty"`
	}{
		me.Name,
		me.Type,
		me.Domain,
	}
	data, _ := json.Marshal(tmp)
	return string(data)
}

func (me *KeyUserAction) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the action",
			Required:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the action. Possible values are `Custom`, `Load` and `Xhr`.",
			Required:    true,
		},
		"domain": {
			Type:        hcl.TypeString,
			Description: "The domain where the action is performed.",
			Optional:    true,
		},
	}
}

func (me *KeyUserAction) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"name":   me.Name,
		"type":   me.Type,
		"domain": me.Domain,
	})
}

func (me *KeyUserAction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"name":   &me.Name,
		"type":   &me.Type,
		"domain": &me.Domain,
	})
}
