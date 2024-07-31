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

// UserGroupUsersPost struct for UserGroupUsersPost
type UserGroupUsersPost struct {
	// The list of usernames to add to the user group
	Usernames []UserGroupUsersPostUsernamesInner `json:"usernames"`
}

// NewUserGroupUsersPost instantiates a new UserGroupUsersPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupUsersPost(usernames []UserGroupUsersPostUsernamesInner) *UserGroupUsersPost {
	this := UserGroupUsersPost{}
	this.Usernames = usernames
	return &this
}

// NewUserGroupUsersPostWithDefaults instantiates a new UserGroupUsersPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupUsersPostWithDefaults() *UserGroupUsersPost {
	this := UserGroupUsersPost{}
	return &this
}

// GetUsernames returns the Usernames field value
func (o *UserGroupUsersPost) GetUsernames() []UserGroupUsersPostUsernamesInner {
	if o == nil {
		var ret []UserGroupUsersPostUsernamesInner
		return ret
	}

	return o.Usernames
}

// GetUsernamesOk returns a tuple with the Usernames field value
// and a boolean to check if the value has been set.
func (o *UserGroupUsersPost) GetUsernamesOk() ([]UserGroupUsersPostUsernamesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usernames, true
}

// SetUsernames sets field value
func (o *UserGroupUsersPost) SetUsernames(v []UserGroupUsersPostUsernamesInner) {
	o.Usernames = v
}

func (o UserGroupUsersPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["usernames"] = o.Usernames
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupUsersPost struct {
	value *UserGroupUsersPost
	isSet bool
}

func (v NullableUserGroupUsersPost) Get() *UserGroupUsersPost {
	return v.value
}

func (v *NullableUserGroupUsersPost) Set(val *UserGroupUsersPost) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupUsersPost) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupUsersPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupUsersPost(val *UserGroupUsersPost) *NullableUserGroupUsersPost {
	return &NullableUserGroupUsersPost{value: val, isSet: true}
}

func (v NullableUserGroupUsersPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupUsersPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


