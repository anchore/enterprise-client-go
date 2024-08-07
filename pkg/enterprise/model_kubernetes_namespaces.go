/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesNamespaces Namespaces defined in Kubernetes
type KubernetesNamespaces struct {
	Namespaces []KubernetesNamespace `json:"namespaces,omitempty"`
}

// NewKubernetesNamespaces instantiates a new KubernetesNamespaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNamespaces() *KubernetesNamespaces {
	this := KubernetesNamespaces{}
	return &this
}

// NewKubernetesNamespacesWithDefaults instantiates a new KubernetesNamespaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNamespacesWithDefaults() *KubernetesNamespaces {
	this := KubernetesNamespaces{}
	return &this
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *KubernetesNamespaces) GetNamespaces() []KubernetesNamespace {
	if o == nil || o.Namespaces == nil {
		var ret []KubernetesNamespace
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNamespaces) GetNamespacesOk() ([]KubernetesNamespace, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *KubernetesNamespaces) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []KubernetesNamespace and assigns it to the Namespaces field.
func (o *KubernetesNamespaces) SetNamespaces(v []KubernetesNamespace) {
	o.Namespaces = v
}

func (o KubernetesNamespaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesNamespaces struct {
	value *KubernetesNamespaces
	isSet bool
}

func (v NullableKubernetesNamespaces) Get() *KubernetesNamespaces {
	return v.value
}

func (v *NullableKubernetesNamespaces) Set(val *KubernetesNamespaces) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNamespaces) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNamespaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNamespaces(val *KubernetesNamespaces) *NullableKubernetesNamespaces {
	return &NullableKubernetesNamespaces{value: val, isSet: true}
}

func (v NullableKubernetesNamespaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNamespaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


