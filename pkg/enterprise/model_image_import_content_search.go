/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.3.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImageImportContentSearch struct for ImageImportContentSearch
type ImageImportContentSearch struct {
	Location ImportPackageLocation `json:"location"`
	ContentSearches []ImportContentSearchElement `json:"content_searches"`
}

// NewImageImportContentSearch instantiates a new ImageImportContentSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImportContentSearch(location ImportPackageLocation, contentSearches []ImportContentSearchElement) *ImageImportContentSearch {
	this := ImageImportContentSearch{}
	this.Location = location
	this.ContentSearches = contentSearches
	return &this
}

// NewImageImportContentSearchWithDefaults instantiates a new ImageImportContentSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportContentSearchWithDefaults() *ImageImportContentSearch {
	this := ImageImportContentSearch{}
	return &this
}

// GetLocation returns the Location field value
func (o *ImageImportContentSearch) GetLocation() ImportPackageLocation {
	if o == nil {
		var ret ImportPackageLocation
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *ImageImportContentSearch) GetLocationOk() (*ImportPackageLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *ImageImportContentSearch) SetLocation(v ImportPackageLocation) {
	o.Location = v
}

// GetContentSearches returns the ContentSearches field value
func (o *ImageImportContentSearch) GetContentSearches() []ImportContentSearchElement {
	if o == nil {
		var ret []ImportContentSearchElement
		return ret
	}

	return o.ContentSearches
}

// GetContentSearchesOk returns a tuple with the ContentSearches field value
// and a boolean to check if the value has been set.
func (o *ImageImportContentSearch) GetContentSearchesOk() ([]ImportContentSearchElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentSearches, true
}

// SetContentSearches sets field value
func (o *ImageImportContentSearch) SetContentSearches(v []ImportContentSearchElement) {
	o.ContentSearches = v
}

func (o ImageImportContentSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["content_searches"] = o.ContentSearches
	}
	return json.Marshal(toSerialize)
}

type NullableImageImportContentSearch struct {
	value *ImageImportContentSearch
	isSet bool
}

func (v NullableImageImportContentSearch) Get() *ImageImportContentSearch {
	return v.value
}

func (v *NullableImageImportContentSearch) Set(val *ImageImportContentSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImportContentSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImportContentSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImportContentSearch(val *ImageImportContentSearch) *NullableImageImportContentSearch {
	return &NullableImageImportContentSearch{value: val, isSet: true}
}

func (v NullableImageImportContentSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImportContentSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


