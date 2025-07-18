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
	"fmt"
)

// checks if the ImportPackageRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportPackageRelationship{}

// ImportPackageRelationship struct for ImportPackageRelationship
type ImportPackageRelationship struct {
	Parent string `json:"parent"`
	Child string `json:"child"`
	Type string `json:"type"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportPackageRelationship ImportPackageRelationship

// NewImportPackageRelationship instantiates a new ImportPackageRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportPackageRelationship(parent string, child string, type_ string) *ImportPackageRelationship {
	this := ImportPackageRelationship{}
	this.Parent = parent
	this.Child = child
	this.Type = type_
	return &this
}

// NewImportPackageRelationshipWithDefaults instantiates a new ImportPackageRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportPackageRelationshipWithDefaults() *ImportPackageRelationship {
	this := ImportPackageRelationship{}
	return &this
}

// GetParent returns the Parent field value
func (o *ImportPackageRelationship) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *ImportPackageRelationship) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *ImportPackageRelationship) SetParent(v string) {
	o.Parent = v
}

// GetChild returns the Child field value
func (o *ImportPackageRelationship) GetChild() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Child
}

// GetChildOk returns a tuple with the Child field value
// and a boolean to check if the value has been set.
func (o *ImportPackageRelationship) GetChildOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Child, true
}

// SetChild sets field value
func (o *ImportPackageRelationship) SetChild(v string) {
	o.Child = v
}

// GetType returns the Type field value
func (o *ImportPackageRelationship) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportPackageRelationship) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImportPackageRelationship) SetType(v string) {
	o.Type = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImportPackageRelationship) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportPackageRelationship) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImportPackageRelationship) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ImportPackageRelationship) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ImportPackageRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportPackageRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parent"] = o.Parent
	toSerialize["child"] = o.Child
	toSerialize["type"] = o.Type
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportPackageRelationship) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parent",
		"child",
		"type",
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

	varImportPackageRelationship := _ImportPackageRelationship{}

	err = json.Unmarshal(data, &varImportPackageRelationship)

	if err != nil {
		return err
	}

	*o = ImportPackageRelationship(varImportPackageRelationship)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parent")
		delete(additionalProperties, "child")
		delete(additionalProperties, "type")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportPackageRelationship struct {
	value *ImportPackageRelationship
	isSet bool
}

func (v NullableImportPackageRelationship) Get() *ImportPackageRelationship {
	return v.value
}

func (v *NullableImportPackageRelationship) Set(val *ImportPackageRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableImportPackageRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableImportPackageRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportPackageRelationship(val *ImportPackageRelationship) *NullableImportPackageRelationship {
	return &NullableImportPackageRelationship{value: val, isSet: true}
}

func (v NullableImportPackageRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportPackageRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


