package storage

import (
	"github.com/dtcookie/dynatrace/api/cluster/v2/envs/storage/retention"
	"github.com/dtcookie/hcl"
)

// Settings represents environment level storage usage and limit information. Not returned if includeStorageInfo param is not true. If skipped when editing via PUT method then already set limits will remain
type Settings struct {
	TransactionTrafficQuota *TransactionTrafficQuota `json:"transactionTrafficQuota"` // Maximum number of newly monitored entry point PurePaths captured per process/minute on environment level. Can be set to any value from 100 to 100000. If skipped when editing via PUT method then already set limit will remain
	UserActionsPerMinute    *UserActionsPerMinute    `json:"userActionsPerMinute"`    // Maximum number of user actions generated per minute on environment level. Can be set to any value from 1 to 2147483646 or left unlimited. If skipped when editing via PUT method then already set limit will remain

	Transactions              *Transactions              `json:"transactionStorage"`        // Transaction storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
	SessionReplayStorage      *SessionReplayStorage      `json:"sessionReplayStorage"`      // Session replay storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
	SymbolFilesFromMobileApps *SymbolFilesFromMobileApps `json:"symbolFilesFromMobileApps"` // Symbol files from mobile apps storage usage and limit information on environment level. If skipped when editing via PUT method then already set limit will remain
	LogMonitoringStorage      *LogMonitoringStorage      `json:"logMonitoringStorage"`      // Log monitoring storage usage and limit information on environment level. Not editable when Log monitoring is not allowed by license or not configured on cluster level. If skipped when editing via PUT method then already set limit will remain

	ServiceRequestLevelRetention *retention.ServiceRequestLevel `json:"serviceRequestLevelRetention"` // Service request level retention settings on environment level. Service code level retention time can't be greater than service request level retention time and both can't exceed one year.If skipped when editing via PUT method then already set limit will remain
	ServiceCodeLevelRetention    *retention.ServiceCodeLevel    `json:"serviceCodeLevelRetention"`    // Service code level retention settings on environment level. Service code level retention time can't be greater than service request level retention time and both can't exceed one year.If skipped when editing via PUT method then already set limit will remain
	RealUserMonitoringRetention  *retention.RealUserMonitoring  `json:"realUserMonitoringRetention"`  // Real user monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
	SyntheticMonitoringRetention *retention.SyntheticMonitoring `json:"syntheticMonitoringRetention"` // Synthetic monitoring retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
	SessionReplayRetention       *retention.SessionReplay       `json:"sessionReplayRetention"`       // Session replay retention settings on environment level. Can be set to any value from 1 to 35 days. If skipped when editing via PUT method then already set limit will remain
	LogMonitoringRetention       *retention.LogMonitoring       `json:"logMonitoringRetention"`       // Log monitoring retention settings on environment level. Not editable when Log monitoring is not allowed by license or not configured on cluster level. Can be set to any value from 5 to 90 days. If skipped when editing via PUT method then already set limit will remain
}

type limits struct {
	Transactions  *int64
	SessionReplay *int64
	SymbolFiles   *int64
	Logs          *int64
}

func (me *limits) IsEmpty() bool {
	return me == nil || (me.Transactions == nil && me.SessionReplay == nil && me.SymbolFiles == nil && me.Logs == nil)
}

func (me *limits) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"transactions": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Transaction storage usage and limit information on environment level in bytes. 0 for unlimited.",
		},
		"session_replay": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Session replay storage usage and limit information on environment level in bytes. 0 for unlimited.",
		},
		"symbol_files": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Session replay storage usage and limit information on environment level in bytes. 0 for unlimited.",
		},
		"logs": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Log monitoring storage usage and limit information on environment level in bytes. Not editable when Log monitoring is not allowed by license or not configured on cluster level. 0 for unlimited.",
		}}
}

func (me *limits) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if me.Transactions != nil {
		if err := properties.Encode("transactions", me.Transactions); err != nil {
			return nil, err
		}
	}
	if me.SessionReplay != nil {
		if err := properties.Encode("session_replay", me.SessionReplay); err != nil {
			return nil, err
		}
	}
	if me.SymbolFiles != nil {
		if err := properties.Encode("symbol_files", me.SymbolFiles); err != nil {
			return nil, err
		}
	}
	if me.Logs != nil {
		if err := properties.Encode("logs", me.Logs); err != nil {
			return nil, err
		}
	}
	return properties, nil
}

func (me *limits) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"transactions":   &me.Transactions,
		"session_replay": &me.SessionReplay,
		"symbol_files":   &me.SymbolFiles,
		"logs":           &me.Logs,
	})
}

type retent struct {
	ServiceRequestLevel int64
	ServiceCodeLevel    int64
	RUM                 int64
	Synthetic           int64
	SessionReplay       int64
	Logs                int64
}

func (me *retent) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"service_request_level": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Service request level retention settings on environment level in days. Service code level retention time can't be greater than service request level retention time and both can't exceed one year",
		},
		"service_code_level": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Service code level retention settings on environment level in days. Service code level retention time can't be greater than service request level retention time and both can't exceed one year",
		},
		"rum": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Real user monitoring retention settings on environment level in days. Can be set to any value from 1 to 35 days",
		},
		"synthetic": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Synthetic monitoring retention settings on environment level in days. Can be set to any value from 1 to 35 days",
		},
		"session_replay": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Session replay retention settings on environment level in days. Can be set to any value from 1 to 35 days",
		},
		"logs": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Log monitoring retention settings on environment level in days. Not editable when Log monitoring is not allowed by license or not configured on cluster level. Can be set to any value from 5 to 90 days",
		}}
}

func (me *retent) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if me.ServiceCodeLevel != 0 {
		if err := properties.Encode("service_code_level", me.ServiceCodeLevel); err != nil {
			return nil, err
		}
	}
	if me.ServiceCodeLevel != 0 {
		if err := properties.Encode("service_request_level", me.ServiceRequestLevel); err != nil {
			return nil, err
		}
	}
	if me.RUM != 0 {
		if err := properties.Encode("rum", me.RUM); err != nil {
			return nil, err
		}
	}
	if me.Synthetic != 0 {
		if err := properties.Encode("synthetic", me.Synthetic); err != nil {
			return nil, err
		}
	}
	if me.SessionReplay != 0 {
		if err := properties.Encode("session_replay", me.SessionReplay); err != nil {
			return nil, err
		}
	}
	if me.Logs != 0 {
		if err := properties.Encode("logs", me.Logs); err != nil {
			return nil, err
		}
	}
	return properties, nil
}

func (me *retent) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]interface{}{
		"service_request_level": &me.ServiceRequestLevel,
		"service_code_level":    &me.ServiceCodeLevel,
		"rum":                   &me.RUM,
		"synthetic":             &me.Synthetic,
		"session_replay":        &me.SessionReplay,
		"logs":                  &me.Logs,
	})
}

func (me *Settings) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"user_actions": {
			Type:        hcl.TypeInt,
			Optional:    true,
			Description: "Maximum number of user actions generated per minute on environment level. Can be set to any value from 1 to 2147483646 or left unlimited by omitting this property",
		},
		"transactions": {
			Type:        hcl.TypeInt,
			Required:    true,
			Description: "Maximum number of newly monitored entry point PurePaths captured per process/minute on environment level. Can be set to any value from 100 to 100000",
		},
		"limits": {
			Type:     hcl.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &hcl.Resource{Schema: new(limits).Schema()},
		},
		"retention": {
			Type:     hcl.TypeList,
			Optional: true,
			MinItems: 1,
			MaxItems: 1,
			Elem:     &hcl.Resource{Schema: new(retent).Schema()},
		},
	}
}

func (me *Settings) MarshalHCL() (map[string]interface{}, error) {
	properties := hcl.Properties{}
	if !me.UserActionsPerMinute.IsEmpty() {
		if err := properties.Encode("user_actions", me.UserActionsPerMinute.MaxLimit); err != nil {
			return nil, err
		}
	}
	if !me.TransactionTrafficQuota.IsEmpty() {
		if err := properties.Encode("transactions", me.TransactionTrafficQuota.MaxLimit); err != nil {
			return nil, err
		}
	}
	vLimits := new(limits)

	if !me.Transactions.IsEmpty() {
		vLimits.Transactions = me.Transactions.MaxLimit
	}
	if !me.SessionReplayStorage.IsEmpty() {
		vLimits.SessionReplay = me.SessionReplayStorage.MaxLimit
	}
	if !me.SymbolFilesFromMobileApps.IsEmpty() {
		vLimits.SymbolFiles = me.SymbolFilesFromMobileApps.MaxLimit
	}
	if !me.LogMonitoringStorage.IsEmpty() {
		vLimits.Logs = me.LogMonitoringStorage.MaxLimit
	}
	if !vLimits.IsEmpty() {
		if err := properties.Encode("limits", vLimits); err != nil {
			return nil, err
		}
	}

	if err := properties.Encode("retention", &retent{
		Logs:                me.LogMonitoringRetention.MaxLimitInDays,
		SessionReplay:       me.SessionReplayRetention.MaxLimitInDays,
		ServiceCodeLevel:    me.ServiceCodeLevelRetention.MaxLimitInDays,
		ServiceRequestLevel: me.ServiceRequestLevelRetention.MaxLimitInDays,
		RUM:                 me.RealUserMonitoringRetention.MaxLimitInDays,
		Synthetic:           me.SyntheticMonitoringRetention.MaxLimitInDays,
	}); err != nil {
		return nil, err
	}

	return properties, nil
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	me.UserActionsPerMinute = new(UserActionsPerMinute)
	if err := decoder.Decode("user_actions", &me.UserActionsPerMinute.MaxLimit); err != nil {
		return err
	}
	me.TransactionTrafficQuota = new(TransactionTrafficQuota)
	if err := decoder.Decode("transactions", &me.TransactionTrafficQuota.MaxLimit); err != nil {
		return err
	}

	vLimits := new(limits)
	if err := decoder.Decode("limits", &vLimits); err != nil {
		return err
	}

	me.Transactions = &Transactions{MaxLimit: vLimits.Transactions}
	me.SessionReplayStorage = &SessionReplayStorage{MaxLimit: vLimits.SessionReplay}
	me.SymbolFilesFromMobileApps = &SymbolFilesFromMobileApps{MaxLimit: vLimits.SymbolFiles}
	me.LogMonitoringStorage = &LogMonitoringStorage{MaxLimit: vLimits.Logs}

	ret := new(retent)
	if err := decoder.Decode("retention", &ret); err != nil {
		return err
	}
	me.LogMonitoringRetention = &retention.LogMonitoring{MaxLimitInDays: ret.Logs}
	me.SessionReplayRetention = &retention.SessionReplay{MaxLimitInDays: ret.SessionReplay}
	me.SyntheticMonitoringRetention = &retention.SyntheticMonitoring{MaxLimitInDays: ret.Synthetic}
	me.RealUserMonitoringRetention = &retention.RealUserMonitoring{MaxLimitInDays: ret.RUM}
	me.ServiceCodeLevelRetention = &retention.ServiceCodeLevel{MaxLimitInDays: ret.ServiceCodeLevel}
	me.ServiceRequestLevelRetention = &retention.ServiceRequestLevel{MaxLimitInDays: ret.ServiceRequestLevel}
	return nil
}
