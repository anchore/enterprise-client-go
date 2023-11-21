/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesInventoryContainers struct for KubernetesInventoryContainers
type KubernetesInventoryContainers struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImageTag string `json:"image_tag"`
	ImageDigest *string `json:"image_digest,omitempty"`
	PodUid string `json:"pod_uid"`
}

// NewKubernetesInventoryContainers instantiates a new KubernetesInventoryContainers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventoryContainers(id string, name string, imageTag string, podUid string) *KubernetesInventoryContainers {
	this := KubernetesInventoryContainers{}
	this.Id = id
	this.Name = name
	this.ImageTag = imageTag
	this.PodUid = podUid
	return &this
}

// NewKubernetesInventoryContainersWithDefaults instantiates a new KubernetesInventoryContainers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryContainersWithDefaults() *KubernetesInventoryContainers {
	this := KubernetesInventoryContainers{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesInventoryContainers) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainers) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesInventoryContainers) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KubernetesInventoryContainers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainers) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesInventoryContainers) SetName(v string) {
	o.Name = v
}

// GetImageTag returns the ImageTag field value
func (o *KubernetesInventoryContainers) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainers) GetImageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *KubernetesInventoryContainers) SetImageTag(v string) {
	o.ImageTag = v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *KubernetesInventoryContainers) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainers) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *KubernetesInventoryContainers) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *KubernetesInventoryContainers) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetPodUid returns the PodUid field value
func (o *KubernetesInventoryContainers) GetPodUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodUid
}

// GetPodUidOk returns a tuple with the PodUid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryContainers) GetPodUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PodUid, true
}

// SetPodUid sets field value
func (o *KubernetesInventoryContainers) SetPodUid(v string) {
	o.PodUid = v
}

func (o KubernetesInventoryContainers) MarshalJSON() ([]byte, error) {
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

type NullableKubernetesInventoryContainers struct {
	value *KubernetesInventoryContainers
	isSet bool
}

func (v NullableKubernetesInventoryContainers) Get() *KubernetesInventoryContainers {
	return v.value
}

func (v *NullableKubernetesInventoryContainers) Set(val *KubernetesInventoryContainers) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventoryContainers) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventoryContainers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventoryContainers(val *KubernetesInventoryContainers) *NullableKubernetesInventoryContainers {
	return &NullableKubernetesInventoryContainers{value: val, isSet: true}
}

func (v NullableKubernetesInventoryContainers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventoryContainers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


