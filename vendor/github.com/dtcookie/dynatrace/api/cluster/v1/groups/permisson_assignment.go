package groups

import "github.com/dtcookie/hcl"

type PermissionAssignments []*PermissionAssignment

func (me *PermissionAssignments) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"grant": {
			Type:        hcl.TypeList,
			Description: "A permission granted to one or multiple environments",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(PermissionAssignment).Schema()},
		},
	}
}

func (me PermissionAssignments) MarshalHCL() (map[string]interface{}, error) {
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
		result["grant"] = entries
	}
	return result, nil
}

func (me *PermissionAssignments) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.DecodeSlice("grant", me); err != nil {
		return err
	}
	return nil
}

type PermissionAssignment struct {
	Permission   Permission
	Environments []string
}

func (me *PermissionAssignment) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"permission": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The permission. Possible values are `VIEWER`, `MANAGE_SETTINGS`, `AGENT_INSTALL`, `LOG_VIEWER`, `VIEW_SENSITIVE_REQUEST_DATA`, `CONFIGURE_REQUEST_CAPTURE_DATA`, `REPLAY_SESSION_DATA`, `REPLAY_SESSION_DATA_WITHOUT_MASKING`, `MANAGE_SECURITY_PROBLEMS` and `MANAGE_SUPPORT_TICKETS`.",
		},
		"environments": {
			Type:        hcl.TypeSet,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
			Optional:    true,
			Description: "The ids of the environments this permission grants the user access to.",
		},
	}
}

func (me *PermissionAssignment) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	return properties.EncodeAll(map[string]interface{}{
		"permission":   me.Permission,
		"environments": me.Environments,
	})
}

func (me *PermissionAssignment) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"permission":   &me.Permission,
		"environments": &me.Environments,
	})
}

type Permission string

var Permissions = struct {
	Viewer                          Permission
	ManageSettings                  Permission
	AgentInstall                    Permission
	LogViewer                       Permission
	ViewSensitiveRequestData        Permission
	ConfigureRequestCaptureData     Permission
	ReplaySessionData               Permission
	ReplaySessionDataWithoutMasking Permission
	ManageSecurityProblems          Permission
	ManageSupportTickets            Permission
}{
	"VIEWER",
	"MANAGE_SETTINGS",
	"AGENT_INSTALL",
	"LOG_VIEWER",
	"VIEW_SENSITIVE_REQUEST_DATA",
	"CONFIGURE_REQUEST_CAPTURE_DATA",
	"REPLAY_SESSION_DATA",
	"REPLAY_SESSION_DATA_WITHOUT_MASKING",
	"MANAGE_SECURITY_PROBLEMS",
	"MANAGE_SUPPORT_TICKETS",
}
