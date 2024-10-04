/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
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

// checks if the UserGroupUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupUser{}

// UserGroupUser struct for UserGroupUser
type UserGroupUser struct {
	// The username of the user
	Username string `json:"username"`
	// The timestamp of when the user was added to the group
	AddedOn *time.Time `json:"added_on,omitempty"`
}

type _UserGroupUser UserGroupUser

// NewUserGroupUser instantiates a new UserGroupUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupUser(username string) *UserGroupUser {
	this := UserGroupUser{}
	this.Username = username
	return &this
}

// NewUserGroupUserWithDefaults instantiates a new UserGroupUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupUserWithDefaults() *UserGroupUser {
	this := UserGroupUser{}
	return &this
}

// GetUsername returns the Username field value
func (o *UserGroupUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserGroupUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserGroupUser) SetUsername(v string) {
	o.Username = v
}

// GetAddedOn returns the AddedOn field value if set, zero value otherwise.
func (o *UserGroupUser) GetAddedOn() time.Time {
	if o == nil || IsNil(o.AddedOn) {
		var ret time.Time
		return ret
	}
	return *o.AddedOn
}

// GetAddedOnOk returns a tuple with the AddedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupUser) GetAddedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddedOn) {
		return nil, false
	}
	return o.AddedOn, true
}

// HasAddedOn returns a boolean if a field has been set.
func (o *UserGroupUser) HasAddedOn() bool {
	if o != nil && !IsNil(o.AddedOn) {
		return true
	}

	return false
}

// SetAddedOn gets a reference to the given time.Time and assigns it to the AddedOn field.
func (o *UserGroupUser) SetAddedOn(v time.Time) {
	o.AddedOn = &v
}

func (o UserGroupUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if !IsNil(o.AddedOn) {
		toSerialize["added_on"] = o.AddedOn
	}
	return toSerialize, nil
}

func (o *UserGroupUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varUserGroupUser := _UserGroupUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroupUser)

	if err != nil {
		return err
	}

	*o = UserGroupUser(varUserGroupUser)

	return err
}

type NullableUserGroupUser struct {
	value *UserGroupUser
	isSet bool
}

func (v NullableUserGroupUser) Get() *UserGroupUser {
	return v.value
}

func (v *NullableUserGroupUser) Set(val *UserGroupUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupUser(val *UserGroupUser) *NullableUserGroupUser {
	return &NullableUserGroupUser{value: val, isSet: true}
}

func (v NullableUserGroupUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


