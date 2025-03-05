/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the UserGroupUsersPostUsernamesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupUsersPostUsernamesInner{}

// UserGroupUsersPostUsernamesInner struct for UserGroupUsersPostUsernamesInner
type UserGroupUsersPostUsernamesInner struct {
	// A username
	Username *string `json:"username,omitempty"`
}

// NewUserGroupUsersPostUsernamesInner instantiates a new UserGroupUsersPostUsernamesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupUsersPostUsernamesInner() *UserGroupUsersPostUsernamesInner {
	this := UserGroupUsersPostUsernamesInner{}
	return &this
}

// NewUserGroupUsersPostUsernamesInnerWithDefaults instantiates a new UserGroupUsersPostUsernamesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupUsersPostUsernamesInnerWithDefaults() *UserGroupUsersPostUsernamesInner {
	this := UserGroupUsersPostUsernamesInner{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserGroupUsersPostUsernamesInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupUsersPostUsernamesInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserGroupUsersPostUsernamesInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserGroupUsersPostUsernamesInner) SetUsername(v string) {
	o.Username = &v
}

func (o UserGroupUsersPostUsernamesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupUsersPostUsernamesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUserGroupUsersPostUsernamesInner struct {
	value *UserGroupUsersPostUsernamesInner
	isSet bool
}

func (v NullableUserGroupUsersPostUsernamesInner) Get() *UserGroupUsersPostUsernamesInner {
	return v.value
}

func (v *NullableUserGroupUsersPostUsernamesInner) Set(val *UserGroupUsersPostUsernamesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupUsersPostUsernamesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupUsersPostUsernamesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupUsersPostUsernamesInner(val *UserGroupUsersPostUsernamesInner) *NullableUserGroupUsersPostUsernamesInner {
	return &NullableUserGroupUsersPostUsernamesInner{value: val, isSet: true}
}

func (v NullableUserGroupUsersPostUsernamesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupUsersPostUsernamesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


