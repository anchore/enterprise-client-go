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

// CorrectionMatch Defines how a particular correction can match depending on type
type CorrectionMatch struct {
	// type of match [supports os, npm, gem, python, java, go]
	Type string `json:"type"`
	// list of field matches that are required in order for this correction to match
	FieldMatches []CorrectionFieldMatch `json:"field_matches,omitempty"`
}

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
	if o == nil || o.FieldMatches == nil {
		var ret []CorrectionFieldMatch
		return ret
	}
	return o.FieldMatches
}

// GetFieldMatchesOk returns a tuple with the FieldMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectionMatch) GetFieldMatchesOk() ([]CorrectionFieldMatch, bool) {
	if o == nil || o.FieldMatches == nil {
		return nil, false
	}
	return o.FieldMatches, true
}

// HasFieldMatches returns a boolean if a field has been set.
func (o *CorrectionMatch) HasFieldMatches() bool {
	if o != nil && o.FieldMatches != nil {
		return true
	}

	return false
}

// SetFieldMatches gets a reference to the given []CorrectionFieldMatch and assigns it to the FieldMatches field.
func (o *CorrectionMatch) SetFieldMatches(v []CorrectionFieldMatch) {
	o.FieldMatches = v
}

func (o CorrectionMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.FieldMatches != nil {
		toSerialize["field_matches"] = o.FieldMatches
	}
	return json.Marshal(toSerialize)
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


