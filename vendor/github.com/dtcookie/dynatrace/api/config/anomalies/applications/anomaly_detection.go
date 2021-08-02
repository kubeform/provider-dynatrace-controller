package applications

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalies/applications/traffic"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/detection"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/failurerate"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/responsetime"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// AnomalyDetection The configuration of anomaly detection for applications.
type AnomalyDetection struct {
	ResponseTimeDegradation *responsetime.Detection `json:"responseTimeDegradation"` // Configuration of response time degradation detection.
	TrafficDrop             *traffic.DropDetection  `json:"trafficDrop"`             // The configuration of traffic drops detection.
	TrafficSpike            *traffic.SpikeDetection `json:"trafficSpike"`            // The configuration of traffic spikes detection.
	FailureRateIncrease     *failurerate.Detection  `json:"failureRateIncrease"`     // Configuration of failure rate increase detection.
	Metadata                *api.ConfigMetadata     `json:"metadata,omitempty"`      // Metadata useful for debugging
}

func (me *AnomalyDetection) getFailureRateIncrease() *failurerate.Detection {
	if me.FailureRateIncrease == nil {
		return &failurerate.Detection{DetectionMode: detection.Modes.DontDetect}
	}
	return me.FailureRateIncrease
}

func (me *AnomalyDetection) getResponseTimeDegradation() *responsetime.Detection {
	if me.ResponseTimeDegradation == nil {
		return &responsetime.Detection{DetectionMode: detection.Modes.DontDetect}
	}
	return me.ResponseTimeDegradation
}

func (me *AnomalyDetection) GetTrafficSpike() *traffic.SpikeDetection {
	if me.TrafficSpike == nil {
		return &traffic.SpikeDetection{Enabled: false}
	}
	return me.TrafficSpike
}

func (me *AnomalyDetection) GetTrafficDrop() *traffic.DropDetection {
	if me.TrafficDrop == nil {
		return &traffic.DropDetection{Enabled: false}
	}
	return me.TrafficDrop
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"traffic": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for anomalies regarding traffic",
			Elem:        &hcl.Resource{Schema: new(traffic.Detection).Schema()},
		},
		"response_time": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of response time degradation detection",
			Elem:        &hcl.Resource{Schema: new(responsetime.Detection).Schema()},
		},
		"failure_rate": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration of failure rate increase detection",
			Elem:        &hcl.Resource{Schema: new(failurerate.Detection).Schema()},
		},
	}
}

func (me *AnomalyDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	trafficDetection := &traffic.Detection{
		Drops:  me.TrafficDrop,
		Spikes: me.TrafficSpike,
	}
	if !trafficDetection.IsEmpty() {
		if marshalled, err := trafficDetection.MarshalHCL(hcl.NewDecoder(decoder, "traffic", 0)); err == nil {
			result["traffic"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.ResponseTimeDegradation != nil {
		if marshalled, err := me.ResponseTimeDegradation.MarshalHCL(hcl.NewDecoder(decoder, "response_time", 0)); err == nil {
			result["response_time"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.FailureRateIncrease != nil {
		if marshalled, err := me.FailureRateIncrease.MarshalHCL(hcl.NewDecoder(decoder, "failure_rate", 0)); err == nil {
			result["failure_rate"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *AnomalyDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("traffic.#"); ok {
		trafficDetection := new(traffic.Detection)
		if err := trafficDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "traffic", 0)); err != nil {
			return err
		}
		me.TrafficDrop = trafficDetection.Drops
		me.TrafficSpike = trafficDetection.Spikes
	}
	if _, ok := decoder.GetOk("response_time.#"); ok {
		me.ResponseTimeDegradation = new(responsetime.Detection)
		if err := me.ResponseTimeDegradation.UnmarshalHCL(hcl.NewDecoder(decoder, "response_time", 0)); err != nil {
			return err
		}
	} else {
		me.ResponseTimeDegradation = &responsetime.Detection{DetectionMode: detection.Modes.DontDetect}
	}
	if _, ok := decoder.GetOk("failure_rate.#"); ok {
		me.FailureRateIncrease = new(failurerate.Detection)
		if err := me.FailureRateIncrease.UnmarshalHCL(hcl.NewDecoder(decoder, "failure_rate", 0)); err != nil {
			return err
		}
	} else {
		me.FailureRateIncrease = &failurerate.Detection{DetectionMode: detection.Modes.DontDetect}
	}
	return nil
}

func (me *AnomalyDetection) MarshalJSON() ([]byte, error) {
	properties := xjson.Properties{}
	if err := properties.MarshalAll(map[string]interface{}{
		"trafficSpike":            me.GetTrafficSpike(),
		"responseTimeDegradation": me.getResponseTimeDegradation(),
		"failureRateIncrease":     me.getFailureRateIncrease(),
		"trafficDrop":             me.GetTrafficDrop(),
		"metadata":                me.Metadata,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *AnomalyDetection) UnmarshalJSON(data []byte) error {
	properties := xjson.Properties{}
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"trafficSpike":            &me.TrafficSpike,
		"responseTimeDegradation": &me.ResponseTimeDegradation,
		"failureRateIncrease":     &me.FailureRateIncrease,
		"trafficDrop":             &me.TrafficDrop,
		"metadata":                &me.Metadata,
	}); err != nil {
		return err
	}
	return nil
}
