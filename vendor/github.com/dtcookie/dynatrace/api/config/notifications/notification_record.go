package notifications

import (
	"encoding/json"

	"github.com/dtcookie/hcl"
)

type NotificationRecord struct {
	NotificationConfig NotificationConfig `json:"-"`
}

func (me *NotificationRecord) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"ansible_tower": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for Ansible Tower Notification",
			Elem:        &hcl.Resource{Schema: new(AnsibleTowerConfig).Schema()},
		},
		"email": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for Email Notification",
			Elem:        &hcl.Resource{Schema: new(EmailConfig).Schema()},
		},
		"hipchat": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for HipChat Notification",
			Elem:        &hcl.Resource{Schema: new(HipChatConfig).Schema()},
		},
		"jira": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for Jira Notification",
			Elem:        &hcl.Resource{Schema: new(JiraConfig).Schema()},
		},
		"ops_genie": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for OpsGenie Notification",
			Elem:        &hcl.Resource{Schema: new(OpsGenieConfig).Schema()},
		},
		"pager_duty": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for PagerDuty Notification",
			Elem:        &hcl.Resource{Schema: new(PagerDutyConfig).Schema()},
		},
		"service_now": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for ServiceNow Notification",
			Elem:        &hcl.Resource{Schema: new(ServiceNowConfig).Schema()},
		},
		"slack": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for Slack Notification",
			Elem:        &hcl.Resource{Schema: new(SlackConfig).Schema()},
		},
		"trello": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for Trello Notification",
			Elem:        &hcl.Resource{Schema: new(TrelloConfig).Schema()},
		},
		"victor_ops": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for VictorOps Notification",
			Elem:        &hcl.Resource{Schema: new(VictorOpsConfig).Schema()},
		},
		"web_hook": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for WebHook Notification",
			Elem:        &hcl.Resource{Schema: new(WebHookConfig).Schema()},
		},
		"xmatters": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for XMatters Notification",
			Elem:        &hcl.Resource{Schema: new(XMattersConfig).Schema()},
		},
		"config": {
			Type:        hcl.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Configuration for Generic Notification",
			Elem:        &hcl.Resource{Schema: new(BaseNotificationConfig).Schema()},
		},
	}
}

func (me *NotificationRecord) MarshalHCL(decoder hcl.Decoder) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	if me.NotificationConfig != nil {
		switch config := me.NotificationConfig.(type) {
		case *AnsibleTowerConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "ansible_tower", 0)); err == nil {
				result["ansible_tower"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *EmailConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "email", 0)); err == nil {
				result["email"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *HipChatConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "hipchat", 0)); err == nil {
				result["hipchat"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *JiraConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "jira", 0)); err == nil {
				result["jira"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *OpsGenieConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "ops_genie", 0)); err == nil {
				result["ops_genie"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *PagerDutyConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "page_duty", 0)); err == nil {
				result["page_duty"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *ServiceNowConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "service_now", 0)); err == nil {
				result["service_now"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *SlackConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "slack", 0)); err == nil {
				result["slack"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *TrelloConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "trello", 0)); err == nil {
				result["trello"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *VictorOpsConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "victor_ops", 0)); err == nil {
				result["victor_ops"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *WebHookConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "web_hook", 0)); err == nil {
				result["web_hook"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		case *XMattersConfig:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "xmatters", 0)); err == nil {
				result["xmatters"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		default:
			if marshalled, err := config.MarshalHCL(hcl.NewDecoder(decoder, "config", 0)); err == nil {
				result["config"] = []interface{}{marshalled}
			} else {
				return nil, err
			}
		}
	}
	return result, nil
}

func (me *NotificationRecord) UnmarshalHCL(decoder hcl.Decoder) error {
	if _, ok := decoder.GetOk("ansible_tower.#"); ok {
		cfg := new(AnsibleTowerConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "ansible_tower", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("email.#"); ok {
		cfg := new(EmailConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "email", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("hipchat.#"); ok {
		cfg := new(HipChatConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "hipchat", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("jira.#"); ok {
		cfg := new(JiraConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "jira", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("ops_genie.#"); ok {
		cfg := new(OpsGenieConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "ops_genie", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("pager_duty.#"); ok {
		cfg := new(PagerDutyConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "pager_duty", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("service_now.#"); ok {
		cfg := new(ServiceNowConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "service_now", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("slack.#"); ok {
		cfg := new(SlackConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "slack", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("trello.#"); ok {
		cfg := new(TrelloConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "trello", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("victor_ops.#"); ok {
		cfg := new(VictorOpsConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "victor_ops", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("web_hook.#"); ok {
		cfg := new(WebHookConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "web_hook", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("xmatters.#"); ok {
		cfg := new(XMattersConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "xmatters", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	if _, ok := decoder.GetOk("config.#"); ok {
		cfg := new(BaseNotificationConfig)
		if err := cfg.UnmarshalHCL(hcl.NewDecoder(decoder, "config", 0)); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	}
	return nil
}

func (me *NotificationRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(me.NotificationConfig)
}

func (me *NotificationRecord) UnmarshalJSON(data []byte) error {
	config := new(BaseNotificationConfig)
	if err := json.Unmarshal(data, &config); err != nil {
		return err
	}
	switch config.Type {
	case Types.Ansibletower:
		cfg := new(AnsibleTowerConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Email:
		cfg := new(EmailConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Hipchat:
		cfg := new(HipChatConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Jira:
		cfg := new(JiraConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.OpsGenie:
		cfg := new(OpsGenieConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.PagerDuty:
		cfg := new(PagerDutyConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.ServiceNow:
		cfg := new(ServiceNowConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Slack:
		cfg := new(SlackConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Trello:
		cfg := new(TrelloConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Victorops:
		cfg := new(VictorOpsConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Webhook:
		cfg := new(WebHookConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	case Types.Xmatters:
		cfg := new(XMattersConfig)
		if err := json.Unmarshal(data, &cfg); err != nil {
			return err
		}
		me.NotificationConfig = cfg
	default:
		me.NotificationConfig = config
	}
	return nil
}
