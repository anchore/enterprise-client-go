/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the RbacManagerIdpUserGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerIdpUserGroup{}

// RbacManagerIdpUserGroup A user group associated with an IdP
type RbacManagerIdpUserGroup struct {
	// The UUID of the user group
	UserGroupUuid string `json:"user_group_uuid"`
	// The name of the user group
	UserGroupName *string `json:"user_group_name,omitempty"`
	// The timestamp when the user group was mapped to the IdP
	MappedOn *time.Time `json:"mapped_on,omitempty"`
}

type _RbacManagerIdpUserGroup RbacManagerIdpUserGroup

// NewRbacManagerIdpUserGroup instantiates a new RbacManagerIdpUserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerIdpUserGroup(userGroupUuid string) *RbacManagerIdpUserGroup {
	this := RbacManagerIdpUserGroup{}
	this.UserGroupUuid = userGroupUuid
	return &this
}

// NewRbacManagerIdpUserGroupWithDefaults instantiates a new RbacManagerIdpUserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerIdpUserGroupWithDefaults() *RbacManagerIdpUserGroup {
	this := RbacManagerIdpUserGroup{}
	return &this
}

// GetUserGroupUuid returns the UserGroupUuid field value
func (o *RbacManagerIdpUserGroup) GetUserGroupUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserGroupUuid
}

// GetUserGroupUuidOk returns a tuple with the UserGroupUuid field value
// and a boolean to check if the value has been set.
func (o *RbacManagerIdpUserGroup) GetUserGroupUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserGroupUuid, true
}

// SetUserGroupUuid sets field value
func (o *RbacManagerIdpUserGroup) SetUserGroupUuid(v string) {
	o.UserGroupUuid = v
}

// GetUserGroupName returns the UserGroupName field value if set, zero value otherwise.
func (o *RbacManagerIdpUserGroup) GetUserGroupName() string {
	if o == nil || IsNil(o.UserGroupName) {
		var ret string
		return ret
	}
	return *o.UserGroupName
}

// GetUserGroupNameOk returns a tuple with the UserGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerIdpUserGroup) GetUserGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserGroupName) {
		return nil, false
	}
	return o.UserGroupName, true
}

// HasUserGroupName returns a boolean if a field has been set.
func (o *RbacManagerIdpUserGroup) HasUserGroupName() bool {
	if o != nil && !IsNil(o.UserGroupName) {
		return true
	}

	return false
}

// SetUserGroupName gets a reference to the given string and assigns it to the UserGroupName field.
func (o *RbacManagerIdpUserGroup) SetUserGroupName(v string) {
	o.UserGroupName = &v
}

// GetMappedOn returns the MappedOn field value if set, zero value otherwise.
func (o *RbacManagerIdpUserGroup) GetMappedOn() time.Time {
	if o == nil || IsNil(o.MappedOn) {
		var ret time.Time
		return ret
	}
	return *o.MappedOn
}

// GetMappedOnOk returns a tuple with the MappedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerIdpUserGroup) GetMappedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MappedOn) {
		return nil, false
	}
	return o.MappedOn, true
}

// HasMappedOn returns a boolean if a field has been set.
func (o *RbacManagerIdpUserGroup) HasMappedOn() bool {
	if o != nil && !IsNil(o.MappedOn) {
		return true
	}

	return false
}

// SetMappedOn gets a reference to the given time.Time and assigns it to the MappedOn field.
func (o *RbacManagerIdpUserGroup) SetMappedOn(v time.Time) {
	o.MappedOn = &v
}

func (o RbacManagerIdpUserGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerIdpUserGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_group_uuid"] = o.UserGroupUuid
	if !IsNil(o.UserGroupName) {
		toSerialize["user_group_name"] = o.UserGroupName
	}
	if !IsNil(o.MappedOn) {
		toSerialize["mapped_on"] = o.MappedOn
	}
	return toSerialize, nil
}

func (o *RbacManagerIdpUserGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_group_uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRbacManagerIdpUserGroup := _RbacManagerIdpUserGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRbacManagerIdpUserGroup)

	if err != nil {
		return err
	}

	*o = RbacManagerIdpUserGroup(varRbacManagerIdpUserGroup)

	return err
}

type NullableRbacManagerIdpUserGroup struct {
	value *RbacManagerIdpUserGroup
	isSet bool
}

func (v NullableRbacManagerIdpUserGroup) Get() *RbacManagerIdpUserGroup {
	return v.value
}

func (v *NullableRbacManagerIdpUserGroup) Set(val *RbacManagerIdpUserGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerIdpUserGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerIdpUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerIdpUserGroup(val *RbacManagerIdpUserGroup) *NullableRbacManagerIdpUserGroup {
	return &NullableRbacManagerIdpUserGroup{value: val, isSet: true}
}

func (v NullableRbacManagerIdpUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerIdpUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


