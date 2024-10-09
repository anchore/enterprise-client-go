/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CorrectionMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorrectionMatch{}

// CorrectionMatch Defines how a particular correction can match depending on type
type CorrectionMatch struct {
	// type of match [supports os, npm, gem, python, java, go]
	Type string `json:"type"`
	// list of field matches that are required in order for this correction to match
	FieldMatches []CorrectionFieldMatch `json:"field_matches,omitempty"`
}

type _CorrectionMatch CorrectionMatch

// NewCorrectionMatch instantiates a new CorrectionMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrectionMatch(type_ string) *CorrectionMatch {
	this := CorrectionMatch{}
	this.Type = type_
	return &this
}

// NewCorrectionMatchWithDefaults instantiates a new CorrectionMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrectionMatchWithDefaults() *CorrectionMatch {
	this := CorrectionMatch{}
	return &this
}

// GetType returns the Type field value
func (o *CorrectionMatch) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CorrectionMatch) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CorrectionMatch) SetType(v string) {
	o.Type = v
}

// GetFieldMatches returns the FieldMatches field value if set, zero value otherwise.
func (o *CorrectionMatch) GetFieldMatches() []CorrectionFieldMatch {
	if o == nil || IsNil(o.FieldMatches) {
		var ret []CorrectionFieldMatch
		return ret
	}
	return o.FieldMatches
}

// GetFieldMatchesOk returns a tuple with the FieldMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectionMatch) GetFieldMatchesOk() ([]CorrectionFieldMatch, bool) {
	if o == nil || IsNil(o.FieldMatches) {
		return nil, false
	}
	return o.FieldMatches, true
}

// HasFieldMatches returns a boolean if a field has been set.
func (o *CorrectionMatch) HasFieldMatches() bool {
	if o != nil && !IsNil(o.FieldMatches) {
		return true
	}

	return false
}

// SetFieldMatches gets a reference to the given []CorrectionFieldMatch and assigns it to the FieldMatches field.
func (o *CorrectionMatch) SetFieldMatches(v []CorrectionFieldMatch) {
	o.FieldMatches = v
}

func (o CorrectionMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorrectionMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.FieldMatches) {
		toSerialize["field_matches"] = o.FieldMatches
	}
	return toSerialize, nil
}

func (o *CorrectionMatch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varCorrectionMatch := _CorrectionMatch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCorrectionMatch)

	if err != nil {
		return err
	}

	*o = CorrectionMatch(varCorrectionMatch)

	return err
}

type NullableCorrectionMatch struct {
	value *CorrectionMatch
	isSet bool
}

func (v NullableCorrectionMatch) Get() *CorrectionMatch {
	return v.value
}

func (v *NullableCorrectionMatch) Set(val *CorrectionMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrectionMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrectionMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrectionMatch(val *CorrectionMatch) *NullableCorrectionMatch {
	return &NullableCorrectionMatch{value: val, isSet: true}
}

func (v NullableCorrectionMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrectionMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


