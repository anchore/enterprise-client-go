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
)

// checks if the RbacManagerPermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerPermission{}

// RbacManagerPermission A grant of specific action against a specific scope and target
type RbacManagerPermission struct {
	// The allowed action. e.g. getImage
	Action *string `json:"action,omitempty"`
	// The target to which the action may be applied. Either a '*' for all or a specific target id
	Target *string `json:"target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RbacManagerPermission RbacManagerPermission

// NewRbacManagerPermission instantiates a new RbacManagerPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerPermission() *RbacManagerPermission {
	this := RbacManagerPermission{}
	return &this
}

// NewRbacManagerPermissionWithDefaults instantiates a new RbacManagerPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerPermissionWithDefaults() *RbacManagerPermission {
	this := RbacManagerPermission{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RbacManagerPermission) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerPermission) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RbacManagerPermission) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RbacManagerPermission) SetAction(v string) {
	o.Action = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *RbacManagerPermission) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerPermission) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *RbacManagerPermission) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *RbacManagerPermission) SetTarget(v string) {
	o.Target = &v
}

func (o RbacManagerPermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerPermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RbacManagerPermission) UnmarshalJSON(data []byte) (err error) {
	varRbacManagerPermission := _RbacManagerPermission{}

	err = json.Unmarshal(data, &varRbacManagerPermission)

	if err != nil {
		return err
	}

	*o = RbacManagerPermission(varRbacManagerPermission)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "target")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRbacManagerPermission struct {
	value *RbacManagerPermission
	isSet bool
}

func (v NullableRbacManagerPermission) Get() *RbacManagerPermission {
	return v.value
}

func (v *NullableRbacManagerPermission) Set(val *RbacManagerPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerPermission(val *RbacManagerPermission) *NullableRbacManagerPermission {
	return &NullableRbacManagerPermission{value: val, isSet: true}
}

func (v NullableRbacManagerPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


