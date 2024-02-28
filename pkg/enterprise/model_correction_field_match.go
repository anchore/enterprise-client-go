/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// CorrectionFieldMatch Defines a particular field name and value to match for a Correction
type CorrectionFieldMatch struct {
	// The package field name to match
	FieldName string `json:"field_name"`
	// The package field value for the corresponding field_name above to match. If field_name corresponds to a list value, this will search the list
	FieldValue string `json:"field_value"`
}

// NewCorrectionFieldMatch instantiates a new CorrectionFieldMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrectionFieldMatch(fieldName string, fieldValue string) *CorrectionFieldMatch {
	this := CorrectionFieldMatch{}
	this.FieldName = fieldName
	this.FieldValue = fieldValue
	return &this
}

// NewCorrectionFieldMatchWithDefaults instantiates a new CorrectionFieldMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrectionFieldMatchWithDefaults() *CorrectionFieldMatch {
	this := CorrectionFieldMatch{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *CorrectionFieldMatch) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *CorrectionFieldMatch) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *CorrectionFieldMatch) SetFieldName(v string) {
	o.FieldName = v
}

// GetFieldValue returns the FieldValue field value
func (o *CorrectionFieldMatch) GetFieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value
// and a boolean to check if the value has been set.
func (o *CorrectionFieldMatch) GetFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldValue, true
}

// SetFieldValue sets field value
func (o *CorrectionFieldMatch) SetFieldValue(v string) {
	o.FieldValue = v
}

func (o CorrectionFieldMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field_name"] = o.FieldName
	}
	if true {
		toSerialize["field_value"] = o.FieldValue
	}
	return json.Marshal(toSerialize)
}

type NullableCorrectionFieldMatch struct {
	value *CorrectionFieldMatch
	isSet bool
}

func (v NullableCorrectionFieldMatch) Get() *CorrectionFieldMatch {
	return v.value
}

func (v *NullableCorrectionFieldMatch) Set(val *CorrectionFieldMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrectionFieldMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrectionFieldMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrectionFieldMatch(val *CorrectionFieldMatch) *NullableCorrectionFieldMatch {
	return &NullableCorrectionFieldMatch{value: val, isSet: true}
}

func (v NullableCorrectionFieldMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrectionFieldMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


