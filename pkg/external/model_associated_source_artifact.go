/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external

import (
	"encoding/json"
)

// AssociatedSourceArtifact Model for an associated source artifact. Composite of the source artifact and its asssociation metadata
type AssociatedSourceArtifact struct {
	ArtifactAssociationMetadata *ArtifactAssociationMetadata `json:"artifact_association_metadata,omitempty"`
	Source *Source `json:"source,omitempty"`
}

// NewAssociatedSourceArtifact instantiates a new AssociatedSourceArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedSourceArtifact() *AssociatedSourceArtifact {
	this := AssociatedSourceArtifact{}
	return &this
}

// NewAssociatedSourceArtifactWithDefaults instantiates a new AssociatedSourceArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedSourceArtifactWithDefaults() *AssociatedSourceArtifact {
	this := AssociatedSourceArtifact{}
	return &this
}

// GetArtifactAssociationMetadata returns the ArtifactAssociationMetadata field value if set, zero value otherwise.
func (o *AssociatedSourceArtifact) GetArtifactAssociationMetadata() ArtifactAssociationMetadata {
	if o == nil || o.ArtifactAssociationMetadata == nil {
		var ret ArtifactAssociationMetadata
		return ret
	}
	return *o.ArtifactAssociationMetadata
}

// GetArtifactAssociationMetadataOk returns a tuple with the ArtifactAssociationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedSourceArtifact) GetArtifactAssociationMetadataOk() (*ArtifactAssociationMetadata, bool) {
	if o == nil || o.ArtifactAssociationMetadata == nil {
		return nil, false
	}
	return o.ArtifactAssociationMetadata, true
}

// HasArtifactAssociationMetadata returns a boolean if a field has been set.
func (o *AssociatedSourceArtifact) HasArtifactAssociationMetadata() bool {
	if o != nil && o.ArtifactAssociationMetadata != nil {
		return true
	}

	return false
}

// SetArtifactAssociationMetadata gets a reference to the given ArtifactAssociationMetadata and assigns it to the ArtifactAssociationMetadata field.
func (o *AssociatedSourceArtifact) SetArtifactAssociationMetadata(v ArtifactAssociationMetadata) {
	o.ArtifactAssociationMetadata = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AssociatedSourceArtifact) GetSource() Source {
	if o == nil || o.Source == nil {
		var ret Source
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociatedSourceArtifact) GetSourceOk() (*Source, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AssociatedSourceArtifact) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given Source and assigns it to the Source field.
func (o *AssociatedSourceArtifact) SetSource(v Source) {
	o.Source = &v
}

func (o AssociatedSourceArtifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ArtifactAssociationMetadata != nil {
		toSerialize["artifact_association_metadata"] = o.ArtifactAssociationMetadata
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableAssociatedSourceArtifact struct {
	value *AssociatedSourceArtifact
	isSet bool
}

func (v NullableAssociatedSourceArtifact) Get() *AssociatedSourceArtifact {
	return v.value
}

func (v *NullableAssociatedSourceArtifact) Set(val *AssociatedSourceArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedSourceArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedSourceArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedSourceArtifact(val *AssociatedSourceArtifact) *NullableAssociatedSourceArtifact {
	return &NullableAssociatedSourceArtifact{value: val, isSet: true}
}

func (v NullableAssociatedSourceArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociatedSourceArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


