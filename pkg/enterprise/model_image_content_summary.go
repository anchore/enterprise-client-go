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

// checks if the ImageContentSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageContentSummary{}

// ImageContentSummary A summary of the content types found in an image with counts
type ImageContentSummary struct {
	Items []ImageContentSummaryItemsInner `json:"items,omitempty"`
}

// NewImageContentSummary instantiates a new ImageContentSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageContentSummary() *ImageContentSummary {
	this := ImageContentSummary{}
	return &this
}

// NewImageContentSummaryWithDefaults instantiates a new ImageContentSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageContentSummaryWithDefaults() *ImageContentSummary {
	this := ImageContentSummary{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ImageContentSummary) GetItems() []ImageContentSummaryItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []ImageContentSummaryItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageContentSummary) GetItemsOk() ([]ImageContentSummaryItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ImageContentSummary) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ImageContentSummaryItemsInner and assigns it to the Items field.
func (o *ImageContentSummary) SetItems(v []ImageContentSummaryItemsInner) {
	o.Items = v
}

func (o ImageContentSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageContentSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableImageContentSummary struct {
	value *ImageContentSummary
	isSet bool
}

func (v NullableImageContentSummary) Get() *ImageContentSummary {
	return v.value
}

func (v *NullableImageContentSummary) Set(val *ImageContentSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableImageContentSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableImageContentSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageContentSummary(val *ImageContentSummary) *NullableImageContentSummary {
	return &NullableImageContentSummary{value: val, isSet: true}
}

func (v NullableImageContentSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageContentSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


