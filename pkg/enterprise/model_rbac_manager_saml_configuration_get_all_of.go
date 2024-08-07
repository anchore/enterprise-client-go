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

// RbacManagerSamlConfigurationGetAllOf struct for RbacManagerSamlConfigurationGetAllOf
type RbacManagerSamlConfigurationGetAllOf struct {
	// List of user groups associated with this IDP (Only for GET operations)
	UserGroups []RbacManagerSamlConfigurationGetAllOfUserGroups `json:"user_groups,omitempty"`
}

// NewRbacManagerSamlConfigurationGetAllOf instantiates a new RbacManagerSamlConfigurationGetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerSamlConfigurationGetAllOf() *RbacManagerSamlConfigurationGetAllOf {
	this := RbacManagerSamlConfigurationGetAllOf{}
	return &this
}

// NewRbacManagerSamlConfigurationGetAllOfWithDefaults instantiates a new RbacManagerSamlConfigurationGetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerSamlConfigurationGetAllOfWithDefaults() *RbacManagerSamlConfigurationGetAllOf {
	this := RbacManagerSamlConfigurationGetAllOf{}
	return &this
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *RbacManagerSamlConfigurationGetAllOf) GetUserGroups() []RbacManagerSamlConfigurationGetAllOfUserGroups {
	if o == nil || o.UserGroups == nil {
		var ret []RbacManagerSamlConfigurationGetAllOfUserGroups
		return ret
	}
	return o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerSamlConfigurationGetAllOf) GetUserGroupsOk() ([]RbacManagerSamlConfigurationGetAllOfUserGroups, bool) {
	if o == nil || o.UserGroups == nil {
		return nil, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *RbacManagerSamlConfigurationGetAllOf) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []RbacManagerSamlConfigurationGetAllOfUserGroups and assigns it to the UserGroups field.
func (o *RbacManagerSamlConfigurationGetAllOf) SetUserGroups(v []RbacManagerSamlConfigurationGetAllOfUserGroups) {
	o.UserGroups = v
}

func (o RbacManagerSamlConfigurationGetAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserGroups != nil {
		toSerialize["user_groups"] = o.UserGroups
	}
	return json.Marshal(toSerialize)
}

type NullableRbacManagerSamlConfigurationGetAllOf struct {
	value *RbacManagerSamlConfigurationGetAllOf
	isSet bool
}

func (v NullableRbacManagerSamlConfigurationGetAllOf) Get() *RbacManagerSamlConfigurationGetAllOf {
	return v.value
}

func (v *NullableRbacManagerSamlConfigurationGetAllOf) Set(val *RbacManagerSamlConfigurationGetAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerSamlConfigurationGetAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerSamlConfigurationGetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerSamlConfigurationGetAllOf(val *RbacManagerSamlConfigurationGetAllOf) *NullableRbacManagerSamlConfigurationGetAllOf {
	return &NullableRbacManagerSamlConfigurationGetAllOf{value: val, isSet: true}
}

func (v NullableRbacManagerSamlConfigurationGetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerSamlConfigurationGetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


