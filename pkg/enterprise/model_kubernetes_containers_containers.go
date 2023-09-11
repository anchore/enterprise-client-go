/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesContainersContainers struct for KubernetesContainersContainers
type KubernetesContainersContainers struct {
	Id string `json:"id"`
	Name string `json:"name"`
	PodId string `json:"pod_id"`
	AccountName string `json:"account_name"`
	Context string `json:"context"`
	ImageTag string `json:"image_tag"`
	ImageDigest string `json:"image_digest"`
}

// NewKubernetesContainersContainers instantiates a new KubernetesContainersContainers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesContainersContainers(id string, name string, podId string, accountName string, context string, imageTag string, imageDigest string) *KubernetesContainersContainers {
	this := KubernetesContainersContainers{}
	this.Id = id
	this.Name = name
	this.PodId = podId
	this.AccountName = accountName
	this.Context = context
	this.ImageTag = imageTag
	this.ImageDigest = imageDigest
	return &this
}

// NewKubernetesContainersContainersWithDefaults instantiates a new KubernetesContainersContainers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesContainersContainersWithDefaults() *KubernetesContainersContainers {
	this := KubernetesContainersContainers{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesContainersContainers) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesContainersContainers) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KubernetesContainersContainers) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesContainersContainers) SetName(v string) {
	o.Name = v
}

// GetPodId returns the PodId field value
func (o *KubernetesContainersContainers) GetPodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodId
}

// GetPodIdOk returns a tuple with the PodId field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetPodIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PodId, true
}

// SetPodId sets field value
func (o *KubernetesContainersContainers) SetPodId(v string) {
	o.PodId = v
}

// GetAccountName returns the AccountName field value
func (o *KubernetesContainersContainers) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *KubernetesContainersContainers) SetAccountName(v string) {
	o.AccountName = v
}

// GetContext returns the Context field value
func (o *KubernetesContainersContainers) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *KubernetesContainersContainers) SetContext(v string) {
	o.Context = v
}

// GetImageTag returns the ImageTag field value
func (o *KubernetesContainersContainers) GetImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetImageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageTag, true
}

// SetImageTag sets field value
func (o *KubernetesContainersContainers) SetImageTag(v string) {
	o.ImageTag = v
}

// GetImageDigest returns the ImageDigest field value
func (o *KubernetesContainersContainers) GetImageDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value
// and a boolean to check if the value has been set.
func (o *KubernetesContainersContainers) GetImageDigestOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ImageDigest, true
}

// SetImageDigest sets field value
func (o *KubernetesContainersContainers) SetImageDigest(v string) {
	o.ImageDigest = v
}

func (o KubernetesContainersContainers) MarshalJSON() ([]byte, error) {
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

type NullableKubernetesContainersContainers struct {
	value *KubernetesContainersContainers
	isSet bool
}

func (v NullableKubernetesContainersContainers) Get() *KubernetesContainersContainers {
	return v.value
}

func (v *NullableKubernetesContainersContainers) Set(val *KubernetesContainersContainers) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesContainersContainers) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesContainersContainers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesContainersContainers(val *KubernetesContainersContainers) *NullableKubernetesContainersContainers {
	return &NullableKubernetesContainersContainers{value: val, isSet: true}
}

func (v NullableKubernetesContainersContainers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesContainersContainers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


