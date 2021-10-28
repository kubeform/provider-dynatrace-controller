package mobile

import "github.com/dtcookie/hcl"

type APIValue UserActionAndSessionProperty

func (me *APIValue) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"key": {
			Type:        hcl.TypeString,
			Description: "The unique key of the mobile session or user action property",
			Required:    true,
		},
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the reported value",
			Optional:    true,
		},
		"type": {
			Type:        hcl.TypeString,
			Description: "The data type of the property. Possible values are `DOUBLE`, `LONG` and `STRING`",
			Required:    true,
		},
		"display_name": {
			Type:        hcl.TypeString,
			Description: "The display name of the property",
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
			Description: "The cleanup rule of the property. Defines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there",
			Optional:    true,
		},
		"aggregation": {
			Type:        hcl.TypeString,
			Description: "The aggregation type of the property. It defines how multiple values of the property are aggregated. Possible values are `SUM`, `MIN`, `MAX`, `FIRST` and `LAST`",
			Optional:    true,
		},
	}
}

func (me *APIValue) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("key", &me.Key); err != nil {
		return err
	}
	if err := decoder.Decode("name", &me.Name); err != nil {
		return err
	}
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("display_name", &me.DisplayName); err != nil {
		return err
	}
	if err := decoder.Decode("store_as_user_action_property", &me.StoreAsUserActionProperty); err != nil {
		return err
	}
	if err := decoder.Decode("store_as_session_property", &me.StoreAsSessionProperty); err != nil {
		return err
	}
	if err := decoder.Decode("cleanup_rule", &me.CleanupRule); err != nil {
		return err
	}
	if err := decoder.Decode("aggregation", &me.Aggregation); err != nil {
		return err
	}
	me.Origin = Origins.API
	me.ServerSideRequestAttribute = nil
	return nil
}

func (me *APIValue) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if err := properties.Encode("key", me.Key); err != nil {
		return nil, err
	}
	if err := properties.Encode("name", *me.Name); err != nil {
		return nil, err
	}
	if err := properties.Encode("type", me.Type); err != nil {
		return nil, err
	}
	if me.DisplayName != nil {
		if err := properties.Encode("display_name", me.DisplayName); err != nil {
			return nil, err
		}
	}
	if me.StoreAsSessionProperty {
		if err := properties.Encode("store_as_session_property", me.StoreAsSessionProperty); err != nil {
			return nil, err
		}
	}
	if me.StoreAsUserActionProperty {
		if err := properties.Encode("store_as_user_action_property", me.StoreAsUserActionProperty); err != nil {
			return nil, err
		}
	}
	if me.CleanupRule != nil && len(*me.CleanupRule) > 0 {
		if err := properties.Encode("cleanup_rule", *me.CleanupRule); err != nil {
			return nil, err
		}
	}
	if me.Aggregation != nil {
		if err := properties.Encode("aggregation", string(*me.Aggregation)); err != nil {
			return nil, err
		}
	}
	return properties, nil
}
