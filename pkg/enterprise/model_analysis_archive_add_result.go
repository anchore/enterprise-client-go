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

// checks if the AnalysisArchiveAddResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisArchiveAddResult{}

// AnalysisArchiveAddResult The result of adding a single digest to the archive
type AnalysisArchiveAddResult struct {
	// The image digest requested to be added
	Digest *string `json:"digest,omitempty"`
	// The status of the archive add operation. Typically either 'archived' or 'error'
	Status *string `json:"status,omitempty"`
	// Details on the status, e.g. the error message
	Detail *string `json:"detail,omitempty"`
}

// NewAnalysisArchiveAddResult instantiates a new AnalysisArchiveAddResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisArchiveAddResult() *AnalysisArchiveAddResult {
	this := AnalysisArchiveAddResult{}
	return &this
}

// NewAnalysisArchiveAddResultWithDefaults instantiates a new AnalysisArchiveAddResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisArchiveAddResultWithDefaults() *AnalysisArchiveAddResult {
	this := AnalysisArchiveAddResult{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *AnalysisArchiveAddResult) GetDigest() string {
	if o == nil || IsNil(o.Digest) {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveAddResult) GetDigestOk() (*string, bool) {
	if o == nil || IsNil(o.Digest) {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *AnalysisArchiveAddResult) HasDigest() bool {
	if o != nil && !IsNil(o.Digest) {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *AnalysisArchiveAddResult) SetDigest(v string) {
	o.Digest = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AnalysisArchiveAddResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveAddResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AnalysisArchiveAddResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AnalysisArchiveAddResult) SetStatus(v string) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *AnalysisArchiveAddResult) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisArchiveAddResult) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *AnalysisArchiveAddResult) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *AnalysisArchiveAddResult) SetDetail(v string) {
	o.Detail = &v
}

func (o AnalysisArchiveAddResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisArchiveAddResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Digest) {
		toSerialize["digest"] = o.Digest
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	return toSerialize, nil
}

type NullableAnalysisArchiveAddResult struct {
	value *AnalysisArchiveAddResult
	isSet bool
}

func (v NullableAnalysisArchiveAddResult) Get() *AnalysisArchiveAddResult {
	return v.value
}

func (v *NullableAnalysisArchiveAddResult) Set(val *AnalysisArchiveAddResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisArchiveAddResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisArchiveAddResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisArchiveAddResult(val *AnalysisArchiveAddResult) *NullableAnalysisArchiveAddResult {
	return &NullableAnalysisArchiveAddResult{value: val, isSet: true}
}

func (v NullableAnalysisArchiveAddResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisArchiveAddResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


