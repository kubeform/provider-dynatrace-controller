package web

import "github.com/dtcookie/hcl"

type UserActionAndSessionProperties []*UserActionAndSessionProperty

func (me *UserActionAndSessionProperties) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"property": {
			Type:        hcl.TypeList,
			Description: "User action and session properties settings",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(UserActionAndSessionProperty).Schema()},
		},
	}
}

func (me UserActionAndSessionProperties) MarshalHCL() (map[string]interface{}, error) {
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
		result["property"] = entries
	}
	return result, nil
}

func (me *UserActionAndSessionProperties) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("property", me); err != nil {
		return err
	}
	return nil
}

type UserActionAndSessionProperty struct {
	DisplayName                *string        `json:"displayName,omitempty"`                // The display name of the property
	Type                       PropertyType   `json:"type"`                                 // The data type of the property. Possible values are `DATE`, `DOUBLE`, `LONG`, `LONG_STRING` and `STRING`.
	Origin                     PropertyOrigin `json:"origin"`                               // The origin of the property. Possible values are `JAVASCRIPT_API`, `META_DATA` and `SERVER_SIDE_REQUEST_ATTRIBUTE`.
	Aggregation                *Aggregation   `json:"aggregation,omitempty"`                // The aggregation type of the property. \n\n  It defines how multiple values of the property are aggregated. Possible values are `AVERAGE`, `FIRST`, `LAST`, `MAXIMUM`, `MINIMUM` and `SUM`.
	StoreAsUserActionProperty  bool           `json:"storeAsUserActionProperty"`            // If `true`, the property is stored as a user action property
	StoreAsSessionProperty     bool           `json:"storeAsSessionProperty"`               // If `true`, the property is stored as a session property
	CleanupRule                *string        `json:"cleanupRule,omitempty"`                // The cleanup rule of the property. \n\nDefines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there
	ServerSideRequestAttribute *string        `json:"serverSideRequestAttribute,omitempty"` // The ID of the request attribute. \n\nOnly applicable when the **origin** is set to `SERVER_SIDE_REQUEST_ATTRIBUTE`
	UniqueID                   int32          `json:"uniqueId"`                             // Unique id among all userTags and properties of this application
	Key                        string         `json:"key"`                                  // Key of the property
	MetaDataID                 *int32         `json:"metadataId,omitempty"`                 // If the origin is `META_DATA`, metaData id of the property
	IgnoreCase                 bool           `json:"ignoreCase,omitempty"`                 // If `true`, the value of this property will always be stored in lower case. Defaults to `false`.
	LongStringLength           *int32         `json:"longStringLength,omitempty"`           // If the `type` is `LONG_STRING`, the max length for this property. Must be a multiple of `100`. Defaults to `200`. Maximum is `1000`.
}

func (me *UserActionAndSessionProperty) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"display_name": {
			Type:        hcl.TypeString,
			Description: "The display name of the property",
			Optional:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The data type of the property. Possible values are `DATE`, `DOUBLE`, `LONG`, `LONG_STRING` and `STRING`.",
			Required:    true,
		},
		"origin": {
			Type:        hcl.TypeString,
			Description: "The origin of the property. Possible values are `JAVASCRIPT_API`, `META_DATA` and `SERVER_SIDE_REQUEST_ATTRIBUTE`.",
			Required:    true,
		},
		"aggregation": {
			Type:        hcl.TypeString,
			Description: "The aggregation type of the property. \n\n  It defines how multiple values of the property are aggregated. Possible values are `AVERAGE`, `FIRST`, `LAST`, `MAXIMUM`, `MINIMUM` and `SUM`.",
			Optional:    true,
		},
		"store_as_user_action_property": {
			Type:        hcl.TypeBool,
			Description: "If `true`, the property is stored as a user action property",
			Optional:    true,
		},
		"store_as_session_property": {
			Type:        hcl.TypeBool,
			Description: "If `true`, the property is stored as a session property",
			Optional:    true,
		},
		"cleanup_rule": {
			Type:        hcl.TypeString,
			Description: "The cleanup rule of the property. \n\nDefines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there",
			Optional:    true,
		},
		"server_side_request_attribute": {
			Type:        hcl.TypeString,
			Description: "The ID of the request attribute. \n\nOnly applicable when the **origin** is set to `SERVER_SIDE_REQUEST_ATTRIBUTE`",
			Optional:    true,
		},
		"id": {
			Type:        hcl.TypeInt,
			Description: "Unique id among all userTags and properties of this application",
			Required:    true,
		},
		"key": {
			Type:        hcl.TypeString,
			Description: "Key of the property",
			Required:    true,
		},
		"metadata_id": {
			Type:        hcl.TypeInt,
			Description: "If the origin is `META_DATA`, metaData id of the property",
			Optional:    true,
		},
		"ignore_case": {
			Type:        hcl.TypeBool,
			Description: "If `true`, the value of this property will always be stored in lower case. Defaults to `false`.",
			Optional:    true,
		},
		"long_string_length": {
			Type:        hcl.TypeInt,
			Description: "If the `type` is `LONG_STRING`, the max length for this property. Must be a multiple of `100`. Defaults to `200`. Maximum is `1000`.",
			Optional:    true,
		},
	}
}

func (me *UserActionAndSessionProperty) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"display_name":                  me.DisplayName,
		"type":                          me.Type,
		"origin":                        me.Origin,
		"aggregation":                   me.Aggregation,
		"store_as_user_action_property": me.StoreAsUserActionProperty,
		"store_as_session_property":     me.StoreAsSessionProperty,
		"cleanup_rule":                  me.CleanupRule,
		"server_side_request_attribute": me.ServerSideRequestAttribute,
		"id":                            me.UniqueID,
		"key":                           me.Key,
		"metadata_id":                   me.MetaDataID,
		"ignore_case":                   me.IgnoreCase,
		"long_string_length":            me.LongStringLength,
	})
}

func (me *UserActionAndSessionProperty) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"display_name":                  &me.DisplayName,
		"type":                          &me.Type,
		"origin":                        &me.Origin,
		"aggregation":                   &me.Aggregation,
		"store_as_user_action_property": &me.StoreAsUserActionProperty,
		"store_as_session_property":     &me.StoreAsSessionProperty,
		"cleanup_rule":                  &me.CleanupRule,
		"server_side_request_attribute": &me.ServerSideRequestAttribute,
		"id":                            &me.UniqueID,
		"key":                           &me.Key,
		"metadata_id":                   &me.MetaDataID,
		"ignore_case":                   &me.IgnoreCase,
		"long_string_length":            &me.LongStringLength,
	})
}
