package monitors

import "github.com/dtcookie/hcl"

// LocalOutagePolicy Local outage handling configuration. \n\n Alert if **affectedLocations** of locations are unable to access the web application **consecutiveRuns** times consecutively
type LocalOutagePolicy struct {
	AffectedLocations *int32 `json:"affectedLocations"` // The number of affected locations to trigger an alert
	ConsecutiveRuns   *int32 `json:"consecutiveRuns"`   // The number of consecutive fails to trigger an alert
}

func (me *LocalOutagePolicy) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"affected_locations": {
			Type:        hcl.TypeInt,
			Description: "The number of affected locations to trigger an alert",
			Required:    true,
		},
		"consecutive_runs": {
			Type:        hcl.TypeInt,
			Description: "The number of consecutive fails to trigger an alert",
			Required:    true,
		},
	}
}

func (me *LocalOutagePolicy) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["affected_locations"] = *me.AffectedLocations
	result["consecutive_runs"] = *me.ConsecutiveRuns
	return result, nil
}

func (me *LocalOutagePolicy) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("affected_locations", &me.AffectedLocations); err != nil {
		return err
	}
	if err := decoder.Decode("consecutive_runs", &me.ConsecutiveRuns); err != nil {
		return err
	}
	return nil
}
