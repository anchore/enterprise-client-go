/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ContentFilesResponse File content listings from images
type ContentFilesResponse struct {
	ImageDigest *string `json:"image_digest,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Content *[]ContentFilesResponseContent `json:"content,omitempty"`
}

// NewContentFilesResponse instantiates a new ContentFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentFilesResponse() *ContentFilesResponse {
	this := ContentFilesResponse{}
	return &this
}

// NewContentFilesResponseWithDefaults instantiates a new ContentFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentFilesResponseWithDefaults() *ContentFilesResponse {
	this := ContentFilesResponse{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ContentFilesResponse) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentFilesResponse) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ContentFilesResponse) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ContentFilesResponse) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ContentFilesResponse) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentFilesResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ContentFilesResponse) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ContentFilesResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ContentFilesResponse) GetContent() []ContentFilesResponseContent {
	if o == nil || o.Content == nil {
		var ret []ContentFilesResponseContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentFilesResponse) GetContentOk() (*[]ContentFilesResponseContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ContentFilesResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []ContentFilesResponseContent and assigns it to the Content field.
func (o *ContentFilesResponse) SetContent(v []ContentFilesResponseContent) {
	o.Content = &v
}

func (o ContentFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableContentFilesResponse struct {
	value *ContentFilesResponse
	isSet bool
}

func (v NullableContentFilesResponse) Get() *ContentFilesResponse {
	return v.value
}

func (v *NullableContentFilesResponse) Set(val *ContentFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentFilesResponse(val *ContentFilesResponse) *NullableContentFilesResponse {
	return &NullableContentFilesResponse{value: val, isSet: true}
}

func (v NullableContentFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


