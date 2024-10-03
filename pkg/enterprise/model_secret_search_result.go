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

// checks if the SecretSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecretSearchResult{}

// SecretSearchResult The retrieved file entry including content (b64 encoded)
type SecretSearchResult struct {
	Path *string `json:"path,omitempty"`
	Matches []RegexContentMatch `json:"matches,omitempty"`
}

// NewSecretSearchResult instantiates a new SecretSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecretSearchResult() *SecretSearchResult {
	this := SecretSearchResult{}
	return &this
}

// NewSecretSearchResultWithDefaults instantiates a new SecretSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecretSearchResultWithDefaults() *SecretSearchResult {
	this := SecretSearchResult{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SecretSearchResult) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSearchResult) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SecretSearchResult) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SecretSearchResult) SetPath(v string) {
	o.Path = &v
}

// GetMatches returns the Matches field value if set, zero value otherwise.
func (o *SecretSearchResult) GetMatches() []RegexContentMatch {
	if o == nil || IsNil(o.Matches) {
		var ret []RegexContentMatch
		return ret
	}
	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretSearchResult) GetMatchesOk() ([]RegexContentMatch, bool) {
	if o == nil || IsNil(o.Matches) {
		return nil, false
	}
	return o.Matches, true
}

// HasMatches returns a boolean if a field has been set.
func (o *SecretSearchResult) HasMatches() bool {
	if o != nil && !IsNil(o.Matches) {
		return true
	}

	return false
}

// SetMatches gets a reference to the given []RegexContentMatch and assigns it to the Matches field.
func (o *SecretSearchResult) SetMatches(v []RegexContentMatch) {
	o.Matches = v
}

func (o SecretSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecretSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Matches) {
		toSerialize["matches"] = o.Matches
	}
	return toSerialize, nil
}

type NullableSecretSearchResult struct {
	value *SecretSearchResult
	isSet bool
}

func (v NullableSecretSearchResult) Get() *SecretSearchResult {
	return v.value
}

func (v *NullableSecretSearchResult) Set(val *SecretSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSecretSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSecretSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecretSearchResult(val *SecretSearchResult) *NullableSecretSearchResult {
	return &NullableSecretSearchResult{value: val, isSet: true}
}

func (v NullableSecretSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecretSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


