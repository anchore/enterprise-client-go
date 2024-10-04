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
	"fmt"
)

// checks if the ImportFileDigest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportFileDigest{}

// ImportFileDigest struct for ImportFileDigest
type ImportFileDigest struct {
	Algorithm string `json:"algorithm"`
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _ImportFileDigest ImportFileDigest

// NewImportFileDigest instantiates a new ImportFileDigest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportFileDigest(algorithm string, value string) *ImportFileDigest {
	this := ImportFileDigest{}
	this.Algorithm = algorithm
	this.Value = value
	return &this
}

// NewImportFileDigestWithDefaults instantiates a new ImportFileDigest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportFileDigestWithDefaults() *ImportFileDigest {
	this := ImportFileDigest{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *ImportFileDigest) GetAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *ImportFileDigest) GetAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *ImportFileDigest) SetAlgorithm(v string) {
	o.Algorithm = v
}

// GetValue returns the Value field value
func (o *ImportFileDigest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ImportFileDigest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ImportFileDigest) SetValue(v string) {
	o.Value = v
}

func (o ImportFileDigest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportFileDigest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportFileDigest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"algorithm",
		"value",
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

	varImportFileDigest := _ImportFileDigest{}

	err = json.Unmarshal(data, &varImportFileDigest)

	if err != nil {
		return err
	}

	*o = ImportFileDigest(varImportFileDigest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportFileDigest struct {
	value *ImportFileDigest
	isSet bool
}

func (v NullableImportFileDigest) Get() *ImportFileDigest {
	return v.value
}

func (v *NullableImportFileDigest) Set(val *ImportFileDigest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportFileDigest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportFileDigest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportFileDigest(val *ImportFileDigest) *NullableImportFileDigest {
	return &NullableImportFileDigest{value: val, isSet: true}
}

func (v NullableImportFileDigest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportFileDigest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


