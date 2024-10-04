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
	"bytes"
	"fmt"
)

// checks if the CorrectionFieldMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorrectionFieldMatch{}

// CorrectionFieldMatch Defines a particular field name and value to match for a Correction
type CorrectionFieldMatch struct {
	// The package field name to match
	FieldName string `json:"field_name"`
	// The package field value for the corresponding field_name above to match. If field_name corresponds to a list value, this will search the list
	FieldValue string `json:"field_value"`
}

type _CorrectionFieldMatch CorrectionFieldMatch

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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorrectionFieldMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field_name"] = o.FieldName
	toSerialize["field_value"] = o.FieldValue
	return toSerialize, nil
}

func (o *CorrectionFieldMatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field_name",
		"field_value",
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

	varCorrectionFieldMatch := _CorrectionFieldMatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCorrectionFieldMatch)

	if err != nil {
		return err
	}

	*o = CorrectionFieldMatch(varCorrectionFieldMatch)

	return err
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


