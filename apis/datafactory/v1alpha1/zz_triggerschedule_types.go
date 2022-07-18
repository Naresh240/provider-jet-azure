/*
Copyright 2022 The Crossplane Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MonthlyObservation struct {
}

type MonthlyParameters struct {

	// +kubebuilder:validation:Optional
	Week *float64 `json:"week,omitempty" tf:"week,omitempty"`

	// +kubebuilder:validation:Required
	Weekday *string `json:"weekday" tf:"weekday,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// +kubebuilder:validation:Optional
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	Hours []*float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// +kubebuilder:validation:Optional
	Minutes []*float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// +kubebuilder:validation:Optional
	Monthly []MonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`
}

type TriggerScheduleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TriggerScheduleParameters struct {

	// +kubebuilder:validation:Optional
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PipelineName *string `json:"pipelineName" tf:"pipeline_name,omitempty"`

	// +kubebuilder:validation:Optional
	PipelineParameters map[string]*string `json:"pipelineParameters,omitempty" tf:"pipeline_parameters,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

// TriggerScheduleSpec defines the desired state of TriggerSchedule
type TriggerScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerScheduleParameters `json:"forProvider"`
}

// TriggerScheduleStatus defines the observed state of TriggerSchedule.
type TriggerScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerSchedule is the Schema for the TriggerSchedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TriggerSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerScheduleSpec   `json:"spec"`
	Status            TriggerScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerScheduleList contains a list of TriggerSchedules
type TriggerScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerSchedule `json:"items"`
}

// Repository type metadata.
var (
	TriggerSchedule_Kind             = "TriggerSchedule"
	TriggerSchedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerSchedule_Kind}.String()
	TriggerSchedule_KindAPIVersion   = TriggerSchedule_Kind + "." + CRDGroupVersion.String()
	TriggerSchedule_GroupVersionKind = CRDGroupVersion.WithKind(TriggerSchedule_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerSchedule{}, &TriggerScheduleList{})
}