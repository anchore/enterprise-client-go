/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// KubernetesInventoryPod struct for KubernetesInventoryPod
type KubernetesInventoryPod struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
	NamespaceUid string `json:"namespace_uid"`
	NodeUid *string `json:"node_uid,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty"`
}

// NewKubernetesInventoryPod instantiates a new KubernetesInventoryPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesInventoryPod(uid string, name string, namespaceUid string) *KubernetesInventoryPod {
	this := KubernetesInventoryPod{}
	this.Uid = uid
	this.Name = name
	this.NamespaceUid = namespaceUid
	return &this
}

// NewKubernetesInventoryPodWithDefaults instantiates a new KubernetesInventoryPod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesInventoryPodWithDefaults() *KubernetesInventoryPod {
	this := KubernetesInventoryPod{}
	return &this
}

// GetUid returns the Uid field value
func (o *KubernetesInventoryPod) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryPod) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *KubernetesInventoryPod) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *KubernetesInventoryPod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryPod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesInventoryPod) SetName(v string) {
	o.Name = v
}

// GetNamespaceUid returns the NamespaceUid field value
func (o *KubernetesInventoryPod) GetNamespaceUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NamespaceUid
}

// GetNamespaceUidOk returns a tuple with the NamespaceUid field value
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryPod) GetNamespaceUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NamespaceUid, true
}

// SetNamespaceUid sets field value
func (o *KubernetesInventoryPod) SetNamespaceUid(v string) {
	o.NamespaceUid = v
}

// GetNodeUid returns the NodeUid field value if set, zero value otherwise.
func (o *KubernetesInventoryPod) GetNodeUid() string {
	if o == nil || o.NodeUid == nil {
		var ret string
		return ret
	}
	return *o.NodeUid
}

// GetNodeUidOk returns a tuple with the NodeUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryPod) GetNodeUidOk() (*string, bool) {
	if o == nil || o.NodeUid == nil {
		return nil, false
	}
	return o.NodeUid, true
}

// HasNodeUid returns a boolean if a field has been set.
func (o *KubernetesInventoryPod) HasNodeUid() bool {
	if o != nil && o.NodeUid != nil {
		return true
	}

	return false
}

// SetNodeUid gets a reference to the given string and assigns it to the NodeUid field.
func (o *KubernetesInventoryPod) SetNodeUid(v string) {
	o.NodeUid = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *KubernetesInventoryPod) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryPod) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *KubernetesInventoryPod) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *KubernetesInventoryPod) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *KubernetesInventoryPod) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesInventoryPod) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *KubernetesInventoryPod) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *KubernetesInventoryPod) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

func (o KubernetesInventoryPod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uid"] = o.Uid
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["namespace_uid"] = o.NamespaceUid
	}
	if o.NodeUid != nil {
		toSerialize["node_uid"] = o.NodeUid
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesInventoryPod struct {
	value *KubernetesInventoryPod
	isSet bool
}

func (v NullableKubernetesInventoryPod) Get() *KubernetesInventoryPod {
	return v.value
}

func (v *NullableKubernetesInventoryPod) Set(val *KubernetesInventoryPod) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesInventoryPod) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesInventoryPod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesInventoryPod(val *KubernetesInventoryPod) *NullableKubernetesInventoryPod {
	return &NullableKubernetesInventoryPod{value: val, isSet: true}
}

func (v NullableKubernetesInventoryPod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesInventoryPod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


