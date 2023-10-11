/*
Anchore API

This is the Anchore API. Provides the external API for users of Anchore Enterprise.

API version: 2.0.0
Contact: dev@anchore.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"encoding/json"
)

// checks if the ECSServices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSServices{}

// ECSServices Services defined in ECS
type ECSServices struct {
	Services []ECSServicesServicesInner `json:"services,omitempty"`
}

// NewECSServices instantiates a new ECSServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSServices() *ECSServices {
	this := ECSServices{}
	return &this
}

// NewECSServicesWithDefaults instantiates a new ECSServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSServicesWithDefaults() *ECSServices {
	this := ECSServices{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ECSServices) GetServices() []ECSServicesServicesInner {
	if o == nil || IsNil(o.Services) {
		var ret []ECSServicesServicesInner
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServices) GetServicesOk() ([]ECSServicesServicesInner, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ECSServices) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ECSServicesServicesInner and assigns it to the Services field.
func (o *ECSServices) SetServices(v []ECSServicesServicesInner) {
	o.Services = v
}

func (o ECSServices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSServices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableECSServices struct {
	value *ECSServices
	isSet bool
}

func (v NullableECSServices) Get() *ECSServices {
	return v.value
}

func (v *NullableECSServices) Set(val *ECSServices) {
	v.value = val
	v.isSet = true
}

func (v NullableECSServices) IsSet() bool {
	return v.isSet
}

func (v *NullableECSServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSServices(val *ECSServices) *NullableECSServices {
	return &NullableECSServices{value: val, isSet: true}
}

func (v NullableECSServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


