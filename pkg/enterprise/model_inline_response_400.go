/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 1.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {
	// ASCII error code from RFC6749
	Error *string `json:"error,omitempty"`
}

// NewInlineResponse400 instantiates a new InlineResponse400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400WithDefaults() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse400) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse400) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse400) SetError(v string) {
	o.Error = &v
}

func (o InlineResponse400) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400 struct {
	value *InlineResponse400
	isSet bool
}

func (v NullableInlineResponse400) Get() *InlineResponse400 {
	return v.value
}

func (v *NullableInlineResponse400) Set(val *InlineResponse400) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400(val *InlineResponse400) *NullableInlineResponse400 {
	return &NullableInlineResponse400{value: val, isSet: true}
}

func (v NullableInlineResponse400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


