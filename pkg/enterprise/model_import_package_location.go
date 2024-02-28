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

// ImportPackageLocation struct for ImportPackageLocation
type ImportPackageLocation struct {
	Path string `json:"path"`
	LayerID *string `json:"layerID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportPackageLocation ImportPackageLocation

// NewImportPackageLocation instantiates a new ImportPackageLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportPackageLocation(path string) *ImportPackageLocation {
	this := ImportPackageLocation{}
	this.Path = path
	return &this
}

// NewImportPackageLocationWithDefaults instantiates a new ImportPackageLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportPackageLocationWithDefaults() *ImportPackageLocation {
	this := ImportPackageLocation{}
	return &this
}

// GetPath returns the Path field value
func (o *ImportPackageLocation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ImportPackageLocation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ImportPackageLocation) SetPath(v string) {
	o.Path = v
}

// GetLayerID returns the LayerID field value if set, zero value otherwise.
func (o *ImportPackageLocation) GetLayerID() string {
	if o == nil || o.LayerID == nil {
		var ret string
		return ret
	}
	return *o.LayerID
}

// GetLayerIDOk returns a tuple with the LayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPackageLocation) GetLayerIDOk() (*string, bool) {
	if o == nil || o.LayerID == nil {
		return nil, false
	}
	return o.LayerID, true
}

// HasLayerID returns a boolean if a field has been set.
func (o *ImportPackageLocation) HasLayerID() bool {
	if o != nil && o.LayerID != nil {
		return true
	}

	return false
}

// SetLayerID gets a reference to the given string and assigns it to the LayerID field.
func (o *ImportPackageLocation) SetLayerID(v string) {
	o.LayerID = &v
}

func (o ImportPackageLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.LayerID != nil {
		toSerialize["layerID"] = o.LayerID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ImportPackageLocation) UnmarshalJSON(bytes []byte) (err error) {
	varImportPackageLocation := _ImportPackageLocation{}

	if err = json.Unmarshal(bytes, &varImportPackageLocation); err == nil {
		*o = ImportPackageLocation(varImportPackageLocation)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "layerID")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportPackageLocation struct {
	value *ImportPackageLocation
	isSet bool
}

func (v NullableImportPackageLocation) Get() *ImportPackageLocation {
	return v.value
}

func (v *NullableImportPackageLocation) Set(val *ImportPackageLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableImportPackageLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableImportPackageLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportPackageLocation(val *ImportPackageLocation) *NullableImportPackageLocation {
	return &NullableImportPackageLocation{value: val, isSet: true}
}

func (v NullableImportPackageLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportPackageLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


