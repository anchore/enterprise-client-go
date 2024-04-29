/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImageImportFileCoordinate struct for ImageImportFileCoordinate
type ImageImportFileCoordinate struct {
	// The path on the filesystem of the file within the given layer
	Path *string `json:"path,omitempty"`
	// The image layer in which the file was found
	LayerID *string `json:"layerID,omitempty"`
}

// NewImageImportFileCoordinate instantiates a new ImageImportFileCoordinate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageImportFileCoordinate() *ImageImportFileCoordinate {
	this := ImageImportFileCoordinate{}
	return &this
}

// NewImageImportFileCoordinateWithDefaults instantiates a new ImageImportFileCoordinate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageImportFileCoordinateWithDefaults() *ImageImportFileCoordinate {
	this := ImageImportFileCoordinate{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ImageImportFileCoordinate) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportFileCoordinate) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ImageImportFileCoordinate) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ImageImportFileCoordinate) SetPath(v string) {
	o.Path = &v
}

// GetLayerID returns the LayerID field value if set, zero value otherwise.
func (o *ImageImportFileCoordinate) GetLayerID() string {
	if o == nil || o.LayerID == nil {
		var ret string
		return ret
	}
	return *o.LayerID
}

// GetLayerIDOk returns a tuple with the LayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageImportFileCoordinate) GetLayerIDOk() (*string, bool) {
	if o == nil || o.LayerID == nil {
		return nil, false
	}
	return o.LayerID, true
}

// HasLayerID returns a boolean if a field has been set.
func (o *ImageImportFileCoordinate) HasLayerID() bool {
	if o != nil && o.LayerID != nil {
		return true
	}

	return false
}

// SetLayerID gets a reference to the given string and assigns it to the LayerID field.
func (o *ImageImportFileCoordinate) SetLayerID(v string) {
	o.LayerID = &v
}

func (o ImageImportFileCoordinate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.LayerID != nil {
		toSerialize["layerID"] = o.LayerID
	}
	return json.Marshal(toSerialize)
}

type NullableImageImportFileCoordinate struct {
	value *ImageImportFileCoordinate
	isSet bool
}

func (v NullableImageImportFileCoordinate) Get() *ImageImportFileCoordinate {
	return v.value
}

func (v *NullableImageImportFileCoordinate) Set(val *ImageImportFileCoordinate) {
	v.value = val
	v.isSet = true
}

func (v NullableImageImportFileCoordinate) IsSet() bool {
	return v.isSet
}

func (v *NullableImageImportFileCoordinate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageImportFileCoordinate(val *ImageImportFileCoordinate) *NullableImageImportFileCoordinate {
	return &NullableImageImportFileCoordinate{value: val, isSet: true}
}

func (v NullableImageImportFileCoordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageImportFileCoordinate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


