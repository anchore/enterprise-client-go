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

// checks if the DeploymentHistoryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryList{}

// DeploymentHistoryList struct for DeploymentHistoryList
type DeploymentHistoryList struct {
	Items []DeploymentHistory `json:"items,omitempty"`
}

// NewDeploymentHistoryList instantiates a new DeploymentHistoryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryList() *DeploymentHistoryList {
	this := DeploymentHistoryList{}
	return &this
}

// NewDeploymentHistoryListWithDefaults instantiates a new DeploymentHistoryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryListWithDefaults() *DeploymentHistoryList {
	this := DeploymentHistoryList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DeploymentHistoryList) GetItems() []DeploymentHistory {
	if o == nil || IsNil(o.Items) {
		var ret []DeploymentHistory
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryList) GetItemsOk() ([]DeploymentHistory, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeploymentHistoryList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DeploymentHistory and assigns it to the Items field.
func (o *DeploymentHistoryList) SetItems(v []DeploymentHistory) {
	o.Items = v
}

func (o DeploymentHistoryList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableDeploymentHistoryList struct {
	value *DeploymentHistoryList
	isSet bool
}

func (v NullableDeploymentHistoryList) Get() *DeploymentHistoryList {
	return v.value
}

func (v *NullableDeploymentHistoryList) Set(val *DeploymentHistoryList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryList(val *DeploymentHistoryList) *NullableDeploymentHistoryList {
	return &NullableDeploymentHistoryList{value: val, isSet: true}
}

func (v NullableDeploymentHistoryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


