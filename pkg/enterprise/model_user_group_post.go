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

// UserGroupPost struct for UserGroupPost
type UserGroupPost struct {
	// The name of the user group
	Name string `json:"name"`
	// The description of the user group
	Description *string `json:"description,omitempty"`
}

// NewUserGroupPost instantiates a new UserGroupPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupPost(name string) *UserGroupPost {
	this := UserGroupPost{}
	this.Name = name
	return &this
}

// NewUserGroupPostWithDefaults instantiates a new UserGroupPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupPostWithDefaults() *UserGroupPost {
	this := UserGroupPost{}
	return &this
}

// GetName returns the Name field value
func (o *UserGroupPost) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroupPost) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroupPost) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserGroupPost) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupPost) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserGroupPost) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserGroupPost) SetDescription(v string) {
	o.Description = &v
}

func (o UserGroupPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupPost struct {
	value *UserGroupPost
	isSet bool
}

func (v NullableUserGroupPost) Get() *UserGroupPost {
	return v.value
}

func (v *NullableUserGroupPost) Set(val *UserGroupPost) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupPost) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupPost(val *UserGroupPost) *NullableUserGroupPost {
	return &NullableUserGroupPost{value: val, isSet: true}
}

func (v NullableUserGroupPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


