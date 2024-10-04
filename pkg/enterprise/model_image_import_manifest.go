/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.7.2
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ImageImportManifest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageImportManifest{}

// ImageImportManifest struct for ImageImportManifest
type ImageImportManifest struct {
	Contents ImportContentDigests `json:"contents"`
	Tags []string `json:"tags"`
	Digest string `json:"digest"`
	// The digest of the image's manifest-list parent if it was accessed from a multi-arch tag where the tag pointed to a manifest-list. This allows preservation of that relationship in the data
	ParentDigest *string `json:"parent_digest,omitempty"`
	// An \"image_id\" as used by Docker if available
	LocalImageId *string `json:"local_image_id,omitempty"`
	OperationUuid string `json:"operation_uuid"`
}

type _ImageImportManifest ImageImportManifest

// NewImageImportManifest instantiates a new ImageImportManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImportManifest(contents ImportContentDigests, tags []string, digest string, operationUuid string) *ImageImportManifest {
	this := ImageImportManifest{}
	this.Contents = contents
	this.Tags = tags
	this.Digest = digest
	this.OperationUuid = operationUuid
	return &this
}

// NewImageImportManifestWithDefaults instantiates a new ImageImportManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportManifestWithDefaults() *ImageImportManifest {
	this := ImageImportManifest{}
	return &this
}

// GetContents returns the Contents field value
func (o *ImageImportManifest) GetContents() ImportContentDigests {
	if o == nil {
		var ret ImportContentDigests
		return ret
	}

	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value
// and a boolean to check if the value has been set.
func (o *ImageImportManifest) GetContentsOk() (*ImportContentDigests, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contents, true
}

// SetContents sets field value
func (o *ImageImportManifest) SetContents(v ImportContentDigests) {
	o.Contents = v
}

// GetTags returns the Tags field value
func (o *ImageImportManifest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ImageImportManifest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ImageImportManifest) SetTags(v []string) {
	o.Tags = v
}

// GetDigest returns the Digest field value
func (o *ImageImportManifest) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *ImageImportManifest) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *ImageImportManifest) SetDigest(v string) {
	o.Digest = v
}

// GetParentDigest returns the ParentDigest field value if set, zero value otherwise.
func (o *ImageImportManifest) GetParentDigest() string {
	if o == nil || IsNil(o.ParentDigest) {
		var ret string
		return ret
	}
	return *o.ParentDigest
}

// GetParentDigestOk returns a tuple with the ParentDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportManifest) GetParentDigestOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDigest) {
		return nil, false
	}
	return o.ParentDigest, true
}

// HasParentDigest returns a boolean if a field has been set.
func (o *ImageImportManifest) HasParentDigest() bool {
	if o != nil && !IsNil(o.ParentDigest) {
		return true
	}

	return false
}

// SetParentDigest gets a reference to the given string and assigns it to the ParentDigest field.
func (o *ImageImportManifest) SetParentDigest(v string) {
	o.ParentDigest = &v
}

// GetLocalImageId returns the LocalImageId field value if set, zero value otherwise.
func (o *ImageImportManifest) GetLocalImageId() string {
	if o == nil || IsNil(o.LocalImageId) {
		var ret string
		return ret
	}
	return *o.LocalImageId
}

// GetLocalImageIdOk returns a tuple with the LocalImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportManifest) GetLocalImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocalImageId) {
		return nil, false
	}
	return o.LocalImageId, true
}

// HasLocalImageId returns a boolean if a field has been set.
func (o *ImageImportManifest) HasLocalImageId() bool {
	if o != nil && !IsNil(o.LocalImageId) {
		return true
	}

	return false
}

// SetLocalImageId gets a reference to the given string and assigns it to the LocalImageId field.
func (o *ImageImportManifest) SetLocalImageId(v string) {
	o.LocalImageId = &v
}

// GetOperationUuid returns the OperationUuid field value
func (o *ImageImportManifest) GetOperationUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationUuid
}

// GetOperationUuidOk returns a tuple with the OperationUuid field value
// and a boolean to check if the value has been set.
func (o *ImageImportManifest) GetOperationUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationUuid, true
}

// SetOperationUuid sets field value
func (o *ImageImportManifest) SetOperationUuid(v string) {
	o.OperationUuid = v
}

func (o ImageImportManifest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageImportManifest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contents"] = o.Contents
	toSerialize["tags"] = o.Tags
	toSerialize["digest"] = o.Digest
	if !IsNil(o.ParentDigest) {
		toSerialize["parent_digest"] = o.ParentDigest
	}
	if !IsNil(o.LocalImageId) {
		toSerialize["local_image_id"] = o.LocalImageId
	}
	toSerialize["operation_uuid"] = o.OperationUuid
	return toSerialize, nil
}

func (o *ImageImportManifest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contents",
		"tags",
		"digest",
		"operation_uuid",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varImageImportManifest := _ImageImportManifest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImageImportManifest)

	if err != nil {
		return err
	}

	*o = ImageImportManifest(varImageImportManifest)

	return err
}

type NullableImageImportManifest struct {
	value *ImageImportManifest
	isSet bool
}

func (v NullableImageImportManifest) Get() *ImageImportManifest {
	return v.value
}

func (v *NullableImageImportManifest) Set(val *ImageImportManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImportManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImportManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImportManifest(val *ImageImportManifest) *NullableImageImportManifest {
	return &NullableImageImportManifest{value: val, isSet: true}
}

func (v NullableImageImportManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImportManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


