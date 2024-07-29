/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// RbacManagerRoleMember A mapping between a username and a role with an account context
type RbacManagerRoleMember struct {
	Username string `json:"username"`
	ForAccount string `json:"for_account"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewRbacManagerRoleMember instantiates a new RbacManagerRoleMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerRoleMember(username string, forAccount string) *RbacManagerRoleMember {
	this := RbacManagerRoleMember{}
	this.Username = username
	this.ForAccount = forAccount
	return &this
}

// NewRbacManagerRoleMemberWithDefaults instantiates a new RbacManagerRoleMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerRoleMemberWithDefaults() *RbacManagerRoleMember {
	this := RbacManagerRoleMember{}
	return &this
}

// GetUsername returns the Username field value
func (o *RbacManagerRoleMember) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMember) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *RbacManagerRoleMember) SetUsername(v string) {
	o.Username = v
}

// GetForAccount returns the ForAccount field value
func (o *RbacManagerRoleMember) GetForAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMember) GetForAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForAccount, true
}

// SetForAccount sets field value
func (o *RbacManagerRoleMember) SetForAccount(v string) {
	o.ForAccount = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerRoleMember) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleMember) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerRoleMember) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerRoleMember) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o RbacManagerRoleMember) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["for_account"] = o.ForAccount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerRoleMember struct {
	value *RbacManagerRoleMember
	isSet bool
}

func (v NullableRbacManagerRoleMember) Get() *RbacManagerRoleMember {
	return v.value
}

func (v *NullableRbacManagerRoleMember) Set(val *RbacManagerRoleMember) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerRoleMember) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerRoleMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerRoleMember(val *RbacManagerRoleMember) *NullableRbacManagerRoleMember {
	return &NullableRbacManagerRoleMember{value: val, isSet: true}
}

func (v NullableRbacManagerRoleMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerRoleMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


