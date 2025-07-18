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
	"bytes"
	"fmt"
)

// checks if the SystemConfigurationPatchInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemConfigurationPatchInner{}

// SystemConfigurationPatchInner struct for SystemConfigurationPatchInner
type SystemConfigurationPatchInner struct {
	Key string `json:"key"`
	Value NullableSystemConfigurationValue `json:"value"`
}

type _SystemConfigurationPatchInner SystemConfigurationPatchInner

// NewSystemConfigurationPatchInner instantiates a new SystemConfigurationPatchInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemConfigurationPatchInner(key string, value NullableSystemConfigurationValue) *SystemConfigurationPatchInner {
	this := SystemConfigurationPatchInner{}
	this.Key = key
	this.Value = value
	return &this
}

// NewSystemConfigurationPatchInnerWithDefaults instantiates a new SystemConfigurationPatchInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemConfigurationPatchInnerWithDefaults() *SystemConfigurationPatchInner {
	this := SystemConfigurationPatchInner{}
	return &this
}

// GetKey returns the Key field value
func (o *SystemConfigurationPatchInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SystemConfigurationPatchInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SystemConfigurationPatchInner) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for SystemConfigurationValue will be returned
func (o *SystemConfigurationPatchInner) GetValue() SystemConfigurationValue {
	if o == nil || o.Value.Get() == nil {
		var ret SystemConfigurationValue
		return ret
	}

	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfigurationPatchInner) GetValueOk() (*SystemConfigurationValue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// SetValue sets field value
func (o *SystemConfigurationPatchInner) SetValue(v SystemConfigurationValue) {
	o.Value.Set(&v)
}

func (o SystemConfigurationPatchInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemConfigurationPatchInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value.Get()
	return toSerialize, nil
}

func (o *SystemConfigurationPatchInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"value",
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

	varSystemConfigurationPatchInner := _SystemConfigurationPatchInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemConfigurationPatchInner)

	if err != nil {
		return err
	}

	*o = SystemConfigurationPatchInner(varSystemConfigurationPatchInner)

	return err
}

type NullableSystemConfigurationPatchInner struct {
	value *SystemConfigurationPatchInner
	isSet bool
}

func (v NullableSystemConfigurationPatchInner) Get() *SystemConfigurationPatchInner {
	return v.value
}

func (v *NullableSystemConfigurationPatchInner) Set(val *SystemConfigurationPatchInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigurationPatchInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigurationPatchInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigurationPatchInner(val *SystemConfigurationPatchInner) *NullableSystemConfigurationPatchInner {
	return &NullableSystemConfigurationPatchInner{value: val, isSet: true}
}

func (v NullableSystemConfigurationPatchInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigurationPatchInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


