package monitors

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/api/config/synthetic/monitors/browser"
	"github.com/dtcookie/hcl"
)

// BrowserSyntheticMonitorUpdate Browser synthetic monitor update. Some fields are inherited from base `SyntheticMonitorUpdate` model
type BrowserSyntheticMonitorUpdate struct {
	SyntheticMonitorUpdate
	KeyPerformanceMetrics *KeyPerformanceMetrics `json:"keyPerformanceMetrics"` // The key performance metrics configuration
	Script                *browser.Script        `json:"script,omitempty"`
}

func (me *BrowserSyntheticMonitorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID                    *string                `json:"entityId,omitempty"`         // The ID of the monitor
		Name                  string                 `json:"name"`                       // The name of the monitor
		Type                  Type                   `json:"type"`                       // Defines the actual set of fields depending on the value. See one of the following objects: \n\n* `BROWSER` -> BrowserSyntheticMonitorUpdate \n* `HTTP` -> HttpSyntheticMonitorUpdate
		FrequencyMin          int32                  `json:"frequencyMin"`               // The frequency of the monitor, in minutes. \n\n You can use one of the following values: `5`, `10`, `15`, `30`, and `60`
		Enabled               bool                   `json:"enabled"`                    // The monitor is enabled (`true`) or disabled (`false`)
		AnomalyDetection      *AnomalyDetection      `json:"anomalyDetection,omitempty"` // Configuration for Anomaly Detection
		Locations             []string               `json:"locations"`                  // A list of locations from which the monitor is executed. \n\n To specify a location, use its entity ID
		Tags                  TagsWithSourceInfo     `json:"tags"`                       // A set of tags assigned to the monitor. \n\n You can specify only the value of the tag here and the `CONTEXTLESS` context and source 'USER' will be added automatically. But preferred option is usage of TagWithSourceDto model
		ManuallyAssignedApps  []string               `json:"manuallyAssignedApps"`       // A set of manually assigned applications
		KeyPerformanceMetrics *KeyPerformanceMetrics `json:"keyPerformanceMetrics"`      // The key performance metrics configuration
		Script                *browser.Script        `json:"script,omitempty"`
	}{
		me.ID,
		me.Name,
		me.GetType(),
		me.FrequencyMin,
		me.Enabled,
		me.AnomalyDetection,
		me.Locations,
		me.GetTags(),
		me.ManuallyAssignedApps,
		me.KeyPerformanceMetrics,
		me.Script,
	})
}

func (me *BrowserSyntheticMonitorUpdate) GetType() Type {
	return Types.Browser
}

func (me *BrowserSyntheticMonitorUpdate) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"name": {
			Type:        hcl.TypeString,
			Description: "The name of the monitor.",
			Required:    true,
		},
		"frequency": {
			Type:        hcl.TypeInt,
			Description: "The frequency of the monitor, in minutes.\n\nYou can use one of the following values: `5`, `10`, `15`, `30`, and `60`.",
			Required:    true,
		},
		"locations": {
			Type:        hcl.TypeSet,
			Description: "A list of locations from which the monitor is executed.\n\nTo specify a location, use its entity ID.",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"enabled": {
			Type:        hcl.TypeBool,
			Description: "The monitor is enabled (`true`) or disabled (`false`).",
			Optional:    true,
		},
		"manually_assigned_apps": {
			Type:        hcl.TypeSet,
			Description: "A set of manually assigned applications.",
			Optional:    true,
			MinItems:    1,
			Elem:        &hcl.Schema{Type: hcl.TypeString},
		},
		"tags": {
			Type:        hcl.TypeList,
			Description: "A set of tags assigned to the monitor.\n\nYou can specify only the value of the tag here and the `CONTEXTLESS` context and source 'USER' will be added automatically.",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(TagsWithSourceInfo).Schema(),
			},
		},
		"anomaly_detection": {
			Type:        hcl.TypeList,
			Description: "The anomaly detection configuration.",
			Optional:    true,
			MaxItems:    1,
			Elem: &hcl.Resource{
				Schema: new(AnomalyDetection).Schema(),
			},
		},
		"key_performance_metrics": {
			Type:        hcl.TypeList,
			Description: "The key performance metrics configuration",
			Optional:    true,
			MaxItems:    1,
			Elem: &hcl.Resource{
				Schema: new(KeyPerformanceMetrics).Schema(),
			},
		},
		"script": {
			Type:        hcl.TypeList,
			Description: "The Browser Script",
			Optional:    true,
			MaxItems:    1,
			Elem:        &hcl.Resource{Schema: new(browser.Script).Schema()},
		},
	}
}

func (me *BrowserSyntheticMonitorUpdate) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}
	result["name"] = me.Name
	result["frequency"] = me.FrequencyMin
	if len(me.Locations) > 0 {
		result["locations"] = me.Locations
	}
	result["enabled"] = me.Enabled
	if len(me.ManuallyAssignedApps) > 0 {
		result["manually_assigned_apps"] = me.ManuallyAssignedApps
	}
	if len(me.Tags) > 0 {
		if marshalled, err := me.Tags.MarshalHCL(); err == nil {
			result["tags"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.AnomalyDetection != nil {
		if marshalled, err := me.AnomalyDetection.MarshalHCL(); err == nil {
			result["anomaly_detection"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.KeyPerformanceMetrics != nil {
		if marshalled, err := me.KeyPerformanceMetrics.MarshalHCL(); err == nil {
			result["key_performance_metrics"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	if me.Script != nil {
		if marshalled, err := me.Script.MarshalHCL(); err == nil {
			result["script"] = []interface{}{marshalled}
		} else {
			return nil, err
		}
	}
	return result, nil
}

func (me *BrowserSyntheticMonitorUpdate) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("name"); ok {
		me.Name = value.(string)
	}
	if value, ok := decoder.GetOk("frequency"); ok {
		me.FrequencyMin = int32(value.(int))
	}
	me.Locations = decoder.GetStringSet("locations")
	if value, ok := decoder.GetOk("enabled"); ok {
		me.Enabled = value.(bool)
	}
	me.ManuallyAssignedApps = decoder.GetStringSet("manually_assigned_apps")
	if _, ok := decoder.GetOk("tags.#"); ok {
		me.Tags = TagsWithSourceInfo{}
		if err := me.Tags.UnmarshalHCL(hcl.NewDecoder(decoder, "tags", 0)); err != nil {
			return err
		}
	}
	if err := decoder.Decode("anomaly_detection", &me.AnomalyDetection); err != nil {
		return err
	}
	if err := decoder.Decode("key_performance_metrics", &me.KeyPerformanceMetrics); err != nil {
		return err
	}
	if err := decoder.Decode("script", &me.Script); err != nil {
		return err
	}
	return nil
}
