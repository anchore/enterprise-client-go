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

// SourceImportContentResponse struct for SourceImportContentResponse
type SourceImportContentResponse struct {
	Digest *string `json:"digest,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewSourceImportContentResponse instantiates a new SourceImportContentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceImportContentResponse() *SourceImportContentResponse {
	this := SourceImportContentResponse{}
	return &this
}

// NewSourceImportContentResponseWithDefaults instantiates a new SourceImportContentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceImportContentResponseWithDefaults() *SourceImportContentResponse {
	this := SourceImportContentResponse{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *SourceImportContentResponse) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportContentResponse) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *SourceImportContentResponse) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *SourceImportContentResponse) SetDigest(v string) {
	o.Digest = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SourceImportContentResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceImportContentResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SourceImportContentResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SourceImportContentResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o SourceImportContentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSourceImportContentResponse struct {
	value *SourceImportContentResponse
	isSet bool
}

func (v NullableSourceImportContentResponse) Get() *SourceImportContentResponse {
	return v.value
}

func (v *NullableSourceImportContentResponse) Set(val *SourceImportContentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceImportContentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceImportContentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceImportContentResponse(val *SourceImportContentResponse) *NullableSourceImportContentResponse {
	return &NullableSourceImportContentResponse{value: val, isSet: true}
}

func (v NullableSourceImportContentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceImportContentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


