/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.7.0
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// ImageFilter Filter for an image list by id, tag, or digest, but not both
type ImageFilter struct {
	Tag *string `json:"tag,omitempty"`
	Digest *string `json:"digest,omitempty"`
}

// NewImageFilter instantiates a new ImageFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageFilter() *ImageFilter {
	this := ImageFilter{}
	return &this
}

// NewImageFilterWithDefaults instantiates a new ImageFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageFilterWithDefaults() *ImageFilter {
	this := ImageFilter{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ImageFilter) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageFilter) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ImageFilter) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ImageFilter) SetTag(v string) {
	o.Tag = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *ImageFilter) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageFilter) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *ImageFilter) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *ImageFilter) SetDigest(v string) {
	o.Digest = &v
}

func (o ImageFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	return json.Marshal(toSerialize)
}

type NullableImageFilter struct {
	value *ImageFilter
	isSet bool
}

func (v NullableImageFilter) Get() *ImageFilter {
	return v.value
}

func (v *NullableImageFilter) Set(val *ImageFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableImageFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableImageFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageFilter(val *ImageFilter) *NullableImageFilter {
	return &NullableImageFilter{value: val, isSet: true}
}

func (v NullableImageFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


