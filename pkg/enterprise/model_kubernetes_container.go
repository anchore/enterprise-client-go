/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesContainer struct for KubernetesContainer
type KubernetesContainer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	PodId string `json:"pod_id"`
	AccountName string `json:"account_name"`
	Context string `json:"context"`
	ImageTag string `json:"image_tag"`
	ImageDigest string `json:"image_digest"`
}

// NewKubernetesContainer instantiates a new KubernetesContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesContainer(id string, name string, podId string, accountName string, context string, imageTag string, imageDigest string) *KubernetesContainer {
	this := KubernetesContainer{}
	this.Id = id
	this.Name = name
	this.PodId = podId
	this.AccountName = accountName
	this.Context = context
	this.ImageTag = imageTag
	this.ImageDigest = imageDigest
	return &this
}

// NewKubernetesContainerWithDefaults instantiates a new KubernetesContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesContainerWithDefaults() *KubernetesContainer {
	this := KubernetesContainer{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesContainer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesContainer) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KubernetesContainer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesContainer) SetName(v string) {
	o.Name = v
}

// GetPodId returns the PodId field value
func (o *KubernetesContainer) GetPodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodId
}

// GetPodIdOk returns a tuple with the PodId field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetPodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodId, true
}

// SetPodId sets field value
func (o *KubernetesContainer) SetPodId(v string) {
	o.PodId = v
}

// GetAccountName returns the AccountName field value
func (o *KubernetesContainer) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *KubernetesContainer) SetAccountName(v string) {
	o.AccountName = v
}

// GetContext returns the Context field value
func (o *KubernetesContainer) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *KubernetesContainer) SetContext(v string) {
	o.Context = v
}

// GetImageTag returns the ImageTag field value
func (o *KubernetesContainer) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *KubernetesContainer) SetImageTag(v string) {
	o.ImageTag = v
}

// GetImageDigest returns the ImageDigest field value
func (o *KubernetesContainer) GetImageDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainer) GetImageDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageDigest, true
}

// SetImageDigest sets field value
func (o *KubernetesContainer) SetImageDigest(v string) {
	o.ImageDigest = v
}

func (o KubernetesContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["pod_id"] = o.PodId
	}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["image_tag"] = o.ImageTag
	}
	if true {
		toSerialize["image_digest"] = o.ImageDigest
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesContainer struct {
	value *KubernetesContainer
	isSet bool
}

func (v NullableKubernetesContainer) Get() *KubernetesContainer {
	return v.value
}

func (v *NullableKubernetesContainer) Set(val *KubernetesContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesContainer(val *KubernetesContainer) *NullableKubernetesContainer {
	return &NullableKubernetesContainer{value: val, isSet: true}
}

func (v NullableKubernetesContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


