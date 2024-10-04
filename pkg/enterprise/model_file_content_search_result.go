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

// checks if the FileContentSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileContentSearchResult{}

// FileContentSearchResult The retrieved file entry including content (b64 encoded)
type FileContentSearchResult struct {
	Path *string `json:"path,omitempty"`
	Matches []RegexContentMatch `json:"matches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FileContentSearchResult FileContentSearchResult

// NewFileContentSearchResult instantiates a new FileContentSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileContentSearchResult() *FileContentSearchResult {
	this := FileContentSearchResult{}
	return &this
}

// NewFileContentSearchResultWithDefaults instantiates a new FileContentSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileContentSearchResultWithDefaults() *FileContentSearchResult {
	this := FileContentSearchResult{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *FileContentSearchResult) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileContentSearchResult) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *FileContentSearchResult) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *FileContentSearchResult) SetPath(v string) {
	o.Path = &v
}

// GetMatches returns the Matches field value if set, zero value otherwise.
func (o *FileContentSearchResult) GetMatches() []RegexContentMatch {
	if o == nil || IsNil(o.Matches) {
		var ret []RegexContentMatch
		return ret
	}
	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileContentSearchResult) GetMatchesOk() ([]RegexContentMatch, bool) {
	if o == nil || IsNil(o.Matches) {
		return nil, false
	}
	return o.Matches, true
}

// HasMatches returns a boolean if a field has been set.
func (o *FileContentSearchResult) HasMatches() bool {
	if o != nil && !IsNil(o.Matches) {
		return true
	}

	return false
}

// SetMatches gets a reference to the given []RegexContentMatch and assigns it to the Matches field.
func (o *FileContentSearchResult) SetMatches(v []RegexContentMatch) {
	o.Matches = v
}

func (o FileContentSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileContentSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Matches) {
		toSerialize["matches"] = o.Matches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FileContentSearchResult) UnmarshalJSON(data []byte) (err error) {
	varFileContentSearchResult := _FileContentSearchResult{}

	err = json.Unmarshal(data, &varFileContentSearchResult)

	if err != nil {
		return err
	}

	*o = FileContentSearchResult(varFileContentSearchResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "matches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFileContentSearchResult struct {
	value *FileContentSearchResult
	isSet bool
}

func (v NullableFileContentSearchResult) Get() *FileContentSearchResult {
	return v.value
}

func (v *NullableFileContentSearchResult) Set(val *FileContentSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFileContentSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFileContentSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileContentSearchResult(val *FileContentSearchResult) *NullableFileContentSearchResult {
	return &NullableFileContentSearchResult{value: val, isSet: true}
}

func (v NullableFileContentSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileContentSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


