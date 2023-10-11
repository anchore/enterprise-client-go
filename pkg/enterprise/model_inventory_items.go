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

// checks if the InventoryItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItems{}

// InventoryItems Inventory report for Images in Use
type InventoryItems struct {
	Items []InventoryItem `json:"items,omitempty"`
}

// NewInventoryItems instantiates a new InventoryItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItems() *InventoryItems {
	this := InventoryItems{}
	return &this
}

// NewInventoryItemsWithDefaults instantiates a new InventoryItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemsWithDefaults() *InventoryItems {
	this := InventoryItems{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *InventoryItems) GetItems() []InventoryItem {
	if o == nil || IsNil(o.Items) {
		var ret []InventoryItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItems) GetItemsOk() ([]InventoryItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *InventoryItems) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InventoryItem and assigns it to the Items field.
func (o *InventoryItems) SetItems(v []InventoryItem) {
	o.Items = v
}

func (o InventoryItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableInventoryItems struct {
	value *InventoryItems
	isSet bool
}

func (v NullableInventoryItems) Get() *InventoryItems {
	return v.value
}

func (v *NullableInventoryItems) Set(val *InventoryItems) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItems) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItems(val *InventoryItems) *NullableInventoryItems {
	return &NullableInventoryItems{value: val, isSet: true}
}

func (v NullableInventoryItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


