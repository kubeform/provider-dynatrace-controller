package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// HTTPHeader The HTTP header.
type HTTPHeader struct {
	Name  string  `json:"name"`            // The name of the HTTP header.
	Value *string `json:"value,omitempty"` // The value of the HTTP header. May contain an empty value.   Required when creating a new notification.  For the **Authorization** header, GET requests return the `null` value.  If you want update a notification configuration with an **Authorization** header which you want to remain intact, set the **Authorization** header with the `null` value.
}

func (me *HTTPHeader) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the HTTP header",
			Required:    true,
		},
		"value": {
			Type:        hcl.TypeString,
			Description: "The value of the HTTP header. May contain an empty value.   Required when creating a new notification.  For the **Authorization** header, GET requests return the `null` value.  If you want update a notification configuration with an **Authorization** header which you want to remain intact, set the **Authorization** header with the `null` value",
			Optional:    true,
		},
	}
}

func (me *HTTPHeader) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	result["name"] = me.Name
	if me.Value != nil {
		result["value"] = *me.Value
	}

	return result, nil
}

func (me *HTTPHeader) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	adapter := hcl.Adapt(decoder)
	me.Value = adapter.GetString("value")
	return nil
}

func (me *HTTPHeader) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"name":  me.Name,
		"value": me.Value,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *HTTPHeader) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"name":  &me.Name,
		"value": &me.Value,
	}); err != nil {
		return err
	}
	return nil
}
