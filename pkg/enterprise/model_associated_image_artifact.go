/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.8.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// AssociatedImageArtifact Model for an associated image artifact. Composites the artifact with the association metadata
type AssociatedImageArtifact struct {
	ArtifactAssociationMetadata *ArtifactAssociationMetadata `json:"artifact_association_metadata,omitempty"`
	Image *ImageArtifact `json:"image,omitempty"`
}

// NewAssociatedImageArtifact instantiates a new AssociatedImageArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedImageArtifact() *AssociatedImageArtifact {
	this := AssociatedImageArtifact{}
	return &this
}

// NewAssociatedImageArtifactWithDefaults instantiates a new AssociatedImageArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedImageArtifactWithDefaults() *AssociatedImageArtifact {
	this := AssociatedImageArtifact{}
	return &this
}

// GetArtifactAssociationMetadata returns the ArtifactAssociationMetadata field value if set, zero value otherwise.
func (o *AssociatedImageArtifact) GetArtifactAssociationMetadata() ArtifactAssociationMetadata {
	if o == nil || o.ArtifactAssociationMetadata == nil {
		var ret ArtifactAssociationMetadata
		return ret
	}
	return *o.ArtifactAssociationMetadata
}

// GetArtifactAssociationMetadataOk returns a tuple with the ArtifactAssociationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedImageArtifact) GetArtifactAssociationMetadataOk() (*ArtifactAssociationMetadata, bool) {
	if o == nil || o.ArtifactAssociationMetadata == nil {
		return nil, false
	}
	return o.ArtifactAssociationMetadata, true
}

// HasArtifactAssociationMetadata returns a boolean if a field has been set.
func (o *AssociatedImageArtifact) HasArtifactAssociationMetadata() bool {
	if o != nil && o.ArtifactAssociationMetadata != nil {
		return true
	}

	return false
}

// SetArtifactAssociationMetadata gets a reference to the given ArtifactAssociationMetadata and assigns it to the ArtifactAssociationMetadata field.
func (o *AssociatedImageArtifact) SetArtifactAssociationMetadata(v ArtifactAssociationMetadata) {
	o.ArtifactAssociationMetadata = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *AssociatedImageArtifact) GetImage() ImageArtifact {
	if o == nil || o.Image == nil {
		var ret ImageArtifact
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedImageArtifact) GetImageOk() (*ImageArtifact, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *AssociatedImageArtifact) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageArtifact and assigns it to the Image field.
func (o *AssociatedImageArtifact) SetImage(v ImageArtifact) {
	o.Image = &v
}

func (o AssociatedImageArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactAssociationMetadata != nil {
		toSerialize["artifact_association_metadata"] = o.ArtifactAssociationMetadata
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableAssociatedImageArtifact struct {
	value *AssociatedImageArtifact
	isSet bool
}

func (v NullableAssociatedImageArtifact) Get() *AssociatedImageArtifact {
	return v.value
}

func (v *NullableAssociatedImageArtifact) Set(val *AssociatedImageArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedImageArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedImageArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedImageArtifact(val *AssociatedImageArtifact) *NullableAssociatedImageArtifact {
	return &NullableAssociatedImageArtifact{value: val, isSet: true}
}

func (v NullableAssociatedImageArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociatedImageArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


