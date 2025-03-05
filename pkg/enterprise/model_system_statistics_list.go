/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the SystemStatisticsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemStatisticsList{}

// SystemStatisticsList struct for SystemStatisticsList
type SystemStatisticsList struct {
	Items []SystemStatistics `json:"items,omitempty"`
}

// NewSystemStatisticsList instantiates a new SystemStatisticsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemStatisticsList() *SystemStatisticsList {
	this := SystemStatisticsList{}
	return &this
}

// NewSystemStatisticsListWithDefaults instantiates a new SystemStatisticsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemStatisticsListWithDefaults() *SystemStatisticsList {
	this := SystemStatisticsList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SystemStatisticsList) GetItems() []SystemStatistics {
	if o == nil || IsNil(o.Items) {
		var ret []SystemStatistics
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemStatisticsList) GetItemsOk() ([]SystemStatistics, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SystemStatisticsList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SystemStatistics and assigns it to the Items field.
func (o *SystemStatisticsList) SetItems(v []SystemStatistics) {
	o.Items = v
}

func (o SystemStatisticsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemStatisticsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSystemStatisticsList struct {
	value *SystemStatisticsList
	isSet bool
}

func (v NullableSystemStatisticsList) Get() *SystemStatisticsList {
	return v.value
}

func (v *NullableSystemStatisticsList) Set(val *SystemStatisticsList) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemStatisticsList) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemStatisticsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemStatisticsList(val *SystemStatisticsList) *NullableSystemStatisticsList {
	return &NullableSystemStatisticsList{value: val, isSet: true}
}

func (v NullableSystemStatisticsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemStatisticsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


