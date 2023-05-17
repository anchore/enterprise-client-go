/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 0.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// RbacManagerTokenResponse An auth token for use in future requests as an Authorization header value of type 'bearer'
type RbacManagerTokenResponse struct {
	// The token content
	Token string `json:"token"`
}

// NewRbacManagerTokenResponse instantiates a new RbacManagerTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerTokenResponse(token string) *RbacManagerTokenResponse {
	this := RbacManagerTokenResponse{}
	this.Token = token
	return &this
}

// NewRbacManagerTokenResponseWithDefaults instantiates a new RbacManagerTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerTokenResponseWithDefaults() *RbacManagerTokenResponse {
	this := RbacManagerTokenResponse{}
	return &this
}

// GetToken returns the Token field value
func (o *RbacManagerTokenResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RbacManagerTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RbacManagerTokenResponse) SetToken(v string) {
	o.Token = v
}

func (o RbacManagerTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerTokenResponse struct {
	value *RbacManagerTokenResponse
	isSet bool
}

func (v NullableRbacManagerTokenResponse) Get() *RbacManagerTokenResponse {
	return v.value
}

func (v *NullableRbacManagerTokenResponse) Set(val *RbacManagerTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerTokenResponse(val *RbacManagerTokenResponse) *NullableRbacManagerTokenResponse {
	return &NullableRbacManagerTokenResponse{value: val, isSet: true}
}

func (v NullableRbacManagerTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


