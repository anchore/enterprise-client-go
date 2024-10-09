/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the RbacManagerIdpUserGroupPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerIdpUserGroupPost{}

// RbacManagerIdpUserGroupPost struct for RbacManagerIdpUserGroupPost
type RbacManagerIdpUserGroupPost struct {
	UserGroupUuids []string `json:"user_group_uuids,omitempty"`
}

// NewRbacManagerIdpUserGroupPost instantiates a new RbacManagerIdpUserGroupPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerIdpUserGroupPost() *RbacManagerIdpUserGroupPost {
	this := RbacManagerIdpUserGroupPost{}
	return &this
}

// NewRbacManagerIdpUserGroupPostWithDefaults instantiates a new RbacManagerIdpUserGroupPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerIdpUserGroupPostWithDefaults() *RbacManagerIdpUserGroupPost {
	this := RbacManagerIdpUserGroupPost{}
	return &this
}

// GetUserGroupUuids returns the UserGroupUuids field value if set, zero value otherwise.
func (o *RbacManagerIdpUserGroupPost) GetUserGroupUuids() []string {
	if o == nil || IsNil(o.UserGroupUuids) {
		var ret []string
		return ret
	}
	return o.UserGroupUuids
}

// GetUserGroupUuidsOk returns a tuple with the UserGroupUuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerIdpUserGroupPost) GetUserGroupUuidsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserGroupUuids) {
		return nil, false
	}
	return o.UserGroupUuids, true
}

// HasUserGroupUuids returns a boolean if a field has been set.
func (o *RbacManagerIdpUserGroupPost) HasUserGroupUuids() bool {
	if o != nil && !IsNil(o.UserGroupUuids) {
		return true
	}

	return false
}

// SetUserGroupUuids gets a reference to the given []string and assigns it to the UserGroupUuids field.
func (o *RbacManagerIdpUserGroupPost) SetUserGroupUuids(v []string) {
	o.UserGroupUuids = v
}

func (o RbacManagerIdpUserGroupPost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerIdpUserGroupPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserGroupUuids) {
		toSerialize["user_group_uuids"] = o.UserGroupUuids
	}
	return toSerialize, nil
}

type NullableRbacManagerIdpUserGroupPost struct {
	value *RbacManagerIdpUserGroupPost
	isSet bool
}

func (v NullableRbacManagerIdpUserGroupPost) Get() *RbacManagerIdpUserGroupPost {
	return v.value
}

func (v *NullableRbacManagerIdpUserGroupPost) Set(val *RbacManagerIdpUserGroupPost) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerIdpUserGroupPost) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerIdpUserGroupPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerIdpUserGroupPost(val *RbacManagerIdpUserGroupPost) *NullableRbacManagerIdpUserGroupPost {
	return &NullableRbacManagerIdpUserGroupPost{value: val, isSet: true}
}

func (v NullableRbacManagerIdpUserGroupPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerIdpUserGroupPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


