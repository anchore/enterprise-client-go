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
)

// ImageSelector A set of selection criteria to match an image by a tagged pull string based on its components, with regex support in each field
type ImageSelector struct {
	// The registry section of a pull string. e.g. with \"docker.io/anchore/anchore-engine:latest\", this is \"docker.io\"
	Registry *string `json:"registry,omitempty"`
	// The repository section of a pull string. e.g. with \"docker.io/anchore/anchore-engine:latest\", this is \"anchore/anchore-engine\"
	Repository *string `json:"repository,omitempty"`
	// The tag-only section of a pull string. e.g. with \"docker.io/anchore/anchore-engine:latest\", this is \"latest\"
	Tag *string `json:"tag,omitempty"`
}

// NewImageSelector instantiates a new ImageSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageSelector() *ImageSelector {
	this := ImageSelector{}
	return &this
}

// NewImageSelectorWithDefaults instantiates a new ImageSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageSelectorWithDefaults() *ImageSelector {
	this := ImageSelector{}
	return &this
}

// GetRegistry returns the Registry field value if set, zero value otherwise.
func (o *ImageSelector) GetRegistry() string {
	if o == nil || o.Registry == nil {
		var ret string
		return ret
	}
	return *o.Registry
}

// GetRegistryOk returns a tuple with the Registry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageSelector) GetRegistryOk() (*string, bool) {
	if o == nil || o.Registry == nil {
		return nil, false
	}
	return o.Registry, true
}

// HasRegistry returns a boolean if a field has been set.
func (o *ImageSelector) HasRegistry() bool {
	if o != nil && o.Registry != nil {
		return true
	}

	return false
}

// SetRegistry gets a reference to the given string and assigns it to the Registry field.
func (o *ImageSelector) SetRegistry(v string) {
	o.Registry = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ImageSelector) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageSelector) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ImageSelector) HasRepository() bool {
	if o != nil && o.Repository != nil {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *ImageSelector) SetRepository(v string) {
	o.Repository = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ImageSelector) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageSelector) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ImageSelector) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ImageSelector) SetTag(v string) {
	o.Tag = &v
}

func (o ImageSelector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Registry != nil {
		toSerialize["registry"] = o.Registry
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableImageSelector struct {
	value *ImageSelector
	isSet bool
}

func (v NullableImageSelector) Get() *ImageSelector {
	return v.value
}

func (v *NullableImageSelector) Set(val *ImageSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableImageSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableImageSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageSelector(val *ImageSelector) *NullableImageSelector {
	return &NullableImageSelector{value: val, isSet: true}
}

func (v NullableImageSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


