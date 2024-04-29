/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesInventoryContainer struct for KubernetesInventoryContainer
type KubernetesInventoryContainer struct {
	// Corresponds to ContainerStatus.containerID in the Kubernetes Spec
	Id string `json:"id"`
	Name string `json:"name"`
	ImageTag string `json:"image_tag"`
	ImageDigest *string `json:"image_digest,omitempty"`
	PodUid string `json:"pod_uid"`
}

// NewKubernetesInventoryContainer instantiates a new KubernetesInventoryContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventoryContainer(id string, name string, imageTag string, podUid string) *KubernetesInventoryContainer {
	this := KubernetesInventoryContainer{}
	this.Id = id
	this.Name = name
	this.ImageTag = imageTag
	this.PodUid = podUid
	return &this
}

// NewKubernetesInventoryContainerWithDefaults instantiates a new KubernetesInventoryContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryContainerWithDefaults() *KubernetesInventoryContainer {
	this := KubernetesInventoryContainer{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesInventoryContainer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesInventoryContainer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KubernetesInventoryContainer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesInventoryContainer) SetName(v string) {
	o.Name = v
}

// GetImageTag returns the ImageTag field value
func (o *KubernetesInventoryContainer) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainer) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *KubernetesInventoryContainer) SetImageTag(v string) {
	o.ImageTag = v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *KubernetesInventoryContainer) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainer) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *KubernetesInventoryContainer) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *KubernetesInventoryContainer) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetPodUid returns the PodUid field value
func (o *KubernetesInventoryContainer) GetPodUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodUid
}

// GetPodUidOk returns a tuple with the PodUid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainer) GetPodUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodUid, true
}

// SetPodUid sets field value
func (o *KubernetesInventoryContainer) SetPodUid(v string) {
	o.PodUid = v
}

func (o KubernetesInventoryContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if true {
		toSerialize["pod_uid"] = o.PodUid
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesInventoryContainer struct {
	value *KubernetesInventoryContainer
	isSet bool
}

func (v NullableKubernetesInventoryContainer) Get() *KubernetesInventoryContainer {
	return v.value
}

func (v *NullableKubernetesInventoryContainer) Set(val *KubernetesInventoryContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventoryContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventoryContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventoryContainer(val *KubernetesInventoryContainer) *NullableKubernetesInventoryContainer {
	return &NullableKubernetesInventoryContainer{value: val, isSet: true}
}

func (v NullableKubernetesInventoryContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventoryContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


