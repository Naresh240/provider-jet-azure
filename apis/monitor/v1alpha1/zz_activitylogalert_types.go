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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Required
	ActionGroupID *string `json:"actionGroupId" tf:"action_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	WebhookProperties map[string]*string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type ActivityLogAlertObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ActivityLogAlertParameters struct {

	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Criteria []CriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Scopes []*string `json:"scopes" tf:"scopes,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CriteriaObservation struct {
}

type CriteriaParameters struct {

	// +kubebuilder:validation:Optional
	Caller *string `json:"caller,omitempty" tf:"caller,omitempty"`

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// +kubebuilder:validation:Optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name,omitempty"`

	// +kubebuilder:validation:Optional
	RecommendationCategory *string `json:"recommendationCategory,omitempty" tf:"recommendation_category,omitempty"`

	// +kubebuilder:validation:Optional
	RecommendationImpact *string `json:"recommendationImpact,omitempty" tf:"recommendation_impact,omitempty"`

	// +kubebuilder:validation:Optional
	RecommendationType *string `json:"recommendationType,omitempty" tf:"recommendation_type,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceProvider *string `json:"resourceProvider,omitempty" tf:"resource_provider,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceHealth []ServiceHealthParameters `json:"serviceHealth,omitempty" tf:"service_health,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	SubStatus *string `json:"subStatus,omitempty" tf:"sub_status,omitempty"`
}

type ServiceHealthObservation struct {
}

type ServiceHealthParameters struct {

	// +kubebuilder:validation:Optional
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// +kubebuilder:validation:Optional
	Locations []*string `json:"locations,omitempty" tf:"locations,omitempty"`

	// +kubebuilder:validation:Optional
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

// ActivityLogAlertSpec defines the desired state of ActivityLogAlert
type ActivityLogAlertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActivityLogAlertParameters `json:"forProvider"`
}

// ActivityLogAlertStatus defines the observed state of ActivityLogAlert.
type ActivityLogAlertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActivityLogAlertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActivityLogAlert is the Schema for the ActivityLogAlerts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActivityLogAlertSpec   `json:"spec"`
	Status            ActivityLogAlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActivityLogAlertList contains a list of ActivityLogAlerts
type ActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActivityLogAlert `json:"items"`
}

// Repository type metadata.
var (
	ActivityLogAlert_Kind             = "ActivityLogAlert"
	ActivityLogAlert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActivityLogAlert_Kind}.String()
	ActivityLogAlert_KindAPIVersion   = ActivityLogAlert_Kind + "." + CRDGroupVersion.String()
	ActivityLogAlert_GroupVersionKind = CRDGroupVersion.WithKind(ActivityLogAlert_Kind)
)

func init() {
	SchemeBuilder.Register(&ActivityLogAlert{}, &ActivityLogAlertList{})
}