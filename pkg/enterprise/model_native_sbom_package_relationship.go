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

// NativeSBOMPackageRelationship struct for NativeSBOMPackageRelationship
type NativeSBOMPackageRelationship struct {
	Parent string `json:"parent"`
	Child string `json:"child"`
	Type string `json:"type"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOMPackageRelationship NativeSBOMPackageRelationship

// NewNativeSBOMPackageRelationship instantiates a new NativeSBOMPackageRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMPackageRelationship(parent string, child string, type_ string) *NativeSBOMPackageRelationship {
	this := NativeSBOMPackageRelationship{}
	this.Parent = parent
	this.Child = child
	this.Type = type_
	return &this
}

// NewNativeSBOMPackageRelationshipWithDefaults instantiates a new NativeSBOMPackageRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNativeSBOMPackageRelationshipWithDefaults() *NativeSBOMPackageRelationship {
	this := NativeSBOMPackageRelationship{}
	return &this
}

// GetParent returns the Parent field value
func (o *NativeSBOMPackageRelationship) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageRelationship) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *NativeSBOMPackageRelationship) SetParent(v string) {
	o.Parent = v
}

// GetChild returns the Child field value
func (o *NativeSBOMPackageRelationship) GetChild() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Child
}

// GetChildOk returns a tuple with the Child field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageRelationship) GetChildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Child, true
}

// SetChild sets field value
func (o *NativeSBOMPackageRelationship) SetChild(v string) {
	o.Child = v
}

// GetType returns the Type field value
func (o *NativeSBOMPackageRelationship) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageRelationship) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NativeSBOMPackageRelationship) SetType(v string) {
	o.Type = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NativeSBOMPackageRelationship) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMPackageRelationship) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NativeSBOMPackageRelationship) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *NativeSBOMPackageRelationship) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o NativeSBOMPackageRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parent"] = o.Parent
	}
	if true {
		toSerialize["child"] = o.Child
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NativeSBOMPackageRelationship) UnmarshalJSON(bytes []byte) (err error) {
	varNativeSBOMPackageRelationship := _NativeSBOMPackageRelationship{}

	if err = json.Unmarshal(bytes, &varNativeSBOMPackageRelationship); err == nil {
		*o = NativeSBOMPackageRelationship(varNativeSBOMPackageRelationship)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "parent")
		delete(additionalProperties, "child")
		delete(additionalProperties, "type")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNativeSBOMPackageRelationship struct {
	value *NativeSBOMPackageRelationship
	isSet bool
}

func (v NullableNativeSBOMPackageRelationship) Get() *NativeSBOMPackageRelationship {
	return v.value
}

func (v *NullableNativeSBOMPackageRelationship) Set(val *NativeSBOMPackageRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableNativeSBOMPackageRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableNativeSBOMPackageRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNativeSBOMPackageRelationship(val *NativeSBOMPackageRelationship) *NullableNativeSBOMPackageRelationship {
	return &NullableNativeSBOMPackageRelationship{value: val, isSet: true}
}

func (v NullableNativeSBOMPackageRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNativeSBOMPackageRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


