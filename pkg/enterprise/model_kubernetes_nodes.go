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

// KubernetesNodes Nodes defined in Kubernetes
type KubernetesNodes struct {
	Namespaces []KubernetesNode `json:"namespaces,omitempty"`
}

// NewKubernetesNodes instantiates a new KubernetesNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodes() *KubernetesNodes {
	this := KubernetesNodes{}
	return &this
}

// NewKubernetesNodesWithDefaults instantiates a new KubernetesNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodesWithDefaults() *KubernetesNodes {
	this := KubernetesNodes{}
	return &this
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *KubernetesNodes) GetNamespaces() []KubernetesNode {
	if o == nil || o.Namespaces == nil {
		var ret []KubernetesNode
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodes) GetNamespacesOk() ([]KubernetesNode, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *KubernetesNodes) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []KubernetesNode and assigns it to the Namespaces field.
func (o *KubernetesNodes) SetNamespaces(v []KubernetesNode) {
	o.Namespaces = v
}

func (o KubernetesNodes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesNodes struct {
	value *KubernetesNodes
	isSet bool
}

func (v NullableKubernetesNodes) Get() *KubernetesNodes {
	return v.value
}

func (v *NullableKubernetesNodes) Set(val *KubernetesNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodes(val *KubernetesNodes) *NullableKubernetesNodes {
	return &NullableKubernetesNodes{value: val, isSet: true}
}

func (v NullableKubernetesNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


