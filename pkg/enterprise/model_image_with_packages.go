/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.6.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImageWithPackages An image record that contains packages
type ImageWithPackages struct {
	Image *ImageReference `json:"image,omitempty"`
	Packages []PackageReference `json:"packages,omitempty"`
}

// NewImageWithPackages instantiates a new ImageWithPackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageWithPackages() *ImageWithPackages {
	this := ImageWithPackages{}
	return &this
}

// NewImageWithPackagesWithDefaults instantiates a new ImageWithPackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithPackagesWithDefaults() *ImageWithPackages {
	this := ImageWithPackages{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ImageWithPackages) GetImage() ImageReference {
	if o == nil || o.Image == nil {
		var ret ImageReference
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageWithPackages) GetImageOk() (*ImageReference, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ImageWithPackages) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given ImageReference and assigns it to the Image field.
func (o *ImageWithPackages) SetImage(v ImageReference) {
	o.Image = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *ImageWithPackages) GetPackages() []PackageReference {
	if o == nil || o.Packages == nil {
		var ret []PackageReference
		return ret
	}
	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageWithPackages) GetPackagesOk() ([]PackageReference, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *ImageWithPackages) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []PackageReference and assigns it to the Packages field.
func (o *ImageWithPackages) SetPackages(v []PackageReference) {
	o.Packages = v
}

func (o ImageWithPackages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	return json.Marshal(toSerialize)
}

type NullableImageWithPackages struct {
	value *ImageWithPackages
	isSet bool
}

func (v NullableImageWithPackages) Get() *ImageWithPackages {
	return v.value
}

func (v *NullableImageWithPackages) Set(val *ImageWithPackages) {
	v.value = val
	v.isSet = true
}

func (v NullableImageWithPackages) IsSet() bool {
	return v.isSet
}

func (v *NullableImageWithPackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageWithPackages(val *ImageWithPackages) *NullableImageWithPackages {
	return &NullableImageWithPackages{value: val, isSet: true}
}

func (v NullableImageWithPackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageWithPackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


