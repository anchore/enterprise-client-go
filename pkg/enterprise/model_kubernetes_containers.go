/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesContainers Containers defined in Kubernetes
type KubernetesContainers struct {
	Containers []KubernetesContainer `json:"containers,omitempty"`
}

// NewKubernetesContainers instantiates a new KubernetesContainers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesContainers() *KubernetesContainers {
	this := KubernetesContainers{}
	return &this
}

// NewKubernetesContainersWithDefaults instantiates a new KubernetesContainers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesContainersWithDefaults() *KubernetesContainers {
	this := KubernetesContainers{}
	return &this
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *KubernetesContainers) GetContainers() []KubernetesContainer {
	if o == nil || o.Containers == nil {
		var ret []KubernetesContainer
		return ret
	}
	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesContainers) GetContainersOk() ([]KubernetesContainer, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *KubernetesContainers) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []KubernetesContainer and assigns it to the Containers field.
func (o *KubernetesContainers) SetContainers(v []KubernetesContainer) {
	o.Containers = v
}

func (o KubernetesContainers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesContainers struct {
	value *KubernetesContainers
	isSet bool
}

func (v NullableKubernetesContainers) Get() *KubernetesContainers {
	return v.value
}

func (v *NullableKubernetesContainers) Set(val *KubernetesContainers) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesContainers) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesContainers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesContainers(val *KubernetesContainers) *NullableKubernetesContainers {
	return &NullableKubernetesContainers{value: val, isSet: true}
}

func (v NullableKubernetesContainers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesContainers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


