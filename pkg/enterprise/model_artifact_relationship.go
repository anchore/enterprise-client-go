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
	"time"
)

// checks if the ArtifactRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactRelationship{}

// ArtifactRelationship A relationship of a specific type between two SDLC artifacts (e.g. container image and source revision). This is and edge in a directed graph where edges are directional from the \"source\" to the \"target\". For example, an edge of type \"contains\" means the source artifact contains the content of the target artifact. 
type ArtifactRelationship struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Source *ArtifactReference `json:"source,omitempty"`
	Target *ArtifactReference `json:"target,omitempty"`
	RelationshipType *RelationshipType `json:"relationship_type,omitempty"`
	Comment *string `json:"comment,omitempty"`
	// User-provided metadata about the relationship
	UserMetadata interface{} `json:"user_metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ArtifactRelationship ArtifactRelationship

// NewArtifactRelationship instantiates a new ArtifactRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactRelationship() *ArtifactRelationship {
	this := ArtifactRelationship{}
	return &this
}

// NewArtifactRelationshipWithDefaults instantiates a new ArtifactRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactRelationshipWithDefaults() *ArtifactRelationship {
	this := ArtifactRelationship{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ArtifactRelationship) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRelationship) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ArtifactRelationship) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ArtifactRelationship) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ArtifactRelationship) GetSource() ArtifactReference {
	if o == nil || IsNil(o.Source) {
		var ret ArtifactReference
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRelationship) GetSourceOk() (*ArtifactReference, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ArtifactRelationship) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given ArtifactReference and assigns it to the Source field.
func (o *ArtifactRelationship) SetSource(v ArtifactReference) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ArtifactRelationship) GetTarget() ArtifactReference {
	if o == nil || IsNil(o.Target) {
		var ret ArtifactReference
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRelationship) GetTargetOk() (*ArtifactReference, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ArtifactRelationship) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given ArtifactReference and assigns it to the Target field.
func (o *ArtifactRelationship) SetTarget(v ArtifactReference) {
	o.Target = &v
}

// GetRelationshipType returns the RelationshipType field value if set, zero value otherwise.
func (o *ArtifactRelationship) GetRelationshipType() RelationshipType {
	if o == nil || IsNil(o.RelationshipType) {
		var ret RelationshipType
		return ret
	}
	return *o.RelationshipType
}

// GetRelationshipTypeOk returns a tuple with the RelationshipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRelationship) GetRelationshipTypeOk() (*RelationshipType, bool) {
	if o == nil || IsNil(o.RelationshipType) {
		return nil, false
	}
	return o.RelationshipType, true
}

// HasRelationshipType returns a boolean if a field has been set.
func (o *ArtifactRelationship) HasRelationshipType() bool {
	if o != nil && !IsNil(o.RelationshipType) {
		return true
	}

	return false
}

// SetRelationshipType gets a reference to the given RelationshipType and assigns it to the RelationshipType field.
func (o *ArtifactRelationship) SetRelationshipType(v RelationshipType) {
	o.RelationshipType = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ArtifactRelationship) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRelationship) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ArtifactRelationship) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ArtifactRelationship) SetComment(v string) {
	o.Comment = &v
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *ArtifactRelationship) GetUserMetadata() interface{} {
	if o == nil || IsNil(o.UserMetadata) {
		var ret interface{}
		return ret
	}
	return o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactRelationship) GetUserMetadataOk() (interface{}, bool) {
	if o == nil || IsNil(o.UserMetadata) {
		return nil, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *ArtifactRelationship) HasUserMetadata() bool {
	if o != nil && !IsNil(o.UserMetadata) {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given interface{} and assigns it to the UserMetadata field.
func (o *ArtifactRelationship) SetUserMetadata(v interface{}) {
	o.UserMetadata = v
}

func (o ArtifactRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.RelationshipType) {
		toSerialize["relationship_type"] = o.RelationshipType
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.UserMetadata) {
		toSerialize["user_metadata"] = o.UserMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ArtifactRelationship) UnmarshalJSON(data []byte) (err error) {
	varArtifactRelationship := _ArtifactRelationship{}

	err = json.Unmarshal(data, &varArtifactRelationship)

	if err != nil {
		return err
	}

	*o = ArtifactRelationship(varArtifactRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "source")
		delete(additionalProperties, "target")
		delete(additionalProperties, "relationship_type")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "user_metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableArtifactRelationship struct {
	value *ArtifactRelationship
	isSet bool
}

func (v NullableArtifactRelationship) Get() *ArtifactRelationship {
	return v.value
}

func (v *NullableArtifactRelationship) Set(val *ArtifactRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactRelationship(val *ArtifactRelationship) *NullableArtifactRelationship {
	return &NullableArtifactRelationship{value: val, isSet: true}
}

func (v NullableArtifactRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


