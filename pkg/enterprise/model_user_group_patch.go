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
	"bytes"
	"fmt"
)

// checks if the UserGroupPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupPatch{}

// UserGroupPatch struct for UserGroupPatch
type UserGroupPatch struct {
	// The description of the user group
	Description string `json:"description"`
}

type _UserGroupPatch UserGroupPatch

// NewUserGroupPatch instantiates a new UserGroupPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupPatch(description string) *UserGroupPatch {
	this := UserGroupPatch{}
	this.Description = description
	return &this
}

// NewUserGroupPatchWithDefaults instantiates a new UserGroupPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupPatchWithDefaults() *UserGroupPatch {
	this := UserGroupPatch{}
	return &this
}

// GetDescription returns the Description field value
func (o *UserGroupPatch) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UserGroupPatch) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UserGroupPatch) SetDescription(v string) {
	o.Description = v
}

func (o UserGroupPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *UserGroupPatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
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

	varUserGroupPatch := _UserGroupPatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGroupPatch)

	if err != nil {
		return err
	}

	*o = UserGroupPatch(varUserGroupPatch)

	return err
}

type NullableUserGroupPatch struct {
	value *UserGroupPatch
	isSet bool
}

func (v NullableUserGroupPatch) Get() *UserGroupPatch {
	return v.value
}

func (v *NullableUserGroupPatch) Set(val *UserGroupPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupPatch(val *UserGroupPatch) *NullableUserGroupPatch {
	return &NullableUserGroupPatch{value: val, isSet: true}
}

func (v NullableUserGroupPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


