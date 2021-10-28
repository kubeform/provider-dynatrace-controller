package quota

import "github.com/dtcookie/hcl"

// Synthetic represents synthetic monitors consumption and quota information on environment level. Not set (and not editable) if neither Synthetic nor DEM units is enabled. If skipped when editing via PUT method then already set quotas will remain
type Synthetic struct {
	MonthlyLimit *int64 `json:"monthlyLimit"` // Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	AnnualLimit  *int64 `json:"annualLimit"`  // Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
}

func (me *Synthetic) IsEmpty() bool {
	return me == nil || (me.MonthlyLimit == nil && me.AnnualLimit == nil)
}

func (me *Synthetic) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"monthly": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Monthly environment quota. Not set if unlimited",
		},
		"annual": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Annual environment quota. Not set if unlimited",
		},
	}
}

func (me *Synthetic) MarshalHCL() (map[string]interface{}, error) {
	return hcl.Properties{}.EncodeAll(map[string]interface{}{
		"monthly": me.MonthlyLimit,
		"annual":  me.AnnualLimit,
	})
}

func (me *Synthetic) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"monthly": &me.MonthlyLimit,
		"annual":  &me.AnnualLimit,
	})
}
