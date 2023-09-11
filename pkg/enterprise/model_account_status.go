/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// AccountStatus A summary of account status
type AccountStatus struct {
	// The status of the account
	State *string `json:"state,omitempty"`
}

// NewAccountStatus instantiates a new AccountStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountStatus() *AccountStatus {
	this := AccountStatus{}
	return &this
}

// NewAccountStatusWithDefaults instantiates a new AccountStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountStatusWithDefaults() *AccountStatus {
	this := AccountStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AccountStatus) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountStatus) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AccountStatus) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AccountStatus) SetState(v string) {
	o.State = &v
}

func (o AccountStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableAccountStatus struct {
	value *AccountStatus
	isSet bool
}

func (v NullableAccountStatus) Get() *AccountStatus {
	return v.value
}

func (v *NullableAccountStatus) Set(val *AccountStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountStatus(val *AccountStatus) *NullableAccountStatus {
	return &NullableAccountStatus{value: val, isSet: true}
}

func (v NullableAccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


