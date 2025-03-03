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

// checks if the NativeSBOMSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NativeSBOMSource{}

// NativeSBOMSource struct for NativeSBOMSource
type NativeSBOMSource struct {
	Type string `json:"type"`
	Target map[string]interface{} `json:"target,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NativeSBOMSource NativeSBOMSource

// NewNativeSBOMSource instantiates a new NativeSBOMSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNativeSBOMSource(type_ string) *NativeSBOMSource {
	this := NativeSBOMSource{}
	this.Type = type_
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
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NativeSBOMSource) SetType(v string) {
	o.Type = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *NativeSBOMSource) GetTarget() map[string]interface{} {
	if o == nil || IsNil(o.Target) {
		var ret map[string]interface{}
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMSource) GetTargetOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Target) {
		return map[string]interface{}{}, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *NativeSBOMSource) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given map[string]interface{} and assigns it to the Target field.
func (o *NativeSBOMSource) SetTarget(v map[string]interface{}) {
	o.Target = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NativeSBOMSource) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMSource) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NativeSBOMSource) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *NativeSBOMSource) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o NativeSBOMSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NativeSBOMSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NativeSBOMSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varNativeSBOMSource := _NativeSBOMSource{}

	err = json.Unmarshal(data, &varNativeSBOMSource)

	if err != nil {
		return err
	}

	*o = NativeSBOMSource(varNativeSBOMSource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "target")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
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


