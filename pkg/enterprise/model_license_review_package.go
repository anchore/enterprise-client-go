/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.12.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the LicenseReviewPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseReviewPackage{}

// LicenseReviewPackage struct for LicenseReviewPackage
type LicenseReviewPackage struct {
	// The name of the package
	Name *string `json:"name,omitempty"`
	// The version of the Package
	Version *string `json:"version,omitempty"`
	// The type of the package
	Type *string `json:"type,omitempty"`
	// The Package URL for the package
	Purl *string `json:"purl,omitempty"`
}

// NewLicenseReviewPackage instantiates a new LicenseReviewPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseReviewPackage() *LicenseReviewPackage {
	this := LicenseReviewPackage{}
	return &this
}

// NewLicenseReviewPackageWithDefaults instantiates a new LicenseReviewPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseReviewPackageWithDefaults() *LicenseReviewPackage {
	this := LicenseReviewPackage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LicenseReviewPackage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseReviewPackage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LicenseReviewPackage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LicenseReviewPackage) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LicenseReviewPackage) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseReviewPackage) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LicenseReviewPackage) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LicenseReviewPackage) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LicenseReviewPackage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseReviewPackage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LicenseReviewPackage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LicenseReviewPackage) SetType(v string) {
	o.Type = &v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *LicenseReviewPackage) GetPurl() string {
	if o == nil || IsNil(o.Purl) {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseReviewPackage) GetPurlOk() (*string, bool) {
	if o == nil || IsNil(o.Purl) {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *LicenseReviewPackage) HasPurl() bool {
	if o != nil && !IsNil(o.Purl) {
		return true
	}

	return false
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *LicenseReviewPackage) SetPurl(v string) {
	o.Purl = &v
}

func (o LicenseReviewPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseReviewPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Purl) {
		toSerialize["purl"] = o.Purl
	}
	return toSerialize, nil
}

type NullableLicenseReviewPackage struct {
	value *LicenseReviewPackage
	isSet bool
}

func (v NullableLicenseReviewPackage) Get() *LicenseReviewPackage {
	return v.value
}

func (v *NullableLicenseReviewPackage) Set(val *LicenseReviewPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseReviewPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseReviewPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseReviewPackage(val *LicenseReviewPackage) *NullableLicenseReviewPackage {
	return &NullableLicenseReviewPackage{value: val, isSet: true}
}

func (v NullableLicenseReviewPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseReviewPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


