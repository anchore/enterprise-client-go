/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.9.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ArtifactAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactAssociationResponse{}

// ArtifactAssociationResponse Response body for an artifact to associate with an application version. Only one artifact type (matching the value of the type field) will be populated.
type ArtifactAssociationResponse struct {
	Source *Source `json:"source,omitempty"`
	Image *ImageArtifact `json:"image,omitempty"`
	ArtifactAssociationMetadata *ArtifactAssociationMetadata `json:"artifact_association_metadata,omitempty"`
}

// NewArtifactAssociationResponse instantiates a new ArtifactAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactAssociationResponse() *ArtifactAssociationResponse {
	this := ArtifactAssociationResponse{}
	return &this
}

// NewArtifactAssociationResponseWithDefaults instantiates a new ArtifactAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactAssociationResponseWithDefaults() *ArtifactAssociationResponse {
	this := ArtifactAssociationResponse{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ArtifactAssociationResponse) GetSource() Source {
	if o == nil || IsNil(o.Source) {
		var ret Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationResponse) GetSourceOk() (*Source, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ArtifactAssociationResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source and assigns it to the Source field.
func (o *ArtifactAssociationResponse) SetSource(v Source) {
	o.Source = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ArtifactAssociationResponse) GetImage() ImageArtifact {
	if o == nil || IsNil(o.Image) {
		var ret ImageArtifact
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationResponse) GetImageOk() (*ImageArtifact, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ArtifactAssociationResponse) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageArtifact and assigns it to the Image field.
func (o *ArtifactAssociationResponse) SetImage(v ImageArtifact) {
	o.Image = &v
}

// GetArtifactAssociationMetadata returns the ArtifactAssociationMetadata field value if set, zero value otherwise.
func (o *ArtifactAssociationResponse) GetArtifactAssociationMetadata() ArtifactAssociationMetadata {
	if o == nil || IsNil(o.ArtifactAssociationMetadata) {
		var ret ArtifactAssociationMetadata
		return ret
	}
	return *o.ArtifactAssociationMetadata
}

// GetArtifactAssociationMetadataOk returns a tuple with the ArtifactAssociationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationResponse) GetArtifactAssociationMetadataOk() (*ArtifactAssociationMetadata, bool) {
	if o == nil || IsNil(o.ArtifactAssociationMetadata) {
		return nil, false
	}
	return o.ArtifactAssociationMetadata, true
}

// HasArtifactAssociationMetadata returns a boolean if a field has been set.
func (o *ArtifactAssociationResponse) HasArtifactAssociationMetadata() bool {
	if o != nil && !IsNil(o.ArtifactAssociationMetadata) {
		return true
	}

	return false
}

// SetArtifactAssociationMetadata gets a reference to the given ArtifactAssociationMetadata and assigns it to the ArtifactAssociationMetadata field.
func (o *ArtifactAssociationResponse) SetArtifactAssociationMetadata(v ArtifactAssociationMetadata) {
	o.ArtifactAssociationMetadata = &v
}

func (o ArtifactAssociationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ArtifactAssociationMetadata) {
		toSerialize["artifact_association_metadata"] = o.ArtifactAssociationMetadata
	}
	return toSerialize, nil
}

type NullableArtifactAssociationResponse struct {
	value *ArtifactAssociationResponse
	isSet bool
}

func (v NullableArtifactAssociationResponse) Get() *ArtifactAssociationResponse {
	return v.value
}

func (v *NullableArtifactAssociationResponse) Set(val *ArtifactAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactAssociationResponse(val *ArtifactAssociationResponse) *NullableArtifactAssociationResponse {
	return &NullableArtifactAssociationResponse{value: val, isSet: true}
}

func (v NullableArtifactAssociationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


