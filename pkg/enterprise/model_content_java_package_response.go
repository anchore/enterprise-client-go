/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ContentJAVAPackageResponse Java package content listings from images
type ContentJAVAPackageResponse struct {
	ImageDigest *string `json:"imageDigest,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Content *[]ContentJAVAPackageResponseContent `json:"content,omitempty"`
}

// NewContentJAVAPackageResponse instantiates a new ContentJAVAPackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentJAVAPackageResponse() *ContentJAVAPackageResponse {
	this := ContentJAVAPackageResponse{}
	return &this
}

// NewContentJAVAPackageResponseWithDefaults instantiates a new ContentJAVAPackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentJAVAPackageResponseWithDefaults() *ContentJAVAPackageResponse {
	this := ContentJAVAPackageResponse{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ContentJAVAPackageResponse) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentJAVAPackageResponse) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ContentJAVAPackageResponse) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ContentJAVAPackageResponse) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ContentJAVAPackageResponse) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentJAVAPackageResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ContentJAVAPackageResponse) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ContentJAVAPackageResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ContentJAVAPackageResponse) GetContent() []ContentJAVAPackageResponseContent {
	if o == nil || o.Content == nil {
		var ret []ContentJAVAPackageResponseContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentJAVAPackageResponse) GetContentOk() (*[]ContentJAVAPackageResponseContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ContentJAVAPackageResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []ContentJAVAPackageResponseContent and assigns it to the Content field.
func (o *ContentJAVAPackageResponse) SetContent(v []ContentJAVAPackageResponseContent) {
	o.Content = &v
}

func (o ContentJAVAPackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageDigest != nil {
		toSerialize["imageDigest"] = o.ImageDigest
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableContentJAVAPackageResponse struct {
	value *ContentJAVAPackageResponse
	isSet bool
}

func (v NullableContentJAVAPackageResponse) Get() *ContentJAVAPackageResponse {
	return v.value
}

func (v *NullableContentJAVAPackageResponse) Set(val *ContentJAVAPackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentJAVAPackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentJAVAPackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentJAVAPackageResponse(val *ContentJAVAPackageResponse) *NullableContentJAVAPackageResponse {
	return &NullableContentJAVAPackageResponse{value: val, isSet: true}
}

func (v NullableContentJAVAPackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentJAVAPackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


