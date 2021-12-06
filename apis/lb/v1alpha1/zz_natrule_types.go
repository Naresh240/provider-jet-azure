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

type NatRuleObservation struct {
	BackendIPConfigurationID *string `json:"backendIpConfigurationId,omitempty" tf:"backend_ip_configuration_id,omitempty"`

	FrontendIPConfigurationID *string `json:"frontendIpConfigurationId,omitempty" tf:"frontend_ip_configuration_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NatRuleParameters struct {

	// +kubebuilder:validation:Required
	BackendPort *int64 `json:"backendPort" tf:"backend_port,omitempty"`

	// +kubebuilder:validation:Optional
	EnableFloatingIP *bool `json:"enableFloatingIp,omitempty" tf:"enable_floating_ip,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// +kubebuilder:validation:Required
	FrontendIPConfigurationName *string `json:"frontendIpConfigurationName" tf:"frontend_ip_configuration_name,omitempty"`

	// +kubebuilder:validation:Required
	FrontendPort *int64 `json:"frontendPort" tf:"frontend_port,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	LoadbalancerID *string `json:"loadbalancerId" tf:"loadbalancer_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// NatRuleSpec defines the desired state of NatRule
type NatRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NatRuleParameters `json:"forProvider"`
}

// NatRuleStatus defines the observed state of NatRule.
type NatRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NatRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NatRule is the Schema for the NatRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type NatRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NatRuleSpec   `json:"spec"`
	Status            NatRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NatRuleList contains a list of NatRules
type NatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NatRule `json:"items"`
}

// Repository type metadata.
var (
	NatRule_Kind             = "NatRule"
	NatRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NatRule_Kind}.String()
	NatRule_KindAPIVersion   = NatRule_Kind + "." + CRDGroupVersion.String()
	NatRule_GroupVersionKind = CRDGroupVersion.WithKind(NatRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NatRule{}, &NatRuleList{})
}