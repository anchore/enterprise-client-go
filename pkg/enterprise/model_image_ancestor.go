/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImageAncestor An summary of an image and it's layers.
type ImageAncestor struct {
	// The digest of the image
	ImageDigest *string `json:"image_digest,omitempty"`
	Tags []string `json:"tags,omitempty"`
	// The full set of layers for this image
	Layers []string `json:"layers,omitempty"`
}

// NewImageAncestor instantiates a new ImageAncestor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAncestor() *ImageAncestor {
	this := ImageAncestor{}
	return &this
}

// NewImageAncestorWithDefaults instantiates a new ImageAncestor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAncestorWithDefaults() *ImageAncestor {
	this := ImageAncestor{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ImageAncestor) GetImageDigest() string {
	if o == nil || o.ImageDigest == nil {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAncestor) GetImageDigestOk() (*string, bool) {
	if o == nil || o.ImageDigest == nil {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ImageAncestor) HasImageDigest() bool {
	if o != nil && o.ImageDigest != nil {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ImageAncestor) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ImageAncestor) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAncestor) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ImageAncestor) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ImageAncestor) SetTags(v []string) {
	o.Tags = v
}

// GetLayers returns the Layers field value if set, zero value otherwise.
func (o *ImageAncestor) GetLayers() []string {
	if o == nil || o.Layers == nil {
		var ret []string
		return ret
	}
	return o.Layers
}

// GetLayersOk returns a tuple with the Layers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAncestor) GetLayersOk() ([]string, bool) {
	if o == nil || o.Layers == nil {
		return nil, false
	}
	return o.Layers, true
}

// HasLayers returns a boolean if a field has been set.
func (o *ImageAncestor) HasLayers() bool {
	if o != nil && o.Layers != nil {
		return true
	}

	return false
}

// SetLayers gets a reference to the given []string and assigns it to the Layers field.
func (o *ImageAncestor) SetLayers(v []string) {
	o.Layers = v
}

func (o ImageAncestor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageDigest != nil {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Layers != nil {
		toSerialize["layers"] = o.Layers
	}
	return json.Marshal(toSerialize)
}

type NullableImageAncestor struct {
	value *ImageAncestor
	isSet bool
}

func (v NullableImageAncestor) Get() *ImageAncestor {
	return v.value
}

func (v *NullableImageAncestor) Set(val *ImageAncestor) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAncestor) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAncestor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAncestor(val *ImageAncestor) *NullableImageAncestor {
	return &NullableImageAncestor{value: val, isSet: true}
}

func (v NullableImageAncestor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAncestor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


