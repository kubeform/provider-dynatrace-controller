package scope

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

type AlertingScopes []AlertingScope

func (me AlertingScopes) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"custom_device_group_name": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for the related custom device group name",
			Elem:        &hcl.Resource{Schema: new(CustomDeviceGroupName).Schema()},
		},
		"entity": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for a monitored entity identifier",
			Elem:        &hcl.Resource{Schema: new(EntityID).Schema()},
		},
		"host_group_name": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for the related host group name",
			Elem:        &hcl.Resource{Schema: new(HostGroupName).Schema()},
		},
		"host_name": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for the related host name",
			Elem:        &hcl.Resource{Schema: new(HostName).Schema()},
		},
		"management_zone": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for a management zone identifier",
			Elem:        &hcl.Resource{Schema: new(ManagementZone).Schema()},
		},
		"name": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for a monitored entity name",
			Elem:        &hcl.Resource{Schema: new(Name).Schema()},
		},
		"process_group_id": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for a process group identifier",
			Elem:        &hcl.Resource{Schema: new(ProcessGroupID).Schema()},
		},
		"process_group_name": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for the related process group name",
			Elem:        &hcl.Resource{Schema: new(ProcessGroupName).Schema()},
		},
		"tag": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A scope filter for tags on entities",
			Elem:        &hcl.Resource{Schema: new(TagFilter).Schema()},
		},
		"scope": {
			Type:        hcl.TypeList,
			Optional:    true,
			Description: "A generic scope filter",
			Elem:        &hcl.Resource{Schema: new(BaseAlertingScope).Schema()},
		},
	}
}

func (me AlertingScopes) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}

	customDeviceGroupNames := []interface{}{}
	entityIDs := []interface{}{}
	hostGroupNames := []interface{}{}
	hostNames := []interface{}{}
	managementZones := []interface{}{}
	names := []interface{}{}
	processGroupIDs := []interface{}{}
	processGroupNames := []interface{}{}
	scopes := []interface{}{}
	tags := []interface{}{}
	for _, scope := range me {
		switch sc := scope.(type) {
		case *CustomDeviceGroupName:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "custom_device_group_name", len(customDeviceGroupNames))); err == nil {
				customDeviceGroupNames = append(customDeviceGroupNames, marshalled)
			} else {
				return nil, err
			}
		case *EntityID:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "entity", len(entityIDs))); err == nil {
				entityIDs = append(entityIDs, marshalled)
			} else {
				return nil, err
			}
		case *HostGroupName:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "host_group_name", len(hostGroupNames))); err == nil {
				hostGroupNames = append(hostGroupNames, marshalled)
			} else {
				return nil, err
			}
		case *HostName:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "host_name", len(hostNames))); err == nil {
				hostNames = append(hostNames, marshalled)
			} else {
				return nil, err
			}
		case *ManagementZone:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "management_zone", len(managementZones))); err == nil {
				managementZones = append(managementZones, marshalled)
			} else {
				return nil, err
			}
		case *Name:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "name", len(names))); err == nil {
				names = append(names, marshalled)
			} else {
				return nil, err
			}
		case *ProcessGroupID:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "process_group_id", len(processGroupIDs))); err == nil {
				processGroupIDs = append(processGroupIDs, marshalled)
			} else {
				return nil, err
			}
		case *ProcessGroupName:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "process_group_name", len(processGroupNames))); err == nil {
				processGroupNames = append(processGroupNames, marshalled)
			} else {
				return nil, err
			}
		case *TagFilter:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "tag", len(tags))); err == nil {
				tags = append(tags, marshalled)
			} else {
				return nil, err
			}
		case *BaseAlertingScope:
			if marshalled, err := sc.MarshalHCL(hcl.NewDecoder(decoder, "scope", len(scopes))); err == nil {
				scopes = append(scopes, marshalled)
			} else {
				return nil, err
			}
		default:
		}
		if len(customDeviceGroupNames) > 0 {
			result["custom_device_group_name"] = customDeviceGroupNames
		}
		if len(entityIDs) > 0 {
			result["entity"] = entityIDs
		}
		if len(hostGroupNames) > 0 {
			result["host_group_name"] = hostGroupNames
		}
		if len(hostNames) > 0 {
			result["host_name"] = hostNames
		}
		if len(managementZones) > 0 {
			result["management_zone"] = managementZones
		}
		if len(names) > 0 {
			result["name"] = names
		}
		if len(processGroupIDs) > 0 {
			result["process_group_id"] = processGroupIDs
		}
		if len(processGroupNames) > 0 {
			result["process_group_name"] = processGroupNames
		}
		if len(tags) > 0 {
			result["tag"] = tags
		}
		if len(scopes) > 0 {
			result["scope"] = scopes
		}
	}
	return result, nil
}

func (me *AlertingScopes) UnmarshalHCL(decoder hcl.Decoder) error {
	nme := AlertingScopes{}
	if result, ok := decoder.GetOk("custom_device_group_name.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(CustomDeviceGroupName)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "custom_device_group_name", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("entity.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(EntityID)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "entity", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("host_group_name.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(HostGroupName)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "host_group_name", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("host_name.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(HostName)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "host_name", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("management_zone.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(ManagementZone)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "management_zone", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("name.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(Name)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "name", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("process_group_id.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(ProcessGroupID)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "process_group_id", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("process_group_name.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(ProcessGroupName)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "process_group_name", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("tag.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(TagFilter)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "tag", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	if result, ok := decoder.GetOk("scope.#"); ok {
		for idx := 0; idx < result.(int); idx++ {
			entry := new(BaseAlertingScope)
			if err := entry.UnmarshalHCL(hcl.NewDecoder(decoder, "scope", idx)); err != nil {
				return err
			}
			nme = append(nme, entry)
		}
	}
	*me = nme
	return nil
}

func (me *AlertingScopes) UnmarshalJSON(data []byte) error {
	dims := AlertingScopes{}
	rawMessages := []json.RawMessage{}
	if err := json.Unmarshal(data, &rawMessages); err != nil {
		return err
	}
	for _, rawMessage := range rawMessages {
		properties := map[string]json.RawMessage{}
		if err := json.Unmarshal(rawMessage, &properties); err != nil {
			return err
		}
		if rawFilterType, found := properties["filterType"]; found {
			var sFilterType string
			if err := json.Unmarshal(rawFilterType, &sFilterType); err != nil {
				return err
			}
			switch sFilterType {
			case string(FilterTypes.CustomDeviceGroupName):
				cfg := new(CustomDeviceGroupName)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.EntityID):
				cfg := new(EntityID)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.HostGroupName):
				cfg := new(HostGroupName)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.HostName):
				cfg := new(HostName)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.ManagementZone):
				cfg := new(ManagementZone)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.Name):
				cfg := new(Name)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.ProcessGroupID):
				cfg := new(ProcessGroupID)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.ProcessGroupName):
				cfg := new(ProcessGroupName)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			case string(FilterTypes.Tag):
				cfg := new(TagFilter)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			default:
				cfg := new(BaseAlertingScope)
				if err := json.Unmarshal(rawMessage, &cfg); err != nil {
					return err
				}
				dims = append(dims, cfg)
			}
		}
		*me = dims
	}
	return nil
}
