package monitors

import "github.com/dtcookie/hcl"

type LoadingTimeThresholds []*LoadingTimeThreshold

func (me *LoadingTimeThresholds) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"threshold": {
			Type:        hcl.TypeList,
			Description: "The list of performance threshold rules",
			Required:    true,
			MinItems:    1,
			Elem:        &hcl.Resource{Schema: new(LoadingTimeThreshold).Schema()},
		},
	}
}

func (me LoadingTimeThresholds) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	entries := []interface{}{}
	if len(me) > 0 {
		for _, entry := range me {
			if marshalled, err := entry.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["threshold"] = entries
	}
	return result, nil
}

func (me *LoadingTimeThresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeSlice("threshold", me)
}

// LoadingTimeThreshold The performance threshold rule
type LoadingTimeThreshold struct {
	Type         LoadingTimeThresholdType `json:"type"`         // The type of the threshold: total loading time or action loading time
	ValueMs      int32                    `json:"valueMs"`      // Notify if monitor takes longer than *X* milliseconds to load
	RequestIndex *int32                   `json:"requestIndex"` // Specify the request to which an ACTION threshold applies
	EventIndex   *int32                   `json:"eventIndex"`   // Specify the event to which an ACTION threshold applies
}

func (me *LoadingTimeThreshold) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"type": {
			Type:        hcl.TypeString,
			Description: "The type of the threshold: `TOTAL` (total loading time) or `ACTION` (action loading time)",
			Optional:    true,
		},
		"value_ms": {
			Type:        hcl.TypeInt,
			Description: "Notify if monitor takes longer than *X* milliseconds to load",
			Required:    true,
		},
		"request_index": {
			Type:        hcl.TypeInt,
			Description: "Specify the request to which an ACTION threshold applies",
			Optional:    true,
		},
		"event_index": {
			Type:        hcl.TypeInt,
			Description: "Specify the event to which an ACTION threshold applies",
			Optional:    true,
		},
	}
}

func (me *LoadingTimeThreshold) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["type"] = string(me.Type)
	result["value_ms"] = int(me.ValueMs)
	if me.RequestIndex != nil {
		result["request_index"] = int(*me.RequestIndex)
	}
	if me.EventIndex != nil {
		result["event_index"] = int(*me.EventIndex)
	}
	return result, nil
}

func (me *LoadingTimeThreshold) UnmarshalHCL(decoder hcl.Decoder) error {
	if err := decoder.Decode("type", &me.Type); err != nil {
		return err
	}
	if err := decoder.Decode("value_ms", &me.ValueMs); err != nil {
		return err
	}
	if err := decoder.Decode("request_index", &me.RequestIndex); err != nil {
		return err
	}
	if err := decoder.Decode("event_index", &me.EventIndex); err != nil {
		return err
	}
	return nil
}

// LoadingTimeThresholdType The type of the threshold: total loading time or action loading time
type LoadingTimeThresholdType string

// LoadingTimeThresholdTypes offers the known enum values
var LoadingTimeThresholdTypes = struct {
	Action LoadingTimeThresholdType
	Total  LoadingTimeThresholdType
}{
	"ACTION",
	"TOTAL",
}
