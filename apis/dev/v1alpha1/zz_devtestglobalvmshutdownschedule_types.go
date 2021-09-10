/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type DevTestGlobalVmShutdownScheduleObservation struct {
}

type DevTestGlobalVmShutdownScheduleParameters struct {
	DailyRecurrenceTime string `json:"dailyRecurrenceTime" tf:"daily_recurrence_time"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	Location string `json:"location" tf:"location"`

	NotificationSettings []NotificationSettingsParameters `json:"notificationSettings" tf:"notification_settings"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Timezone string `json:"timezone" tf:"timezone"`

	VirtualMachineID string `json:"virtualMachineId" tf:"virtual_machine_id"`
}

type NotificationSettingsObservation struct {
}

type NotificationSettingsParameters struct {
	Email *string `json:"email,omitempty" tf:"email"`

	Enabled bool `json:"enabled" tf:"enabled"`

	TimeInMinutes *int64 `json:"timeInMinutes,omitempty" tf:"time_in_minutes"`

	WebhookURL *string `json:"webhookUrl,omitempty" tf:"webhook_url"`
}

// DevTestGlobalVmShutdownScheduleSpec defines the desired state of DevTestGlobalVmShutdownSchedule
type DevTestGlobalVmShutdownScheduleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DevTestGlobalVmShutdownScheduleParameters `json:"forProvider"`
}

// DevTestGlobalVmShutdownScheduleStatus defines the observed state of DevTestGlobalVmShutdownSchedule.
type DevTestGlobalVmShutdownScheduleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DevTestGlobalVmShutdownScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DevTestGlobalVmShutdownSchedule is the Schema for the DevTestGlobalVmShutdownSchedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DevTestGlobalVmShutdownSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestGlobalVmShutdownScheduleSpec   `json:"spec"`
	Status            DevTestGlobalVmShutdownScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevTestGlobalVmShutdownScheduleList contains a list of DevTestGlobalVmShutdownSchedules
type DevTestGlobalVmShutdownScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevTestGlobalVmShutdownSchedule `json:"items"`
}

// Repository type metadata.
var (
	DevTestGlobalVmShutdownScheduleKind             = "DevTestGlobalVmShutdownSchedule"
	DevTestGlobalVmShutdownScheduleGroupKind        = schema.GroupKind{Group: Group, Kind: DevTestGlobalVmShutdownScheduleKind}.String()
	DevTestGlobalVmShutdownScheduleKindAPIVersion   = DevTestGlobalVmShutdownScheduleKind + "." + GroupVersion.String()
	DevTestGlobalVmShutdownScheduleGroupVersionKind = GroupVersion.WithKind(DevTestGlobalVmShutdownScheduleKind)
)

func init() {
	SchemeBuilder.Register(&DevTestGlobalVmShutdownSchedule{}, &DevTestGlobalVmShutdownScheduleList{})
}
