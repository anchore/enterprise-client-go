/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ImageContentSummaryItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageContentSummaryItemsInner{}

// ImageContentSummaryItemsInner struct for ImageContentSummaryItemsInner
type ImageContentSummaryItemsInner struct {
	ContentType *string `json:"content_type,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewImageContentSummaryItemsInner instantiates a new ImageContentSummaryItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageContentSummaryItemsInner() *ImageContentSummaryItemsInner {
	this := ImageContentSummaryItemsInner{}
	return &this
}

// NewImageContentSummaryItemsInnerWithDefaults instantiates a new ImageContentSummaryItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageContentSummaryItemsInnerWithDefaults() *ImageContentSummaryItemsInner {
	this := ImageContentSummaryItemsInner{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ImageContentSummaryItemsInner) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageContentSummaryItemsInner) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ImageContentSummaryItemsInner) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ImageContentSummaryItemsInner) SetContentType(v string) {
	o.ContentType = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ImageContentSummaryItemsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageContentSummaryItemsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ImageContentSummaryItemsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ImageContentSummaryItemsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ImageContentSummaryItemsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageContentSummaryItemsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ImageContentSummaryItemsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ImageContentSummaryItemsInner) SetCount(v int32) {
	o.Count = &v
}

func (o ImageContentSummaryItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageContentSummaryItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableImageContentSummaryItemsInner struct {
	value *ImageContentSummaryItemsInner
	isSet bool
}

func (v NullableImageContentSummaryItemsInner) Get() *ImageContentSummaryItemsInner {
	return v.value
}

func (v *NullableImageContentSummaryItemsInner) Set(val *ImageContentSummaryItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableImageContentSummaryItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableImageContentSummaryItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageContentSummaryItemsInner(val *ImageContentSummaryItemsInner) *NullableImageContentSummaryItemsInner {
	return &NullableImageContentSummaryItemsInner{value: val, isSet: true}
}

func (v NullableImageContentSummaryItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageContentSummaryItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


