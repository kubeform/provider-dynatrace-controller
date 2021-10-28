package quota

import "github.com/dtcookie/hcl"

// DavisDataUnits represents Davis Data Units consumption and quota information on environment level. Not set (and not editable) if Davis data units is not enabled. If skipped when editing via PUT method then already set quotas will remain
type DavisDataUnits struct {
	MonthlyLimit *int64 `json:"monthlyLimit"` // Monthly environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
	AnnualLimit  *int64 `json:"annualLimit"`  // Annual environment quota. Not set if unlimited. When updating via PUT method, skipping this field will set quota unlimited
}

func (me *DavisDataUnits) IsEmpty() bool {
	return me == nil || (me.MonthlyLimit == nil && me.AnnualLimit == nil)
}

func (me *DavisDataUnits) Schema() map[string]*hcl.Schema {
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

func (me *DavisDataUnits) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if me.MonthlyLimit != nil {
		if err := properties.Encode("monthly", me.MonthlyLimit); err != nil {
			return nil, err
		}
	}
	if me.AnnualLimit != nil {
		if err := properties.Encode("annual", me.AnnualLimit); err != nil {
			return nil, err
		}
	}
	return properties, nil
}

func (me *DavisDataUnits) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"monthly": &me.MonthlyLimit,
		"annual":  &me.AnnualLimit,
	})
}
