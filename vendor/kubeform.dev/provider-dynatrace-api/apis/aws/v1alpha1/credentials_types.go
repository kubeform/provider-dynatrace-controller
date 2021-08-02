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

type Credentials struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialsSpec   `json:"spec,omitempty"`
	Status            CredentialsStatus `json:"status,omitempty"`
}

type CredentialsSpecAuthenticationData struct {
	// the access key
	// +optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key"`
	// the ID of the Amazon account
	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// the external ID token for setting an IAM role. You can obtain it with the `GET /aws/iamExternalId` request
	// +optional
	ExternalID *string `json:"externalID,omitempty" tf:"external_id"`
	// the IAM role to be used by Dynatrace to get monitoring data
	// +optional
	IamRole *string `json:"iamRole,omitempty" tf:"iam_role"`
	// the secret access key
	// +optional
	SecretKey *string `json:"secretKey,omitempty" tf:"secret_key"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type CredentialsSpecSupportingServicesToMonitorMonitoredMetrics struct {
	// a list of metric's dimensions names
	// +optional
	Dimensions []string `json:"dimensions,omitempty" tf:"dimensions"`
	// the name of the metric of the supporting service
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// the statistic (aggregation) to be used for the metric. AVG_MIN_MAX value is 3 statistics at once: AVERAGE, MINIMUM and MAXIMUM
	// +optional
	Statistic *string `json:"statistic,omitempty" tf:"statistic"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type CredentialsSpecSupportingServicesToMonitor struct {
	// a list of metrics to be monitored for this service
	// +optional
	// +kubebuilder:validation:MaxItems=10
	MonitoredMetrics []CredentialsSpecSupportingServicesToMonitorMonitoredMetrics `json:"monitoredMetrics,omitempty" tf:"monitored_metrics"`
	// the name of the supporting service
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type CredentialsSpecTagsToMonitor struct {
	// the key of the AWS tag.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
	// the value of the AWS tag
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type CredentialsSpec struct {
	State *CredentialsSpecResource `json:"state,omitempty" tf:"-"`

	Resource CredentialsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CredentialsSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// credentials for the AWS authentication
	AuthenticationData *CredentialsSpecAuthenticationData `json:"authenticationData" tf:"authentication_data"`
	// The name of the credentials
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// The type of the AWS partition
	PartitionType *string `json:"partitionType" tf:"partition_type"`
	// supporting services to be monitored
	// +optional
	// +kubebuilder:validation:MaxItems=10
	SupportingServicesToMonitor []CredentialsSpecSupportingServicesToMonitor `json:"supportingServicesToMonitor,omitempty" tf:"supporting_services_to_monitor"`
	// Monitor only resources which have specified AWS tags (`true`) or all resources (`false`)
	TaggedOnly *bool `json:"taggedOnly" tf:"tagged_only"`
	// AWS tags to be monitored. You can specify up to 10 tags. Only applicable when the **tagged_only** parameter is set to `true`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	TagsToMonitor []CredentialsSpecTagsToMonitor `json:"tagsToMonitor,omitempty" tf:"tags_to_monitor"`
	// Any attributes that aren't yet supported by this provider
	// +optional
	Unknowns *string `json:"unknowns,omitempty" tf:"unknowns"`
}

type CredentialsStatus struct {
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

// CredentialsList is a list of Credentialss
type CredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Credentials CRD objects
	Items []Credentials `json:"items,omitempty"`
}
