/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"time"
)

// ImageImportContentResponse struct for ImageImportContentResponse
type ImageImportContentResponse struct {
	Digest *string `json:"digest,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewImageImportContentResponse instantiates a new ImageImportContentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImportContentResponse() *ImageImportContentResponse {
	this := ImageImportContentResponse{}
	return &this
}

// NewImageImportContentResponseWithDefaults instantiates a new ImageImportContentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportContentResponseWithDefaults() *ImageImportContentResponse {
	this := ImageImportContentResponse{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *ImageImportContentResponse) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportContentResponse) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *ImageImportContentResponse) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *ImageImportContentResponse) SetDigest(v string) {
	o.Digest = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ImageImportContentResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportContentResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ImageImportContentResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ImageImportContentResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o ImageImportContentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableImageImportContentResponse struct {
	value *ImageImportContentResponse
	isSet bool
}

func (v NullableImageImportContentResponse) Get() *ImageImportContentResponse {
	return v.value
}

func (v *NullableImageImportContentResponse) Set(val *ImageImportContentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImportContentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImportContentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImportContentResponse(val *ImageImportContentResponse) *NullableImageImportContentResponse {
	return &NullableImageImportContentResponse{value: val, isSet: true}
}

func (v NullableImageImportContentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImportContentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


