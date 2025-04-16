/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RbacManagerTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerTokenResponse{}

// RbacManagerTokenResponse An auth token for use in future requests as an Authorization header value of type 'bearer'
type RbacManagerTokenResponse struct {
	// The token content
	Token string `json:"token"`
}

type _RbacManagerTokenResponse RbacManagerTokenResponse

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
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RbacManagerTokenResponse) SetToken(v string) {
	o.Token = v
}

func (o RbacManagerTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *RbacManagerTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRbacManagerTokenResponse := _RbacManagerTokenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRbacManagerTokenResponse)

	if err != nil {
		return err
	}

	*o = RbacManagerTokenResponse(varRbacManagerTokenResponse)

	return err
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


