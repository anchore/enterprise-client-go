/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

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
	if o == nil || o.ArtifactId == nil {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetArtifactIdOk() (*string, bool) {
	if o == nil || o.ArtifactId == nil {
		return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *ArtifactReference) HasArtifactId() bool {
	if o != nil && o.ArtifactId != nil {
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
	if o == nil || o.ArtifactType == nil {
		var ret ArtifactType
		return ret
	}
	return *o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactReference) GetArtifactTypeOk() (*ArtifactType, bool) {
	if o == nil || o.ArtifactType == nil {
		return nil, false
	}
	return o.ArtifactType, true
}

// HasArtifactType returns a boolean if a field has been set.
func (o *ArtifactReference) HasArtifactType() bool {
	if o != nil && o.ArtifactType != nil {
		return true
	}

	return false
}

// SetArtifactType gets a reference to the given ArtifactType and assigns it to the ArtifactType field.
func (o *ArtifactReference) SetArtifactType(v ArtifactType) {
	o.ArtifactType = &v
}

func (o ArtifactReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactId != nil {
		toSerialize["artifact_id"] = o.ArtifactId
	}
	if o.ArtifactType != nil {
		toSerialize["artifact_type"] = o.ArtifactType
	}
	return json.Marshal(toSerialize)
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


