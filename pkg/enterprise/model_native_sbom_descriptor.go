/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.1.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NativeSBOMDescriptor struct for NativeSBOMDescriptor
type NativeSBOMDescriptor struct {
	Name string `json:"name"`
	Version string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOMDescriptor NativeSBOMDescriptor

// NewNativeSBOMDescriptor instantiates a new NativeSBOMDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMDescriptor(name string, version string) *NativeSBOMDescriptor {
	this := NativeSBOMDescriptor{}
	this.Name = name
	this.Version = version
	return &this
}

// NewNativeSBOMDescriptorWithDefaults instantiates a new NativeSBOMDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMDescriptorWithDefaults() *NativeSBOMDescriptor {
	this := NativeSBOMDescriptor{}
	return &this
}

// GetName returns the Name field value
func (o *NativeSBOMDescriptor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMDescriptor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NativeSBOMDescriptor) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *NativeSBOMDescriptor) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMDescriptor) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NativeSBOMDescriptor) SetVersion(v string) {
	o.Version = v
}

func (o NativeSBOMDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NativeSBOMDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	varNativeSBOMDescriptor := _NativeSBOMDescriptor{}

	if err = json.Unmarshal(bytes, &varNativeSBOMDescriptor); err == nil {
		*o = NativeSBOMDescriptor(varNativeSBOMDescriptor)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNativeSBOMDescriptor struct {
	value *NativeSBOMDescriptor
	isSet bool
}

func (v NullableNativeSBOMDescriptor) Get() *NativeSBOMDescriptor {
	return v.value
}

func (v *NullableNativeSBOMDescriptor) Set(val *NativeSBOMDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMDescriptor(val *NativeSBOMDescriptor) *NullableNativeSBOMDescriptor {
	return &NullableNativeSBOMDescriptor{value: val, isSet: true}
}

func (v NullableNativeSBOMDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


