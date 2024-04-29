/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesInventoryNode struct for KubernetesInventoryNode
type KubernetesInventoryNode struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Labels *map[string]string `json:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
	KernelVersion *string `json:"kernel_version,omitempty"`
	KubernetesVersion *string `json:"kubernetes_version,omitempty"`
	Arch *string `json:"arch,omitempty"`
	ContainerRuntimeVersion *string `json:"container_runtime_version,omitempty"`
	KubeProxyVersion *string `json:"kube_proxy_version,omitempty"`
	KubeletVersion *string `json:"kubelet_version,omitempty"`
	OperatingSystem *string `json:"operating_system,omitempty"`
}

// NewKubernetesInventoryNode instantiates a new KubernetesInventoryNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventoryNode(uid string, name string) *KubernetesInventoryNode {
	this := KubernetesInventoryNode{}
	this.Uid = uid
	this.Name = name
	return &this
}

// NewKubernetesInventoryNodeWithDefaults instantiates a new KubernetesInventoryNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryNodeWithDefaults() *KubernetesInventoryNode {
	this := KubernetesInventoryNode{}
	return &this
}

// GetUid returns the Uid field value
func (o *KubernetesInventoryNode) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *KubernetesInventoryNode) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *KubernetesInventoryNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesInventoryNode) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *KubernetesInventoryNode) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *KubernetesInventoryNode) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetKernelVersion returns the KernelVersion field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetKernelVersion() string {
	if o == nil || o.KernelVersion == nil {
		var ret string
		return ret
	}
	return *o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKernelVersionOk() (*string, bool) {
	if o == nil || o.KernelVersion == nil {
		return nil, false
	}
	return o.KernelVersion, true
}

// HasKernelVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKernelVersion() bool {
	if o != nil && o.KernelVersion != nil {
		return true
	}

	return false
}

// SetKernelVersion gets a reference to the given string and assigns it to the KernelVersion field.
func (o *KubernetesInventoryNode) SetKernelVersion(v string) {
	o.KernelVersion = &v
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetKubernetesVersion() string {
	if o == nil || o.KubernetesVersion == nil {
		var ret string
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKubernetesVersionOk() (*string, bool) {
	if o == nil || o.KubernetesVersion == nil {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKubernetesVersion() bool {
	if o != nil && o.KubernetesVersion != nil {
		return true
	}

	return false
}

// SetKubernetesVersion gets a reference to the given string and assigns it to the KubernetesVersion field.
func (o *KubernetesInventoryNode) SetKubernetesVersion(v string) {
	o.KubernetesVersion = &v
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasArch() bool {
	if o != nil && o.Arch != nil {
		return true
	}

	return false
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *KubernetesInventoryNode) SetArch(v string) {
	o.Arch = &v
}

// GetContainerRuntimeVersion returns the ContainerRuntimeVersion field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetContainerRuntimeVersion() string {
	if o == nil || o.ContainerRuntimeVersion == nil {
		var ret string
		return ret
	}
	return *o.ContainerRuntimeVersion
}

// GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetContainerRuntimeVersionOk() (*string, bool) {
	if o == nil || o.ContainerRuntimeVersion == nil {
		return nil, false
	}
	return o.ContainerRuntimeVersion, true
}

// HasContainerRuntimeVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasContainerRuntimeVersion() bool {
	if o != nil && o.ContainerRuntimeVersion != nil {
		return true
	}

	return false
}

// SetContainerRuntimeVersion gets a reference to the given string and assigns it to the ContainerRuntimeVersion field.
func (o *KubernetesInventoryNode) SetContainerRuntimeVersion(v string) {
	o.ContainerRuntimeVersion = &v
}

// GetKubeProxyVersion returns the KubeProxyVersion field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetKubeProxyVersion() string {
	if o == nil || o.KubeProxyVersion == nil {
		var ret string
		return ret
	}
	return *o.KubeProxyVersion
}

// GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKubeProxyVersionOk() (*string, bool) {
	if o == nil || o.KubeProxyVersion == nil {
		return nil, false
	}
	return o.KubeProxyVersion, true
}

// HasKubeProxyVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKubeProxyVersion() bool {
	if o != nil && o.KubeProxyVersion != nil {
		return true
	}

	return false
}

// SetKubeProxyVersion gets a reference to the given string and assigns it to the KubeProxyVersion field.
func (o *KubernetesInventoryNode) SetKubeProxyVersion(v string) {
	o.KubeProxyVersion = &v
}

// GetKubeletVersion returns the KubeletVersion field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetKubeletVersion() string {
	if o == nil || o.KubeletVersion == nil {
		var ret string
		return ret
	}
	return *o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKubeletVersionOk() (*string, bool) {
	if o == nil || o.KubeletVersion == nil {
		return nil, false
	}
	return o.KubeletVersion, true
}

// HasKubeletVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKubeletVersion() bool {
	if o != nil && o.KubeletVersion != nil {
		return true
	}

	return false
}

// SetKubeletVersion gets a reference to the given string and assigns it to the KubeletVersion field.
func (o *KubernetesInventoryNode) SetKubeletVersion(v string) {
	o.KubeletVersion = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *KubernetesInventoryNode) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *KubernetesInventoryNode) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

func (o KubernetesInventoryNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.KernelVersion != nil {
		toSerialize["kernel_version"] = o.KernelVersion
	}
	if o.KubernetesVersion != nil {
		toSerialize["kubernetes_version"] = o.KubernetesVersion
	}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.ContainerRuntimeVersion != nil {
		toSerialize["container_runtime_version"] = o.ContainerRuntimeVersion
	}
	if o.KubeProxyVersion != nil {
		toSerialize["kube_proxy_version"] = o.KubeProxyVersion
	}
	if o.KubeletVersion != nil {
		toSerialize["kubelet_version"] = o.KubeletVersion
	}
	if o.OperatingSystem != nil {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesInventoryNode struct {
	value *KubernetesInventoryNode
	isSet bool
}

func (v NullableKubernetesInventoryNode) Get() *KubernetesInventoryNode {
	return v.value
}

func (v *NullableKubernetesInventoryNode) Set(val *KubernetesInventoryNode) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventoryNode) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventoryNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventoryNode(val *KubernetesInventoryNode) *NullableKubernetesInventoryNode {
	return &NullableKubernetesInventoryNode{value: val, isSet: true}
}

func (v NullableKubernetesInventoryNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventoryNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


