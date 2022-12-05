/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.5.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// NativeSBOMSchema struct for NativeSBOMSchema
type NativeSBOMSchema struct {
	Version string `json:"version"`
	Url string `json:"url"`
}

// NewNativeSBOMSchema instantiates a new NativeSBOMSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMSchema(version string, url string) *NativeSBOMSchema {
	this := NativeSBOMSchema{}
	this.Version = version
	this.Url = url
	return &this
}

// NewNativeSBOMSchemaWithDefaults instantiates a new NativeSBOMSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMSchemaWithDefaults() *NativeSBOMSchema {
	this := NativeSBOMSchema{}
	return &this
}

// GetVersion returns the Version field value
func (o *NativeSBOMSchema) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMSchema) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NativeSBOMSchema) SetVersion(v string) {
	o.Version = v
}

// GetUrl returns the Url field value
func (o *NativeSBOMSchema) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMSchema) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NativeSBOMSchema) SetUrl(v string) {
	o.Url = v
}

func (o NativeSBOMSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableNativeSBOMSchema struct {
	value *NativeSBOMSchema
	isSet bool
}

func (v NullableNativeSBOMSchema) Get() *NativeSBOMSchema {
	return v.value
}

func (v *NullableNativeSBOMSchema) Set(val *NativeSBOMSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMSchema(val *NativeSBOMSchema) *NullableNativeSBOMSchema {
	return &NullableNativeSBOMSchema{value: val, isSet: true}
}

func (v NullableNativeSBOMSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


