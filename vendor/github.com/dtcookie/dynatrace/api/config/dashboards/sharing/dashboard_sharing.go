package sharing

import (
	"github.com/dtcookie/hcl"
)

// DashboardSharing represents sharing configuration of the dashboard
type DashboardSharing struct {
	DashboardID  string           `json:"id"`           // The Dynatrace entity ID of the dashboard
	Permissions  SharePermissions `json:"permissions"`  // Access permissions of the dashboard
	PublicAccess *AnonymousAccess `json:"publicAccess"` // Configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard
	Preset       bool             `json:"preset"`       // If `true` the dashboard will be marked as preset
	Enabled      bool             `json:"enabled"`      // The dashboard is shared (`true`) or private (`false`)
}

func (me *DashboardSharing) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"dashboard_id": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The Dynatrace entity ID of the dashboard",
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "The dashboard is shared (`true`) or private (`false`)",
		},
		"preset": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "If `true` the dashboard will be marked as preset",
		},
		"permissions": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(SharePermissions).Schema()},
			Description: "Access permissions of the dashboard",
		},
		"public": {
			Type:        hcl.TypeList,
			Optional:    true,
			MinItems:    1,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(AnonymousAccess).Schema()},
			Description: "Configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard",
		},
	}
}

// MarshalHCL has no documentation
func (me *DashboardSharing) MarshalHCL() (map[string]interface{}, error) {
	var err error
	props := hcl.Properties{}
	if props, err = props.EncodeAll(map[string]interface{}{
		"dashboard_id": me.DashboardID,
		"enabled":      me.Enabled,
		"preset":       me.Preset,
		"permissions":  me.Permissions,
	}); err != nil {
		return nil, err
	}

	if me.PublicAccess != nil && !me.PublicAccess.IsEmpty() {
		if err = props.Encode("public", me.PublicAccess); err != nil {
			return nil, err
		}
	}

	return props, nil
}

// UnmarshalHCL has no documentation
func (me *DashboardSharing) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeAll(map[string]interface{}{
		"dashboard_id": &me.DashboardID,
		"enabled":      &me.Enabled,
		"preset":       &me.Preset,
		"permissions":  &me.Permissions,
		"public":       &me.PublicAccess,
	}); err != nil {
		return err
	}
	if me.PublicAccess == nil {
		me.PublicAccess = &AnonymousAccess{
			ManagementZoneIDs: []string{},
			URLs:              map[string]string{},
		}
	}
	return nil
}
