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

// NativeSBOMSource struct for NativeSBOMSource
type NativeSBOMSource struct {
	Type string `json:"type"`
	Target interface{} `json:"target"`
}

// NewNativeSBOMSource instantiates a new NativeSBOMSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMSource(type_ string, target interface{}) *NativeSBOMSource {
	this := NativeSBOMSource{}
	this.Type = type_
	this.Target = target
	return &this
}

// NewNativeSBOMSourceWithDefaults instantiates a new NativeSBOMSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMSourceWithDefaults() *NativeSBOMSource {
	this := NativeSBOMSource{}
	return &this
}

// GetType returns the Type field value
func (o *NativeSBOMSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMSource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NativeSBOMSource) SetType(v string) {
	o.Type = v
}

// GetTarget returns the Target field value
func (o *NativeSBOMSource) GetTarget() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMSource) GetTargetOk() (*interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *NativeSBOMSource) SetTarget(v interface{}) {
	o.Target = v
}

func (o NativeSBOMSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableNativeSBOMSource struct {
	value *NativeSBOMSource
	isSet bool
}

func (v NullableNativeSBOMSource) Get() *NativeSBOMSource {
	return v.value
}

func (v *NullableNativeSBOMSource) Set(val *NativeSBOMSource) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMSource) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMSource(val *NativeSBOMSource) *NullableNativeSBOMSource {
	return &NullableNativeSBOMSource{value: val, isSet: true}
}

func (v NullableNativeSBOMSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


