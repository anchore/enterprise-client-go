/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

<<<<<<< HEAD
API version: 2.0.0
=======
API version: 0.1.0
>>>>>>> 48fc108 (feat: updated the enterprise ref)
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
	Target *map[string]interface{} `json:"target,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
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
	if o == nil  {
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
	if o == nil || o.Target == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMSource) GetTargetOk() (*map[string]interface{}, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *NativeSBOMSource) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given map[string]interface{} and assigns it to the Target field.
func (o *NativeSBOMSource) SetTarget(v map[string]interface{}) {
	o.Target = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NativeSBOMSource) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NativeSBOMSource) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NativeSBOMSource) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *NativeSBOMSource) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

func (o NativeSBOMSource) MarshalJSON() ([]byte, error) {
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

func (o *NativeSBOMSource) UnmarshalJSON(bytes []byte) (err error) {
	varNativeSBOMSource := _NativeSBOMSource{}

	if err = json.Unmarshal(bytes, &varNativeSBOMSource); err == nil {
		*o = NativeSBOMSource(varNativeSBOMSource)
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


