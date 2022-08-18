/*
Anchore Engine API Server

This is the Anchore Engine API. Provides the primary external API for users of the service.

API version: 0.1.21
Contact: nurmi@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"encoding/json"
)

// PaginatedImageListAllOf struct for PaginatedImageListAllOf
type PaginatedImageListAllOf struct {
	Images *[]ImageWithPackages `json:"images,omitempty"`
}

// NewPaginatedImageListAllOf instantiates a new PaginatedImageListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedImageListAllOf() *PaginatedImageListAllOf {
	this := PaginatedImageListAllOf{}
	return &this
}

// NewPaginatedImageListAllOfWithDefaults instantiates a new PaginatedImageListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedImageListAllOfWithDefaults() *PaginatedImageListAllOf {
	this := PaginatedImageListAllOf{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *PaginatedImageListAllOf) GetImages() []ImageWithPackages {
	if o == nil || o.Images == nil {
		var ret []ImageWithPackages
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedImageListAllOf) GetImagesOk() (*[]ImageWithPackages, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *PaginatedImageListAllOf) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ImageWithPackages and assigns it to the Images field.
func (o *PaginatedImageListAllOf) SetImages(v []ImageWithPackages) {
	o.Images = &v
}

func (o PaginatedImageListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedImageListAllOf struct {
	value *PaginatedImageListAllOf
	isSet bool
}

func (v NullablePaginatedImageListAllOf) Get() *PaginatedImageListAllOf {
	return v.value
}

func (v *NullablePaginatedImageListAllOf) Set(val *PaginatedImageListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedImageListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedImageListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedImageListAllOf(val *PaginatedImageListAllOf) *NullablePaginatedImageListAllOf {
	return &NullablePaginatedImageListAllOf{value: val, isSet: true}
}

func (v NullablePaginatedImageListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedImageListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


