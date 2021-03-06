package sharing

import "github.com/dtcookie/hcl"

// AnonymousAccess represents configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard
type AnonymousAccess struct {
	ManagementZoneIDs []string `json:"managementZoneIds"` // A list of management zones that can display data on the publicly shared dashboard. \n\nSpecify management zone IDs here. For each management zone you specify Dynatrace generates an access link. You can access them in the **urls** list. \n\nTo share the dashboard with its default management zone, use the `default` value
	// READ ONLY
	URLs map[string]string `json:"urls,omitempty"` // A list of URLs for anonymous access to the dashboard. \n\nEach link grants access to data from the specific management zone, listed in the in the **managementZoneIds** list. \n\nThese links are automatically generated by Dynatrace, you can't change them
}

func (me *AnonymousAccess) IsEmpty() bool {
	return len(me.ManagementZoneIDs) == 0 && len(me.URLs) == 0
}

func (me *AnonymousAccess) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"management_zones": {
			Type:        hcl.TypeSet,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Required:    true,
			MinItems:    1,
			Description: "A list of management zones that can display data on the publicly shared dashboard. \n\nSpecify management zone IDs here. For each management zone you specify Dynatrace generates an access link. To share the dashboard with its default management zone, use the `default` value",
		},
	}
}

func (me *AnonymousAccess) MarshalHCL() (map[string]interface{}, error) {
	if len(me.ManagementZoneIDs) > 0 {
		return map[string]interface{}{
			"management_zones": me.ManagementZoneIDs,
		}, nil
	}
	return map[string]interface{}{}, nil
}

func (me *AnonymousAccess) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.Decode("management_zones", &me.ManagementZoneIDs)
}
