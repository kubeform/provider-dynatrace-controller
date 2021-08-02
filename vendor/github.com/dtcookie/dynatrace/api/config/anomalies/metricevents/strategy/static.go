package strategy

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dtcookie/xjson"
)

// MetricEventStaticMonitoringStrategy A static threshold monitoring strategy to alert on hard limits within a given metric. An example is the violation of a critical memory limit.
type Static struct {
	BaseMonitoringStrategy
	AlertCondition        AlertCondition `json:"alertCondition"`                  // The condition for the **threshold** value check: above or below.
	AlertingOnMissingData *bool          `json:"alertingOnMissingData,omitempty"` // If true, also one-minute samples without data are counted as violating samples.
	DealertingSamples     int32          `json:"dealertingSamples"`               // The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	Samples               int32          `json:"samples"`                         // The number of one-minute samples that form the sliding evaluation window.
	Threshold             float64        `json:"threshold"`                       // The value of the static threshold based on the specified unit.
	Unit                  Unit           `json:"unit"`                            // The unit of the threshold, matching the metric definition.
	ViolatingSamples      int32          `json:"violatingSamples"`                // The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event.
}

func (me *Static) GetType() Type {
	return Types.StaticThreshold
}

func (me *Static) Schema() map[string]*hcl.Schema {
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
		"threshold": {
			Type:        hcl.TypeFloat,
			Required:    true,
			Description: "The value of the static threshold based on the specified unit",
		},
		"unit": {
			Type:        hcl.TypeString,
			Required:    true,
			Description: "The unit of the threshold, matching the metric definition",
		},
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "allows for configuring properties that are not explicitly supported by the current version of this provider",
			Optional:    true,
		},
	}
}

func (me *Static) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
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
	result["samples"] = int(me.Samples)
	result["violating_samples"] = int(me.ViolatingSamples)
	result["threshold"] = float64(me.Threshold)
	result["unit"] = string(me.Unit)
	return result, nil
}

func (me *Static) UnmarshalHCL(decoder hcl.Decoder) error {
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
		delete(me.Unknowns, "unit")
		delete(me.Unknowns, "threshold")
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
	if value, ok := decoder.GetOk("samples"); ok {
		me.Samples = int32(value.(int))
	}
	if value, ok := decoder.GetOk("unit"); ok {
		me.Unit = Unit(value.(string))
	}
	if value, ok := decoder.GetOk("threshold"); ok {
		me.Threshold = value.(float64)
	}
	if value, ok := decoder.GetOk("violating_samples"); ok {
		me.ViolatingSamples = int32(value.(int))
	}
	return nil
}

func (me *Static) MarshalJSON() ([]byte, error) {
	properties := xjson.NewProperties(me.Unknowns)
	if err := properties.MarshalAll(map[string]interface{}{
		"type":                  me.GetType(),
		"alertCondition":        me.AlertCondition,
		"alertingOnMissingData": me.AlertingOnMissingData,
		"dealertingSamples":     me.DealertingSamples,
		"threshold":             me.Threshold,
		"unit":                  me.Unit,
		"samples":               me.Samples,
		"violatingSamples":      me.ViolatingSamples,
	}); err != nil {
		return nil, err
	}
	return json.Marshal(properties)
}

func (me *Static) UnmarshalJSON(data []byte) error {
	properties := xjson.NewProperties(me.Unknowns)
	if err := json.Unmarshal(data, &properties); err != nil {
		return err
	}
	if err := properties.UnmarshalAll(map[string]interface{}{
		"type":                  &me.Type,
		"alertCondition":        &me.AlertCondition,
		"alertingOnMissingData": &me.AlertingOnMissingData,
		"dealertingSamples":     &me.DealertingSamples,
		"threshold":             &me.Threshold,
		"unit":                  &me.Unit,
		"samples":               &me.Samples,
		"violatingSamples":      &me.ViolatingSamples,
	}); err != nil {
		return err
	}
	return nil
}

// Unit The unit of the threshold, matching the metric definition.
type Unit string

// Units offers the known enum values
var Units = struct {
	Bit                  Unit
	BitPerHour           Unit
	BitPerMinute         Unit
	BitPerSecond         Unit
	Byte                 Unit
	BytePerHour          Unit
	BytePerMinute        Unit
	BytePerSecond        Unit
	Cores                Unit
	Count                Unit
	Day                  Unit
	DecibelMilliWatt     Unit
	GibiByte             Unit
	Giga                 Unit
	GigaByte             Unit
	Hour                 Unit
	KibiByte             Unit
	KibiBytePerHour      Unit
	KibiBytePerMinute    Unit
	KibiBytePerSecond    Unit
	Kilo                 Unit
	KiloByte             Unit
	KiloBytePerHour      Unit
	KiloBytePerMinute    Unit
	KiloBytePerSecond    Unit
	MebiByte             Unit
	MebiBytePerHour      Unit
	MebiBytePerMinute    Unit
	MebiBytePerSecond    Unit
	Mega                 Unit
	MegaByte             Unit
	MegaBytePerHour      Unit
	MegaBytePerMinute    Unit
	MegaBytePerSecond    Unit
	MicroSecond          Unit
	MilliCores           Unit
	MilliSecond          Unit
	MilliSecondPerMinute Unit
	Minute               Unit
	Month                Unit
	Msu                  Unit
	NanoSecond           Unit
	NanoSecondPerMinute  Unit
	NotApplicable        Unit
	Percent              Unit
	PerHour              Unit
	PerMinute            Unit
	PerSecond            Unit
	Pixel                Unit
	Promille             Unit
	Ratio                Unit
	Second               Unit
	State                Unit
	Unspecified          Unit
	Week                 Unit
	Year                 Unit
}{
	"BIT",
	"BIT_PER_HOUR",
	"BIT_PER_MINUTE",
	"BIT_PER_SECOND",
	"BYTE",
	"BYTE_PER_HOUR",
	"BYTE_PER_MINUTE",
	"BYTE_PER_SECOND",
	"CORES",
	"COUNT",
	"DAY",
	"DECIBEL_MILLI_WATT",
	"GIBI_BYTE",
	"GIGA",
	"GIGA_BYTE",
	"HOUR",
	"KIBI_BYTE",
	"KIBI_BYTE_PER_HOUR",
	"KIBI_BYTE_PER_MINUTE",
	"KIBI_BYTE_PER_SECOND",
	"KILO",
	"KILO_BYTE",
	"KILO_BYTE_PER_HOUR",
	"KILO_BYTE_PER_MINUTE",
	"KILO_BYTE_PER_SECOND",
	"MEBI_BYTE",
	"MEBI_BYTE_PER_HOUR",
	"MEBI_BYTE_PER_MINUTE",
	"MEBI_BYTE_PER_SECOND",
	"MEGA",
	"MEGA_BYTE",
	"MEGA_BYTE_PER_HOUR",
	"MEGA_BYTE_PER_MINUTE",
	"MEGA_BYTE_PER_SECOND",
	"MICRO_SECOND",
	"MILLI_CORES",
	"MILLI_SECOND",
	"MILLI_SECOND_PER_MINUTE",
	"MINUTE",
	"MONTH",
	"MSU",
	"NANO_SECOND",
	"NANO_SECOND_PER_MINUTE",
	"NOT_APPLICABLE",
	"PERCENT",
	"PER_HOUR",
	"PER_MINUTE",
	"PER_SECOND",
	"PIXEL",
	"PROMILLE",
	"RATIO",
	"SECOND",
	"STATE",
	"UNSPECIFIED",
	"WEEK",
	"YEAR",
}
