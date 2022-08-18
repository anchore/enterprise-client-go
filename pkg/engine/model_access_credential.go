/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.21
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// AccessCredential A login credential mapped to a user identity. For password credentials, the username to present for Basic auth is the user's username from the user record
type AccessCredential struct {
	// The type of credential
	Type string `json:"type"`
	// The credential value (e.g. the password)
	Value string `json:"value"`
	// The timestamp of creation of the credential
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewAccessCredential instantiates a new AccessCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessCredential(type_ string, value string) *AccessCredential {
	this := AccessCredential{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAccessCredentialWithDefaults instantiates a new AccessCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessCredentialWithDefaults() *AccessCredential {
	this := AccessCredential{}
	return &this
}

// GetType returns the Type field value
func (o *AccessCredential) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccessCredential) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccessCredential) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *AccessCredential) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AccessCredential) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AccessCredential) SetValue(v string) {
	o.Value = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccessCredential) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessCredential) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccessCredential) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AccessCredential) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o AccessCredential) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAccessCredential struct {
	value *AccessCredential
	isSet bool
}

func (v NullableAccessCredential) Get() *AccessCredential {
	return v.value
}

func (v *NullableAccessCredential) Set(val *AccessCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessCredential(val *AccessCredential) *NullableAccessCredential {
	return &NullableAccessCredential{value: val, isSet: true}
}

func (v NullableAccessCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


