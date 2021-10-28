package quota

import "github.com/dtcookie/hcl"

// UserSessions represents user sessions consumption and quota information on environment level. If skipped when editing via PUT method then already set quotas will remain
type UserSessions struct {
	TotalAnnualLimit  *int64 `json:"totalAnnualLimit"`  // Annual total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	TotalMonthlyLimit *int64 `json:"totalMonthlyLimit"` // Monthly total User sessions environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
}

func (me *UserSessions) IsEmpty() bool {
	return me == nil || (me.TotalAnnualLimit == nil && me.TotalMonthlyLimit == nil)
}

func (me *UserSessions) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"annual": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Annual total User sessions environment quota. Not set if unlimited",
		},
		"monthly": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Monthly total User sessions environment quota. Not set if unlimited",
		},
	}
}

func (me *UserSessions) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"monthly": me.TotalMonthlyLimit,
		"annual":  me.TotalAnnualLimit,
	})
}

func (me *UserSessions) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"monthly": &me.TotalMonthlyLimit,
		"annual":  &me.TotalAnnualLimit,
	})
}
