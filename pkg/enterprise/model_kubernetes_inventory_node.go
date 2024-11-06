/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubernetesInventoryNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesInventoryNode{}

// KubernetesInventoryNode struct for KubernetesInventoryNode
type KubernetesInventoryNode struct {
	Uid string `json:"uid" validate:"regexp=^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"`
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

type _KubernetesInventoryNode KubernetesInventoryNode

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
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
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
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
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
	if o == nil || IsNil(o.KernelVersion) {
		var ret string
		return ret
	}
	return *o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKernelVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KernelVersion) {
		return nil, false
	}
	return o.KernelVersion, true
}

// HasKernelVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKernelVersion() bool {
	if o != nil && !IsNil(o.KernelVersion) {
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
	if o == nil || IsNil(o.KubernetesVersion) {
		var ret string
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKubernetesVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubernetesVersion) {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKubernetesVersion() bool {
	if o != nil && !IsNil(o.KubernetesVersion) {
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
	if o == nil || IsNil(o.Arch) {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetArchOk() (*string, bool) {
	if o == nil || IsNil(o.Arch) {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasArch() bool {
	if o != nil && !IsNil(o.Arch) {
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
	if o == nil || IsNil(o.ContainerRuntimeVersion) {
		var ret string
		return ret
	}
	return *o.ContainerRuntimeVersion
}

// GetContainerRuntimeVersionOk returns a tuple with the ContainerRuntimeVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetContainerRuntimeVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerRuntimeVersion) {
		return nil, false
	}
	return o.ContainerRuntimeVersion, true
}

// HasContainerRuntimeVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasContainerRuntimeVersion() bool {
	if o != nil && !IsNil(o.ContainerRuntimeVersion) {
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
	if o == nil || IsNil(o.KubeProxyVersion) {
		var ret string
		return ret
	}
	return *o.KubeProxyVersion
}

// GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKubeProxyVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubeProxyVersion) {
		return nil, false
	}
	return o.KubeProxyVersion, true
}

// HasKubeProxyVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKubeProxyVersion() bool {
	if o != nil && !IsNil(o.KubeProxyVersion) {
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
	if o == nil || IsNil(o.KubeletVersion) {
		var ret string
		return ret
	}
	return *o.KubeletVersion
}

// GetKubeletVersionOk returns a tuple with the KubeletVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetKubeletVersionOk() (*string, bool) {
	if o == nil || IsNil(o.KubeletVersion) {
		return nil, false
	}
	return o.KubeletVersion, true
}

// HasKubeletVersion returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasKubeletVersion() bool {
	if o != nil && !IsNil(o.KubeletVersion) {
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
	if o == nil || IsNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNode) GetOperatingSystemOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *KubernetesInventoryNode) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *KubernetesInventoryNode) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

func (o KubernetesInventoryNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesInventoryNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uid"] = o.Uid
	toSerialize["name"] = o.Name
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.KernelVersion) {
		toSerialize["kernel_version"] = o.KernelVersion
	}
	if !IsNil(o.KubernetesVersion) {
		toSerialize["kubernetes_version"] = o.KubernetesVersion
	}
	if !IsNil(o.Arch) {
		toSerialize["arch"] = o.Arch
	}
	if !IsNil(o.ContainerRuntimeVersion) {
		toSerialize["container_runtime_version"] = o.ContainerRuntimeVersion
	}
	if !IsNil(o.KubeProxyVersion) {
		toSerialize["kube_proxy_version"] = o.KubeProxyVersion
	}
	if !IsNil(o.KubeletVersion) {
		toSerialize["kubelet_version"] = o.KubeletVersion
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	return toSerialize, nil
}

func (o *KubernetesInventoryNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uid",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubernetesInventoryNode := _KubernetesInventoryNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubernetesInventoryNode)

	if err != nil {
		return err
	}

	*o = KubernetesInventoryNode(varKubernetesInventoryNode)

	return err
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


