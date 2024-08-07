/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// UserGroupRole struct for UserGroupRole
type UserGroupRole struct {
	// The account for this role
	ForAccount string `json:"for_account"`
	Roles []UserGroupRoleRolesInner `json:"roles"`
}

// NewUserGroupRole instantiates a new UserGroupRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupRole(forAccount string, roles []UserGroupRoleRolesInner) *UserGroupRole {
	this := UserGroupRole{}
	this.ForAccount = forAccount
	this.Roles = roles
	return &this
}

// NewUserGroupRoleWithDefaults instantiates a new UserGroupRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupRoleWithDefaults() *UserGroupRole {
	this := UserGroupRole{}
	return &this
}

// GetForAccount returns the ForAccount field value
func (o *UserGroupRole) GetForAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForAccount
}

// GetForAccountOk returns a tuple with the ForAccount field value
// and a boolean to check if the value has been set.
func (o *UserGroupRole) GetForAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForAccount, true
}

// SetForAccount sets field value
func (o *UserGroupRole) SetForAccount(v string) {
	o.ForAccount = v
}

// GetRoles returns the Roles field value
func (o *UserGroupRole) GetRoles() []UserGroupRoleRolesInner {
	if o == nil {
		var ret []UserGroupRoleRolesInner
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *UserGroupRole) GetRolesOk() ([]UserGroupRoleRolesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *UserGroupRole) SetRoles(v []UserGroupRoleRolesInner) {
	o.Roles = v
}

func (o UserGroupRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["for_account"] = o.ForAccount
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupRole struct {
	value *UserGroupRole
	isSet bool
}

func (v NullableUserGroupRole) Get() *UserGroupRole {
	return v.value
}

func (v *NullableUserGroupRole) Set(val *UserGroupRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupRole(val *UserGroupRole) *NullableUserGroupRole {
	return &NullableUserGroupRole{value: val, isSet: true}
}

func (v NullableUserGroupRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


