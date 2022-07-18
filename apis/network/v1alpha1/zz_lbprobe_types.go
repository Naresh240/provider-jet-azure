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

type LBProbeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LoadBalancerRules []*string `json:"loadBalancerRules,omitempty" tf:"load_balancer_rules,omitempty"`
}

type LBProbeParameters struct {

	// +kubebuilder:validation:Optional
	IntervalInSeconds *float64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// +kubebuilder:validation:Required
	LoadbalancerID *string `json:"loadbalancerId" tf:"loadbalancer_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NumberOfProbes *float64 `json:"numberOfProbes,omitempty" tf:"number_of_probes,omitempty"`

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// LBProbeSpec defines the desired state of LBProbe
type LBProbeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBProbeParameters `json:"forProvider"`
}

// LBProbeStatus defines the observed state of LBProbe.
type LBProbeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBProbeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBProbe is the Schema for the LBProbes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type LBProbe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBProbeSpec   `json:"spec"`
	Status            LBProbeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBProbeList contains a list of LBProbes
type LBProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBProbe `json:"items"`
}

// Repository type metadata.
var (
	LBProbe_Kind             = "LBProbe"
	LBProbe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBProbe_Kind}.String()
	LBProbe_KindAPIVersion   = LBProbe_Kind + "." + CRDGroupVersion.String()
	LBProbe_GroupVersionKind = CRDGroupVersion.WithKind(LBProbe_Kind)
)

func init() {
	SchemeBuilder.Register(&LBProbe{}, &LBProbeList{})
}