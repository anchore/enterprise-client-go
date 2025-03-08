/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.10.1
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
	"fmt"
)

// checks if the NativeSBOMPackageLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NativeSBOMPackageLocation{}

// NativeSBOMPackageLocation struct for NativeSBOMPackageLocation
type NativeSBOMPackageLocation struct {
	Path string `json:"path"`
	LayerID *string `json:"layerID,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOMPackageLocation NativeSBOMPackageLocation

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
	if o == nil {
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
	if o == nil || IsNil(o.LayerID) {
		var ret string
		return ret
	}
	return *o.LayerID
}

// GetLayerIDOk returns a tuple with the LayerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageLocation) GetLayerIDOk() (*string, bool) {
	if o == nil || IsNil(o.LayerID) {
		return nil, false
	}
	return o.LayerID, true
}

// HasLayerID returns a boolean if a field has been set.
func (o *NativeSBOMPackageLocation) HasLayerID() bool {
	if o != nil && !IsNil(o.LayerID) {
		return true
	}

	return false
}

// SetLayerID gets a reference to the given string and assigns it to the LayerID field.
func (o *NativeSBOMPackageLocation) SetLayerID(v string) {
	o.LayerID = &v
}

func (o NativeSBOMPackageLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NativeSBOMPackageLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	if !IsNil(o.LayerID) {
		toSerialize["layerID"] = o.LayerID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NativeSBOMPackageLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
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

	varNativeSBOMPackageLocation := _NativeSBOMPackageLocation{}

	err = json.Unmarshal(data, &varNativeSBOMPackageLocation)

	if err != nil {
		return err
	}

	*o = NativeSBOMPackageLocation(varNativeSBOMPackageLocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "path")
		delete(additionalProperties, "layerID")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


