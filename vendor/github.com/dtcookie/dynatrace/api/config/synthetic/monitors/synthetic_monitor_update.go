package monitors

type SyntheticMonitorUpdate struct {
	ID                   *string            `json:"entityId,omitempty"`         // The ID of the monitor
	Name                 string             `json:"name"`                       // The name of the monitor
	Type                 Type               `json:"type"`                       // Defines the actual set of fields depending on the value. See one of the following objects: \n\n* `BROWSER` -> BrowserSyntheticMonitorUpdate \n* `HTTP` -> HttpSyntheticMonitorUpdate
	FrequencyMin         int32              `json:"frequencyMin"`               // The frequency of the monitor, in minutes. \n\n You can use one of the following values: `5`, `10`, `15`, `30`, and `60`
	Enabled              bool               `json:"enabled"`                    // The monitor is enabled (`true`) or disabled (`false`)
	AnomalyDetection     *AnomalyDetection  `json:"anomalyDetection,omitempty"` // Configuration for Anomaly Detection
	Locations            []string           `json:"locations"`                  // A list of locations from which the monitor is executed. \n\n To specify a location, use its entity ID
	Tags                 TagsWithSourceInfo `json:"tags"`                       // A set of tags assigned to the monitor. \n\n You can specify only the value of the tag here and the `CONTEXTLESS` context and source 'USER' will be added automatically. But preferred option is usage of TagWithSourceDto model
	ManuallyAssignedApps []string           `json:"manuallyAssignedApps"`       // A set of manually assigned applications
}

func (me *SyntheticMonitorUpdate) GetTags() TagsWithSourceInfo {
	if me.Tags == nil {
		return TagsWithSourceInfo{}
	}
	return me.Tags
}
