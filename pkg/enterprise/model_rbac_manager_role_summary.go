/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// checks if the RbacManagerRoleSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacManagerRoleSummary{}

// RbacManagerRoleSummary struct for RbacManagerRoleSummary
type RbacManagerRoleSummary struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewRbacManagerRoleSummary instantiates a new RbacManagerRoleSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacManagerRoleSummary() *RbacManagerRoleSummary {
	this := RbacManagerRoleSummary{}
	return &this
}

// NewRbacManagerRoleSummaryWithDefaults instantiates a new RbacManagerRoleSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacManagerRoleSummaryWithDefaults() *RbacManagerRoleSummary {
	this := RbacManagerRoleSummary{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RbacManagerRoleSummary) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleSummary) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RbacManagerRoleSummary) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RbacManagerRoleSummary) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RbacManagerRoleSummary) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RbacManagerRoleSummary) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RbacManagerRoleSummary) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RbacManagerRoleSummary) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacManagerRoleSummary) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RbacManagerRoleSummary) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RbacManagerRoleSummary) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o RbacManagerRoleSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacManagerRoleSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableRbacManagerRoleSummary struct {
	value *RbacManagerRoleSummary
	isSet bool
}

func (v NullableRbacManagerRoleSummary) Get() *RbacManagerRoleSummary {
	return v.value
}

func (v *NullableRbacManagerRoleSummary) Set(val *RbacManagerRoleSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacManagerRoleSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacManagerRoleSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacManagerRoleSummary(val *RbacManagerRoleSummary) *NullableRbacManagerRoleSummary {
	return &NullableRbacManagerRoleSummary{value: val, isSet: true}
}

func (v NullableRbacManagerRoleSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacManagerRoleSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


