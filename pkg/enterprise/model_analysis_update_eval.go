/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the AnalysisUpdateEval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalysisUpdateEval{}

// AnalysisUpdateEval Evaluation Results for an entity (current or last)
type AnalysisUpdateEval struct {
	AnalysisStatus *string `json:"analysis_status,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
	ImageDigest *string `json:"image_digest,omitempty"`
}

// NewAnalysisUpdateEval instantiates a new AnalysisUpdateEval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalysisUpdateEval() *AnalysisUpdateEval {
	this := AnalysisUpdateEval{}
	return &this
}

// NewAnalysisUpdateEvalWithDefaults instantiates a new AnalysisUpdateEval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalysisUpdateEvalWithDefaults() *AnalysisUpdateEval {
	this := AnalysisUpdateEval{}
	return &this
}

// GetAnalysisStatus returns the AnalysisStatus field value if set, zero value otherwise.
func (o *AnalysisUpdateEval) GetAnalysisStatus() string {
	if o == nil || IsNil(o.AnalysisStatus) {
		var ret string
		return ret
	}
	return *o.AnalysisStatus
}

// GetAnalysisStatusOk returns a tuple with the AnalysisStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateEval) GetAnalysisStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AnalysisStatus) {
		return nil, false
	}
	return o.AnalysisStatus, true
}

// HasAnalysisStatus returns a boolean if a field has been set.
func (o *AnalysisUpdateEval) HasAnalysisStatus() bool {
	if o != nil && !IsNil(o.AnalysisStatus) {
		return true
	}

	return false
}

// SetAnalysisStatus gets a reference to the given string and assigns it to the AnalysisStatus field.
func (o *AnalysisUpdateEval) SetAnalysisStatus(v string) {
	o.AnalysisStatus = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *AnalysisUpdateEval) GetAnnotations() interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateEval) GetAnnotationsOk() (interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *AnalysisUpdateEval) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given interface{} and assigns it to the Annotations field.
func (o *AnalysisUpdateEval) SetAnnotations(v interface{}) {
	o.Annotations = v
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *AnalysisUpdateEval) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisUpdateEval) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *AnalysisUpdateEval) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *AnalysisUpdateEval) SetImageDigest(v string) {
	o.ImageDigest = &v
}

func (o AnalysisUpdateEval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalysisUpdateEval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnalysisStatus) {
		toSerialize["analysis_status"] = o.AnalysisStatus
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	return toSerialize, nil
}

type NullableAnalysisUpdateEval struct {
	value *AnalysisUpdateEval
	isSet bool
}

func (v NullableAnalysisUpdateEval) Get() *AnalysisUpdateEval {
	return v.value
}

func (v *NullableAnalysisUpdateEval) Set(val *AnalysisUpdateEval) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalysisUpdateEval) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalysisUpdateEval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalysisUpdateEval(val *AnalysisUpdateEval) *NullableAnalysisUpdateEval {
	return &NullableAnalysisUpdateEval{value: val, isSet: true}
}

func (v NullableAnalysisUpdateEval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalysisUpdateEval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


