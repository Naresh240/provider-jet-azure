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

type DNSSettingsObservation struct {
}

type DNSSettingsParameters struct {

	// +kubebuilder:validation:Required
	DNSServers []string `json:"dnsServers" tf:"dns_servers"`
}

type ExtensionObservation struct {
}

type ExtensionParameters struct {

	// +kubebuilder:validation:Optional
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	ProtectedSettings *string `json:"protectedSettings,omitempty" tf:"protected_settings"`

	// +kubebuilder:validation:Optional
	ProvisionAfterExtensions []string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions"`

	// +kubebuilder:validation:Required
	Publisher string `json:"publisher" tf:"publisher"`

	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Required
	TypeHandlerVersion string `json:"typeHandlerVersion" tf:"type_handler_version"`
}

type IPConfigurationObservation struct {
}

type IPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationGatewayBackendAddressPoolIds []string `json:"applicationGatewayBackendAddressPoolIds,omitempty" tf:"application_gateway_backend_address_pool_ids"`

	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIds []string `json:"applicationSecurityGroupIds,omitempty" tf:"application_security_group_ids"`

	// +kubebuilder:validation:Optional
	LoadBalancerBackendAddressPoolIds []string `json:"loadBalancerBackendAddressPoolIds,omitempty" tf:"load_balancer_backend_address_pool_ids"`

	// +kubebuilder:validation:Optional
	LoadBalancerInboundNatRulesIds []string `json:"loadBalancerInboundNatRulesIds,omitempty" tf:"load_balancer_inbound_nat_rules_ids"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Primary bool `json:"primary" tf:"primary"`

	// +kubebuilder:validation:Optional
	PublicIPAddressConfiguration []PublicIPAddressConfigurationParameters `json:"publicIpAddressConfiguration,omitempty" tf:"public_ip_address_configuration"`

	// +kubebuilder:validation:Required
	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

type NetworkProfileObservation struct {
}

type NetworkProfileParameters struct {

	// +kubebuilder:validation:Optional
	AcceleratedNetworking *bool `json:"acceleratedNetworking,omitempty" tf:"accelerated_networking"`

	// +kubebuilder:validation:Optional
	DNSSettings []DNSSettingsParameters `json:"dnsSettings,omitempty" tf:"dns_settings"`

	// +kubebuilder:validation:Required
	IPConfiguration []IPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration"`

	// +kubebuilder:validation:Optional
	IPForwarding *bool `json:"ipForwarding,omitempty" tf:"ip_forwarding"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id"`

	// +kubebuilder:validation:Required
	Primary bool `json:"primary" tf:"primary"`
}

type OsProfileLinuxConfigSSHKeysObservation struct {
}

type OsProfileLinuxConfigSSHKeysParameters struct {

	// +kubebuilder:validation:Optional
	KeyData *string `json:"keyData,omitempty" tf:"key_data"`

	// +kubebuilder:validation:Required
	Path string `json:"path" tf:"path"`
}

type OsProfileSecretsVaultCertificatesObservation struct {
}

type OsProfileSecretsVaultCertificatesParameters struct {

	// +kubebuilder:validation:Optional
	CertificateStore *string `json:"certificateStore,omitempty" tf:"certificate_store"`

	// +kubebuilder:validation:Required
	CertificateURL string `json:"certificateUrl" tf:"certificate_url"`
}

type OsProfileWindowsConfigAdditionalUnattendConfigObservation struct {
}

type OsProfileWindowsConfigAdditionalUnattendConfigParameters struct {

	// +kubebuilder:validation:Required
	Component string `json:"component" tf:"component"`

	// +kubebuilder:validation:Required
	Content string `json:"content" tf:"content"`

	// +kubebuilder:validation:Required
	Pass string `json:"pass" tf:"pass"`

	// +kubebuilder:validation:Required
	SettingName string `json:"settingName" tf:"setting_name"`
}

type OsProfileWindowsConfigWinrmObservation struct {
}

type OsProfileWindowsConfigWinrmParameters struct {

	// +kubebuilder:validation:Optional
	CertificateURL *string `json:"certificateUrl,omitempty" tf:"certificate_url"`

	// +kubebuilder:validation:Required
	Protocol string `json:"protocol" tf:"protocol"`
}

type PublicIPAddressConfigurationObservation struct {
}

type PublicIPAddressConfigurationParameters struct {

	// +kubebuilder:validation:Required
	DomainNameLabel string `json:"domainNameLabel" tf:"domain_name_label"`

	// +kubebuilder:validation:Required
	IdleTimeout int64 `json:"idleTimeout" tf:"idle_timeout"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type RollingUpgradePolicyObservation struct {
}

type RollingUpgradePolicyParameters struct {

	// +kubebuilder:validation:Optional
	MaxBatchInstancePercent *int64 `json:"maxBatchInstancePercent,omitempty" tf:"max_batch_instance_percent"`

	// +kubebuilder:validation:Optional
	MaxUnhealthyInstancePercent *int64 `json:"maxUnhealthyInstancePercent,omitempty" tf:"max_unhealthy_instance_percent"`

	// +kubebuilder:validation:Optional
	MaxUnhealthyUpgradedInstancePercent *int64 `json:"maxUnhealthyUpgradedInstancePercent,omitempty" tf:"max_unhealthy_upgraded_instance_percent"`

	// +kubebuilder:validation:Optional
	PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches,omitempty" tf:"pause_time_between_batches"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Required
	Capacity int64 `json:"capacity" tf:"capacity"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier"`
}

type StorageProfileDataDiskObservation struct {
}

type StorageProfileDataDiskParameters struct {

	// +kubebuilder:validation:Optional
	Caching *string `json:"caching,omitempty" tf:"caching"`

	// +kubebuilder:validation:Required
	CreateOption string `json:"createOption" tf:"create_option"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`

	// +kubebuilder:validation:Required
	Lun int64 `json:"lun" tf:"lun"`

	// +kubebuilder:validation:Optional
	ManagedDiskType *string `json:"managedDiskType,omitempty" tf:"managed_disk_type"`
}

type StorageProfileImageReferenceObservation struct {
}

type StorageProfileImageReferenceParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id"`

	// +kubebuilder:validation:Optional
	Offer *string `json:"offer,omitempty" tf:"offer"`

	// +kubebuilder:validation:Optional
	Publisher *string `json:"publisher,omitempty" tf:"publisher"`

	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type StorageProfileOsDiskObservation struct {
}

type StorageProfileOsDiskParameters struct {

	// +kubebuilder:validation:Optional
	Caching *string `json:"caching,omitempty" tf:"caching"`

	// +kubebuilder:validation:Required
	CreateOption string `json:"createOption" tf:"create_option"`

	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image"`

	// +kubebuilder:validation:Optional
	ManagedDiskType *string `json:"managedDiskType,omitempty" tf:"managed_disk_type"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type"`

	// +kubebuilder:validation:Optional
	VhdContainers []string `json:"vhdContainers,omitempty" tf:"vhd_containers"`
}

type VirtualMachineScaleSetBootDiagnosticsObservation struct {
}

type VirtualMachineScaleSetBootDiagnosticsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`

	// +kubebuilder:validation:Required
	StorageURI string `json:"storageUri" tf:"storage_uri"`
}

type VirtualMachineScaleSetIdentityObservation struct {
	PrincipalID string `json:"principalId,omitempty" tf:"principal_id"`
}

type VirtualMachineScaleSetIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type VirtualMachineScaleSetObservation struct {
}

type VirtualMachineScaleSetOsProfileLinuxConfigObservation struct {
}

type VirtualMachineScaleSetOsProfileLinuxConfigParameters struct {

	// +kubebuilder:validation:Optional
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" tf:"disable_password_authentication"`

	// +kubebuilder:validation:Optional
	SSHKeys []OsProfileLinuxConfigSSHKeysParameters `json:"sshKeys,omitempty" tf:"ssh_keys"`
}

type VirtualMachineScaleSetOsProfileObservation struct {
}

type VirtualMachineScaleSetOsProfileParameters struct {

	// +kubebuilder:validation:Optional
	AdminPassword *string `json:"adminPassword,omitempty" tf:"admin_password"`

	// +kubebuilder:validation:Required
	AdminUsername string `json:"adminUsername" tf:"admin_username"`

	// +kubebuilder:validation:Required
	ComputerNamePrefix string `json:"computerNamePrefix" tf:"computer_name_prefix"`

	// +kubebuilder:validation:Optional
	CustomData *string `json:"customData,omitempty" tf:"custom_data"`
}

type VirtualMachineScaleSetOsProfileSecretsObservation struct {
}

type VirtualMachineScaleSetOsProfileSecretsParameters struct {

	// +kubebuilder:validation:Required
	SourceVaultID string `json:"sourceVaultId" tf:"source_vault_id"`

	// +kubebuilder:validation:Optional
	VaultCertificates []OsProfileSecretsVaultCertificatesParameters `json:"vaultCertificates,omitempty" tf:"vault_certificates"`
}

type VirtualMachineScaleSetOsProfileWindowsConfigObservation struct {
}

type VirtualMachineScaleSetOsProfileWindowsConfigParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalUnattendConfig []OsProfileWindowsConfigAdditionalUnattendConfigParameters `json:"additionalUnattendConfig,omitempty" tf:"additional_unattend_config"`

	// +kubebuilder:validation:Optional
	EnableAutomaticUpgrades *bool `json:"enableAutomaticUpgrades,omitempty" tf:"enable_automatic_upgrades"`

	// +kubebuilder:validation:Optional
	ProvisionVMAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent"`

	// +kubebuilder:validation:Optional
	Winrm []OsProfileWindowsConfigWinrmParameters `json:"winrm,omitempty" tf:"winrm"`
}

type VirtualMachineScaleSetParameters struct {

	// +kubebuilder:validation:Optional
	AutomaticOsUpgrade *bool `json:"automaticOsUpgrade,omitempty" tf:"automatic_os_upgrade"`

	// +kubebuilder:validation:Optional
	BootDiagnostics []VirtualMachineScaleSetBootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics"`

	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`

	// +kubebuilder:validation:Optional
	Extension []ExtensionParameters `json:"extension,omitempty" tf:"extension"`

	// +kubebuilder:validation:Optional
	HealthProbeID *string `json:"healthProbeId,omitempty" tf:"health_probe_id"`

	// +kubebuilder:validation:Optional
	Identity []VirtualMachineScaleSetIdentityParameters `json:"identity,omitempty" tf:"identity"`

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type"`

	// +kubebuilder:validation:Required
	Location string `json:"location" tf:"location"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NetworkProfile []NetworkProfileParameters `json:"networkProfile" tf:"network_profile"`

	// +kubebuilder:validation:Required
	OsProfile []VirtualMachineScaleSetOsProfileParameters `json:"osProfile" tf:"os_profile"`

	// +kubebuilder:validation:Optional
	OsProfileLinuxConfig []VirtualMachineScaleSetOsProfileLinuxConfigParameters `json:"osProfileLinuxConfig,omitempty" tf:"os_profile_linux_config"`

	// +kubebuilder:validation:Optional
	OsProfileSecrets []VirtualMachineScaleSetOsProfileSecretsParameters `json:"osProfileSecrets,omitempty" tf:"os_profile_secrets"`

	// +kubebuilder:validation:Optional
	OsProfileWindowsConfig []VirtualMachineScaleSetOsProfileWindowsConfigParameters `json:"osProfileWindowsConfig,omitempty" tf:"os_profile_windows_config"`

	// +kubebuilder:validation:Optional
	Overprovision *bool `json:"overprovision,omitempty" tf:"overprovision"`

	// +kubebuilder:validation:Optional
	Plan []VirtualMachineScaleSetPlanParameters `json:"plan,omitempty" tf:"plan"`

	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority"`

	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id"`

	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	// +kubebuilder:validation:Optional
	RollingUpgradePolicy []RollingUpgradePolicyParameters `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy"`

	// +kubebuilder:validation:Optional
	SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty" tf:"single_placement_group"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku"`

	// +kubebuilder:validation:Optional
	StorageProfileDataDisk []StorageProfileDataDiskParameters `json:"storageProfileDataDisk,omitempty" tf:"storage_profile_data_disk"`

	// +kubebuilder:validation:Optional
	StorageProfileImageReference []StorageProfileImageReferenceParameters `json:"storageProfileImageReference,omitempty" tf:"storage_profile_image_reference"`

	// +kubebuilder:validation:Required
	StorageProfileOsDisk []StorageProfileOsDiskParameters `json:"storageProfileOsDisk" tf:"storage_profile_os_disk"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Required
	UpgradePolicyMode string `json:"upgradePolicyMode" tf:"upgrade_policy_mode"`

	// +kubebuilder:validation:Optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type VirtualMachineScaleSetPlanObservation struct {
}

type VirtualMachineScaleSetPlanParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Product string `json:"product" tf:"product"`

	// +kubebuilder:validation:Required
	Publisher string `json:"publisher" tf:"publisher"`
}

// VirtualMachineScaleSetSpec defines the desired state of VirtualMachineScaleSet
type VirtualMachineScaleSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualMachineScaleSetParameters `json:"forProvider"`
}

// VirtualMachineScaleSetStatus defines the observed state of VirtualMachineScaleSet.
type VirtualMachineScaleSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualMachineScaleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineScaleSet is the Schema for the VirtualMachineScaleSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineScaleSetSpec   `json:"spec"`
	Status            VirtualMachineScaleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineScaleSetList contains a list of VirtualMachineScaleSets
type VirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachineScaleSet `json:"items"`
}

// Repository type metadata.
var (
	VirtualMachineScaleSetKind             = "VirtualMachineScaleSet"
	VirtualMachineScaleSetGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualMachineScaleSetKind}.String()
	VirtualMachineScaleSetKindAPIVersion   = VirtualMachineScaleSetKind + "." + GroupVersion.String()
	VirtualMachineScaleSetGroupVersionKind = GroupVersion.WithKind(VirtualMachineScaleSetKind)
)

func init() {
	SchemeBuilder.Register(&VirtualMachineScaleSet{}, &VirtualMachineScaleSetList{})
}