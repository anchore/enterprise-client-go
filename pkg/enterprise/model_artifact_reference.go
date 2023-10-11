/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ArtifactReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactReference{}

// ArtifactReference struct for ArtifactReference
type ArtifactReference struct {
	ArtifactId *string `json:"artifact_id,omitempty"`
	ArtifactType *ArtifactType `json:"artifact_type,omitempty"`
}

// NewArtifactReference instantiates a new ArtifactReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactReference() *ArtifactReference {
	this := ArtifactReference{}
	return &this
}

// NewArtifactReferenceWithDefaults instantiates a new ArtifactReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactReferenceWithDefaults() *ArtifactReference {
	this := ArtifactReference{}
	return &this
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *ArtifactReference) GetArtifactId() string {
	if o == nil || IsNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetArtifactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactId) {
		return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *ArtifactReference) HasArtifactId() bool {
	if o != nil && !IsNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *ArtifactReference) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetArtifactType returns the ArtifactType field value if set, zero value otherwise.
func (o *ArtifactReference) GetArtifactType() ArtifactType {
	if o == nil || IsNil(o.ArtifactType) {
		var ret ArtifactType
		return ret
	}
	return *o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetArtifactTypeOk() (*ArtifactType, bool) {
	if o == nil || IsNil(o.ArtifactType) {
		return nil, false
	}
	return o.ArtifactType, true
}

// HasArtifactType returns a boolean if a field has been set.
func (o *ArtifactReference) HasArtifactType() bool {
	if o != nil && !IsNil(o.ArtifactType) {
		return true
	}

	return false
}

// SetArtifactType gets a reference to the given ArtifactType and assigns it to the ArtifactType field.
func (o *ArtifactReference) SetArtifactType(v ArtifactType) {
	o.ArtifactType = &v
}

func (o ArtifactReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArtifactId) {
		toSerialize["artifact_id"] = o.ArtifactId
	}
	if !IsNil(o.ArtifactType) {
		toSerialize["artifact_type"] = o.ArtifactType
	}
	return toSerialize, nil
}

type NullableArtifactReference struct {
	value *ArtifactReference
	isSet bool
}

func (v NullableArtifactReference) Get() *ArtifactReference {
	return v.value
}

func (v *NullableArtifactReference) Set(val *ArtifactReference) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactReference) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactReference(val *ArtifactReference) *NullableArtifactReference {
	return &NullableArtifactReference{value: val, isSet: true}
}

func (v NullableArtifactReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


