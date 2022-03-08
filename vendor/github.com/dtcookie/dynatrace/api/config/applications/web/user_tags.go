package web

import "github.com/dtcookie/hcl"

type UserTags []*UserTag

func (me *UserTags) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"tag": {
			Type:        hcl.TypeList,
			Description: "User tag settings",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserTag).Schema()},
		},
	}
}

func (me UserTags) MarshalHCL() (map[string]interface{}, error) {
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
		result["tag"] = entries
	}
	return result, nil
}

func (me *UserTags) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("tag", me); err != nil {
		return err
	}
	return nil
}

type UserTag struct {
	UniqueID                   int32   `json:"uniqueId"`                             // A unique ID among all userTags and properties of this application. Minimum value is 1.
	MetaDataID                 *int32  `json:"metadataId,omitempty"`                 // If it's of type metaData, metaData id of the userTag
	CleanUpRule                *string `json:"cleanupRule,omitempty"`                // Cleanup rule expression of the userTag
	ServerSideRequestAttribute *string `json:"serverSideRequestAttribute,omitempty"` // The ID of the RrequestAttribute for the userTag
	IgnoreCase                 bool    `json:"ignoreCase,omitempty"`                 // If `true`, the value of this tag will always be stored in lower case. Defaults to `false`.
}

func (me *UserTag) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"id": {
			Type:        hcl.TypeInt,
			Description: "A unique ID among all userTags and properties of this application. Minimum value is 1.",
			Required:    true,
		},
		"metadata_id": {
			Type:        hcl.TypeInt,
			Description: "If it's of type metaData, metaData id of the userTag",
			Optional:    true,
		},
		"cleanup_rule": {
			Type:        hcl.TypeString,
			Description: "Cleanup rule expression of the userTag",
			Optional:    true,
		},
		"server_side_request_attribute": {
			Type:        hcl.TypeString,
			Description: "The ID of the RrequestAttribute for the userTag",
			Optional:    true,
		},
		"ignore_case": {
			Type:        hcl.TypeBool,
			Description: "If `true`, the value of this tag will always be stored in lower case. Defaults to `false`.",
			Optional:    true,
		},
	}
}

func (me *UserTag) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"id":                            me.UniqueID,
		"metadata_id":                   me.MetaDataID,
		"cleanup_rule":                  me.CleanUpRule,
		"server_side_request_attribute": me.ServerSideRequestAttribute,
		"ignore_case":                   me.IgnoreCase,
	})
}

func (me *UserTag) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"id":                            &me.UniqueID,
		"metadata_id":                   &me.MetaDataID,
		"cleanup_rule":                  &me.CleanUpRule,
		"server_side_request_attribute": &me.ServerSideRequestAttribute,
		"ignore_case":                   &me.IgnoreCase,
	})
}
