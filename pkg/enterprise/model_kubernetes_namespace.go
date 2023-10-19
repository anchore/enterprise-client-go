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

// KubernetesNamespace struct for KubernetesNamespace
type KubernetesNamespace struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	LastSeen *string `json:"last_seen,omitempty"`
}

// NewKubernetesNamespace instantiates a new KubernetesNamespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNamespace(uid string, name string, labels map[string]string, annotations map[string]string) *KubernetesNamespace {
	this := KubernetesNamespace{}
	this.Uid = uid
	this.Name = name
	this.Labels = labels
	this.Annotations = annotations
	return &this
}

// NewKubernetesNamespaceWithDefaults instantiates a new KubernetesNamespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNamespaceWithDefaults() *KubernetesNamespace {
	this := KubernetesNamespace{}
	return &this
}

// GetUid returns the Uid field value
func (o *KubernetesNamespace) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *KubernetesNamespace) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *KubernetesNamespace) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *KubernetesNamespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesNamespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesNamespace) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *KubernetesNamespace) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *KubernetesNamespace) GetLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *KubernetesNamespace) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetAnnotations returns the Annotations field value
func (o *KubernetesNamespace) GetAnnotations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *KubernetesNamespace) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value
func (o *KubernetesNamespace) SetAnnotations(v map[string]string) {
	o.Annotations = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *KubernetesNamespace) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNamespace) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *KubernetesNamespace) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *KubernetesNamespace) SetLastSeen(v string) {
	o.LastSeen = &v
}

func (o KubernetesNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["annotations"] = o.Annotations
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesNamespace struct {
	value *KubernetesNamespace
	isSet bool
}

func (v NullableKubernetesNamespace) Get() *KubernetesNamespace {
	return v.value
}

func (v *NullableKubernetesNamespace) Set(val *KubernetesNamespace) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNamespace) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNamespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNamespace(val *KubernetesNamespace) *NullableKubernetesNamespace {
	return &NullableKubernetesNamespace{value: val, isSet: true}
}

func (v NullableKubernetesNamespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNamespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


