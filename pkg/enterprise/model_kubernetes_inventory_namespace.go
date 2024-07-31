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

// KubernetesInventoryNamespace struct for KubernetesInventoryNamespace
type KubernetesInventoryNamespace struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Labels *map[string]string `json:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewKubernetesInventoryNamespace instantiates a new KubernetesInventoryNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventoryNamespace(uid string, name string) *KubernetesInventoryNamespace {
	this := KubernetesInventoryNamespace{}
	this.Uid = uid
	this.Name = name
	return &this
}

// NewKubernetesInventoryNamespaceWithDefaults instantiates a new KubernetesInventoryNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryNamespaceWithDefaults() *KubernetesInventoryNamespace {
	this := KubernetesInventoryNamespace{}
	return &this
}

// GetUid returns the Uid field value
func (o *KubernetesInventoryNamespace) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespace) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *KubernetesInventoryNamespace) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *KubernetesInventoryNamespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesInventoryNamespace) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *KubernetesInventoryNamespace) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespace) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesInventoryNamespace) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *KubernetesInventoryNamespace) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *KubernetesInventoryNamespace) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryNamespace) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesInventoryNamespace) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *KubernetesInventoryNamespace) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o KubernetesInventoryNamespace) MarshalJSON() ([]byte, error) {
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

type NullableKubernetesInventoryNamespace struct {
	value *KubernetesInventoryNamespace
	isSet bool
}

func (v NullableKubernetesInventoryNamespace) Get() *KubernetesInventoryNamespace {
	return v.value
}

func (v *NullableKubernetesInventoryNamespace) Set(val *KubernetesInventoryNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventoryNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventoryNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventoryNamespace(val *KubernetesInventoryNamespace) *NullableKubernetesInventoryNamespace {
	return &NullableKubernetesInventoryNamespace{value: val, isSet: true}
}

func (v NullableKubernetesInventoryNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventoryNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


