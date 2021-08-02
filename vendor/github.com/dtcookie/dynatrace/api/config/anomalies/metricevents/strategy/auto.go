package strategy

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// Auto An auto-adaptive baseline strategy to detect anomalies within metrics that show a regular change over time, as the baseline is also updated automatically. An example is to detect an anomaly in the number of received network packets or within the number of user actions over time.
type Auto struct {
	BaseMonitoringStrategy
	AlertCondition             AlertCondition `json:"alertCondition"`                  // The condition for the **threshold** value check: above or below.
	AlertingOnMissingData      *bool          `json:"alertingOnMissingData,omitempty"` // If true, also one-minute samples without data are counted as violating samples.
	DealertingSamples          int32          `json:"dealertingSamples"`               // The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	NumberOfSignalFluctuations float64        `json:"numberOfSignalFluctuations"`      // Defines the factor of how many signal fluctuations are valid. Values above the baseline plus the signal fluctuation times the number of tolerated signal fluctuations are alerted.
	Samples                    int32          `json:"samples"`                         // The number of one-minute samples that form the sliding evaluation window.
	ViolatingSamples           int32          `json:"violatingSamples"`                // The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event.
}

func (me *Auto) GetType() Type {
	return Types.AutoAdaptiveBaseline
}

func (me *Auto) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"alert_condition": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The condition for the **threshold** value check: `ABOVE` or `BELOW`",
		},
		"alerting_on_missing_data": {
			Type:        hcl.TypeBool,
			Optional:    true,
			Description: "If true, also one-minute samples without data are counted as violating samples",
		},
		"dealerting_samples": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "The number of one-minute samples within the evaluation window that must go back to normal to close the event",
		},
		"signal_fluctuations": {
			Type:        hcl.TypeFloat,
			Required:    true,
			Description: "Defines the factor of how many signal fluctuations are valid. Values above the baseline plus the signal fluctuation times the number of tolerated signal fluctuations are alerted",
		},
		"samples": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "The number of one-minute samples that form the sliding evaluation window",
		},
		"violating_samples": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Auto) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if len(me.Unknowns) > 0 {
		data, err := json.Marshal(me.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["alert_condition"] = string(me.AlertCondition)
	if me.AlertingOnMissingData != nil {
		result["alerting_on_missing_data"] = *me.AlertingOnMissingData
	}
	result["dealerting_samples"] = int(me.DealertingSamples)
	result["signal_fluctuations"] = float64(me.NumberOfSignalFluctuations)
	result["samples"] = int(me.Samples)
	result["violating_samples"] = int(me.ViolatingSamples)
	return result, nil
}

func (me *Auto) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), me); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &me.Unknowns); err != nil {
			return err
		}
		delete(me.Unknowns, "alertCondition")
		delete(me.Unknowns, "alertingOnMissingData")
		delete(me.Unknowns, "dealertingSamples")
		delete(me.Unknowns, "numberOfSignalFluctuations")
		delete(me.Unknowns, "samples")
		delete(me.Unknowns, "violatingSamples")
		delete(me.Unknowns, "type")

		if len(me.Unknowns) == 0 {
			me.Unknowns = nil
		}
	}
	if value, ok := decoder.GetOk("alert_condition"); ok {
		me.AlertCondition = AlertCondition(value.(string))
	}
	if value, ok := decoder.GetOk("alerting_on_missing_data"); ok {
		me.AlertingOnMissingData = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("dealerting_samples"); ok {
		me.DealertingSamples = int32(value.(int))
	}
	if value, ok := decoder.GetOk("signal_fluctuations"); ok {
		me.NumberOfSignalFluctuations = value.(float64)
	}
	if value, ok := decoder.GetOk("samples"); ok {
		me.Samples = int32(value.(int))
	}
	if value, ok := decoder.GetOk("violating_samples"); ok {
		me.ViolatingSamples = int32(value.(int))
	}
	return nil
}

func (me *Auto) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":                       me.GetType(),
		"alertCondition":             me.AlertCondition,
		"alertingOnMissingData":      me.AlertingOnMissingData,
		"dealertingSamples":          me.DealertingSamples,
		"numberOfSignalFluctuations": me.NumberOfSignalFluctuations,
		"samples":                    me.Samples,
		"violatingSamples":           me.ViolatingSamples,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Auto) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"type":                       &me.Type,
		"alertCondition":             &me.AlertCondition,
		"alertingOnMissingData":      &me.AlertingOnMissingData,
		"dealertingSamples":          &me.DealertingSamples,
		"numberOfSignalFluctuations": &me.NumberOfSignalFluctuations,
		"samples":                    &me.Samples,
		"violatingSamples":           &me.ViolatingSamples,
	}); err != nil {
		return err
	}
	return nil
}
