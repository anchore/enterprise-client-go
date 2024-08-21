/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// AnchoreImageList A list of Anchore Images
type AnchoreImageList struct {
	Items []AnchoreImage `json:"items,omitempty"`
}

// NewAnchoreImageList instantiates a new AnchoreImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchoreImageList() *AnchoreImageList {
	this := AnchoreImageList{}
	return &this
}

// NewAnchoreImageListWithDefaults instantiates a new AnchoreImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchoreImageListWithDefaults() *AnchoreImageList {
	this := AnchoreImageList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AnchoreImageList) GetItems() []AnchoreImage {
	if o == nil || o.Items == nil {
		var ret []AnchoreImage
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreImageList) GetItemsOk() ([]AnchoreImage, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AnchoreImageList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AnchoreImage and assigns it to the Items field.
func (o *AnchoreImageList) SetItems(v []AnchoreImage) {
	o.Items = v
}

func (o AnchoreImageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableAnchoreImageList struct {
	value *AnchoreImageList
	isSet bool
}

func (v NullableAnchoreImageList) Get() *AnchoreImageList {
	return v.value
}

func (v *NullableAnchoreImageList) Set(val *AnchoreImageList) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoreImageList) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoreImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoreImageList(val *AnchoreImageList) *NullableAnchoreImageList {
	return &NullableAnchoreImageList{value: val, isSet: true}
}

func (v NullableAnchoreImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoreImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


