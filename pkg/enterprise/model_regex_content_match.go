/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// RegexContentMatch Match of a named regex on a file
type RegexContentMatch struct {
	// The name associated with the regular expression
	Name *string `json:"name,omitempty"`
	// The regular expression used for the match
	Regex *string `json:"regex,omitempty"`
	// A list of line numbers in the file that matched the regex
	Lines []int32 `json:"lines,omitempty"`
}

// NewRegexContentMatch instantiates a new RegexContentMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegexContentMatch() *RegexContentMatch {
	this := RegexContentMatch{}
	return &this
}

// NewRegexContentMatchWithDefaults instantiates a new RegexContentMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegexContentMatchWithDefaults() *RegexContentMatch {
	this := RegexContentMatch{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegexContentMatch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexContentMatch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegexContentMatch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegexContentMatch) SetName(v string) {
	o.Name = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *RegexContentMatch) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexContentMatch) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *RegexContentMatch) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *RegexContentMatch) SetRegex(v string) {
	o.Regex = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *RegexContentMatch) GetLines() []int32 {
	if o == nil || o.Lines == nil {
		var ret []int32
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegexContentMatch) GetLinesOk() ([]int32, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *RegexContentMatch) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given []int32 and assigns it to the Lines field.
func (o *RegexContentMatch) SetLines(v []int32) {
	o.Lines = v
}

func (o RegexContentMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	return json.Marshal(toSerialize)
}

type NullableRegexContentMatch struct {
	value *RegexContentMatch
	isSet bool
}

func (v NullableRegexContentMatch) Get() *RegexContentMatch {
	return v.value
}

func (v *NullableRegexContentMatch) Set(val *RegexContentMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableRegexContentMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableRegexContentMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegexContentMatch(val *RegexContentMatch) *NullableRegexContentMatch {
	return &NullableRegexContentMatch{value: val, isSet: true}
}

func (v NullableRegexContentMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegexContentMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


