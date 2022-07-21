/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ArtifactAssociationMetadata Metadata for an artifact association to an application version
type ArtifactAssociationMetadata struct {
	// The id of the association between the application version and the artifact
	AssociationId *string `json:"association_id,omitempty"`
	// RFC 3339 formatted UTC timestamp when the artifact was associated with the application version
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// RFC 3339 formatted UTC timestamp when the artifact association was last updated
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

// NewArtifactAssociationMetadata instantiates a new ArtifactAssociationMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactAssociationMetadata() *ArtifactAssociationMetadata {
	this := ArtifactAssociationMetadata{}
	return &this
}

// NewArtifactAssociationMetadataWithDefaults instantiates a new ArtifactAssociationMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactAssociationMetadataWithDefaults() *ArtifactAssociationMetadata {
	this := ArtifactAssociationMetadata{}
	return &this
}

// GetAssociationId returns the AssociationId field value if set, zero value otherwise.
func (o *ArtifactAssociationMetadata) GetAssociationId() string {
	if o == nil || o.AssociationId == nil {
		var ret string
		return ret
	}
	return *o.AssociationId
}

// GetAssociationIdOk returns a tuple with the AssociationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationMetadata) GetAssociationIdOk() (*string, bool) {
	if o == nil || o.AssociationId == nil {
		return nil, false
	}
	return o.AssociationId, true
}

// HasAssociationId returns a boolean if a field has been set.
func (o *ArtifactAssociationMetadata) HasAssociationId() bool {
	if o != nil && o.AssociationId != nil {
		return true
	}

	return false
}

// SetAssociationId gets a reference to the given string and assigns it to the AssociationId field.
func (o *ArtifactAssociationMetadata) SetAssociationId(v string) {
	o.AssociationId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArtifactAssociationMetadata) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationMetadata) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArtifactAssociationMetadata) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArtifactAssociationMetadata) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ArtifactAssociationMetadata) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactAssociationMetadata) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ArtifactAssociationMetadata) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ArtifactAssociationMetadata) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o ArtifactAssociationMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssociationId != nil {
		toSerialize["association_id"] = o.AssociationId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUpdated != nil {
		toSerialize["last_updated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactAssociationMetadata struct {
	value *ArtifactAssociationMetadata
	isSet bool
}

func (v NullableArtifactAssociationMetadata) Get() *ArtifactAssociationMetadata {
	return v.value
}

func (v *NullableArtifactAssociationMetadata) Set(val *ArtifactAssociationMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactAssociationMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactAssociationMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactAssociationMetadata(val *ArtifactAssociationMetadata) *NullableArtifactAssociationMetadata {
	return &NullableArtifactAssociationMetadata{value: val, isSet: true}
}

func (v NullableArtifactAssociationMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactAssociationMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

