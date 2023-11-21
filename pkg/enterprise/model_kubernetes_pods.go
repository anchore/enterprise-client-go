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

// KubernetesPods Pods defined in Kubernetes
type KubernetesPods struct {
	Namespaces []KubernetesPod `json:"namespaces,omitempty"`
}

// NewKubernetesPods instantiates a new KubernetesPods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesPods() *KubernetesPods {
	this := KubernetesPods{}
	return &this
}

// NewKubernetesPodsWithDefaults instantiates a new KubernetesPods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesPodsWithDefaults() *KubernetesPods {
	this := KubernetesPods{}
	return &this
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *KubernetesPods) GetNamespaces() []KubernetesPod {
	if o == nil || o.Namespaces == nil {
		var ret []KubernetesPod
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPods) GetNamespacesOk() ([]KubernetesPod, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *KubernetesPods) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []KubernetesPod and assigns it to the Namespaces field.
func (o *KubernetesPods) SetNamespaces(v []KubernetesPod) {
	o.Namespaces = v
}

func (o KubernetesPods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesPods struct {
	value *KubernetesPods
	isSet bool
}

func (v NullableKubernetesPods) Get() *KubernetesPods {
	return v.value
}

func (v *NullableKubernetesPods) Set(val *KubernetesPods) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesPods) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesPods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesPods(val *KubernetesPods) *NullableKubernetesPods {
	return &NullableKubernetesPods{value: val, isSet: true}
}

func (v NullableKubernetesPods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesPods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


