/*
Copyright 2019 The Crossplane Authors.

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

package v1alpha3

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// DefaultNodeCount is the default node count for a cluster.
	DefaultNodeCount = 1
)

// AKSClusterParameters define the desired state of an Azure Kubernetes Engine
// cluster.
type AKSClusterParameters struct {
	// ResourceGroupName is the name of the resource group that the cluster will
	// be created in
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to a ResourceGroup to retrieve its
	// name
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to a ResourceGroup to
	// retrieve its name
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// Location is the Azure location that the cluster will be created in
	Location string `json:"location"`

	// Version is the Kubernetes version that will be deployed to the cluster
	Version string `json:"version"`

	// VnetSubnetID is the subnet to which the cluster will be deployed.
	// +optional
	VnetSubnetID string `json:"vnetSubnetID,omitempty"`

	// ResourceGroupNameRef - A reference to a Subnet to retrieve its ID
	VnetSubnetIDRef *runtimev1alpha1.Reference `json:"vnetSubnetIDRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to a Subnet to retrieve
	// its ID
	VnetSubnetIDSelector *runtimev1alpha1.Selector `json:"vnetSubnetIDSelector,omitempty"`

	// NodeCount is the number of nodes that the cluster will initially be
	// created with.  This can be scaled over time and defaults to 1.
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=0
	// +optional
	NodeCount *int `json:"nodeCount,omitempty"`

	// NodeVMSize is the name of the worker node VM size, e.g., Standard_B2s,
	// Standard_F2s_v2, etc.
	// +optional
	NodeVMSize string `json:"nodeVMSize"`

	// DNSNamePrefix is the DNS name prefix to use with the hosted Kubernetes
	// API server FQDN. You will use this to connect to the Kubernetes API when
	// managing containers after creating the cluster.
	// +optional
	DNSNamePrefix string `json:"dnsNamePrefix"`

	// DisableRBAC determines whether RBAC will be disabled or enabled in the
	// cluster.
	// +optional
	DisableRBAC bool `json:"disableRBAC,omitempty"`
}

// An AKSClusterSpec defines the desired state of a AKSCluster.
type AKSClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	AKSClusterParameters         `json:",inline"`
}

// An AKSClusterStatus represents the observed state of an AKSCluster.
type AKSClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State is the current state of the cluster.
	State string `json:"state,omitempty"`

	// ProviderID is the external ID to identify this resource in the cloud
	// provider.
	ProviderID string `json:"providerID,omitempty"`

	// Endpoint is the endpoint where the cluster can be reached
	Endpoint string `json:"endpoint"`
}

// +kubebuilder:object:root=true

// An AKSCluster is a managed resource that represents an Azure Kubernetes
// Engine cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="ENDPOINT",type="string",JSONPath=".status.endpoint"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
// +kubebuilder:subresource:status
type AKSCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AKSClusterSpec   `json:"spec"`
	Status AKSClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AKSClusterList contains a list of AKSCluster.
type AKSClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AKSCluster `json:"items"`
}

type VirtualMachineProperties struct {
	// HardwareProfile - Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfile `json:"hardwareProfile,omitempty"`
	// StorageProfile - Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfile `json:"storageProfile,omitempty"`
	// AdditionalCapabilities - Specifies additional capabilities enabled or disabled on the virtual machine.
	//AdditionalCapabilities *AdditionalCapabilities `json:"additionalCapabilities,omitempty"`
	// OsProfile - Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile *OSProfile `json:"osProfile,omitempty"`
	// NetworkProfile - Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	// DiagnosticsProfile - Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	//DiagnosticsProfile *DiagnosticsProfile `json:"diagnosticsProfile,omitempty"`
	// AvailabilitySet - Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	//AvailabilitySet *SubResource `json:"availabilitySet,omitempty"`
	// VirtualMachineScaleSet - Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	//VirtualMachineScaleSet *SubResource `json:"virtualMachineScaleSet,omitempty"`
	// ProximityPlacementGroup - Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	//ProximityPlacementGroup *SubResource `json:"proximityPlacementGroup,omitempty"`
	// Priority - Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01. Possible values include: 'Regular', 'Low', 'Spot'
	//Priority VirtualMachinePriorityTypes `json:"priority,omitempty"`
	// EvictionPolicy - Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview. Possible values include: 'Deallocate', 'Delete'
	//EvictionPolicy VirtualMachineEvictionPolicyTypes `json:"evictionPolicy,omitempty"`
	// BillingProfile - Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	//BillingProfile *BillingProfile `json:"billingProfile,omitempty"`
	// Host - Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	//Host *SubResource `json:"host,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state, which only appears in the response.
	//ProvisioningState *string `json:"provisioningState,omitempty"`
	// InstanceView - READ-ONLY; The virtual machine instance view.
	//InstanceView *VirtualMachineInstanceView `json:"instanceView,omitempty"`
	// LicenseType - Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
	//LicenseType *string `json:"licenseType,omitempty"`
	// VMID - READ-ONLY; Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
	//VMID *string `json:"vmId,omitempty"`
}

// HardwareProfile specifies the hardware settings for the virtual machine.
type HardwareProfile struct {
	// VMSize - Specifies the size of the virtual machine. For more information about virtual machine sizes, see [Sizes for virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-sizes?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> The available VM sizes depend on region and availability set. For a list of available sizes use these APIs:  <br><br> [List all available virtual machine sizes in an availability set](https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes) <br><br> [List all available virtual machine sizes in a region](https://docs.microsoft.com/rest/api/compute/virtualmachinesizes/list) <br><br> [List all available virtual machine sizes for resizing](https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes). Possible values include: 'VirtualMachineSizeTypesBasicA0', 'VirtualMachineSizeTypesBasicA1', 'VirtualMachineSizeTypesBasicA2', 'VirtualMachineSizeTypesBasicA3', 'VirtualMachineSizeTypesBasicA4', 'VirtualMachineSizeTypesStandardA0', 'VirtualMachineSizeTypesStandardA1', 'VirtualMachineSizeTypesStandardA2', 'VirtualMachineSizeTypesStandardA3', 'VirtualMachineSizeTypesStandardA4', 'VirtualMachineSizeTypesStandardA5', 'VirtualMachineSizeTypesStandardA6', 'VirtualMachineSizeTypesStandardA7', 'VirtualMachineSizeTypesStandardA8', 'VirtualMachineSizeTypesStandardA9', 'VirtualMachineSizeTypesStandardA10', 'VirtualMachineSizeTypesStandardA11', 'VirtualMachineSizeTypesStandardA1V2', 'VirtualMachineSizeTypesStandardA2V2', 'VirtualMachineSizeTypesStandardA4V2', 'VirtualMachineSizeTypesStandardA8V2', 'VirtualMachineSizeTypesStandardA2mV2', 'VirtualMachineSizeTypesStandardA4mV2', 'VirtualMachineSizeTypesStandardA8mV2', 'VirtualMachineSizeTypesStandardB1s', 'VirtualMachineSizeTypesStandardB1ms', 'VirtualMachineSizeTypesStandardB2s', 'VirtualMachineSizeTypesStandardB2ms', 'VirtualMachineSizeTypesStandardB4ms', 'VirtualMachineSizeTypesStandardB8ms', 'VirtualMachineSizeTypesStandardD1', 'VirtualMachineSizeTypesStandardD2', 'VirtualMachineSizeTypesStandardD3', 'VirtualMachineSizeTypesStandardD4', 'VirtualMachineSizeTypesStandardD11', 'VirtualMachineSizeTypesStandardD12', 'VirtualMachineSizeTypesStandardD13', 'VirtualMachineSizeTypesStandardD14', 'VirtualMachineSizeTypesStandardD1V2', 'VirtualMachineSizeTypesStandardD2V2', 'VirtualMachineSizeTypesStandardD3V2', 'VirtualMachineSizeTypesStandardD4V2', 'VirtualMachineSizeTypesStandardD5V2', 'VirtualMachineSizeTypesStandardD2V3', 'VirtualMachineSizeTypesStandardD4V3', 'VirtualMachineSizeTypesStandardD8V3', 'VirtualMachineSizeTypesStandardD16V3', 'VirtualMachineSizeTypesStandardD32V3', 'VirtualMachineSizeTypesStandardD64V3', 'VirtualMachineSizeTypesStandardD2sV3', 'VirtualMachineSizeTypesStandardD4sV3', 'VirtualMachineSizeTypesStandardD8sV3', 'VirtualMachineSizeTypesStandardD16sV3', 'VirtualMachineSizeTypesStandardD32sV3', 'VirtualMachineSizeTypesStandardD64sV3', 'VirtualMachineSizeTypesStandardD11V2', 'VirtualMachineSizeTypesStandardD12V2', 'VirtualMachineSizeTypesStandardD13V2', 'VirtualMachineSizeTypesStandardD14V2', 'VirtualMachineSizeTypesStandardD15V2', 'VirtualMachineSizeTypesStandardDS1', 'VirtualMachineSizeTypesStandardDS2', 'VirtualMachineSizeTypesStandardDS3', 'VirtualMachineSizeTypesStandardDS4', 'VirtualMachineSizeTypesStandardDS11', 'VirtualMachineSizeTypesStandardDS12', 'VirtualMachineSizeTypesStandardDS13', 'VirtualMachineSizeTypesStandardDS14', 'VirtualMachineSizeTypesStandardDS1V2', 'VirtualMachineSizeTypesStandardDS2V2', 'VirtualMachineSizeTypesStandardDS3V2', 'VirtualMachineSizeTypesStandardDS4V2', 'VirtualMachineSizeTypesStandardDS5V2', 'VirtualMachineSizeTypesStandardDS11V2', 'VirtualMachineSizeTypesStandardDS12V2', 'VirtualMachineSizeTypesStandardDS13V2', 'VirtualMachineSizeTypesStandardDS14V2', 'VirtualMachineSizeTypesStandardDS15V2', 'VirtualMachineSizeTypesStandardDS134V2', 'VirtualMachineSizeTypesStandardDS132V2', 'VirtualMachineSizeTypesStandardDS148V2', 'VirtualMachineSizeTypesStandardDS144V2', 'VirtualMachineSizeTypesStandardE2V3', 'VirtualMachineSizeTypesStandardE4V3', 'VirtualMachineSizeTypesStandardE8V3', 'VirtualMachineSizeTypesStandardE16V3', 'VirtualMachineSizeTypesStandardE32V3', 'VirtualMachineSizeTypesStandardE64V3', 'VirtualMachineSizeTypesStandardE2sV3', 'VirtualMachineSizeTypesStandardE4sV3', 'VirtualMachineSizeTypesStandardE8sV3', 'VirtualMachineSizeTypesStandardE16sV3', 'VirtualMachineSizeTypesStandardE32sV3', 'VirtualMachineSizeTypesStandardE64sV3', 'VirtualMachineSizeTypesStandardE3216V3', 'VirtualMachineSizeTypesStandardE328sV3', 'VirtualMachineSizeTypesStandardE6432sV3', 'VirtualMachineSizeTypesStandardE6416sV3', 'VirtualMachineSizeTypesStandardF1', 'VirtualMachineSizeTypesStandardF2', 'VirtualMachineSizeTypesStandardF4', 'VirtualMachineSizeTypesStandardF8', 'VirtualMachineSizeTypesStandardF16', 'VirtualMachineSizeTypesStandardF1s', 'VirtualMachineSizeTypesStandardF2s', 'VirtualMachineSizeTypesStandardF4s', 'VirtualMachineSizeTypesStandardF8s', 'VirtualMachineSizeTypesStandardF16s', 'VirtualMachineSizeTypesStandardF2sV2', 'VirtualMachineSizeTypesStandardF4sV2', 'VirtualMachineSizeTypesStandardF8sV2', 'VirtualMachineSizeTypesStandardF16sV2', 'VirtualMachineSizeTypesStandardF32sV2', 'VirtualMachineSizeTypesStandardF64sV2', 'VirtualMachineSizeTypesStandardF72sV2', 'VirtualMachineSizeTypesStandardG1', 'VirtualMachineSizeTypesStandardG2', 'VirtualMachineSizeTypesStandardG3', 'VirtualMachineSizeTypesStandardG4', 'VirtualMachineSizeTypesStandardG5', 'VirtualMachineSizeTypesStandardGS1', 'VirtualMachineSizeTypesStandardGS2', 'VirtualMachineSizeTypesStandardGS3', 'VirtualMachineSizeTypesStandardGS4', 'VirtualMachineSizeTypesStandardGS5', 'VirtualMachineSizeTypesStandardGS48', 'VirtualMachineSizeTypesStandardGS44', 'VirtualMachineSizeTypesStandardGS516', 'VirtualMachineSizeTypesStandardGS58', 'VirtualMachineSizeTypesStandardH8', 'VirtualMachineSizeTypesStandardH16', 'VirtualMachineSizeTypesStandardH8m', 'VirtualMachineSizeTypesStandardH16m', 'VirtualMachineSizeTypesStandardH16r', 'VirtualMachineSizeTypesStandardH16mr', 'VirtualMachineSizeTypesStandardL4s', 'VirtualMachineSizeTypesStandardL8s', 'VirtualMachineSizeTypesStandardL16s', 'VirtualMachineSizeTypesStandardL32s', 'VirtualMachineSizeTypesStandardM64s', 'VirtualMachineSizeTypesStandardM64ms', 'VirtualMachineSizeTypesStandardM128s', 'VirtualMachineSizeTypesStandardM128ms', 'VirtualMachineSizeTypesStandardM6432ms', 'VirtualMachineSizeTypesStandardM6416ms', 'VirtualMachineSizeTypesStandardM12864ms', 'VirtualMachineSizeTypesStandardM12832ms', 'VirtualMachineSizeTypesStandardNC6', 'VirtualMachineSizeTypesStandardNC12', 'VirtualMachineSizeTypesStandardNC24', 'VirtualMachineSizeTypesStandardNC24r', 'VirtualMachineSizeTypesStandardNC6sV2', 'VirtualMachineSizeTypesStandardNC12sV2', 'VirtualMachineSizeTypesStandardNC24sV2', 'VirtualMachineSizeTypesStandardNC24rsV2', 'VirtualMachineSizeTypesStandardNC6sV3', 'VirtualMachineSizeTypesStandardNC12sV3', 'VirtualMachineSizeTypesStandardNC24sV3', 'VirtualMachineSizeTypesStandardNC24rsV3', 'VirtualMachineSizeTypesStandardND6s', 'VirtualMachineSizeTypesStandardND12s', 'VirtualMachineSizeTypesStandardND24s', 'VirtualMachineSizeTypesStandardND24rs', 'VirtualMachineSizeTypesStandardNV6', 'VirtualMachineSizeTypesStandardNV12', 'VirtualMachineSizeTypesStandardNV24'
	VMSize VirtualMachineSizeTypes `json:"vmSize,omitempty"`
}

type StorageProfile struct {
	// ImageReference - Specifies information about the image to use. You can specify information about platform images, marketplace images, or virtual machine images. This element is required when you want to use a platform image, marketplace image, or virtual machine image, but is not used in other creation operations.
	ImageReference *ImageReference `json:"imageReference,omitempty"`
	// OsDisk - Specifies information about the operating system disk used by the virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	OsDisk *OSDisk `json:"osDisk,omitempty"`
	// DataDisks - Specifies the parameters that are used to add a data disk to a virtual machine. <br><br> For more information about disks, see [About disks and VHDs for Azure virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
	// DataDisks *[]DataDisk `json:"dataDisks,omitempty"
}

// OSProfile specifies the operating system settings for the virtual machine. Some of the settings cannot
// be changed once VM is provisioned.
type OSProfile struct {
	// ComputerName - Specifies the host OS name of the virtual machine. <br><br> This name cannot be updated after the VM is created. <br><br> **Max-length (Windows):** 15 characters <br><br> **Max-length (Linux):** 64 characters. <br><br> For naming conventions and restrictions see [Azure infrastructure services implementation guidelines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-infrastructure-subscription-accounts-guidelines?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#1-naming-conventions).
	ComputerName *string `json:"computerName,omitempty"`
	// AdminUsername - Specifies the name of the administrator account. <br><br> This property cannot be updated after the VM is created. <br><br> **Windows-only restriction:** Cannot end in "." <br><br> **Disallowed values:** "administrator", "admin", "user", "user1", "test", "user2", "test1", "user3", "admin1", "1", "123", "a", "actuser", "adm", "admin2", "aspnet", "backup", "console", "david", "guest", "john", "owner", "root", "server", "sql", "support", "support_388945a0", "sys", "test2", "test3", "user4", "user5". <br><br> **Minimum-length (Linux):** 1  character <br><br> **Max-length (Linux):** 64 characters <br><br> **Max-length (Windows):** 20 characters  <br><br><li> For root access to the Linux VM, see [Using root privileges on Linux virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-use-root-privileges?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)<br><li> For a list of built-in system users on Linux that should not be used in this field, see [Selecting User Names for Linux on Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-usernames?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	AdminUsername *string `json:"adminUsername,omitempty"`
	// AdminPassword - Specifies the password of the administrator account. <br><br> **Minimum-length (Windows):** 8 characters <br><br> **Minimum-length (Linux):** 6 characters <br><br> **Max-length (Windows):** 123 characters <br><br> **Max-length (Linux):** 72 characters <br><br> **Complexity requirements:** 3 out of 4 conditions below need to be fulfilled <br> Has lower characters <br>Has upper characters <br> Has a digit <br> Has a special character (Regex match [\W_]) <br><br> **Disallowed values:** "abc@123", "P@$$w0rd", "P@ssw0rd", "P@ssword123", "Pa$$word", "pass@word1", "Password!", "Password1", "Password22", "iloveyou!" <br><br> For resetting the password, see [How to reset the Remote Desktop service or its login password in a Windows VM](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-reset-rdp?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> For resetting root password, see [Manage users, SSH, and check or repair disks on Azure Linux VMs using the VMAccess Extension](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-vmaccess-extension?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json#reset-root-password)
	//AdminPassword *string `json:"adminPassword,omitempty"`
	// CustomData - Specifies a base-64 encoded string of custom data. The base-64 encoded string is decoded to a binary array that is saved as a file on the Virtual Machine. The maximum length of the binary array is 65535 bytes. <br><br> **Note: Do not pass any secrets or passwords in customData property** <br><br> This property cannot be updated after the VM is created. <br><br> customData is passed to the VM to be saved as a file, for more information see [Custom Data on Azure VMs](https://azure.microsoft.com/en-us/blog/custom-data-and-cloud-init-on-windows-azure/) <br><br> For using cloud-init for your Linux VM, see [Using cloud-init to customize a Linux VM during creation](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-using-cloud-init?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
	//CustomData *string `json:"customData,omitempty"`
	// WindowsConfiguration - Specifies Windows operating system settings on the virtual machine.
	WindowsConfiguration *WindowsConfiguration `json:"windowsConfiguration,omitempty"`
	// LinuxConfiguration - Specifies the Linux operating system settings on the virtual machine. <br><br>For a list of supported Linux distributions, see [Linux on Azure-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json) <br><br> For running non-endorsed distributions, see [Information for Non-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
	LinuxConfiguration *LinuxConfiguration `json:"linuxConfiguration,omitempty"`
	// Secrets - Specifies set of certificates that should be installed onto the virtual machine.
	//Secrets *[]VaultSecretGroup `json:"secrets,omitempty"`
	// AllowExtensionOperations - Specifies whether extension operations should be allowed on the virtual machine. <br><br>This may only be set to False when no extensions are present on the virtual machine.
	//AllowExtensionOperations *bool `json:"allowExtensionOperations,omitempty"`
	// RequireGuestProvisionSignal - Specifies whether the guest provision signal is required to infer provision success of the virtual machine.
	//RequireGuestProvisionSignal *bool `json:"requireGuestProvisionSignal,omitempty"`
}

// NetworkProfile specifies the network interfaces of the virtual machine.
type NetworkProfile struct {
	// NetworkInterfaces - Specifies the list of resource Ids for the network interfaces associated with the virtual machine.
	NetworkInterfaces *[]NetworkInterfaceReference `json:"networkInterfaces,omitempty"`
}

type VirtualMachineSizeTypes string

// ImageReference specifies information about the image to use. You can specify information about platform
// images, marketplace images, or virtual machine images. This element is required when you want to use a
// platform image, marketplace image, or virtual machine image, but is not used in other creation
// operations. NOTE: Image reference publisher and offer can only be set when you create the scale set.
type ImageReference struct {
	// Publisher - The image publisher.
	Publisher *string `json:"publisher,omitempty"`
	// Offer - Specifies the offer of the platform image or marketplace image used to create the virtual machine.
	Offer *string `json:"offer,omitempty"`
	// Sku - The image SKU.
	Sku *string `json:"sku,omitempty"`
	// Version - Specifies the version of the platform image or marketplace image used to create the virtual machine. The allowed formats are Major.Minor.Build or 'latest'. Major, Minor, and Build are decimal numbers. Specify 'latest' to use the latest version of an image available at deploy time. Even if you use 'latest', the VM image will not automatically update after deploy time even if a new version becomes available.
	Version *string `json:"version,omitempty"`
	// ExactVersion - READ-ONLY; Specifies in decimal numbers, the version of platform image or marketplace image used to create the virtual machine. This readonly field differs from 'version', only if the value specified in 'version' field is 'latest'.
	//ExactVersion *string `json:"exactVersion,omitempty"`
	// ID - Resource Id
	//ID *string `json:"id,omitempty"`
}

type CachingTypes string
type DiskCreateOptionTypes string
type StorageAccountTypes string

// ManagedDiskParameters the parameters of a managed disk.
type ManagedDiskParameters struct {
	// StorageAccountType - Specifies the storage account type for the managed disk. NOTE: UltraSSD_LRS can only be used with data disks, it cannot be used with OS Disk. Possible values include: 'StorageAccountTypesStandardLRS', 'StorageAccountTypesPremiumLRS', 'StorageAccountTypesStandardSSDLRS', 'StorageAccountTypesUltraSSDLRS'
	StorageAccountType StorageAccountTypes `json:"storageAccountType,omitempty"`
	// DiskEncryptionSet - Specifies the customer managed disk encryption set resource id for the managed disk.
	//DiskEncryptionSet *DiskEncryptionSetParameters `json:"diskEncryptionSet,omitempty"`
	// ID - Resource Id
	//ID *string `json:"id,omitempty"`
}

// OSDisk specifies information about the operating system disk used by the virtual machine. <br><br> For
// more information about disks, see [About disks and VHDs for Azure virtual
// machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-about-disks-vhds?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json).
type OSDisk struct {
	// OsType - This property allows you to specify the type of the OS that is included in the disk if creating a VM from user-image or a specialized VHD. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**. Possible values include: 'Windows', 'Linux'
	//OsType OperatingSystemTypes `json:"osType,omitempty"`
	// EncryptionSettings - Specifies the encryption settings for the OS Disk. <br><br> Minimum api-version: 2015-06-15
	//EncryptionSettings *DiskEncryptionSettings `json:"encryptionSettings,omitempty"`
	// Name - The disk name.
	Name *string `json:"name,omitempty"`
	// Vhd - The virtual hard disk.
	//Vhd *VirtualHardDisk `json:"vhd,omitempty"`
	// Image - The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
	//Image *VirtualHardDisk `json:"image,omitempty"`
	// Caching - Specifies the caching requirements. <br><br> Possible values are: <br><br> **None** <br><br> **ReadOnly** <br><br> **ReadWrite** <br><br> Default: **None** for Standard storage. **ReadOnly** for Premium storage. Possible values include: 'CachingTypesNone', 'CachingTypesReadOnly', 'CachingTypesReadWrite'
	Caching CachingTypes `json:"caching,omitempty"`
	// WriteAcceleratorEnabled - Specifies whether writeAccelerator should be enabled or disabled on the disk.
	//WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty"`
	// DiffDiskSettings - Specifies the ephemeral Disk Settings for the operating system disk used by the virtual machine.
	//DiffDiskSettings *DiffDiskSettings `json:"diffDiskSettings,omitempty"`
	// CreateOption - Specifies how the virtual machine should be created.<br><br> Possible values are:<br><br> **Attach** \u2013 This value is used when you are using a specialized disk to create the virtual machine.<br><br> **FromImage** \u2013 This value is used when you are using an image to create the virtual machine. If you are using a platform image, you also use the imageReference element described above. If you are using a marketplace image, you  also use the plan element previously described. Possible values include: 'DiskCreateOptionTypesFromImage', 'DiskCreateOptionTypesEmpty', 'DiskCreateOptionTypesAttach'
	CreateOption DiskCreateOptionTypes `json:"createOption,omitempty"`
	// DiskSizeGB - Specifies the size of an empty data disk in gigabytes. This element can be used to overwrite the size of the disk in a virtual machine image. <br><br> This value cannot be larger than 1023 GB
	//DiskSizeGB *int32 `json:"diskSizeGB,omitempty"`
	// ManagedDisk - The managed disk parameters.
	ManagedDisk *ManagedDiskParameters `json:"managedDisk,omitempty"`
}

// WindowsConfiguration specifies Windows operating system settings on the virtual machine.
type WindowsConfiguration struct {
	// ProvisionVMAgent - Indicates whether virtual machine agent should be provisioned on the virtual machine. <br><br> When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
	ProvisionVMAgent *bool `json:"provisionVMAgent,omitempty"`
	// EnableAutomaticUpdates - Indicates whether Automatic Updates is enabled for the Windows virtual machine. Default value is true. <br><br> For virtual machine scale sets, this property can be updated and updates will take effect on OS reprovisioning.
	EnableAutomaticUpdates *bool `json:"enableAutomaticUpdates,omitempty"`
	// TimeZone - Specifies the time zone of the virtual machine. e.g. "Pacific Standard Time". <br><br> Possible values can be [TimeZoneInfo.Id](https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.id?#System_TimeZoneInfo_Id) value from time zones returned by [TimeZoneInfo.GetSystemTimeZones](https://docs.microsoft.com/en-us/dotnet/api/system.timezoneinfo.getsystemtimezones).
	TimeZone *string `json:"timeZone,omitempty"`
	// AdditionalUnattendContent - Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.
	//AdditionalUnattendContent *[]AdditionalUnattendContent `json:"additionalUnattendContent,omitempty"`
	// WinRM - Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.
	//WinRM *WinRMConfiguration `json:"winRM,omitempty"`
}

// SSHPublicKey contains information about SSH certificate public key and the path on the Linux VM where
// the public key is placed.
type SSHPublicKey struct {
	// Path - Specifies the full path on the created VM where ssh public key is stored. If the file already exists, the specified key is appended to the file. Example: /home/user/.ssh/authorized_keys
	Path *string `json:"path,omitempty"`
	// KeyData - SSH public key certificate used to authenticate with the VM through ssh. The key needs to be at least 2048-bit and in ssh-rsa format. <br><br> For creating ssh keys, see [Create SSH keys on Linux and Mac for Linux VMs in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-mac-create-ssh-keys?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
	KeyData *string `json:"keyData,omitempty"`
}

// SSHConfiguration SSH configuration for Linux based VMs running on Azure
type SSHConfiguration struct {
	// PublicKeys - The list of SSH public keys used to authenticate with linux based VMs.
	PublicKeys *[]SSHPublicKey `json:"publicKeys,omitempty"`
}

// LinuxConfiguration specifies the Linux operating system settings on the virtual machine. <br><br>For a
// list of supported Linux distributions, see [Linux on Azure-Endorsed
// Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json)
// <br><br> For running non-endorsed distributions, see [Information for Non-Endorsed
// Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
type LinuxConfiguration struct {
	// DisablePasswordAuthentication - Specifies whether password authentication should be disabled.
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty"`
	// SSH - Specifies the ssh key configuration for a Linux OS.
	SSH *SSHConfiguration `json:"ssh,omitempty"`
	// ProvisionVMAgent - Indicates whether virtual machine agent should be provisioned on the virtual machine. <br><br> When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
	ProvisionVMAgent *bool `json:"provisionVMAgent,omitempty"`
}

// NetworkInterfaceReferenceProperties describes a network interface reference properties.
type NetworkInterfaceReferenceProperties struct {
	// Primary - Specifies the primary network interface in case the virtual machine has more than 1 network interface.
	Primary *bool `json:"primary,omitempty"`
}

// NetworkInterfaceReference describes a network interface reference.
type NetworkInterfaceReference struct {
	*NetworkInterfaceReferenceProperties `json:"properties,omitempty"`
	// ID - Resource Id
	ID *string `json:"id,omitempty"`
}

type VirtualMachineSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the VirtualMachine's resource group.
	ResourceGroupName string `json:"resouceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the VirtualMachine's resource
	// group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to the the VirtualMachine's
	// resource group.
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// VirtualMachineProperties - Properties of a virtual machine
	VirtualMachineProperties `json:"properties,omitempty"`

	// Zones - The virtual machine zones.
	Zones *[]string `json:"zones,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

type VirtualMachineStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State of this SecurityGroup.
	State string `json:"state,omitempty"`

	// A Message providing detail about the state of this SecurityGroup, if
	// any.
	Message string `json:"message,omitempty"`

	// ID of this SecurityGroup.
	ID string `json:"id,omitempty"`

	// Etag - A unique read-only string that changes whenever the resource is
	// updated.
	Etag string `json:"etag,omitempty"`

	// ResourceGUID - The GUID of this SecurityGroup.
	ResourceGUID string `json:"resourceGuid,omitempty"`

	// Type of this SecurityGroup.
	Type string `json:"type,omitempty"`
}

// +kubebuilder:object:root=true

// A VirtualMachine is a managed resource that represents a virtual Machine on Azure
// Engine cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="ENDPOINT",type="string",JSONPath=".status.endpoint"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
// +kubebuilder:subresource:status
type VirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualMachineSpec   `json:",spec"`
	Status VirtualMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualMachineList contains a list of VirtualMachine.
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}
