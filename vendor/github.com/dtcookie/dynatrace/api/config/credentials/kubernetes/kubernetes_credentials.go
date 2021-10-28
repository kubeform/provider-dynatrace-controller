package kubernetes

import (
	"encoding/json"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
)

// KubernetesCredentials Configuration for specific Kubernetes credentials.
type KubernetesCredentials struct {
	EventsIntegrationEnabled              *bool                      `json:"eventsIntegrationEnabled,omitempty"`        // The monitoring of events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.  If not set on creation, the `false` value is used.  If the field is omitted during an update, the old value remains unaffected.
	AuthToken                             *string                    `json:"authToken,omitempty"`                       // The service account bearer token for the Kubernetes API server.  Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.  If the field is omitted during an update, the old value remains unaffected.
	Active                                *bool                      `json:"active,omitempty"`                          // The monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	EndpointStatusInfo                    *string                    `json:"endpointStatusInfo,omitempty"`              // The detailed status info of the configured endpoint.
	WorkloadIntegrationEnabled            *bool                      `json:"workloadIntegrationEnabled,omitempty"`      // Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.
	EndpointStatus                        *EndpointStatus            `json:"endpointStatus,omitempty"`                  // The status of the configured endpoint. ASSIGNED: The credentials are assigned to an ActiveGate which is responsible for processing. UNASSIGNED: The credentials are not yet assigned to an ActiveGate so there is currently no processing. DISABLED: The credentials have been disabled by the user. FASTCHECK_AUTH_ERROR: The credentials are invalid. FASTCHECK_TLS_ERROR: The endpoint TLS certificate is invalid. FASTCHECK_NO_RESPONSE: The endpoint did not return a result until the timeout was reached. FASTCHECK_INVALID_ENDPOINT: The endpoint URL was invalid. FASTCHECK_AUTH_LOCKED: The credentials seem to be locked. UNKNOWN: An unknown error occured.
	Label                                 string                     `json:"label"`                                     // The name of the Kubernetes credentials configuration.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.
	ID                                    *string                    `json:"id,omitempty"`                              // The ID of the given credentials configuration.
	Metadata                              *api.ConfigurationMetadata `json:"metadata,omitempty"`                        // Metadata useful for debugging
	CertificateCheckEnabled               *bool                      `json:"certificateCheckEnabled,omitempty"`         // The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.
	EndpointURL                           string                     `json:"endpointUrl"`                               // The URL of the Kubernetes API server.  It must be unique within a Dynatrace environment.  The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.
	HostnameVerificationEnabled           *bool                      `json:"hostnameVerificationEnabled"`               // Verify hostname in certificate against Kubernetes API URL
	PrometheusExportersIntegrationEnabled *bool                      `json:"prometheusExportersIntegrationEnabled"`     // Prometheus exporters integration is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.If the field is omitted during an update, the old value remains unaffected
	EventsFieldSelectors                  []*KubernetesEventPattern  `json:"eventsFieldSelectors,omitempty"`            // Kubernetes event filters based on field-selectors. If set to `null` on creation, no events field selectors are subscribed. If set to `null` on update, no change of stored events field selectors is applied. Set an empty list to clear all events field selectors.
	DavisEventsIntegrationEnabled         *bool                      `json:"davisEventsIntegrationEnabled,omitempty"`   // Inclusion of all Davis relevant events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected
	EventAnalysisAndAlertingEnabled       *bool                      `json:"eventAnalysisAndAlertingEnabled,omitempty"` // Event analysis and alerting is (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected
	Unknowns                              map[string]json.RawMessage `json:"-"`
}

func (kc *KubernetesCredentials) Schema() map[string]*hcl.Schema {
	return map[string]*hcl.Schema{
		"unknowns": {
			Type:        hcl.TypeString,
			Description: "Any attributes that aren't yet supported by this provider",
			Optional:    true,
		},
		"endpoint_url": {
			Type:        hcl.TypeString,
			Description: "The URL of the Kubernetes API server.  It must be unique within a Dynatrace environment.  The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed.",
			Optional:    true,
		},
		"label": {
			Type:        hcl.TypeString,
			Description: "The name of the Kubernetes credentials configuration.  Allowed characters are letters, numbers, whitespaces, and the following characters: `.+-_`. Leading or trailing whitespace is not allowed.",
			Required:    true,
		},
		"events_integration_enabled": {
			Type:        hcl.TypeBool,
			Description: "Monitoring of events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. Event monitoring depends on the active state of this configuration to be true.  If not set on creation, the `false` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"event_analysis_and_alerting_enabled": {
			Type:        hcl.TypeBool,
			Description: "Event analysis and alerting is (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"hostname_verification": {
			Type:        hcl.TypeBool,
			Description: "Verify hostname in certificate against Kubernetes API URL",
			Optional:    true,
		},
		"prometheus_exporters": {
			Type:        hcl.TypeBool,
			Description: "Prometheus exporters integration is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.If the field is omitted during an update, the old value remains unaffected",
			Optional:    true,
		},
		"davis_events_integration_enabled": {
			Type:        hcl.TypeBool,
			Description: "Inclusion of all Davis relevant events is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected",
			Optional:    true,
		},
		"active": {
			Type:        hcl.TypeBool,
			Description: "Monitoring is enabled (`true`) or disabled (`false`) for given credentials configuration.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"workload_integration_enabled": {
			Type:        hcl.TypeBool,
			Description: "Workload and cloud application processing is enabled (`true`) or disabled (`false`) for the Kubernetes cluster. If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"auth_token": {
			Type:        hcl.TypeString,
			Description: "The service account bearer token for the Kubernetes API server.  Submit your token on creation or update of the configuration. For security reasons, GET requests return this field as `null`.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
		},
		"certificate_check_enabled": {
			Type:        hcl.TypeBool,
			Description: "The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
			Default:     true,
		},
		"events_field_selectors": {
			Type:        hcl.TypeList,
			Description: "The check of SSL certificates is enabled (`true`) or disabled (`false`) for the Kubernetes cluster.  If not set on creation, the `true` value is used.  If the field is omitted during an update, the old value remains unaffected.",
			Optional:    true,
			Elem: &hcl.Resource{
				Schema: new(KubernetesEventPattern).Schema(),
			},
		},
	}
}

func (kc *KubernetesCredentials) MarshalHCL() (map[string]interface{}, error) {
	result := map[string]interface{}{}

	if kc.Unknowns != nil {
		data, err := json.Marshal(kc.Unknowns)
		if err != nil {
			return nil, err
		}
		result["unknowns"] = string(data)
	}
	result["active"] = kc.Active
	result["hostname_verification"] = opt.Bool(kc.HostnameVerificationEnabled)
	result["prometheus_exporters"] = opt.Bool(kc.PrometheusExportersIntegrationEnabled)
	result["workload_integration_enabled"] = opt.Bool(kc.WorkloadIntegrationEnabled)
	result["certificate_check_enabled"] = opt.Bool(kc.CertificateCheckEnabled)
	result["events_integration_enabled"] = opt.Bool(kc.EventsIntegrationEnabled)
	result["davis_events_integration_enabled"] = opt.Bool(kc.DavisEventsIntegrationEnabled)
	result["event_analysis_and_alerting_enabled"] = opt.Bool(kc.EventAnalysisAndAlertingEnabled)

	result["label"] = kc.Label
	if kc.AuthToken != nil {
		result["auth_token"] = *kc.AuthToken
	}
	result["endpoint_url"] = kc.EndpointURL
	if kc.EventsFieldSelectors != nil {
		entries := []interface{}{}
		for _, eventPattern := range kc.EventsFieldSelectors {
			if marshalled, err := eventPattern.MarshalHCL(); err == nil {
				entries = append(entries, marshalled)
			} else {
				return nil, err
			}
		}
		result["events_field_selectors"] = entries
	}
	return result, nil
}

func (kc *KubernetesCredentials) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("unknowns"); ok {
		if err := json.Unmarshal([]byte(value.(string)), kc); err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(value.(string)), &kc.Unknowns); err != nil {
			return err
		}
		delete(kc.Unknowns, "events_integration_enabled")
		delete(kc.Unknowns, "active")
		delete(kc.Unknowns, "endpoint_status_info")
		delete(kc.Unknowns, "workload_integration_enabled")
		delete(kc.Unknowns, "endpoint_status")
		delete(kc.Unknowns, "label")
		delete(kc.Unknowns, "id")
		delete(kc.Unknowns, "metadata")
		delete(kc.Unknowns, "auth_token")
		delete(kc.Unknowns, "certificate_check_enabled")
		delete(kc.Unknowns, "endpoint_url")
		delete(kc.Unknowns, "events_field_selectors")
		delete(kc.Unknowns, "hostname_verification")
		delete(kc.Unknowns, "prometheus_exporters")
		delete(kc.Unknowns, "davis_events_integration_enabled")
		delete(kc.Unknowns, "event_analysis_and_alerting_enabled")

		if len(kc.Unknowns) == 0 {
			kc.Unknowns = nil
		}
	}
	if _, value := decoder.GetChange("events_integration_enabled"); value != nil {
		kc.EventsIntegrationEnabled = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("event_analysis_and_alerting_enabled"); value != nil {
		kc.EventAnalysisAndAlertingEnabled = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("active"); value != nil {
		kc.Active = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("workload_integration_enabled"); value != nil {
		kc.WorkloadIntegrationEnabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("label"); ok {
		kc.Label = value.(string)
	}
	if value, ok := decoder.GetOk("id"); ok {
		kc.ID = opt.NewString(value.(string))
	}
	if _, value := decoder.GetChange("hostname_verification"); value != nil {
		kc.HostnameVerificationEnabled = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("prometheus_exporters"); value != nil {
		kc.PrometheusExportersIntegrationEnabled = opt.NewBool(value.(bool))
	}
	if _, value := decoder.GetChange("davis_events_integration_enabled"); value != nil {
		kc.DavisEventsIntegrationEnabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("auth_token"); ok {
		kc.AuthToken = opt.NewString(value.(string))
	}
	if _, value := decoder.GetChange("certificate_check_enabled"); value != nil {
		kc.CertificateCheckEnabled = opt.NewBool(value.(bool))
	}
	if value, ok := decoder.GetOk("endpoint_url"); ok {
		kc.EndpointURL = value.(string)
	}
	if result, ok := decoder.GetOk("events_field_selectors.#"); ok {
		kc.EventsFieldSelectors = []*KubernetesEventPattern{}
		for idx := 0; idx < result.(int); idx++ {
			eventPattern := new(KubernetesEventPattern)
			if err := eventPattern.UnmarshalHCL(hcl.NewDecoder(decoder, "events_field_selectors", idx)); err != nil {
				return err
			}
			kc.EventsFieldSelectors = append(kc.EventsFieldSelectors, eventPattern)
		}
	}
	return nil
}

// UnmarshalJSON provides custom JSON deserialization
func (kc *KubernetesCredentials) UnmarshalJSON(data []byte) error {
	m := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if v, found := m["eventAnalysisAndAlertingEnabled"]; found {
		if err := json.Unmarshal(v, &kc.EventAnalysisAndAlertingEnabled); err != nil {
			return err
		}
	}
	if v, found := m["eventsIntegrationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.EventsIntegrationEnabled); err != nil {
			return err
		}
	}
	if v, found := m["hostnameVerificationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.HostnameVerificationEnabled); err != nil {
			return err
		}
	}
	if v, found := m["prometheusExportersIntegrationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.PrometheusExportersIntegrationEnabled); err != nil {
			return err
		}
	}
	if v, found := m["davisEventsIntegrationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.DavisEventsIntegrationEnabled); err != nil {
			return err
		}
	}

	if v, found := m["active"]; found {
		if err := json.Unmarshal(v, &kc.Active); err != nil {
			return err
		}
	}
	if v, found := m["endpointStatusInfo"]; found {
		if err := json.Unmarshal(v, &kc.EndpointStatusInfo); err != nil {
			return err
		}
	}
	if v, found := m["workloadIntegrationEnabled"]; found {
		if err := json.Unmarshal(v, &kc.WorkloadIntegrationEnabled); err != nil {
			return err
		}
	}
	if v, found := m["endpointStatus"]; found {
		if err := json.Unmarshal(v, &kc.EndpointStatus); err != nil {
			return err
		}
	}
	if v, found := m["label"]; found {
		if err := json.Unmarshal(v, &kc.Label); err != nil {
			return err
		}
	}
	if v, found := m["id"]; found {
		if err := json.Unmarshal(v, &kc.ID); err != nil {
			return err
		}
	}
	if v, found := m["metadata"]; found {
		if err := json.Unmarshal(v, &kc.Metadata); err != nil {
			return err
		}
	}
	if v, found := m["authToken"]; found {
		if err := json.Unmarshal(v, &kc.AuthToken); err != nil {
			return err
		}
	}
	if v, found := m["certificateCheckEnabled"]; found {
		if err := json.Unmarshal(v, &kc.CertificateCheckEnabled); err != nil {
			return err
		}
	}
	if v, found := m["endpointUrl"]; found {
		if err := json.Unmarshal(v, &kc.EndpointURL); err != nil {
			return err
		}
	}
	if v, found := m["eventsFieldSelectors"]; found {
		if err := json.Unmarshal(v, &kc.EventsFieldSelectors); err != nil {
			return err
		}
	}
	delete(m, "eventAnalysisAndAlertingEnabled")
	delete(m, "eventsIntegrationEnabled")
	delete(m, "active")
	delete(m, "endpointStatusInfo")
	delete(m, "workloadIntegrationEnabled")
	delete(m, "endpointStatus")
	delete(m, "label")
	delete(m, "id")
	delete(m, "metadata")
	delete(m, "authToken")
	delete(m, "certificateCheckEnabled")
	delete(m, "endpointUrl")
	delete(m, "eventsFieldSelectors")
	delete(m, "hostnameVerificationEnabled")
	delete(m, "prometheusExportersIntegrationEnabled")
	delete(m, "davisEventsIntegrationEnabled")

	if len(m) > 0 {
		kc.Unknowns = m
	}
	return nil
}

func (kc *KubernetesCredentials) MarshalJSON() ([]byte, error) {
	m := map[string]json.RawMessage{}
	if len(kc.Unknowns) > 0 {
		for k, v := range kc.Unknowns {
			m[k] = v
		}
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(kc.EventAnalysisAndAlertingEnabled))
		if err != nil {
			return nil, err
		}
		m["eventAnalysisAndAlertingEnabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(kc.EventsIntegrationEnabled))
		if err != nil {
			return nil, err
		}
		m["eventsIntegrationEnabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(kc.HostnameVerificationEnabled))
		if err != nil {
			return nil, err
		}
		m["hostnameVerificationEnabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(kc.PrometheusExportersIntegrationEnabled))
		if err != nil {
			return nil, err
		}
		m["prometheusExportersIntegrationEnabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(kc.DavisEventsIntegrationEnabled))
		if err != nil {
			return nil, err
		}
		m["davisEventsIntegrationEnabled"] = rawMessage
	}
	{
		rawMessage, err := json.Marshal(opt.Bool(kc.Active))
		if err != nil {
			return nil, err
		}
		m["active"] = rawMessage
	}
	if kc.EndpointStatusInfo != nil {
		if rawMessage, err := json.Marshal(kc.EndpointStatusInfo); err == nil {
			m["endpointStatusInfo"] = rawMessage
		} else {
			return nil, err
		}
	}
	if rawMessage, err := json.Marshal(opt.Bool(kc.WorkloadIntegrationEnabled)); err == nil {
		m["workloadIntegrationEnabled"] = rawMessage
	} else {
		return nil, err
	}
	if kc.EndpointStatus != nil {
		rawMessage, err := json.Marshal(kc.EndpointStatus)
		if err != nil {
			return nil, err
		}
		m["endpointStatus"] = rawMessage
	}
	rawMessage, err := json.Marshal(kc.Label)
	if err != nil {
		return nil, err
	}
	m["label"] = rawMessage
	if kc.ID != nil {
		rawMessage, err := json.Marshal(kc.ID)
		if err != nil {
			return nil, err
		}
		m["id"] = rawMessage
	}
	if kc.Metadata != nil {
		rawMessage, err := json.Marshal(kc.Metadata)
		if err != nil {
			return nil, err
		}
		m["metadata"] = rawMessage
	}
	if kc.AuthToken != nil {
		rawMessage, err := json.Marshal(kc.AuthToken)
		if err != nil {
			return nil, err
		}
		m["authToken"] = rawMessage
	}
	if rawMessage, err := json.Marshal(opt.Bool(kc.CertificateCheckEnabled)); err == nil {
		m["certificateCheckEnabled"] = rawMessage
	} else {
		return nil, err
	}
	rawMessage, err = json.Marshal(kc.EndpointURL)
	if err != nil {
		return nil, err
	}
	m["endpointUrl"] = rawMessage
	if kc.EventsFieldSelectors != nil {
		rawMessage, err := json.Marshal(kc.EventsFieldSelectors)
		if err != nil {
			return nil, err
		}
		m["eventsFieldSelectors"] = rawMessage
	}
	return json.Marshal(m)
}
