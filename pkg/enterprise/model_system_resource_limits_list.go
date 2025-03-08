/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the SystemResourceLimitsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemResourceLimitsList{}

// SystemResourceLimitsList struct for SystemResourceLimitsList
type SystemResourceLimitsList struct {
	Items []SystemResourceLimit `json:"items,omitempty"`
}

// NewSystemResourceLimitsList instantiates a new SystemResourceLimitsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemResourceLimitsList() *SystemResourceLimitsList {
	this := SystemResourceLimitsList{}
	return &this
}

// NewSystemResourceLimitsListWithDefaults instantiates a new SystemResourceLimitsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemResourceLimitsListWithDefaults() *SystemResourceLimitsList {
	this := SystemResourceLimitsList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SystemResourceLimitsList) GetItems() []SystemResourceLimit {
	if o == nil || IsNil(o.Items) {
		var ret []SystemResourceLimit
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemResourceLimitsList) GetItemsOk() ([]SystemResourceLimit, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SystemResourceLimitsList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SystemResourceLimit and assigns it to the Items field.
func (o *SystemResourceLimitsList) SetItems(v []SystemResourceLimit) {
	o.Items = v
}

func (o SystemResourceLimitsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemResourceLimitsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSystemResourceLimitsList struct {
	value *SystemResourceLimitsList
	isSet bool
}

func (v NullableSystemResourceLimitsList) Get() *SystemResourceLimitsList {
	return v.value
}

func (v *NullableSystemResourceLimitsList) Set(val *SystemResourceLimitsList) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemResourceLimitsList) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemResourceLimitsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemResourceLimitsList(val *SystemResourceLimitsList) *NullableSystemResourceLimitsList {
	return &NullableSystemResourceLimitsList{value: val, isSet: true}
}

func (v NullableSystemResourceLimitsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemResourceLimitsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


