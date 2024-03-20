/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.4.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// ImportSource struct for ImportSource
type ImportSource struct {
	Type string `json:"type"`
	Target map[string]interface{} `json:"target,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportSource ImportSource

// NewImportSource instantiates a new ImportSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSource(type_ string) *ImportSource {
	this := ImportSource{}
	this.Type = type_
	return &this
}

// NewImportSourceWithDefaults instantiates a new ImportSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSourceWithDefaults() *ImportSource {
	this := ImportSource{}
	return &this
}

// GetType returns the Type field value
func (o *ImportSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImportSource) SetType(v string) {
	o.Type = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ImportSource) GetTarget() map[string]interface{} {
	if o == nil || o.Target == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetTargetOk() (map[string]interface{}, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ImportSource) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given map[string]interface{} and assigns it to the Target field.
func (o *ImportSource) SetTarget(v map[string]interface{}) {
	o.Target = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ImportSource) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ImportSource) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ImportSource) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ImportSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ImportSource) UnmarshalJSON(bytes []byte) (err error) {
	varImportSource := _ImportSource{}

	if err = json.Unmarshal(bytes, &varImportSource); err == nil {
		*o = ImportSource(varImportSource)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "target")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportSource struct {
	value *ImportSource
	isSet bool
}

func (v NullableImportSource) Get() *ImportSource {
	return v.value
}

func (v *NullableImportSource) Set(val *ImportSource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSource(val *ImportSource) *NullableImportSource {
	return &NullableImportSource{value: val, isSet: true}
}

func (v NullableImportSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


