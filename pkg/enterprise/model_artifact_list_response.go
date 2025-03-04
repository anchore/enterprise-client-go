/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ArtifactListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactListResponse{}

// ArtifactListResponse The response provided when querying for the artifacts on an application version
type ArtifactListResponse struct {
	AssociatedSourceArtifacts []AssociatedSourceArtifact `json:"associated_source_artifacts,omitempty"`
	AssociatedImageArtifacts []AssociatedImageArtifact `json:"associated_image_artifacts,omitempty"`
}

// NewArtifactListResponse instantiates a new ArtifactListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactListResponse() *ArtifactListResponse {
	this := ArtifactListResponse{}
	return &this
}

// NewArtifactListResponseWithDefaults instantiates a new ArtifactListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactListResponseWithDefaults() *ArtifactListResponse {
	this := ArtifactListResponse{}
	return &this
}

// GetAssociatedSourceArtifacts returns the AssociatedSourceArtifacts field value if set, zero value otherwise.
func (o *ArtifactListResponse) GetAssociatedSourceArtifacts() []AssociatedSourceArtifact {
	if o == nil || IsNil(o.AssociatedSourceArtifacts) {
		var ret []AssociatedSourceArtifact
		return ret
	}
	return o.AssociatedSourceArtifacts
}

// GetAssociatedSourceArtifactsOk returns a tuple with the AssociatedSourceArtifacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactListResponse) GetAssociatedSourceArtifactsOk() ([]AssociatedSourceArtifact, bool) {
	if o == nil || IsNil(o.AssociatedSourceArtifacts) {
		return nil, false
	}
	return o.AssociatedSourceArtifacts, true
}

// HasAssociatedSourceArtifacts returns a boolean if a field has been set.
func (o *ArtifactListResponse) HasAssociatedSourceArtifacts() bool {
	if o != nil && !IsNil(o.AssociatedSourceArtifacts) {
		return true
	}

	return false
}

// SetAssociatedSourceArtifacts gets a reference to the given []AssociatedSourceArtifact and assigns it to the AssociatedSourceArtifacts field.
func (o *ArtifactListResponse) SetAssociatedSourceArtifacts(v []AssociatedSourceArtifact) {
	o.AssociatedSourceArtifacts = v
}

// GetAssociatedImageArtifacts returns the AssociatedImageArtifacts field value if set, zero value otherwise.
func (o *ArtifactListResponse) GetAssociatedImageArtifacts() []AssociatedImageArtifact {
	if o == nil || IsNil(o.AssociatedImageArtifacts) {
		var ret []AssociatedImageArtifact
		return ret
	}
	return o.AssociatedImageArtifacts
}

// GetAssociatedImageArtifactsOk returns a tuple with the AssociatedImageArtifacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactListResponse) GetAssociatedImageArtifactsOk() ([]AssociatedImageArtifact, bool) {
	if o == nil || IsNil(o.AssociatedImageArtifacts) {
		return nil, false
	}
	return o.AssociatedImageArtifacts, true
}

// HasAssociatedImageArtifacts returns a boolean if a field has been set.
func (o *ArtifactListResponse) HasAssociatedImageArtifacts() bool {
	if o != nil && !IsNil(o.AssociatedImageArtifacts) {
		return true
	}

	return false
}

// SetAssociatedImageArtifacts gets a reference to the given []AssociatedImageArtifact and assigns it to the AssociatedImageArtifacts field.
func (o *ArtifactListResponse) SetAssociatedImageArtifacts(v []AssociatedImageArtifact) {
	o.AssociatedImageArtifacts = v
}

func (o ArtifactListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedSourceArtifacts) {
		toSerialize["associated_source_artifacts"] = o.AssociatedSourceArtifacts
	}
	if !IsNil(o.AssociatedImageArtifacts) {
		toSerialize["associated_image_artifacts"] = o.AssociatedImageArtifacts
	}
	return toSerialize, nil
}

type NullableArtifactListResponse struct {
	value *ArtifactListResponse
	isSet bool
}

func (v NullableArtifactListResponse) Get() *ArtifactListResponse {
	return v.value
}

func (v *NullableArtifactListResponse) Set(val *ArtifactListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactListResponse(val *ArtifactListResponse) *NullableArtifactListResponse {
	return &NullableArtifactListResponse{value: val, isSet: true}
}

func (v NullableArtifactListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


