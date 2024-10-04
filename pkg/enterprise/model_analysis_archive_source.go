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

// checks if the AnalysisArchiveSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisArchiveSource{}

// AnalysisArchiveSource An image reference in the analysis archive for the purposes of loading analysis from the archive into th working set
type AnalysisArchiveSource struct {
	// The image digest identify the analysis. Archived analyses are based on digest, tag records are restored as analysis is restored.
	Digest string `json:"digest" validate:"regexp=^sha256:[a-fA-F0-9]{64}$"`
	AdditionalProperties map[string]interface{}
}

type _AnalysisArchiveSource AnalysisArchiveSource

// NewAnalysisArchiveSource instantiates a new AnalysisArchiveSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveSource(digest string) *AnalysisArchiveSource {
	this := AnalysisArchiveSource{}
	this.Digest = digest
	return &this
}

// NewAnalysisArchiveSourceWithDefaults instantiates a new AnalysisArchiveSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveSourceWithDefaults() *AnalysisArchiveSource {
	this := AnalysisArchiveSource{}
	return &this
}

// GetDigest returns the Digest field value
func (o *AnalysisArchiveSource) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveSource) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *AnalysisArchiveSource) SetDigest(v string) {
	o.Digest = v
}

func (o AnalysisArchiveSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisArchiveSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["digest"] = o.Digest

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnalysisArchiveSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"digest",
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

	varAnalysisArchiveSource := _AnalysisArchiveSource{}

	err = json.Unmarshal(data, &varAnalysisArchiveSource)

	if err != nil {
		return err
	}

	*o = AnalysisArchiveSource(varAnalysisArchiveSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "digest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnalysisArchiveSource struct {
	value *AnalysisArchiveSource
	isSet bool
}

func (v NullableAnalysisArchiveSource) Get() *AnalysisArchiveSource {
	return v.value
}

func (v *NullableAnalysisArchiveSource) Set(val *AnalysisArchiveSource) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveSource) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveSource(val *AnalysisArchiveSource) *NullableAnalysisArchiveSource {
	return &NullableAnalysisArchiveSource{value: val, isSet: true}
}

func (v NullableAnalysisArchiveSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


