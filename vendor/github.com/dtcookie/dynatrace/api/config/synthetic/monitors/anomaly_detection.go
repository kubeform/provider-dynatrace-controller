package monitors

import "github.com/dtcookie/hcl"

// AnomalyDetection The anomaly detection configuration
type AnomalyDetection struct {
	OutageHandling        *OutageHandlingPolicy        `json:"outageHandling,omitempty"`
	LoadingTimeThresholds *LoadingTimeThresholdsPolicy `json:"loadingTimeThresholds,omitempty"`
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"outage_handling": {
			Type:        hcl.TypeList,
			Description: "Outage handling configuration",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(OutageHandlingPolicy).Schema(),
			},
		},
		"loading_time_thresholds": {
			Type:        hcl.TypeList,
			Description: "Thresholds for loading times",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(LoadingTimeThresholdsPolicy).Schema(),
			},
		},
	}
}

func (me *AnomalyDetection) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.OutageHandling != nil {
		if marshalled, err := me.OutageHandling.MarshalHCL(); err == nil {
			result["outage_handling"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.LoadingTimeThresholds != nil {
		if marshalled, err := me.LoadingTimeThresholds.MarshalHCL(); err == nil {
			result["loading_time_thresholds"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *AnomalyDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("outage_handling", &me.OutageHandling); err != nil {
		return err
	}
	if err := decoder.Decode("loading_time_thresholds", &me.LoadingTimeThresholds); err != nil {
		return err
	}
	return nil
}
