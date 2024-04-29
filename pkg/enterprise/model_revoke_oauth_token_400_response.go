/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// RevokeOauthToken400Response struct for RevokeOauthToken400Response
type RevokeOauthToken400Response struct {
	// ASCII error code from RFC6749
	Error *string `json:"error,omitempty"`
}

// NewRevokeOauthToken400Response instantiates a new RevokeOauthToken400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeOauthToken400Response() *RevokeOauthToken400Response {
	this := RevokeOauthToken400Response{}
	return &this
}

// NewRevokeOauthToken400ResponseWithDefaults instantiates a new RevokeOauthToken400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeOauthToken400ResponseWithDefaults() *RevokeOauthToken400Response {
	this := RevokeOauthToken400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RevokeOauthToken400Response) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokeOauthToken400Response) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RevokeOauthToken400Response) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *RevokeOauthToken400Response) SetError(v string) {
	o.Error = &v
}

func (o RevokeOauthToken400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableRevokeOauthToken400Response struct {
	value *RevokeOauthToken400Response
	isSet bool
}

func (v NullableRevokeOauthToken400Response) Get() *RevokeOauthToken400Response {
	return v.value
}

func (v *NullableRevokeOauthToken400Response) Set(val *RevokeOauthToken400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeOauthToken400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeOauthToken400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeOauthToken400Response(val *RevokeOauthToken400Response) *NullableRevokeOauthToken400Response {
	return &NullableRevokeOauthToken400Response{value: val, isSet: true}
}

func (v NullableRevokeOauthToken400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeOauthToken400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


