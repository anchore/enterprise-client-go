/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the SystemConfigurationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemConfigurationList{}

// SystemConfigurationList struct for SystemConfigurationList
type SystemConfigurationList struct {
	Items []SystemConfiguration `json:"items,omitempty"`
}

// NewSystemConfigurationList instantiates a new SystemConfigurationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemConfigurationList() *SystemConfigurationList {
	this := SystemConfigurationList{}
	return &this
}

// NewSystemConfigurationListWithDefaults instantiates a new SystemConfigurationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemConfigurationListWithDefaults() *SystemConfigurationList {
	this := SystemConfigurationList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SystemConfigurationList) GetItems() []SystemConfiguration {
	if o == nil || IsNil(o.Items) {
		var ret []SystemConfiguration
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemConfigurationList) GetItemsOk() ([]SystemConfiguration, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SystemConfigurationList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SystemConfiguration and assigns it to the Items field.
func (o *SystemConfigurationList) SetItems(v []SystemConfiguration) {
	o.Items = v
}

func (o SystemConfigurationList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemConfigurationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSystemConfigurationList struct {
	value *SystemConfigurationList
	isSet bool
}

func (v NullableSystemConfigurationList) Get() *SystemConfigurationList {
	return v.value
}

func (v *NullableSystemConfigurationList) Set(val *SystemConfigurationList) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigurationList) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigurationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigurationList(val *SystemConfigurationList) *NullableSystemConfigurationList {
	return &NullableSystemConfigurationList{value: val, isSet: true}
}

func (v NullableSystemConfigurationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigurationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


