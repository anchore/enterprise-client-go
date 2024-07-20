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

// AnchoreErrorCode A description of an anchore error code (name, description)
type AnchoreErrorCode struct {
	// Error code name
	Name *string `json:"name,omitempty"`
	// Description of the error code
	Description *string `json:"description,omitempty"`
}

// NewAnchoreErrorCode instantiates a new AnchoreErrorCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnchoreErrorCode() *AnchoreErrorCode {
	this := AnchoreErrorCode{}
	return &this
}

// NewAnchoreErrorCodeWithDefaults instantiates a new AnchoreErrorCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnchoreErrorCodeWithDefaults() *AnchoreErrorCode {
	this := AnchoreErrorCode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AnchoreErrorCode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreErrorCode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnchoreErrorCode) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AnchoreErrorCode) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AnchoreErrorCode) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreErrorCode) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AnchoreErrorCode) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AnchoreErrorCode) SetDescription(v string) {
	o.Description = &v
}

func (o AnchoreErrorCode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAnchoreErrorCode struct {
	value *AnchoreErrorCode
	isSet bool
}

func (v NullableAnchoreErrorCode) Get() *AnchoreErrorCode {
	return v.value
}

func (v *NullableAnchoreErrorCode) Set(val *AnchoreErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableAnchoreErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableAnchoreErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnchoreErrorCode(val *AnchoreErrorCode) *NullableAnchoreErrorCode {
	return &NullableAnchoreErrorCode{value: val, isSet: true}
}

func (v NullableAnchoreErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnchoreErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


