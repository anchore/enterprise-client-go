/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the UserGroupUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupUsers{}

// UserGroupUsers struct for UserGroupUsers
type UserGroupUsers struct {
	Items []UserGroupUser `json:"items,omitempty"`
}

// NewUserGroupUsers instantiates a new UserGroupUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupUsers() *UserGroupUsers {
	this := UserGroupUsers{}
	return &this
}

// NewUserGroupUsersWithDefaults instantiates a new UserGroupUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupUsersWithDefaults() *UserGroupUsers {
	this := UserGroupUsers{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UserGroupUsers) GetItems() []UserGroupUser {
	if o == nil || IsNil(o.Items) {
		var ret []UserGroupUser
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupUsers) GetItemsOk() ([]UserGroupUser, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UserGroupUsers) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserGroupUser and assigns it to the Items field.
func (o *UserGroupUsers) SetItems(v []UserGroupUser) {
	o.Items = v
}

func (o UserGroupUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableUserGroupUsers struct {
	value *UserGroupUsers
	isSet bool
}

func (v NullableUserGroupUsers) Get() *UserGroupUsers {
	return v.value
}

func (v *NullableUserGroupUsers) Set(val *UserGroupUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupUsers(val *UserGroupUsers) *NullableUserGroupUsers {
	return &NullableUserGroupUsers{value: val, isSet: true}
}

func (v NullableUserGroupUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


