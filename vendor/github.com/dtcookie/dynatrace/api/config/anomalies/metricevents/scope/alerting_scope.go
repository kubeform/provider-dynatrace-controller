package scope

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// AlertingScope A single filter for the alerting scope.
// This is the base version of the filter, depending on the type,
// the actual JSON may contain additional fields.
type AlertingScope interface {
	GetType() FilterType
}

// BaseAlertingScope A single filter for the alerting scope.
// This is the base version of the filter, depending on the type,
// the actual JSON may contain additional fields.
type BaseAlertingScope struct {
	FilterType FilterType                 `json:"filterType"` // Defines the actual set of fields depending on the value. See one of the following objects:  * `ENTITY_ID` -> EntityIdAlertingScope  * `MANAGEMENT_ZONE` -> ManagementZoneAlertingScope  * `TAG` -> TagFilterAlertingScope  * `NAME` -> NameAlertingScope  * `CUSTOM_DEVICE_GROUP_NAME` -> CustomDeviceGroupNameAlertingScope  * `HOST_GROUP_NAME` -> HostGroupNameAlertingScope  * `HOST_NAME` -> HostNameAlertingScope  * `PROCESS_GROUP_ID` -> ProcessGroupIdAlertingScope  * `PROCESS_GROUP_NAME` -> ProcessGroupNameAlertingScope
	Unknowns   map[string]json.RawMessage `json:"-"`
}

func (me *BaseAlertingScope) GetType() FilterType {
	return me.FilterType
}

func (me *BaseAlertingScope) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "Defines the actual set of fields depending on the value",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *BaseAlertingScope) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["type"] = string(me.FilterType)
	return result, nil
}

func (me *BaseAlertingScope) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "filterType")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("type"); ok {
		me.FilterType = FilterType(value.(string))
	}
	return nil
}

func (me *BaseAlertingScope) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"filterType": me.FilterType,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *BaseAlertingScope) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"filterType": &me.FilterType,
	}); err != nil {
		return err
	}
	return nil
}

// FilterType Defines the actual set of fields depending on the value. See one of the following objects:
// * `ENTITY_ID` -> EntityIdAlertingScope
// * `MANAGEMENT_ZONE` -> ManagementZoneAlertingScope
// * `TAG` -> TagFilterAlertingScope
// * `NAME` -> NameAlertingScope
// * `CUSTOM_DEVICE_GROUP_NAME` -> CustomDeviceGroupNameAlertingScope
// * `HOST_GROUP_NAME` -> HostGroupNameAlertingScope
// * `HOST_NAME` -> HostNameAlertingScope
// * `PROCESS_GROUP_ID` -> ProcessGroupIdAlertingScope
// * `PROCESS_GROUP_NAME` -> ProcessGroupNameAlertingScope
type FilterType string

// FilterTypes offers the known enum values
var FilterTypes = struct {
	CustomDeviceGroupName FilterType
	EntityID              FilterType
	HostGroupName         FilterType
	HostName              FilterType
	ManagementZone        FilterType
	Name                  FilterType
	ProcessGroupID        FilterType
	ProcessGroupName      FilterType
	Tag                   FilterType
}{
	"CUSTOM_DEVICE_GROUP_NAME",
	"ENTITY_ID",
	"HOST_GROUP_NAME",
	"HOST_NAME",
	"MANAGEMENT_ZONE",
	"NAME",
	"PROCESS_GROUP_ID",
	"PROCESS_GROUP_NAME",
	"TAG",
}
