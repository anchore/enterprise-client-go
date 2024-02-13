/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.2.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImportDescriptor struct for ImportDescriptor
type ImportDescriptor struct {
	Name string `json:"name"`
	Version string `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _ImportDescriptor ImportDescriptor

// NewImportDescriptor instantiates a new ImportDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportDescriptor(name string, version string) *ImportDescriptor {
	this := ImportDescriptor{}
	this.Name = name
	this.Version = version
	return &this
}

// NewImportDescriptorWithDefaults instantiates a new ImportDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportDescriptorWithDefaults() *ImportDescriptor {
	this := ImportDescriptor{}
	return &this
}

// GetName returns the Name field value
func (o *ImportDescriptor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportDescriptor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImportDescriptor) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *ImportDescriptor) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ImportDescriptor) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ImportDescriptor) SetVersion(v string) {
	o.Version = v
}

func (o ImportDescriptor) MarshalJSON() ([]byte, error) {
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

func (o *ImportDescriptor) UnmarshalJSON(bytes []byte) (err error) {
	varImportDescriptor := _ImportDescriptor{}

	if err = json.Unmarshal(bytes, &varImportDescriptor); err == nil {
		*o = ImportDescriptor(varImportDescriptor)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportDescriptor struct {
	value *ImportDescriptor
	isSet bool
}

func (v NullableImportDescriptor) Get() *ImportDescriptor {
	return v.value
}

func (v *NullableImportDescriptor) Set(val *ImportDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableImportDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableImportDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportDescriptor(val *ImportDescriptor) *NullableImportDescriptor {
	return &NullableImportDescriptor{value: val, isSet: true}
}

func (v NullableImportDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


