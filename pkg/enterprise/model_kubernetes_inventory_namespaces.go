/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesInventoryNamespaces struct for KubernetesInventoryNamespaces
type KubernetesInventoryNamespaces struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Labels *map[string]string `json:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewKubernetesInventoryNamespaces instantiates a new KubernetesInventoryNamespaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventoryNamespaces(uid string, name string) *KubernetesInventoryNamespaces {
	this := KubernetesInventoryNamespaces{}
	this.Uid = uid
	this.Name = name
	return &this
}

// NewKubernetesInventoryNamespacesWithDefaults instantiates a new KubernetesInventoryNamespaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryNamespacesWithDefaults() *KubernetesInventoryNamespaces {
	this := KubernetesInventoryNamespaces{}
	return &this
}

// GetUid returns the Uid field value
func (o *KubernetesInventoryNamespaces) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespaces) GetUidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *KubernetesInventoryNamespaces) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *KubernetesInventoryNamespaces) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespaces) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesInventoryNamespaces) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *KubernetesInventoryNamespaces) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespaces) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesInventoryNamespaces) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *KubernetesInventoryNamespaces) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *KubernetesInventoryNamespaces) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespaces) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesInventoryNamespaces) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *KubernetesInventoryNamespaces) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o KubernetesInventoryNamespaces) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableKubernetesInventoryNamespaces struct {
	value *KubernetesInventoryNamespaces
	isSet bool
}

func (v NullableKubernetesInventoryNamespaces) Get() *KubernetesInventoryNamespaces {
	return v.value
}

func (v *NullableKubernetesInventoryNamespaces) Set(val *KubernetesInventoryNamespaces) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventoryNamespaces) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventoryNamespaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventoryNamespaces(val *KubernetesInventoryNamespaces) *NullableKubernetesInventoryNamespaces {
	return &NullableKubernetesInventoryNamespaces{value: val, isSet: true}
}

func (v NullableKubernetesInventoryNamespaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventoryNamespaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


