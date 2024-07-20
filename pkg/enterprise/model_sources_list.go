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

// SourcesList struct for SourcesList
type SourcesList struct {
	Items []Source `json:"items,omitempty"`
}

// NewSourcesList instantiates a new SourcesList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourcesList() *SourcesList {
	this := SourcesList{}
	return &this
}

// NewSourcesListWithDefaults instantiates a new SourcesList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourcesListWithDefaults() *SourcesList {
	this := SourcesList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SourcesList) GetItems() []Source {
	if o == nil || o.Items == nil {
		var ret []Source
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourcesList) GetItemsOk() ([]Source, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SourcesList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Source and assigns it to the Items field.
func (o *SourcesList) SetItems(v []Source) {
	o.Items = v
}

func (o SourcesList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableSourcesList struct {
	value *SourcesList
	isSet bool
}

func (v NullableSourcesList) Get() *SourcesList {
	return v.value
}

func (v *NullableSourcesList) Set(val *SourcesList) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesList) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesList(val *SourcesList) *NullableSourcesList {
	return &NullableSourcesList{value: val, isSet: true}
}

func (v NullableSourcesList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


