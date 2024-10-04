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

// checks if the AnchoreErrorCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnchoreErrorCode{}

// AnchoreErrorCode A description of an anchore error code (name, description)
type AnchoreErrorCode struct {
	// Error code name
	Name *string `json:"name,omitempty"`
	// Description of the error code
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnchoreErrorCode AnchoreErrorCode

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
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreErrorCode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AnchoreErrorCode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnchoreErrorCode) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AnchoreErrorCode) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AnchoreErrorCode) SetDescription(v string) {
	o.Description = &v
}

func (o AnchoreErrorCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnchoreErrorCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnchoreErrorCode) UnmarshalJSON(data []byte) (err error) {
	varAnchoreErrorCode := _AnchoreErrorCode{}

	err = json.Unmarshal(data, &varAnchoreErrorCode)

	if err != nil {
		return err
	}

	*o = AnchoreErrorCode(varAnchoreErrorCode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


