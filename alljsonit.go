/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
	alertingv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/alerting/v1alpha1"
	applicationv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/application/v1alpha1"
	autotagv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/autotag/v1alpha1"
	awsv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/aws/v1alpha1"
	azurev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/azure/v1alpha1"
	browserv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/browser/v1alpha1"
	calculatedv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/calculated/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/custom/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/dashboard/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/database/v1alpha1"
	diskv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/disk/v1alpha1"
	environmentv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/environment/v1alpha1"
	hostv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/host/v1alpha1"
	httpv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/http/v1alpha1"
	k8sv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/k8s/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/maintenance/v1alpha1"
	managementv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/management/v1alpha1"
	mobilev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/mobile/v1alpha1"
	notificationv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/notification/v1alpha1"
	processgroupv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/processgroup/v1alpha1"
	requestv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/request/v1alpha1"
	resourcev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/resource/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-dynatrace-api/apis/service/v1alpha1"
	slov1alpha1 "kubeform.dev/provider-dynatrace-api/apis/slo/v1alpha1"
	spanv1alpha1 "kubeform.dev/provider-dynatrace-api/apis/span/v1alpha1"
	"kubeform.dev/provider-dynatrace-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var allJsonIt = map[schema.GroupVersionResource]Data{
	{
		Group:    "alerting.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "profiles",
	}: {
		JsonIt:       controllers.GetJSONItr(alertingv1alpha1.GetEncoder(), alertingv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_alerting_profile",
	},
	{
		Group:    "application.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "anomalies",
	}: {
		JsonIt:       controllers.GetJSONItr(applicationv1alpha1.GetEncoder(), applicationv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_application_anomalies",
	},
	{
		Group:    "autotag.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "autotags",
	}: {
		JsonIt:       controllers.GetJSONItr(autotagv1alpha1.GetEncoder(), autotagv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_autotag",
	},
	{
		Group:    "aws.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "credentials",
	}: {
		JsonIt:       controllers.GetJSONItr(awsv1alpha1.GetEncoder(), awsv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_aws_credentials",
	},
	{
		Group:    "azure.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "credentials",
	}: {
		JsonIt:       controllers.GetJSONItr(azurev1alpha1.GetEncoder(), azurev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_azure_credentials",
	},
	{
		Group:    "browser.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "monitors",
	}: {
		JsonIt:       controllers.GetJSONItr(browserv1alpha1.GetEncoder(), browserv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_browser_monitor",
	},
	{
		Group:    "calculated.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "servicemetrics",
	}: {
		JsonIt:       controllers.GetJSONItr(calculatedv1alpha1.GetEncoder(), calculatedv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_calculated_service_metric",
	},
	{
		Group:    "custom.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "anomalies",
	}: {
		JsonIt:       controllers.GetJSONItr(customv1alpha1.GetEncoder(), customv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_custom_anomalies",
	},
	{
		Group:    "custom.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "services",
	}: {
		JsonIt:       controllers.GetJSONItr(customv1alpha1.GetEncoder(), customv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_custom_service",
	},
	{
		Group:    "dashboard.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "dashboards",
	}: {
		JsonIt:       controllers.GetJSONItr(dashboardv1alpha1.GetEncoder(), dashboardv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_dashboard",
	},
	{
		Group:    "dashboard.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "sharings",
	}: {
		JsonIt:       controllers.GetJSONItr(dashboardv1alpha1.GetEncoder(), dashboardv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_dashboard_sharing",
	},
	{
		Group:    "database.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "anomalies",
	}: {
		JsonIt:       controllers.GetJSONItr(databasev1alpha1.GetEncoder(), databasev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_database_anomalies",
	},
	{
		Group:    "disk.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "anomalies",
	}: {
		JsonIt:       controllers.GetJSONItr(diskv1alpha1.GetEncoder(), diskv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_disk_anomalies",
	},
	{
		Group:    "environment.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "environments",
	}: {
		JsonIt:       controllers.GetJSONItr(environmentv1alpha1.GetEncoder(), environmentv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_environment",
	},
	{
		Group:    "host.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "anomalies",
	}: {
		JsonIt:       controllers.GetJSONItr(hostv1alpha1.GetEncoder(), hostv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_host_anomalies",
	},
	{
		Group:    "host.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "namings",
	}: {
		JsonIt:       controllers.GetJSONItr(hostv1alpha1.GetEncoder(), hostv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_host_naming",
	},
	{
		Group:    "http.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "monitors",
	}: {
		JsonIt:       controllers.GetJSONItr(httpv1alpha1.GetEncoder(), httpv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_http_monitor",
	},
	{
		Group:    "k8s.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "credentials",
	}: {
		JsonIt:       controllers.GetJSONItr(k8sv1alpha1.GetEncoder(), k8sv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_k8s_credentials",
	},
	{
		Group:    "maintenance.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "windows",
	}: {
		JsonIt:       controllers.GetJSONItr(maintenancev1alpha1.GetEncoder(), maintenancev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_maintenance_window",
	},
	{
		Group:    "management.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "zones",
	}: {
		JsonIt:       controllers.GetJSONItr(managementv1alpha1.GetEncoder(), managementv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_management_zone",
	},
	{
		Group:    "mobile.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "applications",
	}: {
		JsonIt:       controllers.GetJSONItr(mobilev1alpha1.GetEncoder(), mobilev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_mobile_application",
	},
	{
		Group:    "notification.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "notifications",
	}: {
		JsonIt:       controllers.GetJSONItr(notificationv1alpha1.GetEncoder(), notificationv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_notification",
	},
	{
		Group:    "processgroup.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "namings",
	}: {
		JsonIt:       controllers.GetJSONItr(processgroupv1alpha1.GetEncoder(), processgroupv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_processgroup_naming",
	},
	{
		Group:    "request.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "attributes",
	}: {
		JsonIt:       controllers.GetJSONItr(requestv1alpha1.GetEncoder(), requestv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_request_attribute",
	},
	{
		Group:    "resource.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "attributes",
	}: {
		JsonIt:       controllers.GetJSONItr(resourcev1alpha1.GetEncoder(), resourcev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_resource_attributes",
	},
	{
		Group:    "service.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "anomalies",
	}: {
		JsonIt:       controllers.GetJSONItr(servicev1alpha1.GetEncoder(), servicev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_service_anomalies",
	},
	{
		Group:    "service.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "namings",
	}: {
		JsonIt:       controllers.GetJSONItr(servicev1alpha1.GetEncoder(), servicev1alpha1.GetDecoder()),
		ResourceType: "dynatrace_service_naming",
	},
	{
		Group:    "slo.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "sloes",
	}: {
		JsonIt:       controllers.GetJSONItr(slov1alpha1.GetEncoder(), slov1alpha1.GetDecoder()),
		ResourceType: "dynatrace_slo",
	},
	{
		Group:    "span.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "attributes",
	}: {
		JsonIt:       controllers.GetJSONItr(spanv1alpha1.GetEncoder(), spanv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_span_attribute",
	},
	{
		Group:    "span.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "capturerules",
	}: {
		JsonIt:       controllers.GetJSONItr(spanv1alpha1.GetEncoder(), spanv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_span_capture_rule",
	},
	{
		Group:    "span.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "contextpropagations",
	}: {
		JsonIt:       controllers.GetJSONItr(spanv1alpha1.GetEncoder(), spanv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_span_context_propagation",
	},
	{
		Group:    "span.dynatrace.kubeform.com",
		Version:  "v1alpha1",
		Resource: "entrypoints",
	}: {
		JsonIt:       controllers.GetJSONItr(spanv1alpha1.GetEncoder(), spanv1alpha1.GetDecoder()),
		ResourceType: "dynatrace_span_entry_point",
	},
}

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}
