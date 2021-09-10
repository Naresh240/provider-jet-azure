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

type AppServicePlanObservation struct {
	MaximumNumberOfWorkers int64 `json:"maximumNumberOfWorkers" tf:"maximum_number_of_workers"`
}

type AppServicePlanParameters struct {
	AppServiceEnvironmentID *string `json:"appServiceEnvironmentId,omitempty" tf:"app_service_environment_id"`

	IsXenon *bool `json:"isXenon,omitempty" tf:"is_xenon"`

	Kind *string `json:"kind,omitempty" tf:"kind"`

	Location string `json:"location" tf:"location"`

	MaximumElasticWorkerCount *int64 `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count"`

	Name string `json:"name" tf:"name"`

	PerSiteScaling *bool `json:"perSiteScaling,omitempty" tf:"per_site_scaling"`

	Reserved *bool `json:"reserved,omitempty" tf:"reserved"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Sku []SkuParameters `json:"sku" tf:"sku"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant"`
}

type SkuObservation struct {
}

type SkuParameters struct {
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`

	Size string `json:"size" tf:"size"`

	Tier string `json:"tier" tf:"tier"`
}

// AppServicePlanSpec defines the desired state of AppServicePlan
type AppServicePlanSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppServicePlanParameters `json:"forProvider"`
}

// AppServicePlanStatus defines the observed state of AppServicePlan.
type AppServicePlanStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppServicePlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppServicePlan is the Schema for the AppServicePlans API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppServicePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServicePlanSpec   `json:"spec"`
	Status            AppServicePlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServicePlanList contains a list of AppServicePlans
type AppServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServicePlan `json:"items"`
}

// Repository type metadata.
var (
	AppServicePlanKind             = "AppServicePlan"
	AppServicePlanGroupKind        = schema.GroupKind{Group: Group, Kind: AppServicePlanKind}.String()
	AppServicePlanKindAPIVersion   = AppServicePlanKind + "." + GroupVersion.String()
	AppServicePlanGroupVersionKind = GroupVersion.WithKind(AppServicePlanKind)
)

func init() {
	SchemeBuilder.Register(&AppServicePlan{}, &AppServicePlanList{})
}
