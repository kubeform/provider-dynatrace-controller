package databaseservices

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/detection"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/failurerate"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/load"
	"github.com/dtcookie/dynatrace/api/config/anomalies/common/responsetime"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/xjson"
)

// AnomalyDetection The configuration of the anomaly detection for database services.
type AnomalyDetection struct {
	FailureRateIncrease            *failurerate.Detection      `json:"failureRateIncrease"`            // Configuration of failure rate increase detection.
	LoadDrop                       *load.DropDetection         `json:"loadDrop,omitempty"`             // The configuration of load drops detection.
	LoadSpike                      *load.SpikeDetection        `json:"loadSpike,omitempty"`            // The configuration of load spikes detection.
	ResponseTimeDegradation        *responsetime.Detection     `json:"responseTimeDegradation"`        // Configuration of response time degradation detection.
	DatabaseConnectionFailureCount *ConnectionFailureDetection `json:"databaseConnectionFailureCount"` // Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period.
	Metadata                       *api.ConfigMetadata         `json:"metadata,omitempty"`             // Metadata useful for debugging
}

func (me *AnomalyDetection) getDatabaseConnectionFailureCount() *ConnectionFailureDetection {
	if me.DatabaseConnectionFailureCount == nil {
		return &ConnectionFailureDetection{Enabled: false}
	}
	return me.DatabaseConnectionFailureCount
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

func (me *AnomalyDetection) getLoadSpike() *load.SpikeDetection {
	if me.LoadSpike == nil {
		return &load.SpikeDetection{Enabled: false}
	}
	return me.LoadSpike
}

func (me *AnomalyDetection) getLoadDrop() *load.DropDetection {
	if me.LoadDrop == nil {
		return &load.DropDetection{Enabled: false}
	}
	return me.LoadDrop
}

func (me *AnomalyDetection) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"load": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for anomalies regarding load drops and spikes",
			Elem:        &hcl.Resource{Schema: new(load.Detection).Schema()},
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
		"db_connect_failures": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Parameters of the failed database connections detection.  The alert is triggered when failed connections number exceeds **connectionFailsCount** during any **timePeriodMinutes** minutes period",
			Elem:        &hcl.Resource{Schema: new(ConnectionFailureDetection).Schema()},
		},
	}
}

func (me *AnomalyDetection) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	loadDetection := &load.Detection{
		Drops:  me.LoadDrop,
		Spikes: me.LoadSpike,
	}
	if !loadDetection.IsEmpty() {
		if marshalled, err := loadDetection.MarshalHCL(hcl.NewDecoder(decoder, "load", 0)); err == nil {
			result["load"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}

	if me.ResponseTimeDegradation != nil && string(me.ResponseTimeDegradation.DetectionMode) != string(detection.Modes.DontDetect) {
		if marshalled, err := me.ResponseTimeDegradation.MarshalHCL(hcl.NewDecoder(decoder, "response_time", 0)); err == nil {
			result["response_time"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.FailureRateIncrease != nil && string(me.FailureRateIncrease.DetectionMode) != string(detection.Modes.DontDetect) {
		if marshalled, err := me.FailureRateIncrease.MarshalHCL(hcl.NewDecoder(decoder, "failure_rate", 0)); err == nil {
			result["failure_rate"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.DatabaseConnectionFailureCount != nil && me.DatabaseConnectionFailureCount.Enabled {
		if marshalled, err := me.DatabaseConnectionFailureCount.MarshalHCL(hcl.NewDecoder(decoder, "db_connect_failures", 0)); err == nil {
			result["db_connect_failures"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *AnomalyDetection) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("load.#"); ok {
		loadDetection := new(load.Detection)
		if err := loadDetection.UnmarshalHCL(hcl.NewDecoder(decoder, "load", 0)); err != nil {
			return err
		}
		me.LoadDrop = loadDetection.Drops
		me.LoadSpike = loadDetection.Spikes
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
	if _, ok := decoder.GetOk("db_connect_failures.#"); ok {
		me.DatabaseConnectionFailureCount = new(ConnectionFailureDetection)
		if err := me.DatabaseConnectionFailureCount.UnmarshalHCL(hcl.NewDecoder(decoder, "db_connect_failures", 0)); err != nil {
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
		"loadSpike":                      me.getLoadSpike(),
		"responseTimeDegradation":        me.getResponseTimeDegradation(),
		"failureRateIncrease":            me.getFailureRateIncrease(),
		"loadDrop":                       me.getLoadDrop(),
		"databaseConnectionFailureCount": me.getDatabaseConnectionFailureCount(),
		"metadata":                       me.Metadata,
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
		"loadSpike":                      &me.LoadSpike,
		"responseTimeDegradation":        &me.ResponseTimeDegradation,
		"failureRateIncrease":            &me.FailureRateIncrease,
		"loadDrop":                       &me.LoadDrop,
		"databaseConnectionFailureCount": &me.DatabaseConnectionFailureCount,
		"metadata":                       &me.Metadata,
	}); err != nil {
		return err
	}
	return nil
}
