/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Anomalies struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnomaliesSpec   `json:"spec,omitempty"`
	Status            AnomaliesStatus `json:"status,omitempty"`
}

type AnomaliesSpecFailureRateAuto struct {
	// Absolute increase of failing service calls to trigger an alert, %
	Absolute *int64 `json:"absolute" tf:"absolute"`
	// Relative increase of failing service calls to trigger an alert, %
	Relative *int64 `json:"relative" tf:"relative"`
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type AnomaliesSpecFailureRateThresholds struct {
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers alert
	Sensitivity *string `json:"sensitivity" tf:"sensitivity"`
	// Failure rate during any 5-minute period to trigger an alert, %
	Threshold *int64 `json:"threshold" tf:"threshold"`
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type AnomaliesSpecFailureRate struct {
	// Parameters of failure rate increase auto-detection. Example: If the expected error rate is 1.5%, and you set an absolute increase of 1%, and a relative increase of 50%, the thresholds will be:  Absolute: 1.5% + **1%** = 2.5%  Relative: 1.5% + 1.5% * **50%** = 2.25%
	// +optional
	Auto *AnomaliesSpecFailureRateAuto `json:"auto,omitempty" tf:"auto"`
	// Fixed thresholds for failure rate increase detection
	// +optional
	Thresholds *AnomaliesSpecFailureRateThresholds `json:"thresholds,omitempty" tf:"thresholds"`
}

type AnomaliesSpecResponseTimeAuto struct {
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	Load *string `json:"load" tf:"load"`
	// Alert if the response time degrades by more than *X* milliseconds
	Milliseconds *int64 `json:"milliseconds" tf:"milliseconds"`
	// Alert if the response time degrades by more than *X* %
	Percent *int64 `json:"percent" tf:"percent"`
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	SlowestMilliseconds *int64 `json:"slowestMilliseconds" tf:"slowest_milliseconds"`
	// Alert if the response time of the slowest 10% degrades by more than *X* milliseconds
	SlowestPercent *int64 `json:"slowestPercent" tf:"slowest_percent"`
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type AnomaliesSpecResponseTimeThresholds struct {
	// Minimal service load to detect response time degradation. Response time degradation of services with smaller load won't trigger alerts. Possible values are `FIFTEEN_REQUESTS_PER_MINUTE`, `FIVE_REQUESTS_PER_MINUTE`, `ONE_REQUEST_PER_MINUTE` and `TEN_REQUESTS_PER_MINUTE`
	Load *string `json:"load" tf:"load"`
	// Response time during any 5-minute period to trigger an alert, in milliseconds
	Milliseconds *int64 `json:"milliseconds" tf:"milliseconds"`
	// Sensitivity of the threshold.  With `low` sensitivity, high statistical confidence is used. Brief violations (for example, due to a surge in load) won't trigger alerts.  With `high` sensitivity, no statistical confidence is used. Each violation triggers an alert
	Sensitivity *string `json:"sensitivity" tf:"sensitivity"`
	// Response time of the 10% slowest during any 5-minute period to trigger an alert, in milliseconds
	SlowestMilliseconds *int64 `json:"slowestMilliseconds" tf:"slowest_milliseconds"`
	// allows for configuring properties that are not explicitly supported by the current version of this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type AnomaliesSpecResponseTime struct {
	// Parameters of the response time degradation auto-detection. Violation of **any** criterion triggers an alert
	// +optional
	Auto *AnomaliesSpecResponseTimeAuto `json:"auto,omitempty" tf:"auto"`
	// Fixed thresholds for response time degradation detection
	// +optional
	Thresholds *AnomaliesSpecResponseTimeThresholds `json:"thresholds,omitempty" tf:"thresholds"`
}

type AnomaliesSpecTrafficDrops struct {
	// The detection is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled" tf:"enabled"`
	// Alert if the observed traffic is less than *X* % of the expected value
	// +optional
	Percent *int64 `json:"percent,omitempty" tf:"percent"`
}

type AnomaliesSpecTrafficSpikes struct {
	// The detection is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled" tf:"enabled"`
	// Alert if the observed traffic is less than *X* % of the expected value
	// +optional
	Percent *int64 `json:"percent,omitempty" tf:"percent"`
}

type AnomaliesSpecTraffic struct {
	// The configuration of traffic drops detection
	// +optional
	Drops *AnomaliesSpecTrafficDrops `json:"drops,omitempty" tf:"drops"`
	// The configuration of traffic spikes detection
	// +optional
	Spikes *AnomaliesSpecTrafficSpikes `json:"spikes,omitempty" tf:"spikes"`
}

type AnomaliesSpec struct {
	State *AnomaliesSpecResource `json:"state,omitempty" tf:"-"`

	Resource AnomaliesSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AnomaliesSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration of failure rate increase detection
	// +optional
	FailureRate *AnomaliesSpecFailureRate `json:"failureRate,omitempty" tf:"failure_rate"`
	// Configuration of response time degradation detection
	// +optional
	ResponseTime *AnomaliesSpecResponseTime `json:"responseTime,omitempty" tf:"response_time"`
	// Configuration for anomalies regarding traffic
	// +optional
	Traffic *AnomaliesSpecTraffic `json:"traffic,omitempty" tf:"traffic"`
}

type AnomaliesStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AnomaliesList is a list of Anomaliess
type AnomaliesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Anomalies CRD objects
	Items []Anomalies `json:"items,omitempty"`
}
