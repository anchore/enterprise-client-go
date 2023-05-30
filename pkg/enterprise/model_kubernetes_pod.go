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

// KubernetesPod struct for KubernetesPod
type KubernetesPod struct {
	Id string `json:"id"`
	Name string `json:"name"`
	AccountName string `json:"account_name"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	NodeId *string `json:"node_id,omitempty"`
	NamespaceId *string `json:"namespace_id,omitempty"`
	LastSeen *string `json:"last_seen,omitempty"`
}

// NewKubernetesPod instantiates a new KubernetesPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesPod(id string, name string, accountName string, labels map[string]string, annotations map[string]string) *KubernetesPod {
	this := KubernetesPod{}
	this.Id = id
	this.Name = name
	this.AccountName = accountName
	this.Labels = labels
	this.Annotations = annotations
	return &this
}

// NewKubernetesPodWithDefaults instantiates a new KubernetesPod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesPodWithDefaults() *KubernetesPod {
	this := KubernetesPod{}
	return &this
}

// GetId returns the Id field value
func (o *KubernetesPod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *KubernetesPod) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *KubernetesPod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubernetesPod) SetName(v string) {
	o.Name = v
}

// GetAccountName returns the AccountName field value
func (o *KubernetesPod) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *KubernetesPod) SetAccountName(v string) {
	o.AccountName = v
}

// GetLabels returns the Labels field value
func (o *KubernetesPod) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetLabelsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *KubernetesPod) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetAnnotations returns the Annotations field value
func (o *KubernetesPod) GetAnnotations() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value
func (o *KubernetesPod) SetAnnotations(v map[string]string) {
	o.Annotations = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *KubernetesPod) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *KubernetesPod) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *KubernetesPod) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *KubernetesPod) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *KubernetesPod) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *KubernetesPod) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *KubernetesPod) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPod) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *KubernetesPod) HasLastSeen() bool {
	if o != nil && o.LastSeen != nil {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *KubernetesPod) SetLastSeen(v string) {
	o.LastSeen = &v
}

func (o KubernetesPod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["annotations"] = o.Annotations
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.NamespaceId != nil {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesPod struct {
	value *KubernetesPod
	isSet bool
}

func (v NullableKubernetesPod) Get() *KubernetesPod {
	return v.value
}

func (v *NullableKubernetesPod) Set(val *KubernetesPod) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesPod) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesPod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesPod(val *KubernetesPod) *NullableKubernetesPod {
	return &NullableKubernetesPod{value: val, isSet: true}
}

func (v NullableKubernetesPod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesPod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


