/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.11.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ImageReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageReference{}

// ImageReference A summary of an image identity, including digest, id (if available), and any tags known to have ever been mapped to the digest
type ImageReference struct {
	// The image digest
	ImageDigest *string `json:"image_digest,omitempty"`
	// The image id if available
	ImageId *string `json:"image_id,omitempty"`
	// Timestamp, in rfc3339 format, indicating when the image state became 'analyzed' in Anchore Engine.
	AnalyzedAt *string `json:"analyzed_at,omitempty"`
	TagHistory []TagEntry `json:"tag_history,omitempty"`
}

// NewImageReference instantiates a new ImageReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageReference() *ImageReference {
	this := ImageReference{}
	return &this
}

// NewImageReferenceWithDefaults instantiates a new ImageReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageReferenceWithDefaults() *ImageReference {
	this := ImageReference{}
	return &this
}

// GetImageDigest returns the ImageDigest field value if set, zero value otherwise.
func (o *ImageReference) GetImageDigest() string {
	if o == nil || IsNil(o.ImageDigest) {
		var ret string
		return ret
	}
	return *o.ImageDigest
}

// GetImageDigestOk returns a tuple with the ImageDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageReference) GetImageDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ImageDigest) {
		return nil, false
	}
	return o.ImageDigest, true
}

// HasImageDigest returns a boolean if a field has been set.
func (o *ImageReference) HasImageDigest() bool {
	if o != nil && !IsNil(o.ImageDigest) {
		return true
	}

	return false
}

// SetImageDigest gets a reference to the given string and assigns it to the ImageDigest field.
func (o *ImageReference) SetImageDigest(v string) {
	o.ImageDigest = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ImageReference) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageReference) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageReference) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ImageReference) SetImageId(v string) {
	o.ImageId = &v
}

// GetAnalyzedAt returns the AnalyzedAt field value if set, zero value otherwise.
func (o *ImageReference) GetAnalyzedAt() string {
	if o == nil || IsNil(o.AnalyzedAt) {
		var ret string
		return ret
	}
	return *o.AnalyzedAt
}

// GetAnalyzedAtOk returns a tuple with the AnalyzedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageReference) GetAnalyzedAtOk() (*string, bool) {
	if o == nil || IsNil(o.AnalyzedAt) {
		return nil, false
	}
	return o.AnalyzedAt, true
}

// HasAnalyzedAt returns a boolean if a field has been set.
func (o *ImageReference) HasAnalyzedAt() bool {
	if o != nil && !IsNil(o.AnalyzedAt) {
		return true
	}

	return false
}

// SetAnalyzedAt gets a reference to the given string and assigns it to the AnalyzedAt field.
func (o *ImageReference) SetAnalyzedAt(v string) {
	o.AnalyzedAt = &v
}

// GetTagHistory returns the TagHistory field value if set, zero value otherwise.
func (o *ImageReference) GetTagHistory() []TagEntry {
	if o == nil || IsNil(o.TagHistory) {
		var ret []TagEntry
		return ret
	}
	return o.TagHistory
}

// GetTagHistoryOk returns a tuple with the TagHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageReference) GetTagHistoryOk() ([]TagEntry, bool) {
	if o == nil || IsNil(o.TagHistory) {
		return nil, false
	}
	return o.TagHistory, true
}

// HasTagHistory returns a boolean if a field has been set.
func (o *ImageReference) HasTagHistory() bool {
	if o != nil && !IsNil(o.TagHistory) {
		return true
	}

	return false
}

// SetTagHistory gets a reference to the given []TagEntry and assigns it to the TagHistory field.
func (o *ImageReference) SetTagHistory(v []TagEntry) {
	o.TagHistory = v
}

func (o ImageReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageDigest) {
		toSerialize["image_digest"] = o.ImageDigest
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.AnalyzedAt) {
		toSerialize["analyzed_at"] = o.AnalyzedAt
	}
	if !IsNil(o.TagHistory) {
		toSerialize["tag_history"] = o.TagHistory
	}
	return toSerialize, nil
}

type NullableImageReference struct {
	value *ImageReference
	isSet bool
}

func (v NullableImageReference) Get() *ImageReference {
	return v.value
}

func (v *NullableImageReference) Set(val *ImageReference) {
	v.value = val
	v.isSet = true
}

func (v NullableImageReference) IsSet() bool {
	return v.isSet
}

func (v *NullableImageReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageReference(val *ImageReference) *NullableImageReference {
	return &NullableImageReference{value: val, isSet: true}
}

func (v NullableImageReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


