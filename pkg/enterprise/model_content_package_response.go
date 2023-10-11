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

// checks if the ContentPackageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentPackageResponse{}

// ContentPackageResponse Package content listings from images
type ContentPackageResponse struct {
	ImageDigest *string `json:"image_digest,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	Content []ContentPackageResponseContentInner `json:"content,omitempty"`
}

// NewContentPackageResponse instantiates a new ContentPackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentPackageResponse() *ContentPackageResponse {
	this := ContentPackageResponse{}
	return &this
}

// NewContentPackageResponseWithDefaults instantiates a new ContentPackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentPackageResponseWithDefaults() *ContentPackageResponse {
	this := ContentPackageResponse{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ContentPackageResponse) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentPackageResponse) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ContentPackageResponse) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ContentPackageResponse) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *ContentPackageResponse) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentPackageResponse) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *ContentPackageResponse) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *ContentPackageResponse) SetContentType(v string) {
	o.ContentType = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ContentPackageResponse) GetContent() []ContentPackageResponseContentInner {
	if o == nil || IsNil(o.Content) {
		var ret []ContentPackageResponseContentInner
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentPackageResponse) GetContentOk() ([]ContentPackageResponseContentInner, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ContentPackageResponse) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given []ContentPackageResponseContentInner and assigns it to the Content field.
func (o *ContentPackageResponse) SetContent(v []ContentPackageResponseContentInner) {
	o.Content = v
}

func (o ContentPackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentPackageResponse) ToMap() (map[string]interface{}, error) {
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

type NullableContentPackageResponse struct {
	value *ContentPackageResponse
	isSet bool
}

func (v NullableContentPackageResponse) Get() *ContentPackageResponse {
	return v.value
}

func (v *NullableContentPackageResponse) Set(val *ContentPackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentPackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentPackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentPackageResponse(val *ContentPackageResponse) *NullableContentPackageResponse {
	return &NullableContentPackageResponse{value: val, isSet: true}
}

func (v NullableContentPackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentPackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


