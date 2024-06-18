/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// UserGroupRoles struct for UserGroupRoles
type UserGroupRoles struct {
	// The list of accounts and all its roles which are configured in the user group
	Items []UserGroupRole `json:"items,omitempty"`
}

// NewUserGroupRoles instantiates a new UserGroupRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupRoles() *UserGroupRoles {
	this := UserGroupRoles{}
	return &this
}

// NewUserGroupRolesWithDefaults instantiates a new UserGroupRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupRolesWithDefaults() *UserGroupRoles {
	this := UserGroupRoles{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UserGroupRoles) GetItems() []UserGroupRole {
	if o == nil || o.Items == nil {
		var ret []UserGroupRole
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupRoles) GetItemsOk() ([]UserGroupRole, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UserGroupRoles) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserGroupRole and assigns it to the Items field.
func (o *UserGroupRoles) SetItems(v []UserGroupRole) {
	o.Items = v
}

func (o UserGroupRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupRoles struct {
	value *UserGroupRoles
	isSet bool
}

func (v NullableUserGroupRoles) Get() *UserGroupRoles {
	return v.value
}

func (v *NullableUserGroupRoles) Set(val *UserGroupRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupRoles(val *UserGroupRoles) *NullableUserGroupRoles {
	return &NullableUserGroupRoles{value: val, isSet: true}
}

func (v NullableUserGroupRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


