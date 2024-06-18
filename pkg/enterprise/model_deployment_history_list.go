/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

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
	if o == nil || o.Items == nil {
		var ret []DeploymentHistory
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryList) GetItemsOk() ([]DeploymentHistory, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeploymentHistoryList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DeploymentHistory and assigns it to the Items field.
func (o *DeploymentHistoryList) SetItems(v []DeploymentHistory) {
	o.Items = v
}

func (o DeploymentHistoryList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
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


