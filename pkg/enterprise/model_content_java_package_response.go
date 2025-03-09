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

// checks if the ContentJavaPackageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentJavaPackageResponse{}

// ContentJavaPackageResponse Java package content listings from images
type ContentJavaPackageResponse struct {
	ImageDigest *string `json:"image_digest,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Content []JavaPackageContent `json:"content,omitempty"`
}

// NewContentJavaPackageResponse instantiates a new ContentJavaPackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentJavaPackageResponse() *ContentJavaPackageResponse {
	this := ContentJavaPackageResponse{}
	return &this
}

// NewContentJavaPackageResponseWithDefaults instantiates a new ContentJavaPackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentJavaPackageResponseWithDefaults() *ContentJavaPackageResponse {
	this := ContentJavaPackageResponse{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ContentJavaPackageResponse) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentJavaPackageResponse) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ContentJavaPackageResponse) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ContentJavaPackageResponse) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ContentJavaPackageResponse) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentJavaPackageResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ContentJavaPackageResponse) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ContentJavaPackageResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ContentJavaPackageResponse) GetContent() []JavaPackageContent {
	if o == nil || IsNil(o.Content) {
		var ret []JavaPackageContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentJavaPackageResponse) GetContentOk() ([]JavaPackageContent, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ContentJavaPackageResponse) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []JavaPackageContent and assigns it to the Content field.
func (o *ContentJavaPackageResponse) SetContent(v []JavaPackageContent) {
	o.Content = v
}

func (o ContentJavaPackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentJavaPackageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableContentJavaPackageResponse struct {
	value *ContentJavaPackageResponse
	isSet bool
}

func (v NullableContentJavaPackageResponse) Get() *ContentJavaPackageResponse {
	return v.value
}

func (v *NullableContentJavaPackageResponse) Set(val *ContentJavaPackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentJavaPackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentJavaPackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentJavaPackageResponse(val *ContentJavaPackageResponse) *NullableContentJavaPackageResponse {
	return &NullableContentJavaPackageResponse{value: val, isSet: true}
}

func (v NullableContentJavaPackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentJavaPackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


