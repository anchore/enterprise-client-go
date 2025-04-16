/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
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

// checks if the RbacManagerRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerRole{}

// RbacManagerRole Role definition
type RbacManagerRole struct {
	// The name of the role
	Name string `json:"name"`
	// A role description for humans
	Description *string `json:"description,omitempty"`
	Permissions []RbacManagerPermission `json:"permissions,omitempty"`
	// Are the permissions of this role modifiable by users (including admin users)
	Immutable *bool `json:"immutable,omitempty"`
	// The timestamp when the role was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The timestamp of the last update to the role metadata itself
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

type _RbacManagerRole RbacManagerRole

// NewRbacManagerRole instantiates a new RbacManagerRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerRole(name string) *RbacManagerRole {
	this := RbacManagerRole{}
	this.Name = name
	return &this
}

// NewRbacManagerRoleWithDefaults instantiates a new RbacManagerRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerRoleWithDefaults() *RbacManagerRole {
	this := RbacManagerRole{}
	return &this
}

// GetName returns the Name field value
func (o *RbacManagerRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RbacManagerRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RbacManagerRole) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RbacManagerRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RbacManagerRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RbacManagerRole) SetDescription(v string) {
	o.Description = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RbacManagerRole) GetPermissions() []RbacManagerPermission {
	if o == nil || IsNil(o.Permissions) {
		var ret []RbacManagerPermission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRole) GetPermissionsOk() ([]RbacManagerPermission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RbacManagerRole) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []RbacManagerPermission and assigns it to the Permissions field.
func (o *RbacManagerRole) SetPermissions(v []RbacManagerPermission) {
	o.Permissions = v
}

// GetImmutable returns the Immutable field value if set, zero value otherwise.
func (o *RbacManagerRole) GetImmutable() bool {
	if o == nil || IsNil(o.Immutable) {
		var ret bool
		return ret
	}
	return *o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRole) GetImmutableOk() (*bool, bool) {
	if o == nil || IsNil(o.Immutable) {
		return nil, false
	}
	return o.Immutable, true
}

// HasImmutable returns a boolean if a field has been set.
func (o *RbacManagerRole) HasImmutable() bool {
	if o != nil && !IsNil(o.Immutable) {
		return true
	}

	return false
}

// SetImmutable gets a reference to the given bool and assigns it to the Immutable field.
func (o *RbacManagerRole) SetImmutable(v bool) {
	o.Immutable = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerRole) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRole) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerRole) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerRole) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RbacManagerRole) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRole) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RbacManagerRole) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RbacManagerRole) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o RbacManagerRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.Immutable) {
		toSerialize["immutable"] = o.Immutable
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return toSerialize, nil
}

func (o *RbacManagerRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varRbacManagerRole := _RbacManagerRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRbacManagerRole)

	if err != nil {
		return err
	}

	*o = RbacManagerRole(varRbacManagerRole)

	return err
}

type NullableRbacManagerRole struct {
	value *RbacManagerRole
	isSet bool
}

func (v NullableRbacManagerRole) Get() *RbacManagerRole {
	return v.value
}

func (v *NullableRbacManagerRole) Set(val *RbacManagerRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerRole(val *RbacManagerRole) *NullableRbacManagerRole {
	return &NullableRbacManagerRole{value: val, isSet: true}
}

func (v NullableRbacManagerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


