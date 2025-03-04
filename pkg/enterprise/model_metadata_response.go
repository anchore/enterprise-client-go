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

// checks if the MetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataResponse{}

// MetadataResponse Generic wrapper for metadata listings from images
type MetadataResponse struct {
	ImageDigest *string `json:"image_digest,omitempty"`
	MetadataType *string `json:"metadata_type,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewMetadataResponse instantiates a new MetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataResponse() *MetadataResponse {
	this := MetadataResponse{}
	return &this
}

// NewMetadataResponseWithDefaults instantiates a new MetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataResponseWithDefaults() *MetadataResponse {
	this := MetadataResponse{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *MetadataResponse) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataResponse) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *MetadataResponse) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *MetadataResponse) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetMetadataType returns the MetadataType field value if set, zero value otherwise.
func (o *MetadataResponse) GetMetadataType() string {
	if o == nil || IsNil(o.MetadataType) {
		var ret string
		return ret
	}
	return *o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataResponse) GetMetadataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataType) {
		return nil, false
	}
	return o.MetadataType, true
}

// HasMetadataType returns a boolean if a field has been set.
func (o *MetadataResponse) HasMetadataType() bool {
	if o != nil && !IsNil(o.MetadataType) {
		return true
	}

	return false
}

// SetMetadataType gets a reference to the given string and assigns it to the MetadataType field.
func (o *MetadataResponse) SetMetadataType(v string) {
	o.MetadataType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataResponse) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataResponse) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MetadataResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *MetadataResponse) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o MetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if !IsNil(o.MetadataType) {
		toSerialize["metadata_type"] = o.MetadataType
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableMetadataResponse struct {
	value *MetadataResponse
	isSet bool
}

func (v NullableMetadataResponse) Get() *MetadataResponse {
	return v.value
}

func (v *NullableMetadataResponse) Set(val *MetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataResponse(val *MetadataResponse) *NullableMetadataResponse {
	return &NullableMetadataResponse{value: val, isSet: true}
}

func (v NullableMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


