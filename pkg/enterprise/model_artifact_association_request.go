/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ArtifactAssociationRequest Request body for an artifact to associate with an application version
type ArtifactAssociationRequest struct {
	// The type of the artifact
	ArtifactType string `json:"artifact_type"`
	// A json with key-pair values to query on
	ArtifactKeys interface{} `json:"artifact_keys"`
}

// NewArtifactAssociationRequest instantiates a new ArtifactAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactAssociationRequest(artifactType string, artifactKeys interface{}) *ArtifactAssociationRequest {
	this := ArtifactAssociationRequest{}
	this.ArtifactType = artifactType
	this.ArtifactKeys = artifactKeys
	return &this
}

// NewArtifactAssociationRequestWithDefaults instantiates a new ArtifactAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactAssociationRequestWithDefaults() *ArtifactAssociationRequest {
	this := ArtifactAssociationRequest{}
	return &this
}

// GetArtifactType returns the ArtifactType field value
func (o *ArtifactAssociationRequest) GetArtifactType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationRequest) GetArtifactTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *ArtifactAssociationRequest) SetArtifactType(v string) {
	o.ArtifactType = v
}

// GetArtifactKeys returns the ArtifactKeys field value
func (o *ArtifactAssociationRequest) GetArtifactKeys() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ArtifactKeys
}

// GetArtifactKeysOk returns a tuple with the ArtifactKeys field value
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationRequest) GetArtifactKeysOk() (interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArtifactKeys, true
}

// SetArtifactKeys sets field value
func (o *ArtifactAssociationRequest) SetArtifactKeys(v interface{}) {
	o.ArtifactKeys = v
}

func (o ArtifactAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["artifact_type"] = o.ArtifactType
	}
	if true {
		toSerialize["artifact_keys"] = o.ArtifactKeys
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactAssociationRequest struct {
	value *ArtifactAssociationRequest
	isSet bool
}

func (v NullableArtifactAssociationRequest) Get() *ArtifactAssociationRequest {
	return v.value
}

func (v *NullableArtifactAssociationRequest) Set(val *ArtifactAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactAssociationRequest(val *ArtifactAssociationRequest) *NullableArtifactAssociationRequest {
	return &NullableArtifactAssociationRequest{value: val, isSet: true}
}

func (v NullableArtifactAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


