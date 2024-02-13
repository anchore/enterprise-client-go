/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// LocalAnalysisSource struct for LocalAnalysisSource
type LocalAnalysisSource struct {
	Digest *string `json:"digest,omitempty"`
}

// NewLocalAnalysisSource instantiates a new LocalAnalysisSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalAnalysisSource() *LocalAnalysisSource {
	this := LocalAnalysisSource{}
	return &this
}

// NewLocalAnalysisSourceWithDefaults instantiates a new LocalAnalysisSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalAnalysisSourceWithDefaults() *LocalAnalysisSource {
	this := LocalAnalysisSource{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *LocalAnalysisSource) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalAnalysisSource) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *LocalAnalysisSource) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *LocalAnalysisSource) SetDigest(v string) {
	o.Digest = &v
}

func (o LocalAnalysisSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	return json.Marshal(toSerialize)
}

type NullableLocalAnalysisSource struct {
	value *LocalAnalysisSource
	isSet bool
}

func (v NullableLocalAnalysisSource) Get() *LocalAnalysisSource {
	return v.value
}

func (v *NullableLocalAnalysisSource) Set(val *LocalAnalysisSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalAnalysisSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalAnalysisSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalAnalysisSource(val *LocalAnalysisSource) *NullableLocalAnalysisSource {
	return &NullableLocalAnalysisSource{value: val, isSet: true}
}

func (v NullableLocalAnalysisSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalAnalysisSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


