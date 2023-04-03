/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NativeSBOMPackageLocation struct for NativeSBOMPackageLocation
type NativeSBOMPackageLocation struct {
	Path string `json:"path"`
	LayerID *string `json:"layerID,omitempty"`
}

// NewNativeSBOMPackageLocation instantiates a new NativeSBOMPackageLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMPackageLocation(path string) *NativeSBOMPackageLocation {
	this := NativeSBOMPackageLocation{}
	this.Path = path
	return &this
}

// NewNativeSBOMPackageLocationWithDefaults instantiates a new NativeSBOMPackageLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMPackageLocationWithDefaults() *NativeSBOMPackageLocation {
	this := NativeSBOMPackageLocation{}
	return &this
}

// GetPath returns the Path field value
func (o *NativeSBOMPackageLocation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageLocation) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *NativeSBOMPackageLocation) SetPath(v string) {
	o.Path = v
}

// GetLayerID returns the LayerID field value if set, zero value otherwise.
func (o *NativeSBOMPackageLocation) GetLayerID() string {
	if o == nil || o.LayerID == nil {
		var ret string
		return ret
	}
	return *o.LayerID
}

// GetLayerIDOk returns a tuple with the LayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageLocation) GetLayerIDOk() (*string, bool) {
	if o == nil || o.LayerID == nil {
		return nil, false
	}
	return o.LayerID, true
}

// HasLayerID returns a boolean if a field has been set.
func (o *NativeSBOMPackageLocation) HasLayerID() bool {
	if o != nil && o.LayerID != nil {
		return true
	}

	return false
}

// SetLayerID gets a reference to the given string and assigns it to the LayerID field.
func (o *NativeSBOMPackageLocation) SetLayerID(v string) {
	o.LayerID = &v
}

func (o NativeSBOMPackageLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.LayerID != nil {
		toSerialize["layerID"] = o.LayerID
	}
	return json.Marshal(toSerialize)
}

type NullableNativeSBOMPackageLocation struct {
	value *NativeSBOMPackageLocation
	isSet bool
}

func (v NullableNativeSBOMPackageLocation) Get() *NativeSBOMPackageLocation {
	return v.value
}

func (v *NullableNativeSBOMPackageLocation) Set(val *NativeSBOMPackageLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMPackageLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMPackageLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMPackageLocation(val *NativeSBOMPackageLocation) *NullableNativeSBOMPackageLocation {
	return &NullableNativeSBOMPackageLocation{value: val, isSet: true}
}

func (v NullableNativeSBOMPackageLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMPackageLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


