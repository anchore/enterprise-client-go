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

// ImageRef A reference to an image
type ImageRef struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

// NewImageRef instantiates a new ImageRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageRef(type_ string, value string) *ImageRef {
	this := ImageRef{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewImageRefWithDefaults instantiates a new ImageRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageRefWithDefaults() *ImageRef {
	this := ImageRef{}
	return &this
}

// GetType returns the Type field value
func (o *ImageRef) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImageRef) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImageRef) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ImageRef) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ImageRef) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ImageRef) SetValue(v string) {
	o.Value = v
}

func (o ImageRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableImageRef struct {
	value *ImageRef
	isSet bool
}

func (v NullableImageRef) Get() *ImageRef {
	return v.value
}

func (v *NullableImageRef) Set(val *ImageRef) {
	v.value = val
	v.isSet = true
}

func (v NullableImageRef) IsSet() bool {
	return v.isSet
}

func (v *NullableImageRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageRef(val *ImageRef) *NullableImageRef {
	return &NullableImageRef{value: val, isSet: true}
}

func (v NullableImageRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


