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

// checks if the ImageContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageContent{}

// ImageContent A metadata content record for a specific image, containing different content type entries
type ImageContent struct {
	Metadata *ImageContentMetadata `json:"metadata,omitempty"`
}

// NewImageContent instantiates a new ImageContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageContent() *ImageContent {
	this := ImageContent{}
	return &this
}

// NewImageContentWithDefaults instantiates a new ImageContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageContentWithDefaults() *ImageContent {
	this := ImageContent{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImageContent) GetMetadata() ImageContentMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ImageContentMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageContent) GetMetadataOk() (*ImageContentMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImageContent) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ImageContentMetadata and assigns it to the Metadata field.
func (o *ImageContent) SetMetadata(v ImageContentMetadata) {
	o.Metadata = &v
}

func (o ImageContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableImageContent struct {
	value *ImageContent
	isSet bool
}

func (v NullableImageContent) Get() *ImageContent {
	return v.value
}

func (v *NullableImageContent) Set(val *ImageContent) {
	v.value = val
	v.isSet = true
}

func (v NullableImageContent) IsSet() bool {
	return v.isSet
}

func (v *NullableImageContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageContent(val *ImageContent) *NullableImageContent {
	return &NullableImageContent{value: val, isSet: true}
}

func (v NullableImageContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


